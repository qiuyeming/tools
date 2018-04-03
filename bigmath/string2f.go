package bigmath

import (
	"fmt"
)

//ScientificS2F 科学计数法的string to float64
func ScientificS2F(digit string) (float64, error) {
	var fdes float64
	_, err := fmt.Sscanf(digit, "%f", &fdes)
	if err != nil {
		return fdes, err
	}
	return fdes, nil
}
