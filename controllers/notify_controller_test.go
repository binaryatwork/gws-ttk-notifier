package controllers

import (
	"bytes"
	"fmt"
	"github.com/carnellj/notifier/di"
	"github.com/carnellj/notifier/service/notifierservice"
	"github.com/facebookgo/inject"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"errors"
	"github.com/carnellj/notifier/api"
)

type MockGoodTwilioTextNotify struct{}
type MockErrorTwilioTextNotify struct{}

func (t *MockGoodTwilioTextNotify) Notify(Target string, Message string) (StatusCode int, Error error) {
	return http.StatusOK, nil
}

func (t *MockErrorTwilioTextNotify) Notify(Target string, Message string) (StatusCode int, Error error) {
	return http.StatusBadRequest, errors.New("Bad request returned from Twilio")
}

func mockServices(notifier api.NotifierApi) *di.ServerServices {
	var g inject.Graph
	var s di.ServerServices
	var a notifierservice.NotifierService

	err0 := g.Provide(
		&inject.Object{Value: &s},
		&inject.Object{Value: &a},
		&inject.Object{Value: notifier},
	)

	if err0 != nil {
		log.Fatalf("Unable to setup Provider")
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return &s
}

func buildJSONRequest() []byte {
	jsonBlob := []byte(`{
        "type":"twilio",
        "target":"9202651560",
        "message":"This is a message"
    }`)
	return jsonBlob
}

func TestGoodNotify(t *testing.T) {
	formatter := render.New(render.Options{IndentJSON: true})
	mux := http.NewServeMux()
	services := mockServices(&MockGoodTwilioTextNotify{})

	controller := NotifyController(formatter, services)
	mux.HandleFunc("/notify", controller)

	writer := httptest.NewRecorder()
	jsonBlob := buildJSONRequest()

	request, err := http.NewRequest("POST", "/notify", bytes.NewBuffer(jsonBlob))

	if err != nil {
		t.Error("Unable to parse twilio JSON Blob")
	}

	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code > 299 {
		t.Errorf("Error returned by notify call.  Expected 200, received %d\n ", writer.Code)
	}
}

func TestErrorNotify(t *testing.T) {
	formatter := render.New(render.Options{IndentJSON: true})
	mux := http.NewServeMux()
	services := mockServices(&MockErrorTwilioTextNotify{})

	controller := NotifyController(formatter, services)
	mux.HandleFunc("/notify", controller)

	writer := httptest.NewRecorder()
	jsonBlob := buildJSONRequest()

	request, err := http.NewRequest("POST", "/notify", bytes.NewBuffer(jsonBlob))

	if err != nil {
		t.Error("Unable to parse twilio JSON Blob")
	}

	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusBadRequest {
		t.Errorf("Error returned by notify call.  Expected 400, received %d\n ", writer.Code)
	}
}
