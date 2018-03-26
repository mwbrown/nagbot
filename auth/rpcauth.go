package auth

import (
	"github.com/mwbrown/nagbot/nbproto"
	"golang.org/x/net/context"
)

type NagbotCreds struct {
	Token string
}

func (c *NagbotCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	m := make(map[string]string)
	m[nbproto.TOKEN_METADATA_KEY] = c.Token
	return m, nil
}

func (c *NagbotCreds) RequireTransportSecurity() bool {
	// TODO: this needs to be changed to true for prod
	return false
}
