package main

import (
	"strings"
)

var alphabet = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
var base = len(alphabet)
var domain = "madtiny.com/"

func Encode(urlID int) string {
	url := ""
	for urlID > 0 {
		url = string(alphabet[urlID%base]) + url
		urlID /= base
	}

	return domain + url
}

func Decode(decode string) int {
	short := decode[len(domain):]
	number := 0
	for _, c := range short {
		number = number*base + strings.Index(alphabet, string(c))
	}
	return number
}
