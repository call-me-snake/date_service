package httpServer

import (
	"encoding/json"
	"net/http"
)

type time64Message struct {
	Time64 float64 `json:"Time64"`
}

type timeStringMessage struct {
	Time string `json:"Time"`
}

type addTimeRequest struct {
	Time64 float64 `json:"Time64"`
	Delta  float64 `json:"Delta"`
}

type errorResponce struct {
	Message string `json:"Message"`
	ErrCode int    `json:"ErrCode"`
}

type correctTimeResponce struct {
	Message string `json:"Message"`
}

func makeErrResponce(userMessage string, errCode int, w http.ResponseWriter) {
	res, _ := json.Marshal(errorResponce{Message: userMessage, ErrCode: errCode})
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(errCode)
	w.Write(res)
}
