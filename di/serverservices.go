package di

import (
	"github.com/carnellj/notifier/service/notifierservice"
	"github.com/carnellj/notifier/utils"
)

type ServerServices struct {
	Config          *utils.Config                    `inject:""`
	NotifierService *notifierservice.NotifierService `inject:""`
}
