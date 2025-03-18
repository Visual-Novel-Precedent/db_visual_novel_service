package model

import (
	"db_novel_service/pkg/config"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// MockService реализует интерфейс ServiceGetter для тестов
type MockService struct {
	GetDBFunc     func() *gorm.DB
	GetLoggerFunc func() *zerolog.Logger
	GetRouterFunc func() *mux.Router
	GetConfigFunc func() *config.Config
}

func (m *MockService) GetDB() *gorm.DB {
	return m.GetDBFunc()
}

func (m *MockService) GetLogger() *zerolog.Logger {
	return m.GetLoggerFunc()
}

func (m *MockService) GetRouter() *mux.Router {
	return m.GetRouterFunc()
}

func (m *MockService) GetConfig() *config.Config {
	return m.GetConfigFunc()
}
