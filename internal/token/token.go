package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2

	// since it is being used locally for backend api
	// we would use symmetric encrytpion to encrypt the token payload
	symmetricKey []byte
}

func NewPasetoMaker(symmetrickey string) (Maker, error) {
	if len(symmetrickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key seiz: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetrickey),
	}

	return maker, nil
}

// Create Token creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(user_id int64, duration time.Duration) (string, error) {
	payload, err := NewPayload(user_id, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		
		return nil, fmt.Errorf("couldn't decrypt %v", ErrInvalidToken)
	}

	err = payload.Valid()
	if err != nil {
		fmt.Print("line 51, token.go %s", payload)
		return nil, fmt.Errorf("payload not valid %v", err)
	}

	return payload, nil

}
