package goBootcamp

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetFinalBalance(t *testing.T) {
	ans := GetFinalBalance()
	assert.Equal(t, ans, int64(900))
}
