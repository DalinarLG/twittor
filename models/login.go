package models

//ResponseLogin returns token
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
