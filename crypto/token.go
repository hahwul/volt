package crypto

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// MakeRandomToken
func MakeRandomTokenSha256(key string) string {
	now := time.Now()
	nanos := now.UnixNano()
	sum := sha256.Sum256([]byte(strconv.FormatInt(nanos, 10) + key))
	data := fmt.Sprintf("%x", string(sum[:]))
	return data
}
