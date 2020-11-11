package model

import (
	"fmt"
	"time"
)

//Itime - интерфейс для получения текущего времени сервера
type Itime interface {
	Now() time.Time
	Correct(t time.Time)
}

//MyTime - структура, реализующая ITime
type MyTime struct {
	timeDelta time.Duration
}

//Now - получение времени сервера
func (myTime *MyTime) Now() time.Time {
	fmt.Println("now")
	fmt.Println(myTime.timeDelta)
	fmt.Println(time.Now().Add(myTime.timeDelta).Format(time.RFC1123))
	return time.Now().Add(myTime.timeDelta)
}

//Correct - корректировка времени сервера
func (myTime *MyTime) Correct(t time.Time) {
	fmt.Println("correct")
	fmt.Println(myTime.timeDelta)
	fmt.Println("t:+ " + t.Format(time.RFC1123))
	fmt.Println("now: " + time.Now().Format(time.RFC1123))
	myTime.timeDelta = t.Sub(time.Now())
}
