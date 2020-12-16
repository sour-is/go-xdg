package xdg_test

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/sour-is/go-xdg"
)

func TestPaths(t *testing.T) {
	is := is.New(t)

	sep := string(os.PathListSeparator)

	is.Equal(xdg.DataHome.String(), testDataHome)     // testDataHome
	is.Equal(xdg.DataDirs.String(), testDataDirs)     // testDataDirs
	is.Equal(xdg.ConfigHome.String(), testConfigHome) // testConfigHome
	is.Equal(xdg.ConfigDirs.String(), testConfigDirs) // testConfigDirs
	is.Equal(xdg.CacheHome.String(), testCacheHome)   // testCacheHome
	is.Equal(xdg.Runtime.String(), testRuntime)       // testRuntime
	is.Equal(xdg.ConfigHome.String(), testConfigHome) // testConfigHome

	is.Equal(xdg.UserHome.String(), testUserHome)
	is.Equal(xdg.UserData.String(), testDataHome+sep+testDataDirs)
	is.Equal(xdg.UserConfig.String(), testConfigHome+sep+testConfigDirs)

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

	os.Setenv("XDG_DATA_HOME", testXDGDataHome)
	os.Setenv("XDG_DATA_DIRS", testXDGDataDirs)
	os.Setenv("XDG_CONFIG_HOME", testXDGConfigHome)
	os.Setenv("XDG_CONFIG_DIRS", testXDGConfigDirs)
	os.Setenv("XDG_CACHE_HOME", testXDGCacheHome)
	os.Setenv("XDG_RUNTIME_DIR", testXDGRuntime)

	os.Setenv("XDG_DESKTOP_DIR", testXDGDesktop)
	os.Setenv("XDG_DOWNLOAD_DIR", testXDGDownload)
	os.Setenv("XDG_DOCUMENTS_DIR", testXDGDocuments)
	os.Setenv("XDG_MUSIC_DIR", testXDGMusic)
	os.Setenv("XDG_PICTURES_DIR", testXDGPictures)
	os.Setenv("XDG_VIDEOS_DIR", testXDGVideos)
	os.Setenv("XDG_TEMPLATES_DIR", testXDGTemplates)
	os.Setenv("XDG_PUBLICSHARE_DIR", testXDGPublic)

	for i, s := range xdg.DataHome.First().Segments() {
		t.Logf("%d - %T %#v", i, s, s)
	}

	is.Equal(xdg.DataHome.String(), testXDGDataHome)
	is.Equal(xdg.DataDirs.String(), testXDGDataDirs)
	is.Equal(xdg.ConfigHome.String(), testXDGConfigHome)
	is.Equal(xdg.ConfigDirs.String(), testXDGConfigDirs)
	is.Equal(xdg.CacheHome.String(), testXDGCacheHome)
	is.Equal(xdg.Runtime.String(), testXDGRuntime)

	is.Equal(xdg.UserHome.String(), testUserHome)
	is.Equal(xdg.UserData.String(), testXDGDataHome+sep+testXDGDataDirs)
	is.Equal(xdg.UserConfig.String(), testXDGConfigHome+sep+testXDGConfigDirs)

	is.Equal(xdg.UserDesktop.String(), testXDGDesktop)
	is.Equal(xdg.UserDocuments.String(), testXDGDocuments)
	is.Equal(xdg.UserDownload.String(), testXDGDownload)
	is.Equal(xdg.UserMusic.String(), testXDGMusic)
	is.Equal(xdg.UserPictures.String(), testXDGPictures)
	is.Equal(xdg.UserVideos.String(), testXDGVideos)
	is.Equal(xdg.UserTemplates.String(), testXDGTemplates)
	is.Equal(xdg.UserPublic.String(), testXDGPublic)

	is.Equal(xdg.Applications.String(), testXDGApplications)
	is.Equal(xdg.Fonts.String(), testXDGFonts)
}
