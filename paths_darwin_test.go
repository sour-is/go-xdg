// +build darwin

package xdg_test

import (
	"os"
)

var (
	testUserHome = os.Getenv("HOME")

	testDataHome   = testUserHome + "/Library/Application Support"
	testDataDirs   = "/Library/Application Support"
	testConfigHome = testUserHome + "/Library/Preferences"
	testConfigDirs = "/Library/Preferences"
	testCacheHome  = testUserHome + "/Library/Caches"
	testRuntime    = testUserHome + "/Library/Application Support"

	testDesktop   = testUserHome + "/Desktop"
	testDownload  = testUserHome + "/Downloads"
	testDocuments = testUserHome + "/Documents"
	testMusic     = testUserHome + "/Music"
	testPictures  = testUserHome + "/Pictures"
	testVideos    = testUserHome + "/Videos"
	testTemplates = testUserHome + "/Templates"
	testPublic    = testUserHome + "/Public"

	testApplications = testUserHome + "/Applications:/Applications"
	testFonts        = testUserHome + "/Library/Fonts:/Library/Fonts:/System/Library/Fonts:/Network/Library/Fonts"

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

	testXDGApplications = testApplications
	testXDGFonts        = testFonts
)

// /xdg/data/Applications:/xdg/data/Applications:/xdg/opt/data/Applications:/Users/jonathan.lundy/Applications:/Applications !=
// /Users/jonathan.lundy/Applications:/Applications
