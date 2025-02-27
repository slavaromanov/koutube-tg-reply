package koutube_conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter_ConvertVideoURL(t *testing.T) {
	// inputAndExpec
	table := map[string]string{
		" ": "",
		`ВИА "Лыбiдi" нахуярело нового

https://www.youtube.com/watch?v=SYkjVGrO5Ug`: "https://koutu.be/SYkjVGrO5Ug",
		"https://www.youtube.com/watch?v=abc123XYZ":                   "https://koutu.be/abc123XYZ",
		"https://m.youtube.com/watch?v=abc123XYZ":                     "https://koutu.be/abc123XYZ",
		"https://youtu.be/abc123XYZ":                                  "https://koutu.be/abc123XYZ",
		"https://www.youtube.com/shorts/xyz789ABC":                    "https://koutu.be/shorts/xyz789ABC",
		"https://m.youtube.com/shorts/xyz789ABC":                      "https://koutu.be/shorts/xyz789ABC",
		"https://youtube.com/shorts/xyz789ABC?feature=share":          "https://koutu.be/shorts/xyz789ABC",
		"https://www.youtube.com/embed/lmno567PQR":                    "https://koutu.be/lmno567PQR",
		"https://www.youtube.com/watch?v=abc123XYZ&list=PL1234567890": "https://koutu.be/abc123XYZ",
		"https://m.youtube.com/watch?v=abc123XYZ&list=PL1234567890":   "https://koutu.be/abc123XYZ",
		"https://www.youtube.com/watch?v=abc123XYZ&t=60s":             "https://koutu.be/abc123XYZ",
		"https://youtu.be/abc123XYZ?t=60":                             "https://koutu.be/abc123XYZ",
		"https://www.youtube.com/c/ChannelName":                       "",
		"https://m.youtube.com/c/ChannelName":                         "",
		"https://youtube.com/@ChannelName":                            "",
		"https://www.youtube.com/live/abcd1234XYZ":                    "",
	}
	c := NewConverter()
	for k, v := range table {
		_, res := c.ConvertVideoURL(k)
		assert.Equal(t, res, v)
	}
}
