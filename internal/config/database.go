package config

import (
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/project"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type DBConfiguration struct {
	DBDriver, DBName, Username, Password, Host, Port string
	LogMode                                          bool
}

func (s *Server) setupDatabase(dbconfig DBConfiguration) (*gorm.DB, error) {
	db, err := connectToDatabase(dbconfig)

	db.SingularTable(true)

	// Optional: Create tables
	db.DropTable(&project.Project{})
	db.CreateTable(&project.Project{})

	return db, err
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
