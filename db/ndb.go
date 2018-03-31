//
// This file contains the top level API for interfacing
// with the GNORM-generated code.
//

package ndb

import (
	"errors"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/mwbrown/nagbot/config"
	"github.com/mwbrown/nagbot/db/nbsql"

	"github.com/gobuffalo/packr"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
)

type SchemaLoader struct {
	schemaBox packr.Box
}

// Open returns a connection object to a Nagbot database,
// based on the configuration set in the application.
func Open() (nbsql.DB, error) {
	connStr := makeConnStr()
	return sql.Open("postgres", connStr)
}

var (
	// VerifyDbOptions errors

	SqlHostMissingError = errors.New("PostgreSQL server hostname not set.")
	SqlPortMissingError = errors.New("PostgreSQL server port not set or invalid.")
	SelUserMissingError = errors.New("PostgreSQL server username not set.")
	SqlPassMissingError = errors.New("PostgreSQL server password not set.")
	SqlDbMissingError   = errors.New("PostgreSQL server db name not set.")

	// SchemaLoader errors

	SchemaVersionRangeError error = errors.New("Schema version out of range.")
	SchemaFileNotFoundError error = errors.New("Schema definition not found.") // Receiving this indicates the build is not correct.
)

// Checks to see if the required pgsql options are present (and set).
func VerifyDbOptions() error {

	var err *multierror.Error

	if s := viper.GetString(nbconfig.CFG_KEY_PGSQL_HOST); len(s) == 0 {
		err = multierror.Append(err, SqlHostMissingError)
	}

	if s := viper.GetInt(nbconfig.CFG_KEY_PGSQL_PORT); s < 0 || s > 65535 {
		err = multierror.Append(err, SqlPortMissingError)
	}

	if s := viper.GetString(nbconfig.CFG_KEY_PGSQL_USER); len(s) == 0 {
		err = multierror.Append(err, SelUserMissingError)
	}

	if s := viper.GetString(nbconfig.CFG_KEY_PGSQL_PASS); len(s) == 0 {
		err = multierror.Append(err, SqlPassMissingError)
	}

	if s := viper.GetString(nbconfig.CFG_KEY_PGSQL_DB); len(s) == 0 {
		err = multierror.Append(err, SqlDbMissingError)
	}

	return err.ErrorOrNil()
}

func makeConnStr() string {
	host := viper.GetString(nbconfig.CFG_KEY_PGSQL_HOST)
	port := viper.GetInt(nbconfig.CFG_KEY_PGSQL_PORT)
	user := viper.GetString(nbconfig.CFG_KEY_PGSQL_USER)
	pass := viper.GetString(nbconfig.CFG_KEY_PGSQL_PASS)

	str := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, pass)

	return str
}

func NewSchemaLoader() (*SchemaLoader, error) {
	sl := &SchemaLoader{}

	sl.schemaBox = packr.NewBox("../schema")

	return sl, nil
}

func (sl *SchemaLoader) getSchemaFile(filename string) (string, error) {
	contents, err := sl.schemaBox.MustString(filename)

	if err != nil || len(contents) == 0 {
		return "", SchemaFileNotFoundError
	}

	return contents, nil
}

func (sl *SchemaLoader) getSchemaFormatted(version int, format string) (string, error) {
	if version <= 0 || version > LATEST_SCHEMA_VERSION {
		return "", SchemaVersionRangeError
	}

	filename := fmt.Sprintf(format, version)
	return sl.getSchemaFile(filename)
}

func (sl *SchemaLoader) GetSchemaResetScript() (string, error) {
	return sl.getSchemaFile("reset.sql")
}

func (sl *SchemaLoader) GetSchemaInitScript(version int) (string, error) {
	// Schema init scripts are named "xxxxx.sql", with just the version number as the file name.
	return sl.getSchemaFormatted(version, "%05d.sql")
}

func (sl *SchemaLoader) GetSchemaUpgradeScript(version int) (string, error) {
	// Schema upgrade scripts are named "xxxxx-upgrade.sql", with the version number in the file name.
	return sl.getSchemaFormatted(version, "%05d-upgrade.sql")
}
