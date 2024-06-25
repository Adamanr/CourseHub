package main

import (
	"CourseHub/internal/router/api"
	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"net"
	"net/http"
)

func main() {

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil
	r := chi.NewRouter()

	curseHub := api.NewCourseHub()
	api.HandlerFromMux(curseHub, r)

	h := middleware.OapiRequestValidator(swagger)(r)

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", "8091"),
	}

	log.Fatal(s.ListenAndServe())
}
