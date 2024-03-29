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
	"github.com/TheOguzhan/Drone-Mobile-App-Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
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

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Use(middleware.FiberContextFromMiddleware())

	app.Use(middleware.NewJWT())

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

	app.Post("/file-upload", routes.HandleFileUploads)

	app.Static("/static-files", "./static-files")

	app.Get("/qr-codes/:code", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"code": c.Params("code"),
		})
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	app.Listen(":3000")
}
