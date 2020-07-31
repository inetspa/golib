package util

import (
	"crypto/sha1"
	"encoding/json"
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

func GenOtpCode(n int) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func HashStruct(arr interface{}) [20]byte {
	var arrBytes []byte
	jsonBytes, _ := json.Marshal(arr)
	arrBytes = append(arrBytes, jsonBytes...)
	return sha1.Sum(arrBytes)
}

func AtoI(s string, v int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return v
	}
	return i
}

func AtoF(s string, v float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return v
	}
	return f
}

func SetIfDiff(a interface{}, b interface{}) interface{} {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return a
	}
	if reflect.ValueOf(a) != reflect.ValueOf(b) {
		return b
	}
	return a
}
