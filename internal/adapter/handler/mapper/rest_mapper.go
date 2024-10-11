package mapper

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		return err
	}

	return nil
}

type ProblemDetail struct {
	Type       string      `json:"type" `
	Status     int         `json:"status" `
	Detail     string      `json:"detail" `
	Instance   string      `json:"instance,omitempty"`
	Extensions interface{} `json:"extensions,omitempty"`
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := ProblemDetail{
		Type:   "error",
		Status: statusCode,
		Detail: err.Error(),
	}

	return WriteJSON(w, statusCode, payload)
}
