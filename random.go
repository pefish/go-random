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

func (randomInstance *RandomClass) GetUniqueIdBytes() [16]byte {
	return uuid.NewV4()
}

func (randomInstance *RandomClass) GetUniqueIdString() string {
	return fmt.Sprintf(`%s`, uuid.NewV4())
}

func (randomInstance *RandomClass) RandomFromStringSlice(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}

func (randomInstance *RandomClass) MustRandomInt(start int, end int) int {
	re, err := randomInstance.RandomInt(start, end)
	if err != nil {
		panic(err)
	}
	return re
}

func (randomInstance *RandomClass) RandomInt(start int, end int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	if end < start {
		return 0, errors.New(`end must gte start`)
	}
	return rand.Intn(end-start) + start, nil
}

func (randomInstance *RandomClass) MustRandomInt64(start int64, end int64) int64 {
	re, err := randomInstance.RandomInt64(start, end)
	if err != nil {
		panic(err)
	}
	return re
}

func (randomInstance *RandomClass) RandomInt64(start int64, end int64) (int64, error) {
	rand.Seed(time.Now().UnixNano())
	if end < start {
		return 0, errors.New(`end must gte start`)
	}
	return rand.Int63n(end-start) + start, nil
}

func (randomInstance *RandomClass) MustRandomDuration(start int64, end int64) time.Duration {
	re, err := randomInstance.RandomDuration(start, end)
	if err != nil {
		panic(err)
	}
	return re
}

func (randomInstance *RandomClass) RandomDuration(start int64, end int64) (time.Duration, error) {
	rand.Seed(time.Now().UnixNano())
	if end < start {
		return 0, errors.New(`end must gte start`)
	}
	return time.Duration(rand.Int63n(end-start) + start), nil
}

func (randomInstance *RandomClass) GetRandomString(count int32) string {
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

func (randomInstance *RandomClass) GetRandomNumberStr(count int32) string {
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
