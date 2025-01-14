package main

import (
	"db_novel_service/cmd/service/model"
	admin_ "db_novel_service/internal/transport/handlers/admin "
	"db_novel_service/internal/transport/handlers/chapter"
	"db_novel_service/internal/transport/handlers/character"
	"db_novel_service/internal/transport/handlers/node"
	player_ "db_novel_service/internal/transport/handlers/player"
	"db_novel_service/internal/transport/handlers/request"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

const (
	PORT = "APP_PORT"
)

func main() {
	service := run()

	service.Router.HandleFunc("/create-chapter", func(w http.ResponseWriter, r *http.Request) {
		handler := chapter.CreateChapterHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/update-chapter", func(w http.ResponseWriter, r *http.Request) {
		handler := chapter.UpdateChapterHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/get-chapter", func(w http.ResponseWriter, r *http.Request) {
		handler := chapter.GetChaptersByUserIdHandler(service.DB)
		handler.ServeHTTP(w, r)
	})

	service.Router.HandleFunc("/create-node", func(w http.ResponseWriter, r *http.Request) {
		handler := node.CreateNodeHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/update-node", func(w http.ResponseWriter, r *http.Request) {
		handler := node.UpdateNodeHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/delete-node", func(w http.ResponseWriter, r *http.Request) {
		handler := node.DeleteNodeHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/get-nodes", func(w http.ResponseWriter, r *http.Request) {
		handler := node.GetNodeByChapterIdHandler(service.DB)
		handler.ServeHTTP(w, r)
	})

	service.Router.HandleFunc("/admin-authorization", func(w http.ResponseWriter, r *http.Request) {
		handler := admin_.AdminAuthorisationHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/admin-authorization", func(w http.ResponseWriter, r *http.Request) {
		handler := admin_.AdminRegistrationHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/admin-authorization", func(w http.ResponseWriter, r *http.Request) {
		handler := admin_.ChangeAdminHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/admin-update-chapter-progress", func(w http.ResponseWriter, r *http.Request) {
		handler := admin_.AdminChapterProgressHandler(service.DB)
		handler.ServeHTTP(w, r)
	})

	service.Router.HandleFunc("/player-authorization", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.PlayerAuthorisationHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/player-registration", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.PlayerRegistrationHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/player-update-chapter-progress", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.PlayerChapterProgressHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/player-update", func(w http.ResponseWriter, r *http.Request) {
		handler := player_.ChangePlayerRequestHandler(service.DB)
		handler.ServeHTTP(w, r)
	})

	service.Router.HandleFunc("/approve-request", func(w http.ResponseWriter, r *http.Request) {
		handler := request.ApproveRequestHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/create-request", func(w http.ResponseWriter, r *http.Request) {
		handler := request.CreateRequestHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/my-requests", func(w http.ResponseWriter, r *http.Request) {
		handler := request.GetMyRequestHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/to-me-requests", func(w http.ResponseWriter, r *http.Request) {
		handler := request.GetReceivedRequestHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/reject-requests", func(w http.ResponseWriter, r *http.Request) {
		handler := request.RejectRequestHandler(service.DB)
		handler.ServeHTTP(w, r)
	})

	service.Router.HandleFunc("/create-character", func(w http.ResponseWriter, r *http.Request) {
		handler := character.CreateCharacterHandler(service.DB)
		handler.ServeHTTP(w, r)
	})
	service.Router.HandleFunc("/update-character", func(w http.ResponseWriter, r *http.Request) {
		handler := character.UpdateCharacterHandler(service.DB)
		handler.ServeHTTP(w, r)
	})

	port := os.Getenv(PORT)
	if port == "" {
		port = "8080" // Default port if not set
	}

	// Запуск сервера
	http.ListenAndServe(port, service.Router)
}

func run() *model.Service {
	service := model.NewService()

	service.Log.Info().Msg("service is created ")

	return service
}
