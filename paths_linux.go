// +build linux

package xdg

var (
	defaultDataHome   = ParsePath("~/.local/share")
	defaultDataDirs   = ParseDirs("/usr/local/share:/usr/share")
	defaultConfigHome = ParsePath("~/.config")
	defaultConfigDirs = ParseDirs("/etc/xdg")
	defaultCacheHome  = ParsePath("~/.cache")
	defaultRuntime    = ParsePath("/run/user/UID")

	defaultDesktop   = ParsePath("~/Desktop")
	defaultDownload  = ParsePath("~/Downloads")
	defaultDocuments = ParsePath("~/Documents")
	defaultMusic     = ParsePath("~/Music")
	defaultPictures  = ParsePath("~/Pictures")
	defaultVideos    = ParsePath("~/Videos")
	defaultTemplates = ParsePath("~/Templates")
	defaultPublic    = ParsePath("~/Public")

	defaultApplicationDirs = Merge(
		ParseDirs("$XDG_DATA_HOME/applications:~/.local/share/applications:/usr/local/share/applications:/usr/share/applications"),
		AddSuffix(EnvDefault("XDG_DATA_DIRS", nil), "applications"),
	)
	defaultFontDirs = Merge(
		ParseDirs("$XDG_DATA_HOME/fonts:~/.local/share/fonts:/usr/local/share/fonts:/usr/share/fonts"),
		AddSuffix(EnvDefault("XDG_DATA_DIRS", nil), "fonts"),
	)
)
