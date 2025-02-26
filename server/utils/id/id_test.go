package id

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	id := UuidSimple()
	fmt.Println(id)
}
