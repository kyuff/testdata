package generate

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

var (
	charList  = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	charCount = len(charList)
)

func String(r *rand.Rand, typ reflect.Type, size uint16) reflect.Value {
	var b []byte
	for i := 0; i < int(size); i++ {
		b = append(b, charList[r.IntN(charCount)])
	}

	s := fmt.Sprintf("%s-%s", typ.Name(), string(b))

	return reflect.ValueOf(s)
}
