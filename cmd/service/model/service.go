package model

import (
	"db_novel_service/pkg/config"
	"db_novel_service/pkg/db"
	"db_novel_service/pkg/log"
	router2 "db_novel_service/pkg/router"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Service struct {
	Log    *zerolog.Logger
	Router *mux.Router
	DB     *gorm.DB
	Config *config.Config
}

func NewService() *Service {
	logger := log.NewLogger()

	router := router2.NewRouter()

	db, err := db.InitDB()

	if err != nil {
		logger.Error().Msg("error to get db")
	}

	config := config.NewConfig()

	return &Service{
		Log:    logger,
		Router: router,
		DB:     db,
		Config: config,
	}
}
