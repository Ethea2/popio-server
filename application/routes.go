package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Ethea2/popio-server/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/user", loadUserRoutes)
	router.Route("/proof", loadProofRoutes)
	router.Route("/bill", loadBillsRoutes)
	router.Route("/group", loadGroupRoutes)

	return router
}

func loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{}

	router.Get("/{id}", userHandler.GetUser)
	router.Post("/login", userHandler.Login)
	router.Post("/logout", userHandler.Logout)
	router.Post("/signup", userHandler.SignUp)
}

func loadProofRoutes(router chi.Router) {
	proofHandler := &handler.Proof{}

	router.Get("/group/{id}", proofHandler.GetProofByGroup)
	router.Get("/user/{id}", proofHandler.GetUserProof)
	router.Post("/group/create", proofHandler.CreateForGroup)
	router.Post("/user/create", proofHandler.CreateForGroup)
	router.Delete("/{id}", proofHandler.Delete)
	router.Patch("/{id}", proofHandler.Edit)
}

func loadBillsRoutes(router chi.Router) {
	billsHandler := &handler.Bills{}

	router.Get("/group/{id}", billsHandler.GetBillsByGroup)
	router.Get("/user/{id}", billsHandler.GetBillsByUser)
	router.Post("/group", billsHandler.AddBillForGroup)
	router.Post("/user", billsHandler.AddBillForUser)
	router.Delete("/user/{id}", billsHandler.DeleteBillForUser)
	router.Delete("/group/{id}", billsHandler.DeleteBillForGroup)
}

func loadGroupRoutes(router chi.Router) {
	groupHandler := &handler.Group{}

	router.Get("/{id}", groupHandler.GetGroupsForUser)
	router.Get("/", groupHandler.GetGroup)
	router.Post("/", groupHandler.CreateGroup)
	router.Delete("/{id}", groupHandler.DeleteGroup)
	router.Patch("/{id}", groupHandler.EditGroup)
}
