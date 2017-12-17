package main

import (
	"errors"
	"strings"
)

type stringService struct{}

// Uppercase 能夠將字串給轉為大寫。
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// Count 能夠計算字串的長度。
func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty 當輸入字串為空時，ErrEmpty 會被回傳。
var ErrEmpty = errors.New("Empty string")
