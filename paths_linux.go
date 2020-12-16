// +build linux

package xdg

var (
	defaultDataHome   = ParsePath("~/.local/share")
	defaultDataDirs   = ParseDirs("/usr/local/share:/usr/share")
	defaultConfigHome = ParsePath("~/.config")
	defaultConfigDirs = ParseDirs("/etc/xdg")
	defaultCacheHome  = ParsePath("~/.cache")
	defaultRuntime    = ParsePath("/run/user/$UID")

	defaultDesktop   = ParsePath("~/Desktop")
	defaultDownload  = ParsePath("~/Downloads")
	defaultDocuments = ParsePath("~/Documents")
	defaultMusic     = ParsePath("~/Music")
	defaultPictures  = ParsePath("~/Pictures")
	defaultVideos    = ParsePath("~/Videos")
	defaultTemplates = ParsePath("~/Templates")
	defaultPublic    = ParsePath("~/Public")

	defaultApplicationDirs = Merge(
		AddSuffix(EnvDefault("XDG_DATA_HOME", nil), "applications"),
		AddSuffix(defaultDataHome, "applications"),
		AddSuffix(defaultDataDirs, "applications"),
		AddSuffix(EnvDefault("XDG_DATA_DIRS", nil), "applications"),
	)
	defaultFontDirs = Merge(
		AddSuffix(EnvDefault("XDG_DATA_HOME", nil), "fonts"),
		ParsePath("~/.fonts"),
		AddSuffix(defaultDataHome, "fonts"),
		AddSuffix(defaultDataDirs, "fonts"),
		AddSuffix(EnvDefault("XDG_DATA_DIRS", nil), "fonts"),
	)
)
