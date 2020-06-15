// +build linux

package xdg_test

var (
	testDataHome   = "~/Library/Application Support"
	testDataDirs   = "/Library/Application Support"
	testConfigHome = "~/Library/Preferences"
	testConfigDirs = "/Library/Preferences"
	testCacheHome  = "~/Library/Caches"
	testRuntime    = "~/Library/Application Support"

	testDesktop   = "~/Desktop"
	testDownload  = "~/Downloads"
	testDocuments = "~/Documents"
	testMusic     = "~/Music"
	testPictures  = "~/Pictures"
	testVideos    = "~/Videos"
	testTemplates = "~/Templates"
	testPublic    = "~/Public"

	testApplicationDirs = "/Applications"
	testFontDirs        = "$XDG_DATA_HOME/fonts:~/.fonts:~/.local/share/fonts:/usr/local/share/fonts:/usr/share/fonts:~/.fonts",
)
