package internal

import "errors"

var (
	ErrBadRequest        = errors.New("bad request")
	ErrGetEasyJSON       = errors.New("err get easyjson")
	ErrJSONUnexpectedEnd = errors.New("unexpected end of JSON input")
)
