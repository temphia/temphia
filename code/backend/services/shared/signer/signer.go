package signer

import (
	"crypto/sha1"
	"encoding/json"
	"errors"

	"github.com/hako/branca"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"

	"golang.org/x/crypto/pbkdf2"
)

var (
	ErrWrongTenant = errors.New("wrong tenant")
)

type Signer struct {
	masterKey    []byte
	signers      map[string]*branca.Branca
	globalSigner *branca.Branca
}

func (a *Signer) signer(namespace string) *branca.Branca {
	sig, ok := a.signers[namespace]
	if !ok {
		// race condition is fine :>
		sig = a.generator(namespace)
		a.signers[namespace] = sig
		return sig
	}

	return sig
}

func (a *Signer) generator(salt string) *branca.Branca {
	return branca.NewBranca(
		string(pbkdf2.Key(a.masterKey, []byte(salt), 100, 32, sha1.New)),
	)
}

func (a *Signer) Sign(namespace string, payload any) (string, error) {
	return sign(a.signer(namespace), payload)
}

// caller should verify tenant
func (a *Signer) Parse(namespace string, token string, target any) error {
	return parse(a.signer(namespace), token, target)
}

func (a *Signer) GlobalSignRaw(payload string) (string, error) {
	return a.globalSigner.EncodeToString(string(payload))
}

func (a *Signer) GlobalParseRaw(token string) (string, error) {
	return a.globalSigner.DecodeToString(token)
}

func (a *Signer) SignRaw(namespace string, payload string) (string, error) {
	return a.signer(namespace).EncodeToString(payload)
}

func (a *Signer) ParseRaw(namespace string, token string) (string, error) {
	return a.signer(namespace).DecodeToString(token)
}

// private

func parse(signer *branca.Branca, token string, dest any) error {
	str, err := signer.DecodeToString(token)
	if err != nil {
		return err
	}

	return json.Unmarshal(kosher.Byte(str), dest)
}

func sign(signer *branca.Branca, o any) (string, error) {
	out, err := json.Marshal(o)
	if err != nil {
		return "", nil
	}

	return signer.EncodeToString(kosher.Str(out))
}
