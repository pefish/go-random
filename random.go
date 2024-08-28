package go_random

import (
	"errors"
	"math/rand"
	"time"
)

func RandomFromSlice[T any](slice []T) (index int, value T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index = r.Intn(len(slice))
	value = slice[index]
	return
}

func RandomCountFromSlice[T any](slice []T, count int) []T {
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	values := make([]T, 0)
	for i := 0; i < count; i++ {
		index, value := RandomFromSlice(newSlice)
		values = append(values, value)
		newSlice = append(newSlice[:index], newSlice[index+1:]...)
	}
	return values
}

func MustRandomInt(start int, end int) int {
	re, err := RandomInt(start, end)
	if err != nil {
		panic(err)
	}
	return re
}

func RandomInt(start int, end int) (int, error) {
	randInstance := rand.New(rand.NewSource(time.Now().UnixNano()))
	if end <= start {
		return 0, errors.New(`end must gt start`)
	}
	return randInstance.Intn(end-start) + start, nil
}

func MustRandomInt64(start int64, end int64) int64 {
	re, err := RandomInt64(start, end)
	if err != nil {
		panic(err)
	}
	return re
}

func RandomInt64(start int64, end int64) (int64, error) {
	randInstance := rand.New(rand.NewSource(time.Now().UnixNano()))
	if end <= start {
		return 0, errors.New(`end must gt start`)
	}
	return randInstance.Int63n(end-start) + start, nil
}

func MustRandomFloat64(start float64, end float64) float64 {
	result, err := RandomFloat64(start, end)
	if err != nil {
		panic(err)
	}
	return result
}

func RandomFloat64(start float64, end float64) (float64, error) {
	randInstance := rand.New(rand.NewSource(time.Now().UnixNano()))
	if end <= start {
		return 0, errors.New(`end must gt start`)
	}
	return randInstance.Float64()*(end-start) + start, nil
}

func MustRandomString(count int32) string {
	result, err := RandomString(count)
	if err != nil {
		panic(err)
	}
	return result
}

func RandomString(count int32) (string, error) {
	return RandomStringFromDic("_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", count)
}

func MustRandomStringFromDic(dictionary string, count int32) string {
	r, err := RandomStringFromDic(dictionary, count)
	if err != nil {
		panic(err)
	}
	return r
}

func RandomStringFromDic(dictionary string, count int32) (string, error) {
	b := make([]byte, count)
	l := len(dictionary)

	_, err := rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)

	if err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = dictionary[v%byte(l)]
	}

	return string(b), nil
}

func MustRandomBytes(count int32) []byte {
	result, err := RandomBytes(count)
	if err != nil {
		panic(err)
	}
	return result
}

func RandomBytes(count int32) ([]byte, error) {
	b := make([]byte, count)

	_, err := rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomNumberStr(count int32) (string, error) {
	return RandomStringFromDic("0123456789", count)
}

func MustRandomNumberStr(count int32) string {
	result, err := RandomNumberStr(count)
	if err != nil {
		panic(err)
	}
	return result
}
