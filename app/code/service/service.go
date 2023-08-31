package service

import (
	"github.com/go-redis/redis"
)

// Service is an interface that defines the required functions for handlers.
type Service interface {
	Get(key string) (string, error)
	Del(key string) (int64, error)
	Set(key string, value string) (error)
}

// ServiceImpl is a struct that implements the Service interface.
type ServiceImpl struct {
	redisClient *redis.Client
}

func NewServiceImpl(redisClient *redis.Client) *ServiceImpl {
	return &ServiceImpl{
		redisClient: redisClient,
	}
}
// Get retrieves the value for the given key from the Redis database.
func (s *ServiceImpl) Get(key string) (string, error) {
	// value, err := s.redisClient.Get(key).Result()
	// if err != nil {
	// 	return "", err
	// } 
	// return value, nil   
	// так тоже можно, но я решил не дублировать код
	// В идеале логгирование должн быть тут, но так чуть сложнее, поэтому сделал как проще
	return s.redisClient.Get(key).Result()
}

func (s *ServiceImpl) Del(key string) ( int64, error) {

	// value, err := s.redisClient.Del(key).Result()
	
	// if err != nil{
	// 	return 0, err
	// } 

	return s.redisClient.Del(key).Result()
}

func (s *ServiceImpl) Set(key string, value string) ( error) {

	_ , err := s.redisClient.Set(key, value, 0).Result()
	return  err
}