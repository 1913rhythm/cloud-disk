package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGetUUID(t *testing.T) {
	v4 := uuid.NewV4().String()
	fmt.Println(v4)
}
