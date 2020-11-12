package httpServer

import (
	"encoding/json"
	"fmt"
	"github.com/call-me-snake/date_service/internal/convert"
	"github.com/call-me-snake/date_service/internal/model"
	"log"
	"net/http"
	"time"
)

const (
	badRequestMessage    = "Некорректные входные данные"
	internalErrorMessage = "Внутренняя ошибка сервера"
)

func timeNow(serverTime model.Itime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeInfo := time64Message{}
		timeInfo.Time64 = convert.TimeToFloat64(serverTime.Now())
		resp, err := json.Marshal(timeInfo)
		if err != nil {
			log.Printf("timeNow: %v", err)
			makeErrResponce(internalErrorMessage, http.StatusInternalServerError, w)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("timeNow: %v", err)
		}
	}
}

func timeToString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toStringRequest := &time64Message{}
		err := json.NewDecoder(r.Body).Decode(toStringRequest)
		if err != nil {
			log.Printf("timeToString: %v", err)
			makeErrResponce(badRequestMessage, http.StatusBadRequest, w)
			return
		}

		ttime, err := convert.Float64ToTime(toStringRequest.Time64)
		if err != nil {
			makeErrResponce(fmt.Sprintf("Параметр запроса Time64: %f имеет некорректый формат для конвертации", toStringRequest.Time64), http.StatusBadRequest, w)
			return
		}
		timeInfo := timeStringMessage{}
		timeInfo.Time = ttime.Format(time.RFC1123)
		resp, err := json.Marshal(timeInfo)
		if err != nil {
			log.Printf("timeToString: %v", err)
			makeErrResponce(internalErrorMessage, http.StatusInternalServerError, w)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("timeToString: %v", err)
		}
	}
}

//addTime - ручка добавления времени. Параметр запроса:
//{"Time64":201111.1800,"Delta":-0.000101}
//Time64 имеет стандартные ограничения по мм, дд, ч, мин, сек
//Delta не принимает годы, месяцы, имеет ограничение дд<=99 , стандартные ограничения по ч, мин, сек
func addTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &addTimeRequest{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			log.Printf("addTime: %v", err)
			makeErrResponce(badRequestMessage, http.StatusBadRequest, w)
			return
		}

		ttime, err := convert.Float64ToTime(request.Time64)
		if err != nil {
			makeErrResponce(fmt.Sprintf("Параметр запроса Time64: %f имеет некорректый формат для конвертации", request.Time64), http.StatusBadRequest, w)
			return
		}

		delta, err := convert.Float64ToDuration(request.Delta)
		if err != nil {
			makeErrResponce(fmt.Sprintf("Параметр запроса Delta: %f имеет некорректый формат для конвертации", request.Delta), http.StatusBadRequest, w)
			return
		}
		resultTime := ttime.Add(delta)
		resultTime64 := convert.TimeToFloat64(resultTime)
		timeInfo := time64Message{Time64: resultTime64}
		resp, err := json.Marshal(timeInfo)
		if err != nil {
			log.Printf("addTime: %v", err)
			makeErrResponce(internalErrorMessage, http.StatusInternalServerError, w)
			return
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("addTime: %v", err)
		}
	}
}

func correctTime(serverTime model.Itime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &time64Message{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			log.Printf("correctTime: %v", err)
			makeErrResponce(badRequestMessage, http.StatusBadRequest, w)
			return
		}

		ttime, err := convert.Float64ToTime(request.Time64)
		if err != nil {
			makeErrResponce(fmt.Sprintf("Параметр запроса Time64: %f имеет некорректый формат для конвертации", request.Time64), http.StatusBadRequest, w)
			return
		}
		serverTime.Correct(ttime)
		successMessage := correctTimeResponce{Message: "Время сервера успешно скорректировано"}
		resp, _ := json.Marshal(successMessage)
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("addTime: %v", err)
		}
	}
}
