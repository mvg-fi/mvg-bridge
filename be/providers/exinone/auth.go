package exinone

import (
	"crypto/ed25519"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mvg-fi/common/uuid"
)

const (
	UserMePath = "/users/me"
	RefMePath  = "/referral_statistics/me"
)

func SignExinToken(uid, sid, privateKey, method, uri, body string) (string, error) {
	expire := time.Now().UTC().Add(time.Hour * 24 * 30 * 3)
	sum := sha256.Sum256([]byte(method + uri + body))

	claims := jwt.MapClaims{
		"uid": uid,
		"sid": sid,
		"iat": time.Now().UTC().Unix(),
		"exp": expire.Unix(),
		"jti": uuid.NewV4(),
		"sig": hex.EncodeToString(sum[:]),
		"scp": "FULL",
		"aud": CLIENTID,
	}
	priv, err := base64.RawURLEncoding.DecodeString(privateKey)
	if err != nil {
		block, _ := pem.Decode([]byte(privateKey))
		if block == nil {
			return "", fmt.Errorf("bad RSA private pem format %s", privateKey)
		}
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
		return token.SignedString(key)
	}
	// more validate the private key
	if len(priv) != 64 {
		return "", fmt.Errorf("bad ed25519 private key %s", priv)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(ed25519.PrivateKey(priv))
}
