package models

type Review struct {
	ClientID int    `json:"client_id"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment,omitempty"`
}
