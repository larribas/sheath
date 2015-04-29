package commands

import (
    "log"
    "net/http"

    "github.com/medbook/sheathe/domain"
    "github.com/medbook/sheathe/storage/redis"
    "github.com/medbook/sheathe/notifiers/file_log"
)

// TODO Make it depend on a sort of configuration Golang file
var linkRepository domain.LinkRepository = redis.NewRedisLinkRepository()
var notifier domain.EventNotifier = file_log.NewFileLogEventNotifier()

func CreateLink(rawUrl string, r *http.Request) (linkStub string, err error) {
    link, err := domain.NewLink(rawUrl)
    if err != nil {
        return
    }

    err = linkRepository.Store(link)
    if err != nil {
        log.Printf("The following error occured while storing a link: %s\n", err.Error())
        return
    }

    notifier.Notify(domain.NewDomainEvent("LinkCreated", link, r))
    linkStub = link.Stub
    return
}