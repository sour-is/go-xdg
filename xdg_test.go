package xdg_test

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/sour-is/go-xdg"
)

func TestXDG(t *testing.T) {
	is := is.New(t)

	os.Setenv("HOME", "/home/user")

	is.Equal(xdg.UserHome.String(), "/home/user")
}
