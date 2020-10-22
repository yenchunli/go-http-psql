package main

import (
	"fmt"
	"net/http"
	"github.com/go-http-psql/pkg/user"
	"log"
	"github.com/jackc/pgx/v4"
	"os"
	"context"
)

func main() {
	// Create PostgreSQL Connection
	// Using pgx
	var conn *pgx.Conn
	var err error
	conn, err = pgx.Connect(context.Background(), "postgresql://localhost/hello?user=hello&password=Hello123@")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	// Create User Service
	user_svc := user.NewService(conn)

	fmt.Println("===Server Start===")

	http.HandleFunc("/login", user_svc.Login)

	http_err := http.ListenAndServe(":9090", nil)

	if http_err != nil{
		log.Fatal("Error")
	}
}
