package server

import (
	"errors"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"

	"github.com/mwbrown/nagbot/auth"
	"github.com/mwbrown/nagbot/config"
	"github.com/mwbrown/nagbot/db"
	"github.com/mwbrown/nagbot/db/nbsql"
	"github.com/mwbrown/nagbot/db/nbsql/config"
	"github.com/mwbrown/nagbot/db/nbsql/users"
	"github.com/mwbrown/nagbot/nbproto"
)

var (
	NotImplementedError = errors.New("Not Implemented")
	AuthenticationError = errors.New("Not Authorized")
	LoginRejectedError  = errors.New("Login rejected.")

	// Server configuration errors
	ServerAddrMissingError   = errors.New("Server listen address not set.")
	ServerPortMissingError   = errors.New("Server port not set or invalid.")
	ServerSecretMissingError = errors.New("Server secret not set.")
)

type NagbotServer struct {
	addr   string
	db     nbsql.DB
	secret []byte
	auth   *auth.Authenticator
}

//
// Helper functions
//

func checkDbVersion(db nbsql.DB) error {

	cfgRows, err := nbsql_config.Query(db, nil)
	if err != nil {
		return err
	}

	if len(cfgRows) != 1 {
		return fmt.Errorf("Invalid number of config rows returned (%d)", len(cfgRows))
	}

	ver := cfgRows[0].SchemaVer
	if ver > ndb.LATEST_SCHEMA_VERSION {
		return fmt.Errorf("Database version (%d) is newer than server (%d). Server upgrade required.", ver, ndb.LATEST_SCHEMA_VERSION)
	} else if ver < ndb.LATEST_SCHEMA_VERSION {
		return fmt.Errorf("Database version (%d) is older than server (%d). Database upgrade required.", ver, ndb.LATEST_SCHEMA_VERSION)
	}

	return nil
}

//
// GRPC Handlers
//

func (nb *NagbotServer) Login(ctx context.Context, req *nbproto.LoginRequest) (*nbproto.LoginResponse, error) {

	var remoteIp string

	// FIXME: There is currently no protection against repeated (or parallel) login attempts.
	//        It may be useful to create a separate login service that can perform both user
	//        authentication and user creation with the appropriate precautions.

	if peer, ok := peer.FromContext(ctx); ok {
		remoteIp = peer.Addr.String()
	} else {
		remoteIp = "<unknown>"
	}

	log.Printf("[%s] Login request for '%s'", remoteIp, req.Username)

	// Attempt to locate the user.
	users, err := nbsql_users.Query(nb.db, nbsql_users.UsernameCol.Equals(req.Username))
	if err != nil || len(users) != 1 {
		log.Printf("User \"%s\" not found (err=%v).", req.Username, err)
		return nil, LoginRejectedError
	}

	user := users[0]

	// Validate the password.
	calcPwHash := auth.UserPasswordHash(req.Password, user.PwSalt)

	if user.PwHash != calcPwHash {
		log.Printf("User \"%s\" password invalid.", req.Username)
		return nil, LoginRejectedError
	}

	sessId := user.NextSessID
	token, err := nb.auth.CreateToken(user.ID, user.NextSessID)
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return nil, LoginRejectedError
	}

	// Update the user's outstanding session count.
	user.NextSessID += 1

	err = nbsql_users.Update(nb.db, user)
	if err != nil {
		log.Printf("Error updating user session ID: %v", err)
	}

	log.Printf("Created login token for %s, id=%d", req.Username, sessId)
	return &nbproto.LoginResponse{Token: token}, nil
}

func (nb *NagbotServer) Logout(ctx context.Context, req *nbproto.LogoutRequest) (*nbproto.LogoutResponse, error) {
	log.Printf("Logout request.")

	user, err := nb.auth.RequireAuth(ctx)
	if err != nil {
		return nil, AuthenticationError
	}

	log.Printf("Logout request for user %d (%s)\n", user.ID, user.Username)
	return nil, NotImplementedError
}

func (nb *NagbotServer) CheckLogin(ctx context.Context, req *nbproto.CheckLoginRequest) (*nbproto.CheckLoginResponse, error) {
	log.Printf("Check login request.")

	user, err := nb.auth.RequireAuth(ctx)
	if err != nil {
		log.Println("Token auth failed:", err)
		return nil, AuthenticationError
	}

	log.Printf("Valid login check for user %s\n", user.Username)
	return &nbproto.CheckLoginResponse{}, nil
}

func (nb *NagbotServer) AddTaskDef(ctx context.Context, req *nbproto.AddTaskDefRequest) (*nbproto.AddTaskDefResponse, error) {
	return nil, NotImplementedError
}

func (nb *NagbotServer) DelTaskDef(ctx context.Context, req *nbproto.DelTaskDefRequest) (*nbproto.DelTaskDefResponse, error) {
	return nil, NotImplementedError
}

//
// Public API
//

func NewServer() (*NagbotServer, error) {

	var err *multierror.Error

	host := viper.GetString(nbconfig.CFG_KEY_RPC_LISTEN)
	port := viper.GetInt(nbconfig.CFG_KEY_RPC_PORT)
	secret := viper.GetString(nbconfig.CFG_KEY_RPC_SECRET)

	if len(host) == 0 {
		err = multierror.Append(err, ServerAddrMissingError)
	}

	if port < 0 || port > 65535 {
		err = multierror.Append(err, ServerPortMissingError)
	}

	if len(secret) == 0 {
		err = multierror.Append(err, ServerSecretMissingError)
	}

	if err != nil {
		return nil, err.ErrorOrNil()
	}

	db, e := ndb.Open()
	if e != nil {
		return nil, e
	}

	// Load the current schema version and ensure it is up-to-date.
	e = checkDbVersion(db)
	if e != nil {
		return nil, e
	}

	listenAddr := fmt.Sprintf("%s:%d", host, port)

	server := &NagbotServer{
		addr:   listenAddr,
		db:     db,
		secret: []byte(secret),
	}

	server.auth = auth.NewAuthenticator(server, server.secret)

	return server, nil
}

func (nb *NagbotServer) GetDB() nbsql.DB {
	return nb.db
}

func (nb *NagbotServer) ServeGrpc() error {

	log.Printf("Starting Nagbot Server.\n")

	sock, err := net.Listen("tcp", nb.addr)
	if err != nil {
		return err
	}

	log.Printf("Listening on %v", nb.addr)

	grpcServer := grpc.NewServer()
	nbproto.RegisterNagbotServer(grpcServer, nb)

	// TODO: Launch this in its own goroutine, listen for ctrl-c for proper cleanup.
	return grpcServer.Serve(sock)
}
