package domain

type RedisService interface {
	Get(key string) (string, error)
}