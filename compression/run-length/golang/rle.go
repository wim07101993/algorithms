package run_length

import (
	"strconv"
	"strings"
)

func Encode(s string) string {
	encoded := strings.Builder{}

	for i := 0; i < len(s); i++ {
		count := 1
		char := s[i]

		for i < len(s)-1 && s[i+1] == char {
			count++
			i++
		}

		encoded.WriteString(strconv.Itoa(count))
		encoded.WriteByte(char)
		if i < len(s)-1 {
			encoded.WriteRune(',')
		}
	}

	return encoded.String()
}

func Decode(s string) (string, error) {
	decoded := strings.Builder{}

	split := strings.Split(s, ",")
	for _, pair := range split {
		count, err := strconv.Atoi(pair[0 : len(pair)-1])
		if err != nil {
			return "", err
		}
		char := pair[len(pair)-1]

		for i := 0; i < count; i++ {
			decoded.WriteByte(char)
		}
	}

	return decoded.String(), nil
}
