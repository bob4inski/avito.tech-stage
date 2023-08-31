package service

import (
	"github.com/go-redis/redis"
)

// Service is an interface that defines the required functions for handlers.
type Service interface {
	Get(key string) (string, error)
	Del(key string) (int64, error)
	Set(key string, value string) (string, error)
}

// ServiceImpl is a struct that implements the Service interface.
type ServiceImpl struct {
	RedisClient *redis.Client
}

// Get retrieves the value for the given key from the Redis database.
func (s *ServiceImpl) Get(key string) (string, error) {
	value, err := s.RedisClient.Get(key).Result()
	if err != nil {
		return "", err
	} 
	return value, nil
}

func (s *ServiceImpl) Del(key string) ( int64, error) {

	value, err := s.RedisClient.Del(key).Result()
	
	if err != nil{
		return 0, err
	} 

	return value, nil
}

func (s *ServiceImpl) Set(key string, value string) (string, error) {

	setResult, err := s.RedisClient.Set(key, value, 0).Result()

	if err != nil {
		return "", err
	}
	return setResult, nil
}