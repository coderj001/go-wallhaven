package src

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "go-wallheven" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
