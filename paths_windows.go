// +build windows

package xdg

var (
	defaultDataHome   = ParsePath(`%LOCALAPPDATA%`)
	defaultDataDirs   = ParseDirs(`%APPDATA%\Roaming;%ProgramData%`)
	defaultConfigHome = ParsePath(`%LOCALAPPDATA%`)
	defaultConfigDirs = ParseDirs(`%ProgramData%`)
	defaultCacheHome  = ParsePath(`%LOCALAPPDATA%\cache`)
	defaultStateHome  = ParsePath(`%LOCALAPPDATA%\state`)
	defaultRuntime    = ParsePath(`%LOCALAPPDATA%`)

	defaultDesktop   = ParsePath(`%USERPROFILE%\Desktop`)
	defaultDownload  = ParsePath(`%USERPROFILE%\Downloads`)
	defaultDocuments = ParsePath(`%USERPROFILE%\Documents`)
	defaultMusic     = ParsePath(`%USERPROFILE%\Music`)
	defaultPictures  = ParsePath(`%USERPROFILE%\Pictures`)
	defaultVideos    = ParsePath(`%USERPROFILE%\Videos`)
	defaultTemplates = ParsePath(`%USERPROFILE%\Templates`)
	defaultPublic    = ParsePath(`%USERPROFILE%\Public`)

	defaultApplicationDirs = ParseDirs(`%APPDATA%\Roaming\Microsoft\Windows\Start Menu\Programs`)
	defaultFontDirs        = ParseDirs(`%windir%\Fonts;%LOCALAPPDATA%\Microsoft\Windows\Fonts`)
)
