package crons

import (
	"testing"

	"github.com/pkg/errors"
)

func TestInfo(t *testing.T) {
	l := NewLogger()
	l.Info("test")

	l.Error(errors.New("test error"), "test")
}
