package nbconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	InvalidTokenTypeError = errors.New("Token stored in Viper as incorrect type.")
	BadBytesReaderError   = errors.New("Could not create reader for file data.")
)

func getConfigDir() (string, error) {
	// FIXME: This is currently the Linux XDG-style config path...
	//        as a result it is not where we should put it on Windows
	return homedir.Expand("~/.config/nagbot")
}

func getConfigFilename(baseDir string) string {
	return path.Join(baseDir, "auth.json")
}

// SaveAuthFile writes the current Nagbot RPC token to a file that can be read by later
// invocations of the command-line utility. This is used primarily with the login command
// handler, to save the resulting token upon a successful login.
func SaveAuthFile() error {

	token, ok := viper.Get(CFG_KEY_RPC_TOKEN).(string)
	if !ok {
		return InvalidTokenTypeError
	}

	// Build the JSON data to save to the config file.
	keymap := map[string]string{CFG_KEY_RPC_TOKEN: token}
	data, err := json.MarshalIndent(keymap, "", "    ")
	data = append(data, '\n')
	if err != nil {
		return err
	}

	// Create the config directory if it doesn't exist.
	confdir, err := getConfigDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(confdir, 0700)
	if err != nil {
		return err
	}

	// Attempt to save the file
	filename := getConfigFilename(confdir)
	return ioutil.WriteFile(filename, data, 0600)
}

// LoadAuthFile loads a file called "auth.json" from the Nagbot user
// configuration directory (on Linux, this is ~/.config/nagbot/). This
// file contains just the RPC token and is written when using the
// `nagbot login` command, to store the active authentication token
// separately from the rest of the static configuration.
func LoadAuthFile() error {

	confdir, err := getConfigDir()
	if err != nil {
		return err
	}

	filename := getConfigFilename(confdir)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	dataRdr := bytes.NewReader(data)
	if dataRdr == nil {
		return BadBytesReaderError
	}

	viper.SetConfigType("json")
	err = viper.ReadConfig(dataRdr)

	return err
}
