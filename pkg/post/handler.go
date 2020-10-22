package post

import (
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
)

type Service interface {
	ShowPosts(w http.ResponseWriter, req *http.Request, _ httprouter.Params) ()
	ShowPostById(w http.ResponseWriter, req *http.Request, ps httprouter.Params) ()
	CreatePost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) ()
	EditPost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) ()
	DeletePost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) ()
}

type service struct {
	conn *pgx.Conn
}

func NewService(conn *pgx.Conn) Service {
	return &service {
		conn: conn,
	}
}

// ShowPosts "/posts"
func(s *service) ShowPosts(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "ShowPosts")
}

// ShowPostById "/posts/:id"
func(s *service) ShowPostById(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "ShowPostById")
}

// CreatePost "/posts"
func(s *service) CreatePost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "CreatePost")
}

// EditPost "/posts/:id"
func(s *service) EditPost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "EditPost")
}

// DeletePost "/posts/:id"
func(s *service) DeletePost(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "DeletePost")
}


