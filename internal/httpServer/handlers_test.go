package httpServer

import (
	"bytes"
	"encoding/json"
	mock_model "github.com/call-me-snake/date_service/internal/model/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//TestTimeNow - тест получения времени сервера
func TestTimeNow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testLoc, _ := time.LoadLocation("Local")
	testTime := time.Date(2008, 01, 01, 02, 30, 00, 0, testLoc)
	validateFloat64 := 80101.0230
	validate, _ := json.Marshal(time64Message{Time64: validateFloat64})
	mockTimer := mock_model.NewMockItime(ctrl)
	mockTimer.EXPECT().Now().Return(testTime)
	req, err := http.NewRequest("GET", "/time/now", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(timeNow(mockTimer))
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, validate, rr.Body.Bytes())
}

//TestTimeToString - тест успешной конвертации времени
func TestTimeToString(t *testing.T) {
	float64ToConvert := 80101.0230
	reqBody, _ := json.Marshal(time64Message{Time64: float64ToConvert})

	validateLoc, _ := time.LoadLocation("Local")
	validateTime := time.Date(2008, 01, 01, 02, 30, 0, 0, validateLoc)
	validateMessage := timeStringMessage{Time: validateTime.Format(time.RFC1123)}
	validate, _ := json.Marshal(validateMessage)

	req, err := http.NewRequest("GET", "/time/string", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(timeToString())
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, validate, rr.Body.Bytes())
}

//TestTimeToString - тест успешного добавления времени
func TestAddTime(t *testing.T) {
	float64ToConvert := 80101.0259
	delta64ToConvert := 0.0001
	reqBody, _ := json.Marshal(addTimeRequest{Time64: float64ToConvert, Delta: delta64ToConvert})
	validateMessage := time64Message{Time64: 80101.03}
	validate, _ := json.Marshal(validateMessage)

	req, err := http.NewRequest("GET", "/time/add", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTime())
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, validate, rr.Body.Bytes())
}

//TestCorrectTime - тест успешной смены времени сервера
func TestCorrectTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	float64ToConvert := 201111.18
	loc, _ := time.LoadLocation("Local")
	timeConverted := time.Date(2020, 11, 11, 18, 00, 00, 00, loc)
	reqBody, _ := json.Marshal(time64Message{Time64: float64ToConvert})
	validate, _ := json.Marshal(correctTimeResponce{Message: "Время сервера успешно скорректировано"})

	mockTimer := mock_model.NewMockItime(ctrl)
	mockTimer.EXPECT().Correct(timeConverted)

	req, err := http.NewRequest("POST", "/time/correct", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(correctTime(mockTimer))
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, validate, rr.Body.Bytes())

}
