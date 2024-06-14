package web

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	telemetrymiddleware "github.com/go-chi/telemetry"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {

	logger := httplog.NewLogger("clean-arch", httplog.Options{
		JSON: true,
	})
	s.Router.Use(middleware.Logger)
	s.Router.Use(httplog.RequestLogger(logger, []string{"/ping"}))
	s.Router.Use(telemetrymiddleware.Collector(telemetrymiddleware.Config{
		AllowAny: true,
	}, []string{"/v1"}))
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
	fmt.Println("Starting web server on port", s.WebServerPort)
}
