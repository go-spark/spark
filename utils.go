package spark

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	b := make([]byte, length+2)
	r.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
