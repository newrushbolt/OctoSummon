package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/newrushbolt/OctoSummon/config"
	"github.com/newrushbolt/OctoSummon/logger"
)

func alertHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debugf("Got request:\n%+v", r)
	alertMsgData, err := processAlertMsg(r)
	if err != nil {
		logger.Logger.Errorf("Cannot get alert data:\n%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger.Debugf("AlertsData:\t%+v", alertMsgData)
	if alertMsgData.CommonLabels.Severity == "page" && alertMsgData.Status == "firing" {
		logger.Logger.Infof("A DISASTER on host <%s>:\t%s", alertMsgData.CommonLabels.Instance, alertMsgData.CommonAnnotations.Summary)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debugf("Got request:\n%+v", r)
	if validateContentType(w, r) && validateMethod(w, r) {
		http.Error(w, "Use /alerts for alerting, or go away", http.StatusNotFound)
	}
}

func Start(config.MainConfig) {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(errorHandler)
	r.HandleFunc("/alerts", alertHandler).
		Methods("POST").
		Headers("Content-Type", "application/json")

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		localErr := fmt.Errorf("HTTP server is down:\n%v", err)
		logger.Logger.Fatal(localErr)
	}
}
