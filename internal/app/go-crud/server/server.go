package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drone/routes"
	"github.com/tk-yuet/go-crud/internal/app/go-crud/config"
	"github.com/tk-yuet/go-crud/internal/app/go-crud/controllers"
	"github.com/tk-yuet/go-crud/internal/app/go-crud/database"
)

type Server struct {
	Mux *routes.RouteMux
	Env *controllers.ControllerEnv
}

func NewServer() *Server {
	env := controllers.NewControllerEnv()
	return &Server{
		Mux: nil,
		Env: env,
	}
}

func (s *Server) ConnectToDB() {
	dbConfig := config.NewMySqlConfig("root", "root", "db")
	dbClient := database.NewMySqlClient(dbConfig)
	s.Env.DbClient = dbClient

	s.Env.DbClient.Connect()
	s.Env.DbClient.UseTestDbAndTable()
}

func (s *Server) DisconnectToDB() {
	s.Env.DbClient.Disconnect()
}

func Start() {
	port := 8090
	ser := NewServer()
	ser.ConnectToDB()
	defer ser.DisconnectToDB()

	ser.constructRoutes()

	http.Handle("/", ser.Mux)
	log.Printf("Listening On http://localhost:%d ...\n", port)
	http.ListenAndServe(":8090", nil)
}

func (ser *Server) constructRoutes() {
	mux := routes.New()
	env := ser.Env

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<div> Index </div>")
	})

	mux.Get("/signup", logging(env.SignUp))
	mux.Get("/signin", logging(env.SignIn))

	mux.Post("/tasks", logging(env.AuthMiddleware(env.CreateTask)))
	mux.Get("/tasks/:id", logging(env.AuthMiddleware(env.ShowTask)))
	mux.Put("/tasks/:id", logging(env.AuthMiddleware(env.UpdateTask)))
	mux.Del("/tasks/:id", logging(env.AuthMiddleware(env.DeleteTask)))

	ser.Mux = mux
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
