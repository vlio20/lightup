package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(password string) (string, error)
	CheckHash(password, hash string) bool
}

type HasherImpl struct {
}

func New() Hasher {
	return &HasherImpl{}
}

func (impl *HasherImpl) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (impl *HasherImpl) CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
