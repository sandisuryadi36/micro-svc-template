package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/sandisuryadi36/micro-svc-template/server/pb"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbMain    *gorm.DB
	dbMainSQL *sql.DB
)

func startDBConnection() {
	log.Printf("Starting Db Connections...")

	initDBMain()

}

func initDBMain() {
	log.Printf("Main Db - Connecting")
	var err error
	dbMain, err = gorm.Open(postgres.Open(GetEnv("DB_DSN", "")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connect to DB main: %v", err)
		os.Exit(1)
		return
	}

	dbMainSQL, err = dbMain.DB()
	if err != nil {
		log.Fatalf("Error cannot initiate connection to DB main: %v", err)
		os.Exit(1)
		return
	}

	dbMainSQL.SetMaxIdleConns(0)
	dbMainSQL.SetMaxOpenConns(0)

	err = dbMainSQL.Ping()
	if err != nil {
		log.Fatalf("Cannot ping DB main: %v", err)
		os.Exit(1)
		return
	}

	log.Printf("Main Db - Connected")
}

func closeDBMain() {
	log.Print("Closing DB Main Connection ... ")
	if err := dbMainSQL.Close(); err != nil {
		log.Fatalf("Error on disconnection with DB Main : %v", err)
	}
	log.Println("Closing DB Main Success")
}

func migrateDB() error {
	initDBMain()
	defer closeDBMain()

	migrator := dbMain.Migrator()
	if migrator.HasTable(
		// List table ORM from proto gorm
		&pb.ExampleORM{},
	) {
		log.Println("Table already exists, no migration needed")
		return nil
	}

	log.Println("Migration process begin...")
	if err := dbMain.AutoMigrate(
		// List table from proto gorm
		&pb.ExampleORM{},
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
		os.Exit(1)
	}

	log.Println("Migration process finished...")

	return nil
}
