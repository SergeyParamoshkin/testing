package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	assert := assert.New(t)
	testSet := []struct {
		in  int64
		out int64
	}{
		{112312, 0},
		{7, 0},
		{3, 0},
		{1, 0},
		{2, 1},
	}

	for _, tt := range testSet {
		assert.LessOrEqual(tt.out, GetRandomNumber(tt.in), "test ")
	}

}
