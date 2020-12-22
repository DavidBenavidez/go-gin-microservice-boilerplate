package config

import (
	"github.com/davidbenavidez/chi-gorm/internal/log"
	. "github.com/davidbenavidez/chi-gorm/internal/project"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfiguration struct {
	DBDriver, DBName, Username, Password, Host, Port string
	LogMode                                          bool
}

func (s *Server) setupDatabase(config DBConfiguration) (*gorm.DB, error) {
	dbconfig := makeDBConfig(config)

	db, err := connectToDatabase(dbconfig)

	db.SingularTable(true)

	// Optional: Create tables
	db.DropTable(&Project{})
	db.CreateTable(&Project{})

	return db, err
}

func makeDBConfig(config DBConfiguration) DBConfiguration {
	dbConfig := DBConfiguration{
		Username: config.Username,
		Password: config.Password,
		Host:     config.Host,
		Port:     config.Port,
		DBName:   config.DBName,
		LogMode:  config.LogMode,
	}

	return dbConfig
}

func connectToDatabase(config DBConfiguration) (*gorm.DB, error) {

	connectionString := config.Username +
		":" +
		config.Password +
		"@tcp(" +
		config.Host +
		":" +
		config.Port +
		")/" +
		config.DBName +
		"?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
		return nil, err
	}

	log.Infoln("Connected to database")
	return db, nil
}
