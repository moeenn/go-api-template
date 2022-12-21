package password

import (
	"testing"
)

var params = Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func TestPasswordValidVerify(t *testing.T) {
	cleartext := "password"
	hash, err := Hash(cleartext, &params)
	if err != nil {
		t.Errorf("password hashing failed: %v", err)
	}

	isValid, err := Verify(cleartext, hash)
	if err != nil {
		t.Errorf("password verification failed: %v", err)
	}

	if !isValid {
		t.Errorf("expected valid match, got !isValid")
	}
}

func TestPasswordVerificationFailure(t *testing.T) {
	cleartext := "password"
	hash, err := Hash(cleartext, &params)
	if err != nil {
		t.Errorf("password hashing failed: %v", err)
	}

	isValid, err := Verify("some random password", hash)
	if err != nil {
		t.Errorf("password verification failed: %v", err)
	}

	if isValid {
		t.Errorf("expected no match, got isValid")
	}
}
