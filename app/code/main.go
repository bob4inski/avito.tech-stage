package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"crypto/tls"
	// "context"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

// Service is an interface that defines the required functions for handlers.
type Service interface {
	Get(key string) (string, error)
	Del(key string) (string, error)
	Set(key string, value string) (string, error)
}

// ServiceImpl is a struct that implements the Service interface.
type ServiceImpl struct {
	redisClient *redis.Client
}

// Get retrieves the value for the given key from the Redis database.
func (s *ServiceImpl) Get(key string) (string, error) {
	value, err := s.redisClient.Get(key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func (s *ServiceImpl) Del(key string) (string, error) {

	_, err := s.redisClient.Del(key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		return "", err
	}

	return "ok", nil
}

func (s *ServiceImpl) Set(key string, value string) (string, error) {

	set_result, err := s.redisClient.Set(key, value, 0).Result()

	if err != nil {
		return "", err
	}
	return set_result, nil
}

// Server is a struct that represents the HTTP server.
type Server struct {
	service Service
}

// NewServer creates a new instance of the Server struct.
func NewServer(service Service) *Server {
	return &Server{
		service: service,
	}
}

// GetHandler handles the GET request for retrieving a value from the Redis database.
func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Отсутствует аргумент 'key'", http.StatusBadRequest)
		return
	}

	value, err := s.service.Get(key)
	if err != nil {
		if err.Error() == "Key not found" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Page Not Found")
		} else {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s : %s", key, value)
}

func (s *Server) DelHandler(w http.ResponseWriter, r *http.Request) {
	var del_payload map[string]string

	err := json.NewDecoder(r.Body).Decode(&del_payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(del_payload) == 1 {
		for key, _ := range del_payload {
			_, err := s.service.Del(key)
			if err == redis.Nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "key does not exist")
				return
			} else if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Wrong request")
		// fmt.Fprint(w, "цеа")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "DELETE request processed")

}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {

	var payload map[string]string

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(payload) == 1 {
		for key, value := range payload {

			set_result, err := s.service.Set(key, value)

			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error while set", http.StatusInternalServerError)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s : %s", "Запись в базу данных прошла успешно", set_result)
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Wrong request")
	}
}

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

	service := &ServiceImpl{
		redisClient: rdb,
	}

	server := NewServer(service)

	r := chi.NewRouter()
	r.Get("/get", server.GetHandler)
	r.Delete("/del", server.DelHandler)
	r.Post("/set", server.SetHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
