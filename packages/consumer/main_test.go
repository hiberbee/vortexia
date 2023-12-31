package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvVariable_Existing(t *testing.T) {
	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")

	value := loadEnvVariable("TEST_VAR")
	assert.Equal(t, "test_value", value, "The loaded value should match the set environment variable.")
}

// Additional tests for non-existing variables, and getKafkaReader would follow a similar pattern.
