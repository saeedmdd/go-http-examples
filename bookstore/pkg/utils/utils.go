package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func ParseBody(request *http.Request, X any) {
	if body, err := io.ReadAll(request.Body); err == nil {
		if error := json.Unmarshal(body, X); error != nil {
			return
		}
	}
}
func ResponseBody(writer http.ResponseWriter, data interface{}, status string, message string, code int) {
	if status == "" {
		status = "ok"
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	response := response{Code: code, Status: status, Message: message, Data: data}
	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	writer.Write(result)
}
