// +build windows

package xdg_test

var (
	testDataHome   = "%LOCALAPPDATA%"
	testDataDirs   = `%APPDATA%\Roaming:%PROGRAMDATA%`
	testConfigHome = "%LOCALAPPDATA%"
	testConfigDirs = "%PROGRAMDATA%"
	testCacheHome  = "%LOCALAPPDATA%\cache"
	testRuntime    = "%LOCALAPPDATA%"

	testDesktop   = "%USERPROFILE%/Desktop"
	testDownload  = "%USERPROFILE%/Downloads"
	testDocuments = "%USERPROFILE%/Documents"
	testMusic     = "%USERPROFILE%/Music"
	testPictures  = "%USERPROFILE%/Pictures"
	testVideos    = "%USERPROFILE%/Videos"
	testTemplates = "%USERPROFILE%/Templates"
	testPublic    = "%USERPROFILE%/Public"

	testApplicationDirs = "%APPDATA%\Roaming\Microsoft\Windows\Start Menu\Programs"
	testFontDirs        = "%windir%\Fonts:%LOCALAPPDATA%\Microsoft\Windows\Fonts"
)
