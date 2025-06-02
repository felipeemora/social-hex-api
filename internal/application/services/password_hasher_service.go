package services

import "golang.org/x/crypto/bcrypt"

func Compare(hash []byte, text string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(text))
	return err == nil
}

func Hash(text string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, err
}
