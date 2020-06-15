package xdg

import (
	"os"

	"github.com/sour-is/xdg/vfs"
)

var fs vfs.FileSystem = &vfs.OsFS{}

func SetFS(newFS vfs.FileSystem) {
	fs = newFS
}

const (
	EnvDataHome   = Env("XDG_DATA_HOME")
	EnvDataDirs   = Env("XDG_DATA_DIRS")
	EnvConfigHome = Env("XDG_CONFIG_HOME")
	EnvConfigDirs = Env("XDG_CONFIG_DIRS")
	EnvCacheHome  = Env("XDG_CACHE_HOME")
	EnvRuntime    = Env("XDG_RUNTIME_DIR")

	EnvDesktopDir     = Env("XDG_DESKTOP_DIR")
	EnvDownloadDir    = Env("XDG_DOWNLOAD_DIR")
	EnvDocumentsDir   = Env("XDG_DOCUMENTS_DIR")
	EnvMusicDir       = Env("XDG_MUSIC_DIR")
	EnvPicturesDir    = Env("XDG_PICTURES_DIR")
	EnvVideosDir      = Env("XDG_VIDEOS_DIR")
	EnvTemplatesDir   = Env("XDG_TEMPLATES_DIR")
	EnvPublicShareDir = Env("XDG_PUBLICSHARE_DIR")
)

var (
	DataHome   Dirs = EnvPath(EnvDataHome, defaultDataHome)
	DataDirs   Dirs = EnvDirs(EnvDataDirs, defaultDataDirs)
	ConfigHome Dirs = EnvPath(EnvConfigHome, defaultConfigHome)
	ConfigDirs Dirs = EnvDirs(EnvConfigDirs, defaultConfigDirs)
	CacheHome  Dirs = EnvPath(EnvCacheHome, defaultCacheHome)
	Runtime    Dirs = EnvPath(EnvRuntime, defaultRuntime)

	UserHome   Dirs = DirList{homeFn}
	UserData   Dirs = PrependDir(DataHome, DataDirs)
	UserConfig Dirs = PrependDir(ConfigHome, ConfigDirs)

	UserDesktop   Dirs = EnvPath(EnvDesktopDir, defaultDesktop)
	UserDownload  Dirs = EnvPath(EnvDownloadDir, defaultDownload)
	UserDocuments Dirs = EnvPath(EnvDocumentsDir, defaultDocuments)
	UserMusic     Dirs = EnvPath(EnvMusicDir, defaultMusic)
	UserPictures  Dirs = EnvPath(EnvPicturesDir, defaultPictures)
	UserVideos    Dirs = EnvPath(EnvVideosDir, defaultVideos)
	UserTemplates Dirs = EnvPath(EnvTemplatesDir, defaultTemplates)
	UserPublic    Dirs = EnvPath(EnvPublicShareDir, defaultPublic)

	Applications Dirs = defaultApplicationDirs
	Fonts        Dirs = defaultFontDirs
)

var homeFn pathFn = func() []Segment {
	home, err := os.UserHomeDir()
	if err != nil {
		return path{Str("")}
	}
	return path{Str(home)}
}
