package xdg

import (
	"os"

	"github.com/sour-is/go-xdg/vfs"
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
	EnvStateHome  = Env("XDG_STATE_HOME")
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
	DataHome   Dirs = EnvDefault(EnvDataHome, defaultDataHome)
	DataDirs   Dirs = EnvDefault(EnvDataDirs, defaultDataDirs)
	ConfigHome Dirs = EnvDefault(EnvConfigHome, defaultConfigHome)
	ConfigDirs Dirs = EnvDefault(EnvConfigDirs, defaultConfigDirs)
	CacheHome  Dirs = EnvDefault(EnvCacheHome, defaultCacheHome)
	StateHome  Dirs = EnvDefault(EnvStateHome, defaultStateHome)
	Runtime    Dirs = EnvDefault(EnvRuntime, defaultRuntime)

	UserHome   Dirs = homeFn
	UserData   Dirs = Merge(DataHome, DataDirs)
	UserConfig Dirs = Merge(ConfigHome, ConfigDirs)

	UserDesktop   Dirs = EnvDefault(EnvDesktopDir, defaultDesktop)
	UserDownload  Dirs = EnvDefault(EnvDownloadDir, defaultDownload)
	UserDocuments Dirs = EnvDefault(EnvDocumentsDir, defaultDocuments)
	UserMusic     Dirs = EnvDefault(EnvMusicDir, defaultMusic)
	UserPictures  Dirs = EnvDefault(EnvPicturesDir, defaultPictures)
	UserVideos    Dirs = EnvDefault(EnvVideosDir, defaultVideos)
	UserTemplates Dirs = EnvDefault(EnvTemplatesDir, defaultTemplates)
	UserPublic    Dirs = EnvDefault(EnvPublicShareDir, defaultPublic)

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
