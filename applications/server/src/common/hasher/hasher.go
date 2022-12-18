package hasher

import (
	"golang.org/x/crypto/bcrypt"
	"hash/fnv"
)

type Hasher interface {
	Hash(password string) (string, error)
	CheckHash(password, hash string) bool
	HashStringToFloat(str string, seed string) (float32, error)
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

func (impl *HasherImpl) CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (impl *HasherImpl) HashStringToFloat(str string, seed string) (float32, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(str))

	if err != nil {
		return 0, err
	}

	_, err = h.Write([]byte(seed))

	if err != nil {
		return 0, err
	}

	return float32(h.Sum32()) / float32(1<<32), nil
}
