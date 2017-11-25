package di

import (
	"github.com/carnellj/gws-ttk-notifier/service/notifierservice"
	"github.com/carnellj/gws-ttk-notifier/utils"
)

type ServerServices struct {
	Config          *utils.Config                    `inject:""`
	NotifierService *notifierservice.NotifierService `inject:""`
}
