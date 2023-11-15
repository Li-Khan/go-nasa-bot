package entity

import "strings"

type Apod struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func (a *Apod) Normalize() *Apod {
	a.Copyright = strings.TrimSpace(a.Copyright)
	a.Date = strings.TrimSpace(a.Date)
	a.Explanation = strings.TrimSpace(a.Explanation)
	a.Hdurl = strings.TrimSpace(a.Hdurl)
	a.MediaType = strings.TrimSpace(a.MediaType)
	a.ServiceVersion = strings.TrimSpace(a.ServiceVersion)
	a.Title = strings.TrimSpace(a.Title)
	a.URL = strings.TrimSpace(a.URL)
	return a
}
