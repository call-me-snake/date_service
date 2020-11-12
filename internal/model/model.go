package model

import (
	"time"
)

//Itime - интерфейс для получения текущего времени сервера
type Itime interface {
	Now() time.Time
	Correct(t time.Time)
}

//TimeData - структура, реализующая ITime
type TimeData struct {
	timeDelta time.Duration
}

//Now - получение времени сервера
func (timeData *TimeData) Now() time.Time {
	//fmt.Println("now")
	//fmt.Println(timeData.timeDelta)
	//fmt.Println(time.Now().Add(timeData.timeDelta).Format(time.RFC1123))
	return time.Now().Add(timeData.timeDelta)
}

//Correct - корректировка времени сервера
func (timeData *TimeData) Correct(t time.Time) {
	//fmt.Println("correct")
	//fmt.Println(timeData.timeDelta)
	//fmt.Println("t:+ " + t.Format(time.RFC1123))
	//fmt.Println("now: " + time.Now().Format(time.RFC1123))
	timeData.timeDelta = t.Sub(time.Now())
}
