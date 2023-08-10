package port

import "net/http"

type CommentHandlerInterface interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	GetAllComment(w http.ResponseWriter, r *http.Request)
}