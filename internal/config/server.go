package config

import (
	"net/http"

	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/log"
	"github.com/davidbenavidez/go-chi-mysql-boilerplate/internal/project"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Server struct {
	router *chi.Mux
	db     *gorm.DB
	Port   string
}

type Config struct {
	Server   Server
	Database DBConfiguration
}

var config *Config

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Setup handlers
func (s *Server) handleProject() http.HandlerFunc {
	return project.GetProjects(s.db)
}

/*
	func (s *Server) handleSomething() http.HandlerFunc {
		// Optional: Prepare Something
		return func(w http.ResponseWriter, r *http.Request) {
			// Do something
		}
	}
*/

func init() {
	// Get config from config file
	viper.AddConfigPath("./static")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func SetupServer() (*Server, string, error) {
	var err error
	s := &Server{}

	// Setup DB connection
	s.db, err = s.setupDatabase(config.Database)

	if err != nil {
		return nil, "", err
	}

	// inject handlers to routes
	s.setupRoutes()

	port := ":" + config.Server.Port
	return s, port, nil
}
