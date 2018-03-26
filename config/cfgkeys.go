package nbconfig

import (
	"github.com/spf13/viper"
)

// Constants for configuration keys.
const (
	CFG_KEY_PGSQL_HOST = "pgsql_host" // Config key for PostgreSQL host (server only).
	CFG_KEY_PGSQL_PORT = "pgsql_port" // Config key for PostgreSQL port (server only).
	CFG_KEY_PGSQL_USER = "pgsql_user" // Config key for PostgreSQL user (server only).
	CFG_KEY_PGSQL_PASS = "pgsql_pass" // Config key for PostgreSQL password (server only).
	CFG_KEY_PGSQL_DB   = "pgsql_db"   // Config key for PostgreSQL database name (server only).
	CFG_KEY_RPC_SECRET = "rpc_secret" // Config key for RPC secret for JWT tokens (server only).
	CFG_KEY_RPC_LISTEN = "rpc_listen" // Config key for local RPC listen/bind address (server only).
	CFG_KEY_RPC_HOST   = "rpc_host"   // Config key for RPC port (client and server).
	CFG_KEY_RPC_PORT   = "rpc_port"   // Config key for remote RPC server (client only).
	CFG_KEY_RPC_TOKEN  = "rpc_token"  // Config key for RPC authentication token (client only).
)

// Constants for configuration defaults.
const (
	CFG_PGSQL_HOST_DEFAULT = "127.0.0.1" // Defaults to a PostgreSQL server that listens on localhost.
	CFG_PGSQL_PORT_DEFAULT = 5432        // Default PostgreSQL port.
	CFG_PGSQL_USER_DEFAULT = "nagbot"    // Default PostgreSQL username.
	CFG_PGSQL_DB_DEFAULT   = "nagbot"    // Default PostgreSQL database name.
	CFG_RPC_LISTEN_DEFAULT = "0.0.0.0"   // Defaults to listening on all interfaces.
	CFG_RPC_HOST_DEFAULT   = "127.0.0.1" // Defaults to a Nagbot server listening on localhost.
	CFG_RPC_PORT_DEFAULT   = 50051       // Default port (TODO: change from GRPC example port).
)

func setViperDefaults() {
	viper.SetDefault(CFG_KEY_PGSQL_HOST, CFG_PGSQL_HOST_DEFAULT)
	viper.SetDefault(CFG_KEY_PGSQL_PORT, CFG_PGSQL_PORT_DEFAULT)
	viper.SetDefault(CFG_KEY_PGSQL_USER, CFG_PGSQL_USER_DEFAULT)
	viper.SetDefault(CFG_KEY_PGSQL_DB, CFG_PGSQL_DB_DEFAULT)
	viper.SetDefault(CFG_KEY_RPC_LISTEN, CFG_RPC_LISTEN_DEFAULT)
	viper.SetDefault(CFG_KEY_RPC_HOST, CFG_RPC_HOST_DEFAULT)
	viper.SetDefault(CFG_KEY_RPC_PORT, CFG_RPC_PORT_DEFAULT)
}

func init() {
	setViperDefaults()
}
