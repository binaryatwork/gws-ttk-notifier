package notifierservice

import (
	"github.com/carnellj/notifier/api"
	"log"
)

type NotifierService struct {
	Notifier api.NotifierApi `inject:""`
}

func (a *NotifierService) Notify(Type string, Target string, Message string) (int, error) {
	if a.Notifier == nil {
		log.Fatalf("Dependency is nil")
	}

	status, err := a.Notifier.Notify(Target, Message)
	return status, err
}
