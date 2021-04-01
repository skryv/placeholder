package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Main_Panics(t *testing.T) {
	assert.Panics(t, func() { main() })
}
