package generate

import (
	"math/rand/v2"
	"reflect"
	"time"
)

func Time(rand *rand.Rand, typ reflect.Type) reflect.Value {
	now := time.Now()
	return reflect.ValueOf(time.Date(
		now.Year()+rand.IntN(3),     //year
		time.Month(1+rand.IntN(12)), // Month
		rand.IntN(30)+1,             // day
		rand.IntN(24),               // hour
		rand.IntN(60),               // min
		rand.IntN(60),               // sec
		rand.IntN(1000),             // nano-sec
		time.UTC,
	))
}
