package redis

import (
    "log"
    "errors"

    "github.com/fzzy/radix/redis"
    "github.com/medbook/sheathe/domain"
)

const REDIS_PREFIX = "sheathe#"

type RedisLinkRepository struct {
    client *redis.Client
}

func (r RedisLinkRepository) Client() (*redis.Client) {
    var err error

    if r.client == nil {
		// TODO Take this to the configuration file
        r.client, err = redis.Dial("tcp", "redis:6379")
        if err != nil {
            log.Fatalf("Couldn't connect to Redis. Error: %s\n", err.Error())
        }
    }

    return r.client
}

func (r RedisLinkRepository) Store(l domain.Link) (error) {
    reply := r.Client().Cmd("SET", REDIS_PREFIX + l.Stub, l.Original.String(), "NX")
    if reply.Type == redis.NilReply {
        return errors.New("A link with such Hash already exists")
    }

    return nil
}

func (r RedisLinkRepository) Find(hash string) (domain.Link, error) {
    reply := r.Client().Cmd("GET", REDIS_PREFIX + hash)
    if reply.Type != redis.NilReply {
        link, _ := domain.NewLink(reply.String())
        return link, nil
    } else {
        return domain.Link{}, errors.New("There is no link with such Hash")
    }
}


func NewRedisLinkRepository() *RedisLinkRepository {
    return &RedisLinkRepository{}
}
