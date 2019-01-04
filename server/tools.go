package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/newrushbolt/OctoSummon/logger"
)

func sliceContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func validateMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "POST" {
		return true
	}
	http.Error(w, "Wrong method, must be POST", http.StatusMethodNotAllowed)
	return false
}

func validateContentType(w http.ResponseWriter, r *http.Request) bool {
	targetString := "application/json"
	keyData, keyExists := r.Header["Content-Type"]
	if keyExists {
		if sliceContainsString(keyData, targetString) {
			return true
		}
	}
	http.Error(w, "Wrong Content-Type, must be application/json", http.StatusUnsupportedMediaType)
	return false
}

func processAlertMsg(r *http.Request) (AlertMsg, error) {
	var err error
	var alertData AlertMsg
	maxDataSize := int64(1048576)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxDataSize))
	if err != nil {
		logger.Logger.Warnf("Cannot read request body:\n%+v\n%v", &r, err)
		return alertData, err
	}
	logger.Logger.Debugf("Unparsed data:\t%s", body)

	err = json.Unmarshal(body, &alertData)
	if err != nil {
		err = fmt.Errorf("Cannot parse JSON:\n%s\n%v", body, err)
		return alertData, err
	}
	fVersion, err := strconv.ParseFloat(alertData.StrVersion, 64)
	alertData.Version = fVersion
	if err != nil || alertData.Version < 4.0 || alertData.Version >= 5.0 {
		err = fmt.Errorf("Unsupported AlertManager version(need 4.x):\t%v", alertData.StrVersion)
	} else if len(alertData.Alerts) < 1 {
		err = fmt.Errorf("No alerts found:\n%+v", alertData)
	}
	return alertData, err
}
