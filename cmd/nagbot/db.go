package main

import (
	"fmt"
	"os"

	"github.com/mwbrown/nagbot/auth"
	"github.com/mwbrown/nagbot/ndb"
	"github.com/mwbrown/nagbot/ndb/nbsql/config"
	"github.com/mwbrown/nagbot/ndb/nbsql/users"

	"github.com/spf13/cobra"
)

var dbRootCmd = &cobra.Command{
	Use:   "db",
	Short: "Performs Nagbot database operations.",
}

var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a fresh Nagbot database.",
	Run:   dbInitHandler,
}

var dbUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades an existing Nagbot database.",
	Run:   dbUpgradeHandler,
}

func init() {

	dbInitCmd.PersistentFlags().Bool("force", false, "Forces deletion of an existing database's data and schema.")

	dbRootCmd.AddCommand(dbInitCmd, dbUpgradeCmd)
	rootCmd.AddCommand(dbRootCmd)
}

func verifyDbOptions(exitOnError bool) error {

	fmt.Println("Checking DB configuration...")
	e := ndb.VerifyDbOptions()

	if e != nil {

		fmt.Println("DB configuration error found:", e)

		if exitOnError {
			os.Exit(1)
		}
	}

	return e
}

func dbInitHandler(cmd *cobra.Command, args []string) {
	verifyDbOptions(true)

	// Depending on whether this command is forced or not, we need to know
	// if the database has already been initialized.
	var dbExists bool
	dbForce, err := cmd.PersistentFlags().GetBool("force")

	// The previous line should not ever return an error.
	if err != nil {
		panic(err)
	}

	db, err := ndb.Open()
	if err != nil {
		fmt.Println("Could not open database connection:", err)
	}

	// Attempt to read the configuration value from the database.
	configRows, err := nbsql_config.Query(db, nil)
	if err == nil {
		dbExists = true

		if len(configRows) == 1 {
			cfg := configRows[0]
			fmt.Printf("Found existing database with schema version %d.\n", cfg.SchemaVer)
		} else {
			fmt.Printf("Warning: found %d config row(s) in database.")
		}
	}

	loader, err := ndb.NewSchemaLoader()
	if err != nil {
		fmt.Println("Could not load schema information:", err)
		os.Exit(1)
	}

	if dbExists {
		if !dbForce {
			fmt.Println("Database already exists. Re-run `db init` with -force to delete existing data.")
			os.Exit(1)
		}

		resetScript, err := loader.GetSchemaResetScript()

		if err != nil {
			fmt.Println("Could not load schema reset script:", err)
			os.Exit(1)
		}

		fmt.Println("Deleting existing database...")
		_, err = db.Query(resetScript)
		if err != nil {
			fmt.Println("Error running reset script:", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Creating initial schema (version %d)\n", ndb.LATEST_SCHEMA_VERSION)
	schemaScript, err := loader.GetSchemaInitScript(ndb.LATEST_SCHEMA_VERSION)

	_, err = db.Query(schemaScript)
	if err != nil {
		fmt.Println("Error running init script:", err)
		os.Exit(1)
	}

	fmt.Println("Schema created. Please enter initial admin user information.")
	user, pass, err := readUserPass(true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate a random salt value.
	salt, err := auth.GenerateSalt()
	if err != nil {
		fmt.Println("Could not generate salt value:", err)
		os.Exit(1)
	}

	// Generate the initial password hash.
	hash := auth.UserPasswordHash(pass, salt)

	// Create the initial user from the given information.
	r := &nbsql_users.Row{
		ID:         1, // TODO: the GNORM code needs to be modified to allow not specifying this or min/next sessid
		Username:   user,
		PwHash:     hash,
		PwSalt:     salt,
		IsAdmin:    true,
		IsEnabled:  true,
		MinSessID:  1,
		NextSessID: 1,
	}

	fmt.Println("Creating admin user...")
	if err := nbsql_users.Insert(db, r); err != nil {
		fmt.Println("Could not insert: ", err)
	}

	fmt.Println("Database initialized!")
}

func dbUpgradeHandler(cmd *cobra.Command, args []string) {
	verifyDbOptions(true)

	var currSchema int

	db, err := ndb.Open()
	if err != nil {
		panic(err)
	}

	// Attempt to read the configuration value from the database.
	configRows, err := nbsql_config.Query(db, nil)
	if err != nil {
		fmt.Println("Could not find database config row:", err)
		os.Exit(1)
	}

	if len(configRows) != 1 {
		fmt.Printf("Error: found %d config row(s) in database.")
		os.Exit(1)
	}

	currSchema = configRows[0].SchemaVer
	fmt.Println("Tool schema version: ", ndb.LATEST_SCHEMA_VERSION)
	fmt.Println("Remote DB schema:    ", currSchema)

	if currSchema >= ndb.LATEST_SCHEMA_VERSION {
		fmt.Println("Skipping upgrade, database is already at latest or newer.")
		os.Exit(0) // FIXME: Is this an appropriate exit for an up-to-date database?
	}

	loader, err := ndb.NewSchemaLoader()
	if err != nil {
		fmt.Println("Could not load schema information:", err)
		os.Exit(1)
	}

	// Perform incremental migration scripts until at the latest version.
	for currSchema < ndb.LATEST_SCHEMA_VERSION {
		nextSchema := currSchema + 1

		fmt.Printf("Attempting schema upgrade from %d to %d...", currSchema, nextSchema)

		upgradeScript, err := loader.GetSchemaUpgradeScript(nextSchema)
		if err != nil {
			fmt.Println("error.")
			fmt.Printf("Could not load upgrade script for schema %d: %v\n", nextSchema, err)
			os.Exit(1)
		}

		_, err = db.Query(upgradeScript)
		if err != nil {
			fmt.Println("error.")
			fmt.Printf("Could not upgrade to schema %d: %v\n", nextSchema, err)
			os.Exit(1)
		}

		fmt.Println("done.")
		currSchema = nextSchema
	}

	fmt.Println("Database upgrade complete!")
}
