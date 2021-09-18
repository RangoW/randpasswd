package randpasswd

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

const (
	randBytesLen = 1 << 5
	adecimal     = 1 << 6 // 64decimal
)

// password scope: 0-9,a-z,A-Z,#,!
func generate_symbols() []byte {
	passwd_symbols := make([]byte, 0, adecimal)
	// fill
	for i := '0'; i <= '9'; i++ {
		passwd_symbols = append(passwd_symbols, byte(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		passwd_symbols = append(passwd_symbols, byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		passwd_symbols = append(passwd_symbols, byte(i))
	}
	passwd_symbols = append(passwd_symbols, []byte{'#', '!'}...)

	return passwd_symbols
}

func Generate(length uint) (string, error) {
	if length > randBytesLen*8/6 {
		return "", fmt.Errorf("the length must be less than 42")
	}

	randBytes := make([]byte, randBytesLen) // 32Bytes => 256 BitString
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	randBitString := ""
	for i := range randBytes {
		randBitString += fmt.Sprintf("%08b", randBytes[i])
	}

	var result []byte
	passwdSymbols := generate_symbols()
	for i := uint(0); i < length; i++ {
		index, _ := strconv.ParseInt(randBitString[i*6:(i+1)*6], 2, 64)
		result = append(result, passwdSymbols[index])
	}

	return string(result), nil
}
