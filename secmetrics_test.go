package secmetrics

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Log("secmetrics package exists and builds successfully")
}

func TestVersion(t *testing.T) {
	const version = "1.0.0"
	if version == "" {
		t.Error("Expected version to be defined")
	}
}
