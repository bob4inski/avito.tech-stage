package main

import (
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"log"
	"main/handlers"
	"main/service"
	"net/http"
)

func main() {
	// ctx := context.Background()

	// tlsConfig := &tls.Config{
	// 	InsecureSkipVerify: true, // Set this to true if using a self-signed certificate
	// }

	options := &redis.Options{
		Addr:     "redis:6379",
		Password: "biba",
		// TLSConfig: tlsConfig,
	}
	rdb := redis.NewClient(options)

	serviceImpl := &service.ServiceImpl{
		RedisClient: rdb,
	}

	server := handlers.NewServer(serviceImpl)

	r := chi.NewRouter()
	r.Get("/get", server.GetHandler)
	r.Delete("/del", server.DelHandler)
	r.Post("/set", server.SetHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
