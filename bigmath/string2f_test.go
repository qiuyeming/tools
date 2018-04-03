package bigmath

import (
	"fmt"
	"testing"

	"github.com/testify/assert"
)

func TestScientificS2F(t *testing.T) {
	f, err := ScientificS2F("1321423e8")
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, 1.321423e+14, f)
	assert.Equal(t, 1.321423e+13, f)
}
