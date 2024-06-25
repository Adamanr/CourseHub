package main

import (
	api2 "CourseHub/internal/handler/api"
	"context"
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"

	"log"
	"log/slog"
	"net"
	"net/http"
)

func connectDB() (*pgx.Conn, error) {
	url := "postgres://postgres:admin21@localhost:5431/courseHub"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		slog.Error("Error connecting to database", err)
		return nil, err
	}

	return conn, nil
}

func main() {

	swagger, err := api2.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil
	r := chi.NewRouter()

	conn, err := connectDB()
	if err != nil {
		panic(err)
	}

	curseHub := api2.NewCourseHub(conn)
	r.Use(middleware2.Logger)
	api2.HandlerFromMux(curseHub, r)

	h := middleware.OapiRequestValidator(swagger)(r)

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", "8091"),
	}

	log.Fatal(s.ListenAndServe())
}
