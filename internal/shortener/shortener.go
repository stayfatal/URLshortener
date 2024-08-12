package shortener

import (
	"crypto/rand"
	"encoding/hex"
)

const linkLen = 6

func GenerateShortLink(checker func(link string) (bool, error)) (string, error) {
	var ok bool = true
	var link string
	var err error
	for ok {
		link, err = generateRandomString()
		if err != nil {
			return "", err
		}

		ok, err = checker(link)
		if err != nil {
			return "", err
		}
	}
	return link, nil
}

func generateRandomString() (string, error) {
	bytes := make([]byte, linkLen/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
