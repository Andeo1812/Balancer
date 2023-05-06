package internal

//go:generate easyjson  -disallow_unknown_fields specification.go

type EchoResponse struct {
	Body string `json:"body"`
}

type ErrResponse struct {
	ErrMassage string `json:"error"`
}
