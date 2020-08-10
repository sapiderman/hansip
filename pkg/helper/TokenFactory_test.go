package helper

import (
	"testing"
	"time"
)

var (
	signKey    = "thisisatestsignkey"
	signMethod = "HS256"
	issuer     = "anIssuer"
	subject    = "aSubject"
	audience   = []string{"aud1", "aud2"}
	issuedAt   = time.Date(2010, 1, 1, 1, 0, 0, 0, time.Local)
	notBefore  = time.Date(2019, 1, 1, 1, 0, 0, 0, time.Local)
	expiry     = time.Date(2030, 1, 1, 1, 0, 0, 0, time.Local)
	additional = map[string]interface{}{
		"type": "access",
	}
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiYXVkMSIsImF1ZDIiXSwiZXhwIjoxODkzNDM0NDAwLCJpYXQiOjEyNjIyODI0MDAsImlzcyI6ImFuSXNzdWVyIiwibmJmIjoxNTQ2Mjc5MjAwLCJzdWIiOiJhU3ViamVjdCIsInR5cGUiOiJhY2Nlc3MifQ.ve00z6ZL4o26pnfkPpG5xjrT4dyUCq0x4Fn4Fjl-Cco"
)

func TestCreateJWTStringToken(t *testing.T) {
	tok, err := CreateJWTStringToken(signKey, signMethod, issuer, subject, audience, issuedAt, notBefore, expiry, additional)
	if err != nil {
		t.Errorf("got %s", err)
		t.Fail()
	}
	if tok != token {
		t.Errorf("token not match\n%s\ngot\n%s", token, tok)
		t.Fail()
	}
}

func TestReadJWTStringToken(t *testing.T) {
	iss, sub, aud, iat, nbf, exp, add, err := ReadJWTStringToken(true, signKey, signMethod, token)
	if err != nil {
		t.Errorf("got %s", err)
		t.Fail()
	}
	if iss != issuer {
		t.Errorf("expect issuer %s but %s", issuer, iss)
	}
	if sub != subject {
		t.Errorf("expect subject %s but %s", subject, sub)
	}
	if aud[0] != audience[0] {
		t.Errorf("expect audience[0] %s but %s", audience[0], aud[0])
	}
	if aud[1] != audience[1] {
		t.Errorf("expect audience[1] %s but %s", audience[1], aud[1])
	}
	if nbf != notBefore {
		t.Errorf("expect notBefore %s but %s", notBefore, nbf)
	}
	if iat != issuedAt {
		t.Errorf("expect issuedAt %s but %s", issuedAt, iat)
	}
	if exp != expiry {
		t.Errorf("expect expiry %s but %s", expiry, exp)
	}
	if additional["type"] != add["type"] {
		t.Errorf("expect type %s but %s", additional["type"], add["type"])
	}
}