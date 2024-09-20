package applications

import (
	"context"
	"fmt"
	handlers "keploy-test/handler"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	// This  is the unit testing
	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handlers.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.DeleteById)
}

func (a *App) Start(ctx context.Context) error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("router.go Handler ::: PORT environment variable is not set")
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}

	ch := make(chan error, 1)

	fmt.Println("Server running at port", port)

	go func() {
		e := server.ListenAndServe()
		if e != nil {
			ch <- fmt.Errorf("failed to server: %w", e)
		}
		close(ch)
	}()

	ctx.Done()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

	return nil
}
