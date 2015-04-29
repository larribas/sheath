package file_log

import (
    "os"
    "log"

    "github.com/medbook/sheathe/domain"
)

type FileLogEventNotifier struct {
    log *log.Logger
}

const LogDir = "var/logs"

func NewFileLogEventNotifier() *FileLogEventNotifier {
    err := os.MkdirAll(LogDir, 0777)
    if err != nil {
        log.Println("Failed to create log directory due to error:", err)
    }

    file, err := os.OpenFile("var/logs/event.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Println("Failed to open log file due to error:", err)
    }

    return &FileLogEventNotifier{log: log.New(file, "EVENT: ", 0)}
}

func (n *FileLogEventNotifier) Notify(event *domain.DomainEvent) {
    n.log.Printf("%s registered at %d. IP %s for link %s\n", event.Type, event.Time, event.IP, event.Link)
}