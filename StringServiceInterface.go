package main

// IStringService 提供了處理字串的相關功能。
type IStringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
