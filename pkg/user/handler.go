package user

import (
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
)

type Service interface {
	Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) ()
	Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) ()
}

type service struct {
	conn *pgx.Conn
}

func NewService(conn *pgx.Conn) Service {
	return &service {
		conn: conn,
	}
}

// Login "/login"
func(s *service) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Login")
}

// Logout "/logout"
func(s *service) Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Logout")
}
