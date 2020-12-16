// +build linux

package xdg_test

import (
	"fmt"
	"os"
	"strconv"
)

var (
	testUserHome = os.Getenv("HOME")

	testDataHome   = testUserHome + "/.local/share"
	testDataDirs   = "/usr/local/share:/usr/share"
	testConfigHome = testUserHome + "/.config"
	testConfigDirs = "/etc/xdg"
	testCacheHome  = testUserHome + "/.cache"
	testRuntime    = "/run/user/" + strconv.Itoa(os.Getuid())

	testDesktop   = testUserHome + "/Desktop"
	testDownload  = testUserHome + "/Downloads"
	testDocuments = testUserHome + "/Documents"
	testMusic     = testUserHome + "/Music"
	testPictures  = testUserHome + "/Pictures"
	testVideos    = testUserHome + "/Videos"
	testTemplates = testUserHome + "/Templates"
	testPublic    = testUserHome + "/Public"

	testApplications = testUserHome + "/.local/share/applications:/usr/local/share/applications:/usr/share/applications"
	testFonts        = testUserHome + "/.fonts:" + testUserHome + "/.local/share/fonts:/usr/local/share/fonts:/usr/share/fonts"

	testXDGDataHome   = "/xdg/data"
	testXDGDataDirs   = "/xdg/data:/xdg/opt/data"
	testXDGConfigHome = "/xdg/config"
	testXDGConfigDirs = "/xdg/config:/xdg/opt/config"
	testXDGCacheHome  = "/xdg/cache"
	testXDGRuntime    = "/xdg/run"

	testXDGDesktop   = "/xdg/desktop"
	testXDGDownload  = "/xdg/download"
	testXDGDocuments = "/xdg/documents"
	testXDGMusic     = "/xdg/music"
	testXDGPictures  = "/xdg/pictures"
	testXDGVideos    = "/xdg/videos"
	testXDGTemplates = "/xdg/templates"
	testXDGPublic    = "/xdg/public"

	testXDGApplications = "/xdg/data/applications:" + testUserHome + "/.local/share/applications:/usr/local/share/applications:/usr/share/applications:/xdg/data/applications:/xdg/opt/data/applications"
	testXDGFonts        = "/xdg/data/fonts:" + testUserHome + "/.fonts:" + testUserHome + "/.local/share/fonts:/usr/local/share/fonts:/usr/share/fonts:/xdg/data/fonts:/xdg/opt/data/fonts"
)

func init() {
	os.Setenv("UID", fmt.Sprint(os.Getuid()))
}

// /xdg/data/applications:/root/.local/share/applications:/usr/local/share/applications:/usr/share/applications:/xdg/data/applications:/xdg/opt/data/applications !=
// /xdg/data/applications: /.local/share/applications:/usr/local/share/applications:/usr/share/applications:/xdg/data/applications:/xdg/opt/data/applications
