package main

import (
	"fmt"
	"net/http"
	"github.com/go-http-psql/pkg/user"
	"github.com/go-http-psql/pkg/post"
	"log"
	"github.com/jackc/pgx/v4"
	"os"
	"context"
	"github.com/julienschmidt/httprouter"
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

	// Create Service
	user_svc := user.NewService(conn)
	post_svc := post.NewService(conn)

	fmt.Println("===Server Start===")

	router := httprouter.New()

	router.POST("/api/v1/login", user_svc.Login)
	router.GET("/api/v1/login", user_svc.Logout)

	router.GET("/api/v1/posts", post_svc.ShowPosts)
	router.GET("/api/v1/posts/:id", post_svc.ShowPostById)
	router.POST("/api/v1/posts", post_svc.CreatePost)
	router.PUT("/api/v1/posts/:id", post_svc.EditPost)
	router.DELETE("/api/v1/posts/:id", post_svc.DeletePost)

	http_err := http.ListenAndServe(":9090", router)

	if http_err != nil{
		log.Fatal("Error")
	}
}
