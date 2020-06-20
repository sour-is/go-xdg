// +build darwin

package xdg

var (
	defaultDataHome   = ParsePath("~/Library/Application Support")
	defaultDataDirs   = ParseDirs("/Library/Application Support")
	defaultConfigHome = ParsePath("~/Library/Preferences")
	defaultConfigDirs = ParseDirs("/Library/Preferences")
	defaultCacheHome  = ParsePath("~/Library/Caches")
	defaultRuntime    = ParsePath("~/Library/Application Support")

	defaultDesktop   = ParsePath("~/Desktop")
	defaultDownload  = ParsePath("~/Downloads")
	defaultDocuments = ParsePath("~/Documents")
	defaultMusic     = ParsePath("~/Music")
	defaultPictures  = ParsePath("~/Pictures")
	defaultVideos    = ParsePath("~/Videos")
	defaultTemplates = ParsePath("~/Templates")
	defaultPublic    = ParsePath("~/Public")

	defaultApplicationDirs = ParseDirs("/Applications")
	defaultFontDirs        = Merge(
		ParseDirs("~/Library/Fonts:/Library/Fonts:/System/Library/Fonts:/Network/Library/Fonts"),
		AddSuffix(EnvDefault("XDG_DATA_DIRS", nil), "fonts"),
	)
)
