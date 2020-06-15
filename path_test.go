package xdg_test

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/sour-is/xdg"
)

func TestPaths(t *testing.T) {
	is := is.New(t)

	os.Setenv("HOME", "/home/user")

	is.Equal(xdg.DataHome.String(), testDataHome)
	is.Equal(xdg.DataDirs.String(), testDataDirs)
	is.Equal(xdg.ConfigHome.String(), testConfigHome)
	is.Equal(xdg.ConfigDirs.String(), testConfigDirs)
	is.Equal(xdg.CacheHome.String(), testCacheHome)
	is.Equal(xdg.Runtime.String(), testRuntime)
	is.Equal(xdg.ConfigHome.String(), testConfigHome)

	is.Equal(xdg.UserHome.String(), "/home/user")
	is.Equal(xdg.UserData.String(), testDataHome+":"+testDataDirs)
	is.Equal(xdg.UserConfig.String(), testConfigHome+":"+testConfigDirs)

	is.Equal(xdg.UserDesktop.String(), testDesktop)
	is.Equal(xdg.UserDocuments.String(), testDocuments)
	is.Equal(xdg.UserDownload.String(), testDownload)
	is.Equal(xdg.UserMusic.String(), testMusic)
	is.Equal(xdg.UserPictures.String(), testPictures)
	is.Equal(xdg.UserVideos.String(), testVideos)
	is.Equal(xdg.UserTemplates.String(), testTemplates)
	is.Equal(xdg.UserPublic.String(), testPublic)

	is.Equal(xdg.Applications.String(), testApplications)
	is.Equal(xdg.Fonts.String(), testFonts)

	os.Setenv("XDG_DATA_HOME", "/xdg/data")
	os.Setenv("XDG_DATA_DIRS", "/xdg/data:/xdg/opt/data")
	os.Setenv("XDG_CONFIG_HOME", "/xdg/config")
	os.Setenv("XDG_CONFIG_DIRS", "/xdg/config:/xdg/opt/config")
	os.Setenv("XDG_CACHE_HOME", "/xdg/cache")
	os.Setenv("XDG_RUNTIME_DIR", "/xdg/run")

	os.Setenv("XDG_DESKTOP_DIR", "/xdg/desktop")
	os.Setenv("XDG_DOWNLOAD_DIR", "/xdg/download")
	os.Setenv("XDG_DOCUMENTS_DIR", "/xdg/documents")
	os.Setenv("XDG_MUSIC_DIR", "/xdg/music")
	os.Setenv("XDG_PICTURES_DIR", "/xdg/pictures")
	os.Setenv("XDG_VIDEOS_DIR", "/xdg/videos")
	os.Setenv("XDG_TEMPLATES_DIR", "/xdg/templates")
	os.Setenv("XDG_PUBLICSHARE_DIR", "/xdg/public")

	is.Equal(xdg.DataHome.String(), "/xdg/data")
	is.Equal(xdg.DataDirs.String(), "/xdg/data:/xdg/opt/data")
	is.Equal(xdg.ConfigHome.String(), "/xdg/config")
	is.Equal(xdg.ConfigDirs.String(), "/xdg/config:/xdg/opt/config")
	is.Equal(xdg.CacheHome.String(), "/xdg/cache")
	is.Equal(xdg.Runtime.String(), "/xdg/run")

	is.Equal(xdg.UserHome.String(), "/home/user")
	is.Equal(xdg.UserData.String(), "/xdg/data:/xdg/data:/xdg/opt/data")
	is.Equal(xdg.UserConfig.String(), "/xdg/config:/xdg/config:/xdg/opt/config")

	is.Equal(xdg.UserDesktop.String(), "/xdg/desktop")
	is.Equal(xdg.UserDocuments.String(), "/xdg/documents")
	is.Equal(xdg.UserDownload.String(), "/xdg/download")
	is.Equal(xdg.UserMusic.String(), "/xdg/music")
	is.Equal(xdg.UserPictures.String(), "/xdg/pictures")
	is.Equal(xdg.UserVideos.String(), "/xdg/videos")
	is.Equal(xdg.UserTemplates.String(), "/xdg/templates")
	is.Equal(xdg.UserPublic.String(), "/xdg/public")

	is.Equal(xdg.Applications.String(), testApplications)
	is.Equal(xdg.Fonts.String(), testFonts+":/xdg/data/fonts:/xdg/opt/data/fonts")
}
