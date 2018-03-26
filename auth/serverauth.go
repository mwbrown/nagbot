package auth

import (
	"crypto/sha256"
	"errors"
	"time"

	"github.com/mwbrown/nagbot/db/nbsql"
	"github.com/mwbrown/nagbot/db/nbsql/users"
	"github.com/mwbrown/nagbot/nbproto"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type Authenticator struct {
	dbProvider AuthDbProvider
	secretKey  []byte
}

type AuthDbProvider interface {
	GetDB() nbsql.DB
}

var (
	NoDatabaseError         = errors.New("No database was given for user lookup.")
	NoMetadataProvidedError = errors.New("No metadata provided with context.")
	NoTokenProvidedError    = errors.New("No auth token provided.")
	TokenMetadataError      = errors.New("Invalid data given as token metadata.")
	TokenInvalidError       = errors.New("Token is not valid.")
	TokenSignatureError     = errors.New("Token signature is not valid.")
	TokenSigningMethodError = errors.New("Token has incorrect signature method.")
	TokenClaimsFormatError  = errors.New("Claims are not in correct format.")
	TokenClaimsMissingError = errors.New("Required claims are not present.")
	TokenClaimsInvalidError = errors.New("Claims did not validate properly.")
	TokenClaimsTypeError    = errors.New("Claim(s) are not correct type.")
	NoSuchUserError         = errors.New("User not found in database.")
	NoSecretKeyError        = errors.New("Secret key was not provided.")
)

// NewAuthenticator returns an Authenticator instance given a database provider.
func NewAuthenticator(dbProvider AuthDbProvider, secret []byte) *Authenticator {

	// Ensure we own this copy of the authentication key.
	copySecret := make([]byte, len(secret))
	copy(copySecret, secret)

	return &Authenticator{
		dbProvider: dbProvider,
		secretKey:  copySecret,
	}
}

// Retrieves the user specified by the
func getUserById(id int, db nbsql.DB) *nbsql_users.Row {
	result, err := nbsql_users.Query(db, nbsql_users.IDCol.Equals(id))
	if err != nil {
		return nil
	}

	if len(result) != 1 {
		return nil
	}

	return result[0]
}

func (auth *Authenticator) CreateToken(userId int, sessId int) (string, error) {
	timestamp := time.Now().Unix()
	const lifetime = 7 * 24 * 60 * 60 // TODO: configurable expiry times

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NagbotClaims{
		UserId:  userId,
		Issued:  timestamp,
		Expires: timestamp + lifetime,
		SessId:  sessId,
	})

	tokenStr, err := token.SignedString(auth.secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// RequireAuth is called on an incoming context for a GRPC request for which a valid
// session token is required. It checks both the presence and validity of the token,
// and if successful, returns the user that is authenticated for this call.
func (auth *Authenticator) RequireAuth(ctx context.Context) (u *nbsql_users.Row, e error) {

	// Check for the presence of the token.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, NoMetadataProvidedError
	}

	tokenRaw, ok := md[nbproto.TOKEN_METADATA_KEY]
	if !ok {
		return nil, NoTokenProvidedError
	}

	// Metadata provides an array. However, this can only have one entry.
	if len(tokenRaw) != 1 {
		return nil, TokenMetadataError
	}

	// Validate the token and its signature using our custom claims type.
	token, err := jwt.ParseWithClaims(tokenRaw[0], &NagbotClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, TokenSigningMethodError
		}

		if len(auth.secretKey) == 0 || auth.secretKey == nil {
			return nil, NoSecretKeyError
		}

		return auth.secretKey, nil
	})

	if err != nil {
		return nil, err // TODO wrap err
	}

	// Validate the claims.
	claims, ok := token.Claims.(*NagbotClaims)
	if !ok {
		return nil, TokenClaimsFormatError
	}

	// Is the token valid?
	if !token.Valid {
		return nil, TokenClaimsInvalidError
	}

	db := auth.dbProvider.GetDB()
	if db == nil {
		return nil, NoDatabaseError
	}

	// Get the user to return.
	user := getUserById(claims.UserId, db)
	if user == nil {
		return nil, NoSuchUserError
	}

	return user, nil
}

func UserPasswordHash(password string, salt string) string {

	hash := sha256.New()

	hash.Write([]byte(password))
	hash.Write([]byte(salt))

	return string(hash.Sum(nil))
}
