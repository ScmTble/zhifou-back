package jwtx

import (
	"fmt"
	"testing"
	"time"
)

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA1Nzg2ODUsImlhdCI6MTY4MDQ5MjI4NSwidWlkIjo5OTk5fQ.h0zfH3ks4TmGO6JaZxWjpuCAPocNZacUISbfgJDbwiM
func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("wad", time.Now().Unix(), 86400, "wadawdawd")
	fmt.Println(token)
	fmt.Println(err)
}

func TestParseToken(t *testing.T) {
	token, err := ParseToken("wad", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA1Nzg2ODUsImlhdCI6MTY4MDQ5MjI4NSwidWlkIjo5OTk5fQ.h0zfH3ks4TmGO6JaZxWjpuCAPocNZacUISbfgJDbwiM")
	fmt.Println(token)
	fmt.Println(err)
}
