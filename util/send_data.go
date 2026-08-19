package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, status int) {

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter ,statusCode int , msg string){
w.WriteHeader(statusCode)
json.NewEncoder(w).Encode(msg)
}