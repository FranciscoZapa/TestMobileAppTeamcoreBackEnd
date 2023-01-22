package models

type Response struct {
	Title      string     `json:"titulo"`
	Day        string     `json:"dia"`
	Info       []Question `json:"info"`
	ApiVersion int        `json:"api_version"`
}
