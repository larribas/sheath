package commands

import (
    "log"
    "net/http"

    "github.com/medbook/sheathe/domain"
)

func RetrieveLink(link_id string, r *http.Request) (originalLink string, err error) {
    link, err := linkRepository.Find(link_id)
    if err != nil {
        // TODO Control only some causes
        log.Printf("The following error occured while retrieving a link: %s\n", err.Error())
        return
    }

    notifier.Notify(domain.NewDomainEvent("LinkRetrieved", link, r))
    originalLink = link.Original.String()
    return
}