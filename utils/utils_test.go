package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParamsOrDefault(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"a", "b"}, ParamsOrDefault("a/b", 2), "should work well")
	assert.Equal([]string{"a", "b"}, ParamsOrDefault("/a/b", 2), "should work well")
	assert.Equal([]string{"a", "b", ""}, ParamsOrDefault("/a/b", 3), "should work well")
	assert.Equal([]string{"a", "b"}, ParamsOrDefault("/a/b/c", 2), "should work well")
}

func TestStringOrDefault(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("val", StringOrDefault("val", "haha"), "should work well")
	assert.Equal("haha", StringOrDefault("", "haha"), "should work well")
}

func TestIsOneOf(t *testing.T) {
	assert := assert.New(t)
	arr := []string{"a", "b", "c"}
	assert.True(IsOneOf(arr, "a"))
	assert.True(IsOneOf(arr, "b"))
	assert.True(IsOneOf(arr, "c"))
	assert.False(IsOneOf(arr, "d"))
}
