package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sour-is/xdg/vfs"
)

type Dirs interface {
	Paths() []Path
	String() string
	First() Path
}

type DirList []Path

var _ Dirs = (DirList)(nil)

func (d DirList) String() string {
	paths := d.Paths()
	lis := make([]string, len(paths))
	for i := range paths {
		lis[i] = paths[i].String()
	}

	return strings.Join(lis, string(os.PathListSeparator))
}

func (d DirList) Paths() []Path {
	return []Path(d)
}

func (d DirList) First() Path {
	lis := d.Paths()
	if len(lis) == 0 {
		return path{Str("")}
	}
	return lis[0]
}

type DirFn func() []Path

var _ Dirs = (*DirFn)(nil)

func (d DirFn) String() string {
	var lis DirList = d.Paths()
	return lis.String()
}

func (d DirFn) Paths() []Path {
	var lis []Path = d()
	return lis
}

func (d DirFn) First() Path {
	lis := d.Paths()
	if len(lis) == 0 {
		return path{Str("")}
	}
	return lis[0]
}

func Find(d Dirs, name string) (string, error) {
	for _, p := range AddDirSuffix(d, name).Paths() {
		path := p.String()
		if exists(path) {
			return path, nil
		}
	}

	return "", fmt.Errorf("could not locate `%s` in any of the following paths: %s",
		filepath.Base(name), d)
}

func FindAll(d Dirs, name string) ([]string, error) {
	paths := AddDirSuffix(d, name).Paths()
	foundPaths := make([]string, 0, len(paths))

	for _, p := range paths {
		path := p.String()
		if exists(path) {
			foundPaths = append(foundPaths, path)
		}
	}

	if len(foundPaths) == 0 {
		return nil, fmt.Errorf("could not locate `%s` in any of the following paths: %s",
			filepath.Base(name), d)
	}

	return foundPaths, nil
}

func Create(d Dirs, name string) (vfs.File, error) {
	p := d.First()

	_, err := createPath(p)
	if err != nil {
		return nil, err
	}

	f, err := fs.Create(AddPathSuffix(p, name).String())
	if err != nil {
		return nil, err
	}

	return f, f.Chmod(0600)

}

func NewDirs(a Path, lis ...Path) Dirs {

	return DirList(append([]Path{a}, lis...))
}

func ParseDirs(s string) Dirs {
	sp := strings.Split(s, string(os.PathListSeparator))
	lis := make([]Path, 0, len(sp))

	for i := range sp {
		lis = append(lis, ParsePath(sp[i]))
	}

	return DirList(lis)
}

func AddDirSuffix(d Dirs, s string) Dirs {
	return DirFn(func() []Path {
		paths := d.Paths()
		lis := make([]Path, len(paths))
		for i := range paths {
			lis[i] = path(append(paths[i].Segments(), Str(s)))
		}

		return lis
	})
}

func EnvDirs(env Env, defaultDirs Dirs) Dirs {
	return DirFn(func() []Path {
		if env.IsSet() {
			return ParseDirs(env.String()).Paths()
		}

		if defaultDirs == nil {
			return []Path(nil)
		}

		return defaultDirs.Paths()
	})
}

func PrependDir(pre Dirs, rest Dirs) Dirs {
	return DirFn(func() []Path {
		return append(pre.Paths(), rest.Paths()...)
	})
}
