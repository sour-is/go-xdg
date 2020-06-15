package vfs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileSystem interface {
	Stat(name string) (os.FileInfo, error)
	IsExist(err error) bool
	MkdirAll(path string, perm os.FileMode) error
	Create(name string) (File, error)
}

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
	Chmod(mode os.FileMode) error
}

var _ File = (*os.File)(nil)

// osFS implements fileSystem using the local disk.
type OsFS struct{}

var _ FileSystem = (*OsFS)(nil)

func (*OsFS) Stat(name string) (os.FileInfo, error)        { return os.Stat(name) }
func (*OsFS) IsExist(err error) bool                       { return os.IsExist(err) }
func (*OsFS) MkdirAll(path string, perm os.FileMode) error { return os.MkdirAll(path, perm) }
func (*OsFS) Create(name string) (File, error)             { return os.Create(name) }

type vfsFile struct {
	name  string
	isDir bool
	mode  os.FileMode
}

var _ File = (*vfsFile)(nil)
var _ os.FileInfo = (*vfsFile)(nil)

func (f *vfsFile) Name() string       { return filepath.Base(f.name) }
func (f *vfsFile) Size() int64        { return 0 }
func (f *vfsFile) Mode() os.FileMode  { return f.mode }
func (f *vfsFile) ModTime() time.Time { return time.Now() }
func (f *vfsFile) IsDir() bool        { return f.isDir }
func (f *vfsFile) Sys() interface{}   { return nil }

func (f *vfsFile) Chmod(mode os.FileMode) error      { f.mode = mode; return nil }
func (f *vfsFile) Close() error                      { return nil }
func (f *vfsFile) Read([]byte) (int, error)          { return 0, nil }
func (f *vfsFile) ReadAt([]byte, int64) (int, error) { return 0, nil }
func (f *vfsFile) Seek(int64, int) (int64, error)    { return 0, nil }
func (f *vfsFile) Stat() (os.FileInfo, error)        { return f, nil }

func (f *vfsFile) String() string { return fmt.Sprintf("%04o %T %s", f.mode, f.isDir, f.name) }

type VFS struct {
	files []*vfsFile
}

var _ FileSystem = (*VFS)(nil)

func (fs *VFS) Stat(name string) (os.FileInfo, error) {
	for _, f := range fs.files {
		if f.name == name {
			return f, nil
		}
	}

	return nil, os.ErrNotExist
}
func (*VFS) IsExist(err error) bool { return os.IsExist(err) }
func (fs *VFS) MkdirAll(path string, perm os.FileMode) error {
	sp := strings.Split(path, string(os.PathSeparator))
	parts := make([]string, 0, len(sp))
	for _, part := range sp {
		if part == "" {
			part = "/"
		}

		parts = append(parts, part)
		path = filepath.Join(parts...)
		f, err := fs.Stat(path)
		if !os.IsNotExist(err) {
			if f.IsDir() {
				continue
			} else {
				return fmt.Errorf("File exists %s", path)
			}
		}
		dir := &vfsFile{name: path, isDir: true, mode: 0700}
		fs.files = append(fs.files, dir)

	}
	return nil
}
func (fs *VFS) Create(name string) (File, error) {
	sp := strings.Split(filepath.Dir(name), string(os.PathSeparator))
	parts := make([]string, 0, len(sp))
	for _, part := range sp {
		if part == "" {
			part = "/"
		}

		parts = append(parts, part)
		path := filepath.Join(parts...)

		f, err := fs.Stat(path)
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("directory does not exist: %s", path)
		}

		if f.IsDir() {
			continue
		} else {
			return nil, fmt.Errorf("File exists %s", path)
		}
	}

	file := &vfsFile{name: name, isDir: false, mode: 0666}
	fs.files = append(fs.files, file)
	return file, nil
}
func (fs *VFS) String() string {
	s := fmt.Sprintln("Filesystem:")
	for _, file := range fs.files {
		s += fmt.Sprintln(file)
	}
	return s
}
