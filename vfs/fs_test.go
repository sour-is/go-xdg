package vfs_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sour-is/go-xdg/vfs"
)

func TestCreateFile(t *testing.T) {
	is := is.New(t)

	var fs vfs.FileSystem = &vfs.VFS{}

	file, err := fs.Create("/newfile/path")
	is.True(err != nil)  // should return error
	is.True(file == nil) // file should be nil

	err = fs.MkdirAll("/home", 0700)
	is.NoErr(err)

	file, err = fs.Create("/newfile")
	is.NoErr(err)
	is.True(file != nil) // file should have created

	file, err = fs.Create("/newfile/path")
	is.True(err != nil)
	is.True(file == nil) // no file should be created

	t.Logf("%s", fs)
}
