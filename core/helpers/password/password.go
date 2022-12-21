package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var (
	ErrInvalidHash         = errors.New("incorrect encoded hash format")
	ErrIncompatibleVersion = errors.New("incompatible argon2 version")
)

/* generate random salt of specified length */
func generateSalt(length uint32) (salt []byte, err error) {
	salt = make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return []byte{}, err
	}

	return salt, nil
}

/* generate hash bytes using cleartext and salt */
func generateHash(cleartext string, salt []byte, p *Params) (hash []byte, err error) {
	hash = argon2.IDKey(
		[]byte(cleartext),
		salt,
		p.Iterations,
		p.Memory,
		p.Parallelism,
		p.KeyLength,
	)

	return hash, nil
}

/* generate encoded password representation for storage in the database */
func Hash(cleartext string, p *Params) (string, error) {
	salt, err := generateSalt(p.SaltLength)
	if err != nil {
		return "", err
	}

	hash, err := generateHash(cleartext, salt, p)
	if err != nil {
		return "", err
	}

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.Memory,
		p.Iterations,
		p.Parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

/* extract encoded data from the store password representation */
func decodeHashComponents(encodedHash string) (p *Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return &Params{}, []byte{}, []byte{}, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return &Params{}, []byte{}, []byte{}, err
	}
	if version != argon2.Version {
		return &Params{}, []byte{}, []byte{}, ErrIncompatibleVersion
	}

	p = &Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return &Params{}, []byte{}, []byte{}, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return &Params{}, []byte{}, []byte{}, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return &Params{}, []byte{}, []byte{}, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}

/* check if encoded hash and cleartext password match */
func Verify(cleartext, encodedHash string) (bool, error) {
	p, salt, hash, err := decodeHashComponents(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey(
		[]byte(cleartext),
		salt,
		p.Iterations,
		p.Memory,
		p.Parallelism,
		p.KeyLength,
	)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}
