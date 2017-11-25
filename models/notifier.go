package models

type NotifierRequest struct {
	Type    string `json:"type"`
	Target  string `json:"target"`
	Message string `json:"message"`
}

type NotifierResponse struct {
	Message string `json:"message"`
}
