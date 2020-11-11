package convert

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TestTimeToFloat64 - тест конвертации время - f64
func TestTimeToFloat64(t *testing.T) {
	testLoc, _ := time.LoadLocation("Local")
	testTimeSlice := []time.Time{
		time.Date(1997, 2, 20, 18, 40, 21, 1000000, testLoc),
		time.Date(2000, 1, 1, 0, 0, 0, 0, testLoc),
		time.Date(2020, 12, 31, 23, 59, 59, 999, testLoc),
	}

	verifySlice := []float64{
		970220.184021001,
		101.0,
		201231.235959999,
	}

	for i, val := range testTimeSlice {
		fmt.Printf("Test %d", i+1)
		f := TimeToFloat64(val)
		//избегаю ошибки округления
		if f-verifySlice[i] >= 0.000000001 {
			t.Errorf("Not equal:\nexpected: %f\nactual: %f", verifySlice[i], f)
		}
	}
}

//TestFloat64ToTimeSuccess - тест успешной конвертации f64 - время
func TestFloat64ToTimeSuccess(t *testing.T) {
	testFloatSlice := []float64{
		970220.184021001,
		101.0,
		201231.235959999,
	}

	testLoc, _ := time.LoadLocation("Local")
	verifySlice := []time.Time{
		time.Date(1997, 2, 20, 18, 40, 21, 1000000, testLoc),
		time.Date(2000, 1, 1, 0, 0, 0, 0, testLoc),
		time.Date(2020, 12, 31, 23, 59, 59, 999000000, testLoc),
	}
	for i, val := range testFloatSlice {
		fmt.Printf("Test %d: value %f\n", i+1, val)
		ti, err := Float64ToTime(val)
		assert.Nil(t, err)
		//избегаю ошибки округления
		if ti.Sub(verifySlice[i]) >= 1000000 {
			tiStr := ti.Format(time.RFC3339Nano)
			verifyStr := verifySlice[i].Format(time.RFC3339Nano)
			t.Errorf("Not equal:\nexpected: %s\nactual: %s", verifyStr, tiStr)
		}
	}
}

//TestFloat64ToTimeFail - тест неудачной конвертации f64 - время
func TestFloat64ToTimeFail(t *testing.T) {
	testFloatSlice := []float64{
		-970220.184021001, //отрицательное
		1970220.184021001, //переполнение лет
		971320.184021001,  //переполнение месяцев
		971240.184021001,  //переполнение дней
		971131.184021001,  //31 день в ноябре
		530229.0,          //29 дней в феврале невисокосного
		0.184021001,       //нулевая целая
		970220.24,         //переполнение часов
		970220.0060,       //переполнение минут
		970220.0000601,    //переполнение секунд
	}

	for i, val := range testFloatSlice {
		fmt.Printf("Test %d\n", i+1)
		_, err := Float64ToTime(val)
		assert.Error(t, err)
	}
}

//TestFloat64ToDurationSuccess - тест успешной конвертации f64 - time.Duration
func TestFloat64ToDurationSuccess(t *testing.T) {
	testFloatSlice := []float64{
		1.0,
		0.23,
		0.0001,
		0.000054500,
		0.010101,
	}

	verifySlice := []time.Duration{
		24 * time.Hour,
		23 * time.Hour,
		time.Minute,
		54*time.Second + 500*time.Millisecond,
		time.Hour + time.Minute + time.Second,
	}

	for i, val := range testFloatSlice {
		fmt.Printf("Test %d\n", i+1)

		d, err := Float64ToDuration(val)
		assert.Nil(t, err)
		assert.Equal(t, verifySlice[i], d)
	}
}

//TestFloat64ToDurationFail - тест неудачной конвертации f64 - time.Duration
func TestFloat64ToDurationFail(t *testing.T) {
	testFloatSlice := []float64{
		100.0,
		0.24,
		0.0060,
		0.000060500,
	}
	for i, val := range testFloatSlice {
		fmt.Printf("Test %d\n", i+1)

		_, err := Float64ToDuration(val)
		assert.Error(t, err)
	}
}
