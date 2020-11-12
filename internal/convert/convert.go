package convert

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

const defaultFStringLength = 16

//TimeToFloat64 - перевод time.Time в float64
func TimeToFloat64(ttime time.Time) (ftime float64) {
	s := ttime.Format("060102.150405.999")
	if len(s) > 13{
		s = s[:13] + s[14:]
	}
	ftime, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Print(err)
	}
	return
}

//Float64ToTime - перевод float64 в time.Time
func Float64ToTime(ftime float64) (ttime time.Time, err error) {
	if ftime <= 0.0 {
		return time.Now(), errors.New("convert.Float64ToTime: отрицательное число")
	}
	fstring := fmt.Sprintf("%.9f", ftime)
	if len(fstring) > defaultFStringLength {
		return time.Now(), fmt.Errorf("convert.Float64ToTime: переполнение: %s", fstring)
	} else if len(fstring) < defaultFStringLength {
		delta := defaultFStringLength - len(fstring)
		sl0 := make([]byte, delta, delta)
		for i := 0; i < delta; i++ {
			sl0[i] = '0'
		}
		fstring = string(sl0) + fstring
	}
	ttime, err = time.Parse("060102.150405.999", fstring[:13]+"."+fstring[13:])
	if err != nil {
		log.Println(err)
	} else {
		loc, _ := time.LoadLocation("Local")
		ttime = time.Date(ttime.Year(), ttime.Month(), ttime.Day(), ttime.Hour(), ttime.Minute(), ttime.Second(), ttime.Nanosecond(), loc)
	}
	return
}

//Float64ToDuration - перевод float64 в time.Duration
//так  как я не знаю точное количество дней в месяце и году, а в задании о смещении сказано общими фразами, я ограничил
//смещение днями
func Float64ToDuration(fduration float64) (tduration time.Duration, err error) {
	var add bool = true
	if fduration < 0 {
		fduration = -fduration
		add = false
	}
	days := int(fduration)
	if days > 99 {
		err = errors.New("Некорректный формат числа, дни")
		return
	}

	_, fractional := math.Modf(fduration)
	fstring := fmt.Sprintf("%0.9f", fractional)
	hours, _ := strconv.Atoi(fstring[2:4])
	if hours > 23 {
		err = errors.New("Некорректный формат числа, часы")
		return
	}
	minutes, _ := strconv.Atoi(fstring[4:6])
	if minutes > 59 {
		err = errors.New("Некорректный формат числа, минуты")
		return
	}
	seconds, _ := strconv.Atoi(fstring[6:8])
	if seconds > 59 {
		err = errors.New("Некорректный формат числа, секунды")
		return
	}
	mseconds, _ := strconv.Atoi(fstring[8:])

	tduration += 24*time.Hour*time.Duration(days) + time.Hour*time.Duration(hours)
	tduration += time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds) + time.Millisecond*time.Duration(mseconds)
	if !add {
		tduration *= -1
	}
	return
}
