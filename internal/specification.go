package internal

//go:generate easyjson  -disallow_unknown_fields specification.go

//easyjson:json
type EchoResponse struct {
	Body string `json:"body"`
}

//easyjson:json
type ErrResponse struct {
	ErrMassage string `json:"error"`
}
