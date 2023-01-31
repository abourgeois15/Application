package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DB_User   string `yaml:"MYSQL_USER"`
	DB_Driver string `yaml:"MYSQL_DRIVER"`
	DB_Name   string `yaml:"MYSQL_DATABASE"`
	DB_Pass   string `yaml:"MYSQL_ROOT_PASSWORD"`
	DB_Host   string `yaml:"MYSQL_HOST"`
	DB_Port   string `yaml:"MYSQL_PORT"`
}

type Config_Test struct {
	DB_User   string `yaml:"MYSQL_USER_TEST"`
	DB_Driver string `yaml:"MYSQL_DRIVER_TEST"`
	DB_Name   string `yaml:"MYSQL_DATABASE_TEST"`
	DB_Pass   string `yaml:"MYSQL_ROOT_PASSWORD_TEST"`
	DB_Host   string `yaml:"MYSQL_HOST_TEST"`
	DB_Port   string `yaml:"MYSQL_PORT_TEST"`
}

type Queries struct {
	db *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

var (
	db *sql.DB
)

func ConnectMySQLDB() {

	confContent, err := os.ReadFile("c:\\Users\\bor6rt\\go\\Application\\api\\conf.yaml")
	if err != nil {
		panic(err)
	}

	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}

	dbDriver := conf.DB_Driver
	dbUser := conf.DB_User
	dbPass := conf.DB_Pass
	dbName := conf.DB_Name
	dbHost := conf.DB_Host
	dbPort := conf.DB_Port
	fmt.Println("Connecting to database...")

	d, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)

	limit := 15
	// check that the database is reachable; try at least 3 times to connect
	for i := 0; i <= limit; i++ {
		err := d.Ping()
		if err != nil && i == limit {
			_ = fmt.Errorf("couldn't connect to database after %d tries: %s", i, err)
			break
		} else if err != nil {
			log.Info("Couldn't connect to database, retrying in 1 second ...")
			time.Sleep(5 * time.Second)
		} else {
			log.Info("Successfully connected to database")
			break
		}
	}

	if err != nil {
		panic(err.Error())
	}

	db = d
}

func ConnectMySQLDBTest() {
	confContent, err := os.ReadFile("c:\\Users\\bor6rt\\go\\Application\\api\\conf.yaml")
	if err != nil {
		panic(err)
	}

	conf := &Config_Test{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}

	dbDriver := conf.DB_Driver
	dbUser := conf.DB_User
	dbPass := conf.DB_Pass
	dbName := conf.DB_Name
	dbHost := conf.DB_Host
	dbPort := conf.DB_Port
	fmt.Println("Connecting to database...")

	d, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)

	limit := 15
	// check that the database is reachable; try at least 3 times to connect
	for i := 0; i <= limit; i++ {
		err := d.Ping()
		if err != nil && i == limit {
			_ = fmt.Errorf("couldn't connect to database after %d tries: %s", i, err)
			break
		} else if err != nil {
			log.Info("Couldn't connect to database, retrying in 1 second ...")
			time.Sleep(5 * time.Second)
		} else {
			log.Info("Successfully connected to database")
			break
		}
	}

	if err != nil {
		panic(err.Error())
	}

	db = d
}

func GetMySQLDB() (*sql.DB, error) {
	return db, nil
}
