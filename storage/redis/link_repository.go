package redis

import (
    "errors"

    "github.com/fzzy/radix/redis"
    "github.com/medbook/sheathe/domain"
)

// Redis LinkRepository implementation
const REDIS_PREFIX = "sheathe#"

type RedisLinkRepository struct {
    client *redis.Client
}

func (r RedisLinkRepository) Store(l domain.Link) (error) {
    reply := r.client.Cmd("SET", REDIS_PREFIX + l.Stub, l.Original.String(), "NX")
    if reply.Type == redis.NilReply {
        return errors.New("A link with such Hash already exists")
    }

    return nil
}

func (r RedisLinkRepository) Find(hash string) (domain.Link, error) {
    reply := r.client.Cmd("GET", REDIS_PREFIX + hash)
    if reply.Type != redis.NilReply {
        link, _ := domain.NewLink(reply.String())
        return link, nil
    } else {
        return domain.Link{}, errors.New("There is no link with such Hash")
    }
}


func NewRedisLinkRepository() *RedisLinkRepository {
    // TODO Explore the possibility of a lazy connection
    // TODO Handle errors
    cl, _ := redis.Dial("tcp", "localhost:6379")
    return &RedisLinkRepository{client: cl}
}
