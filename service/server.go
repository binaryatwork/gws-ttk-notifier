package service

import (
	"github.com/carnellj/notifier/controllers"
	"fmt"
	"github.com/carnellj/notifier/api/twilio"
	"github.com/carnellj/notifier/di"
	"github.com/carnellj/notifier/service/notifierservice"
	"github.com/carnellj/notifier/utils"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"os"
)

func loadConfig() *utils.Config {
	config := utils.Config{}
	config.TwilioAuth = os.Getenv("TWILIO_AUTH")
	config.TwilioSid = os.Getenv( "TWILIO_SID")
	config.TwilioNumber = os.Getenv("TWILIO_NUMBER")

	return &config
}

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{IndentJSON: true})

	n := negroni.Classic()
	mx := mux.NewRouter()

	var g inject.Graph
	var c = loadConfig()
	var s di.ServerServices
	var a notifierservice.NotifierService
	var t twilio.TwilioTextNotify
	err0 := g.Provide(
		&inject.Object{Value: c},
		&inject.Object{Value: &s},
		&inject.Object{Value: &a},
		&inject.Object{Value: &t},
	)

	if err0 != nil {
		log.Fatalf("Unable to setup provider.  Existing service")
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	initRoutes(mx, formatter, &s)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, services *di.ServerServices) {
	mx.HandleFunc("/health/check", controllers.HealthCheckController(formatter)).Methods("GET")
	mx.HandleFunc("/notify", controllers.NotifyController(formatter, services)).Methods("POST")
}
