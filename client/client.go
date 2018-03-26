package client

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/mwbrown/nagbot/auth"
	"github.com/mwbrown/nagbot/config"
	"github.com/mwbrown/nagbot/nbproto"

	"github.com/spf13/viper"

	"github.com/hashicorp/go-multierror"
)

var (
	RpcHostMissingError = errors.New("Nagbot server hostname not set.")
	RpcPortMissingError = errors.New("Nagbot server port not set or invalid.")
)

type Client struct {
	rpc        nbproto.NagbotClient
	conn       *grpc.ClientConn
	serverAddr string
	authToken  string
}

func NewClient() (*Client, error) {

	var err *multierror.Error

	// Retrieve the configuration parameters for the server.
	host := viper.GetString(nbconfig.CFG_KEY_RPC_HOST)
	port := viper.GetInt(nbconfig.CFG_KEY_RPC_PORT)
	token := viper.GetString(nbconfig.CFG_KEY_RPC_TOKEN)

	if len(host) == 0 {
		err = multierror.Append(err, RpcHostMissingError)
	}

	if port < 0 || port > 65535 {
		err = multierror.Append(err, RpcPortMissingError)
	}

	if err != nil {
		return nil, err.ErrorOrNil()
	}

	// Set up the server address to be used with Open later.
	c := &Client{
		serverAddr: fmt.Sprintf("%s:%d", host, port),
		authToken:  token,
	}

	return c, nil
}

func (c *Client) Open() error {

	var dialOptions []grpc.DialOption

	// TODO: Add TLS support, simultaneously change credentials to require TLS
	dialOptions = append(dialOptions, grpc.WithInsecure())

	if len(c.authToken) != 0 {
		creds := &auth.NagbotCreds{Token: c.authToken}
		dialOptions = append(dialOptions, grpc.WithPerRPCCredentials(creds))
	}

	conn, err := grpc.Dial(c.serverAddr, dialOptions...)
	if err != nil {
		return err
	}

	c.conn = conn
	c.rpc = nbproto.NewNagbotClient(conn)

	return nil
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	c.conn = nil
	c.rpc = nil
}

func (c *Client) Login(username string, password string) (string, error) {
	resp, err := c.rpc.Login(context.Background(), &nbproto.LoginRequest{Username: username, Password: password})

	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (c *Client) Logout() error {
	_, err := c.rpc.Logout(context.Background(), &nbproto.LogoutRequest{})
	return err
}

func (c *Client) CheckLogin() error {
	_, err := c.rpc.CheckLogin(context.Background(), &nbproto.CheckLoginRequest{})
	return err
}
