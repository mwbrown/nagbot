package auth

import (
	"encoding/json"
	"errors"

	"github.com/mwbrown/nagbot/config"
	"github.com/mwbrown/nagbot/nbproto"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"io/ioutil"
	"os"
	"path"
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

var (
	InvalidTokenTypeError = errors.New("Token stored in Viper as incorrect type.")
)

// SaveToFile writes the current Nagbot RPC token to a file that can be read by later
// invocations of the command-line utility. This is used primarily with the login command
// handler, to save the resulting token upon a successful login.
func SaveToFile() error {

	token, ok := viper.Get(nbconfig.CFG_KEY_RPC_TOKEN).(string)
	if !ok {
		return InvalidTokenTypeError
	}

	// Build the JSON data to save to the config file.
	keymap := map[string]string{nbconfig.CFG_KEY_RPC_TOKEN: token}
	data, err := json.MarshalIndent(keymap, "", "    ")
	data = append(data, '\n')
	if err != nil {
		return err
	}

	// Create the config directory if it doesn't exist.
	// FIXME: This is currently the Linux XDG-style config path...
	//        as a result it is not where we should put it on Windows
	confdir, err := homedir.Expand("~/.config/nagbot")
	if err != nil {
		return err
	}

	err = os.MkdirAll(confdir, 0700)
	if err != nil {
		return err
	}

	// Attempt to save the file
	filename := path.Join(confdir, "auth.json")
	return ioutil.WriteFile(filename, data, 0600)
}
