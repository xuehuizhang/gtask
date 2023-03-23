package client

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	now := time.Now()
	err := Get("http://127.0.0.1:9098/name", nil)
	if err != nil {
		fmt.Println(err)
	}
	since := time.Since(now)
	fmt.Println(since.Seconds())
}
