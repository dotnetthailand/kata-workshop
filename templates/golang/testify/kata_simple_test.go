package kata

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// TODO: Add in GitHub Action

func TestCalculate(t *testing.T) {
    assert := assert.New(t)

    var tests = []struct {
        input    int
        expected int
    }{
        {2, 5},
        {-1, 1},
        {0, 2},
        {-5, -3},
    }

    for _, test := range tests {
        assert.Equal(Calculate(test.input), test.expected)
    }
}

// Calculate returns x + 2.
func Calculate(x int) (result int) {
    result = x + 2
    return result
}