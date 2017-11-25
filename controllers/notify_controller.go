package controllers

import (
	"encoding/json"
	"github.com/carnellj/gws-ttk-notifier/di"
	"github.com/carnellj/gws-ttk-notifier/models"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

func encodeHTTPResponse(w http.ResponseWriter, notifierResponse models.NotifierResponse, formatter *render.Render, status int, err error) {

	//Note:  In Go you need to set the status right away on the write to the Response.  If you write anything else to the header
	//GO will set the Header to 200 and you will not be able to change it.  This will occur even if you set the WriteHeader to 200
	//later in the GO.
	if err != nil {
		log.Printf("Error encountered while calling the notifier service.  Status code %s, Error %s", status, err)
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&notifierResponse); err != nil {
		log.Printf("Error encountered: %s \n", err.Error())
		formatter.JSON(w, http.StatusBadRequest,
			struct{ Status string }{err.Error()})
	}
}

func NotifyController(formatter *render.Render, services *di.ServerServices) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		notifierRequest := models.NotifierRequest{}
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&notifierRequest)

		if err != nil {
			log.Printf("Error encountered: %s \n", err)
			formatter.JSON(w, http.StatusBadRequest,
				struct{ Status string }{err.Error()})
		}

		status, err := services.NotifierService.Notify(notifierRequest.Type, notifierRequest.Target, notifierRequest.Message)
		notifierResponse := &models.NotifierResponse{}
		notifierResponse.Message = "Successfully Processed Response"
		encodeHTTPResponse(w, *notifierResponse, formatter, status, err)

	}
}
