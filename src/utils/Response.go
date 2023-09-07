package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ResponseWriter http.ResponseWriter
	StatusCode int
	Data any
}

func (response *Response) Send() {
	jsonData, _ := json.Marshal(response.Data)

	response.ResponseWriter.Header().Set("Content-Type", "application/json")
	response.ResponseWriter.WriteHeader(response.StatusCode)
	response.ResponseWriter.Write(jsonData)
}