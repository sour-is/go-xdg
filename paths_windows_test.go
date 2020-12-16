// +build windows

package xdg_test

import "os"

var (
	testDataHome   = os.Getenv(`LOCALAPPDATA`)
	testDataDirs   = os.Getenv(`APPDATA`) + `\Roaming;` + os.Getenv(`PROGRAMDATA`)
	testConfigHome = os.Getenv(`LOCALAPPDATA`)
	testConfigDirs = os.Getenv(`PROGRAMDATA`)
	testCacheHome  = os.Getenv(`LOCALAPPDATA`) + `\cache`
	testRuntime    = os.Getenv(`LOCALAPPDATA`)

	testUserHome  = os.Getenv(`USERPROFILE`)
	testDesktop   = os.Getenv(`USERPROFILE`) + `\Desktop`
	testDownload  = os.Getenv(`USERPROFILE`) + `\Downloads`
	testDocuments = os.Getenv(`USERPROFILE`) + `\Documents`
	testMusic     = os.Getenv(`USERPROFILE`) + `\Music`
	testPictures  = os.Getenv(`USERPROFILE`) + `\Pictures`
	testVideos    = os.Getenv(`USERPROFILE`) + `\Videos`
	testTemplates = os.Getenv(`USERPROFILE`) + `\Templates`
	testPublic    = os.Getenv(`USERPROFILE`) + `\Public`

	testApplications = os.Getenv(`APPDATA`) + `\Roaming\Microsoft\Windows\Start Menu\Programs`
	testFonts        = os.Getenv(`windir`) + `\Fonts;` + os.Getenv(`LOCALAPPDATA`) + `\Microsoft\Windows\Fonts`

	testXDGDataHome   = `C:\xdg\data`
	testXDGDataDirs   = `C:\xdg\data:C:\xdg\opt\data`
	testXDGConfigHome = `C:\xdg\config`
	testXDGConfigDirs = `C:\xdg\config:\xdg\opt\config`
	testXDGCacheHome  = `C:\xdg\cache`
	testXDGRuntime    = `C:\xdg\run`

	testXDGDesktop   = `C:\xdg\desktop`
	testXDGDownload  = `C:\xdg\download`
	testXDGDocuments = `C:\xdg\documents`
	testXDGMusic     = `C:\xdg\music`
	testXDGPictures  = `C:\xdg\pictures`
	testXDGVideos    = `C:\xdg\videos`
	testXDGTemplates = `C:\xdg\templates`
	testXDGPublic    = `C:\xdg\public`

	testXDGApplications = testApplications
	testXDGFonts        = testFonts
)
