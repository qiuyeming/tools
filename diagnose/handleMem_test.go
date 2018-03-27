package tools

import (
	"testing"
	"time"
)

func TestSampling(t *testing.T) {
	for {
		Sampling(time.Now().Format("150405"))
		time.Sleep(1e9)
	}

}
