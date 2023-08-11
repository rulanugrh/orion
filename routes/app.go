package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/orion/configs"
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/handler/port"
	"github.com/rulanugrh/orion/middleware"
)

func Run(event port.EventHandlerInterface, comment port.CommentHandlerInterface, user port.UserHandlerInterface) {
	db := configs.GetConnection()
	db.AutoMigrate(&domain.CommentEntity{}, &domain.EventEntity{}, &domain.UserEntity{})

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/api/v1/user/auth", user.Register).Methods("POST")
	router.HandleFunc("/api/v1/user/login", user.Login).Methods("POST")
	router.HandleFunc("/api/v1/user/refreshToken", user.RefreshToken).Methods("POST")

	routesHandler := router.PathPrefix("/api/v1/").Subrouter()
	routesHandler.Use(middleware.JWTVerify)
	routesHandler.Use(commonMiddleware)

	// routing for event
	routesHandler.HandleFunc("/event/createEvent", event.CreateEvent).Methods("POST")
	routesHandler.HandleFunc("/event/getEvent", event.GetEvent).Methods("GET")
	routesHandler.HandleFunc("/event/getEvent/{id}", event.GetEventByID).Methods("GET")
	routesHandler.HandleFunc("/event/updateEvent/{id}", event.UpdateEvent).Methods("PUT")
	routesHandler.HandleFunc("/event/deleteEvent/{id}", event.DeleteEvent).Methods("DELETE")

	// routing for comment
	routesHandler.HandleFunc("/comment/createComment", comment.CreateComment).Methods("POST")
	routesHandler.HandleFunc("/comment/listAllComment", comment.GetAllComment).Methods("GET")

	// routing for user
	routesHandler.HandleFunc("/user/updateUser/{id}", user.UpdateAccount).Methods("PUT")
	routesHandler.HandleFunc("/user/deleteUser/{id}", user.DeleteAccount).Methods("DELETE")
	routesHandler.HandleFunc("/user/detailUser/{id}", user.DetailUser).Methods("GET")
	routesHandler.HandleFunc("/user/joinEvent", user.JoinEvent).Methods("POST")
	routesHandler.HandleFunc("/user/refreshToken", user.RefreshToken).Methods("POST")

	config := configs.GetConfig()
	host := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Hport)
	server := http.Server{
		Addr:    host,
		Handler: router,
	}

	log.Printf("Server running on %v", host)
	server.ListenAndServe()
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Context-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
