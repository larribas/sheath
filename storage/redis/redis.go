// Package redis implements a LinkRepository that stores and retrieves links from a Redis database
package redis

import (
	"os"
	"fmt"
	"log"

	"github.com/fzzy/radix/redis"
	"github.com/larribas/sheath/application/domain"
)

var (
	host = getEnvOrDefault("SHEATH_REDIS_HOST", "localhost")
	port = getEnvOrDefault("SHEATH_REDIS_PORT", "6379")
)

// RedisPrefix is used in all calls to redis to provide a particular namespace for Sheath,
// in the event that the same Redis cluster is being shared by multiple applications
const RedisPrefix = "github.com/larribas/sheath#"

// LinkRepository connects to a Redis database via a lazily instantiated client,
// and uses it to store and retrieve links under its own namespace
type LinkRepository struct {
	client *redis.Client
}

// NewLinkRepository returns a new instance of LinkRepository,
// ensuring it complies with the corresponding domain interface
func NewLinkRepository() domain.LinkRepository {
	return &LinkRepository{}
}

// Store saves the specified Link to the Redis database
func (r LinkRepository) Store(l *domain.Link) error {
	reply := r.getClient().Cmd("SET", RedisPrefix+l.Stub, l.Original.String(), "NX")
	if reply.Type == redis.NilReply {
		return domain.ErrLinkAlreadyExists(l.Stub)
	}

	return nil
}

// Find retrieves the Link corresponding to the supplied stub from the Redis database
func (r LinkRepository) Find(stub string) (*domain.Link, error) {
	reply := r.getClient().Cmd("GET", RedisPrefix+stub)
	if reply.Type != redis.NilReply {
		link, _ := domain.NewLink(reply.String())
		return link, nil
	}

	return &domain.Link{}, domain.ErrLinkNotFound(stub)
}

func (r LinkRepository) getClient() *redis.Client {
	var err error

	if r.client == nil {
		r.client, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			log.Fatalf("Couldn't connect to Redis. Error: %s\n", err.Error())
		}
	}

	return r.client
}

func getEnvOrDefault(key string, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}

	return env
}
