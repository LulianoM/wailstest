package internal

import (
	"database/sql"
	"fmt"
	"log"

	"{{module_name}}/src/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host        string `env:"DATABASE_HOST"`
	User        string `env:"DATABASE_USER"`
	Password    string `env:"DATABASE_PASSWORD"`
	Name        string `env:"DATABASE_NAME"`
	Port        string `env:"DATABASE_PORT"`
	IsMigration bool   `env:"IS_MIGRATION"`
}

type DBClient struct {
	defaultDB *gorm.DB
	Conn      *gorm.DB
}

var (
	globalDB *DBClient
)

func Init(conf *DBConfig) (*DBClient, error) {
	if globalDB != nil {
		return globalDB, nil
	}

	dbConfig := gorm.Config{}

	dialect := postgres.Open(conf.getDSN())
	db, err := gorm.Open(dialect, &dbConfig)
	if err != nil {
		panic(err)
	}

	globalDB = &DBClient{
		defaultDB: db,
		Conn:      db,
	}

	return globalDB, nil
}

func InitDatabaseClient() (dbc *DBClient) {
	cfg := GetConfig()

	dbc, err := Init(&cfg.Database)
	if err != nil {
		log.Panicf("Failed to connect to database. Error: %s", err.Error())
	}

	tables := database.Tables

	err = dbc.DropTables(tables...)
	if err != nil {
		log.Panicf("Failed to DropTable database. Error: %s", err.Error())
	}

	dbc.Conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	dbc.Conn.Exec("CREATE OR REPLACE FUNCTION unaccent_string(text) RETURNS text LANGUAGE sql IMMUTABLE STRICT AS $$SELECT TRANSLATE ($1, 'âãäåāăąáàÁÂÃÄÅĀĂĄÀèéêëēĕėęěÉÊĒĔĖĘĚìíîïìĩīĭÌÍÎÏÌĨĪĬóôõöōŏőÒÓÔÕÖŌŎŐùúûüũūŭůÙÚÛÜŨŪŬŮçÇ',	'aaaaaaaaaaaaaaaaaaeeeeeeeeeeeeeeeeiiiiiiiiiiiiiiiiooooooooooooooouuuuuuuuuuuuuuuucc')$$")

	err = dbc.MigrateTables(tables...)
	if err != nil {
		fmt.Printf("Failed to migrate database. Error: %s\n", err.Error())
		log.Panicf("Failed to migrate database. Error: %s", err.Error())
	}

	return
}

func (this *DBConfig) getDSN() string {
	host := this.Host
	port := this.Port

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		this.User,
		this.Password,
		this.Name,
		port,
	)
}

func (this *DBClient) MigrateTables(tables ...interface{}) error {
	return this.Conn.Migrator().AutoMigrate(tables...)
}

func (this *DBClient) GetDB() (*sql.DB, error) {
	return this.Conn.DB()
}

func (this *DBClient) DropTables(tables ...interface{}) error {
	for _, table := range tables {
		if err := this.Conn.Migrator().DropTable(table); err != nil {
			return err
		}
	}
	return nil
}
