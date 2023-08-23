package server

import "main/server/internal/handler"

func (s *Server) setupRoutes() {
	s.Router.Get("/_healthz", handler.Healthz)

	// http://localhost:8081/user/get?id=1
	s.Router.Get("/user/{id}", handler.UserSubscriptions(s.DB))
	s.Router.Get("/subscriptions", handler.Subscriptions(s.DB))

	s.Router.Post("/subscribe", handler.Subscribe(s.DB))
}
