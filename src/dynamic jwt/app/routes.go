package main

//import (
//	"github.com/go-chi/chi"
//	"github.com/monirz/upskill/middleware"
//	handler "github.com/upskill/upskill/handlers"
//)
//
//func (s *Server) SetupRoutes() {
//
//	userHandler := handler.NewUserHandler(s.DB)
//	workShopHandler := handler.NewWorkShopHandler(s.DB)
//	nwsHandler := handler.NewNewsLetterHandler(s.DB)
//	s.router.Post("/api/login", userHandler.Login)
//	s.router.Post("/api/register", userHandler.CreateUser)
//	s.router.Post("/api/newsletter", nwsHandler.SubscribeNewsLetter)
//	s.router.Post("/api/validate", userHandler.GetUserByEmail)
//
//	s.router.Get("/api/workshops", workShopHandler.GetWorkShops)
//	s.router.Get("/api/workshops/{id}", workShopHandler.GetWorkShopByID)
//
//	s.router.Group(func(r chi.Router) {
//		r.Use(middleware.JWTAuth)
//		r.Get("/api/users/{id}", userHandler.GetUserByID)
//		r.Post("/api/workshops", workShopHandler.Create)
//		r.Delete("/api/workshops/{id}", workShopHandler.CancelWorkShop)
//		r.Patch("/api/workshops/{id}", workShopHandler.UpdateWorkShop)
//		s.router.Get("/api/newsletter", nwsHandler.GetNewsLetterEmails)
//
//	})
//}