package main

import (
	"log"

	"github.com/bisal-dd/go-backend-boilerplate/graph/loaders"
	resolver "github.com/bisal-dd/go-backend-boilerplate/graph/resolver"
	"github.com/bisal-dd/go-backend-boilerplate/pkg/db"
	"github.com/bisal-dd/go-backend-boilerplate/pkg/redis"
	"github.com/bisal-dd/go-backend-boilerplate/pkg/rmq"
	"github.com/bisal-dd/go-backend-boilerplate/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	 err := godotenv.Load()
	 if err != nil {
	   log.Fatal("Error loading .env file")
	 }
	 database := db.Init()
	 cacheRedis, queueRedis, err := redis.Init()
	 if err != nil {
		log.Fatal(err)
	 }
	 if err := rmq.InitEmailQueue(queueRedis); err != nil {
		log.Fatal(err)
	 }
	 dependencyResolver := resolver.InitializeResolver(cacheRedis, database)
	 
	log.Printf("connect to http://localhost:%d/graphql for GraphQL playground", 9000)
	r := gin.Default()
	r.GET("/graphql", routes.PlaygroundHandler())
	r.Use(GinContextToContextMiddleware())
	r.Use(loaders.LoaderMiddleware(database))
	r.POST("/query", routes.GraphqlHandler(dependencyResolver))
	r.GET("/generate-upload-url", routes.HandlePresignedURL)
	r.Run()
	
}