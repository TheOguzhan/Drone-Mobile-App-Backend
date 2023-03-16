package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/ent"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/graph"
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	_ "github.com/TheOguzhan/Drone-Mobile-App-Backend/ent/runtime"
	_ "github.com/lib/pq"
)

const defaultPort = "3000"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD")))

	if err != nil {
		log.Fatalf("Failed opening connection to posgres: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &Resolver{
			client: client,
		},
		Directives: graph.DirectiveRoot{},
		Complexity: graph.ComplexityRoot{},
	}))

	app := fiber.New()

	app.Use(middleware.FiberContextFromMiddleware())
	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})

	playground_n := playground.Handler("GraphQL playground", "/query")

	app.Post("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler)(c.Context())
		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playground_n)(c.Context())
		return nil
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	app.Listen(":3000")
}
