package go_random

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type RandomClass struct {
}

var Random = RandomClass{}

func (this *RandomClass) GetUniqueIdBytes() [16]byte {
	return uuid.NewV4()
}

func (this *RandomClass) GetUniqueIdString() string {
	return fmt.Sprintf(`%s`, uuid.NewV4())
}

func (this *RandomClass) RandomFromStringSlice(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}

func (this *RandomClass) RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	if end < start {
		panic(errors.New(`end must gte start`))
	}
	return rand.Intn(end-start) + start
}

func (this *RandomClass) RandomInt64(start int64, end int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if end < start {
		panic(errors.New(`end must gte start`))
	}
	return rand.Int63n(end-start) + start
}

func (this *RandomClass) RandomDuration(start int64, end int64) time.Duration {
	rand.Seed(time.Now().UnixNano())
	if end < start {
		panic(errors.New(`end must gte start`))
	}
	return time.Duration(rand.Int63n(end-start) + start)
}

func (this *RandomClass) GetRandomString(count int32) string {
	dictionary := "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, count)
	l := len(dictionary)

	_, err := rand.Read(b)

	if err != nil {
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}

func (this *RandomClass) GetRandomNumberStr(count int32) string {
	dictionary := "0123456789"
	b := make([]byte, count)
	l := len(dictionary)

	_, err := rand.Read(b)

	if err != nil {
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}

func (this *RandomClass) Test() string {
	fmt.Println(rand.New(rand.NewSource(12)).Intn(2))
	return `11`
}
