package models

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

type HttpErrorMessage struct {
	Errors []HttpError `json:"errors"`
}
