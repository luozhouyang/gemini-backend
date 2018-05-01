package main

import (
	"testing"
	_ "backend/db"
	"github.com/stretchr/testify/assert"
	"fmt"
)

// TODO(luozhouyang) test failed, fix that
func TestQueryByUuid(t *testing.T) {
	a, err := QueryByUuid("123456")
	if err != nil {
		assert.Nil(t, a)
		fmt.Println(err.Error())
		return
	}
	assert.Equal(t, 0, a.Id)
	assert.Equal(t, "", a.Uuid)
}
