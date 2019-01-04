package server

import "time"

type AlertMsg struct {
	Alerts            []SingleAlert     `json:"alerts"`
	CommonAnnotations AlertAbbotations  `json:"commonAnnotations"`
	CommonLabels      AlertLabels       `json:"commonLabels"`
	ExternalURL       string            `json:"externalURL"`
	GroupKey          string            `json:"groupKey"`
	GroupLabels       map[string]string `json:"groupLabels"`
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	StrVersion        string            `json:"version"`
	Version           float64
}

type SingleAlert struct {
	Status       string           `json:"status"`
	Labels       AlertLabels      `json:"labels"`
	Annotations  AlertAbbotations `json:"annotations"`
	StartsAt     time.Time        `json:"startsAt"`
	EndsAt       time.Time        `json:"endsAt"`
	GeneratorURL string           `json:"generatorURL"`
}

type AlertLabels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Monitor   string `json:"monitor"`
	Severity  string `json:"severity"`
}

type AlertAbbotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
