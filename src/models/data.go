package models

type Data struct {
	Date string     `json:"date"`
	Data []Question `json:"data"`
}
