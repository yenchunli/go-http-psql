package user

import (
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v4"
)

type Service interface {
	Login(w http.ResponseWriter, req *http.Request) ()
	Logout(w http.ResponseWriter, req *http.Request) ()
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
func(s *service) Login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Login")
}

// Logout "/logout"
func(s *service) Logout(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Logout")
}
