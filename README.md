# xdg
XDG Base Directory Specification for Go

Provides an implementation of the [XDG Base Directory Specification](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).
The specification defines a set of standard paths for storing application files,
including data and configuration files. For portability and flexibility reasons,
applications should use the XDG defined locations instead of hardcoding paths.
The package also includes the locations of well known user directories.
The current implementation supports Windows, Mac OS and most flavors of Unix.

For more information regarding the XDG Base Directory Specification see:
https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html

For more information regarding the XDG user directories see:
https://wiki.archlinux.org/index.php/XDG_user_directories

## Installation
    go get github.com/adrg/xdg

## Default locations

The package defines sensible defaults for XDG variables which are empty or not
present in the environment.

#### XDG Base Directory

|                 | Unix                                | Mac OS                          | Windows                                 |
| :---            | :---                                | :-----                          | :---                                    |
| XDG_DATA_HOME   | `~/.local/share`                    | `~/Library/Application Support` | `%LOCALAPPDATA%`                        |
| XDG_DATA_DIRS   | `/usr/local/share`<br/>`/usr/share` | `/Library/Application Support`  | `%APPDATA%\Roaming`<br/>`%PROGRAMDATA%` |
| XDG_CONFIG_HOME | `~/.config`                         | `~/Library/Preferences`         | `%LOCALAPPDATA%`                        |
| XDG_CONFIG_DIRS | `/etc/xdg`                          | `/Library/Preferences`          | `%PROGRAMDATA%`                         |
| XDG_CACHE_HOME  | `~/.cache`                          | `~/Library/Caches`              | `%LOCALAPPDATA%\cache`                  |
| XDG_RUNTIME_DIR | `/run/user/UID`                     | `~/Library/Application Support` | `%LOCALAPPDATA%`                        |

#### XDG user directories

|                     | Unix          | Mac OS        | Windows                   |
| :---                | :---          | :-----        | :---                      |
| XDG_DESKTOP_DIR     | `~/Desktop`   | `~/Desktop`   | `%USERPROFILE%/Desktop`   |
| XDG_DOWNLOAD_DIR    | `~/Downloads` | `~/Downloads` | `%USERPROFILE%/Downloads` |
| XDG_DOCUMENTS_DIR   | `~/Documents` | `~/Documents` | `%USERPROFILE%/Documents` |
| XDG_MUSIC_DIR       | `~/Music`     | `~/Music`     | `%USERPROFILE%/Music`     |
| XDG_PICTURES_DIR    | `~/Pictures`  | `~/Pictures`  | `%USERPROFILE%/Pictures`  |
| XDG_VIDEOS_DIR      | `~/Videos`    | `~/Movies`    | `%USERPROFILE%/Videos`    |
| XDG_TEMPLATES_DIR   | `~/Templates` | `~/Templates` | `%USERPROFILE%/Templates` |
| XDG_PUBLICSHARE_DIR | `~/Public`    | `~/Public`    | `%PUBLIC%`                |

#### Non-standard directories

Application directories

```
Unix:
- $XDG_DATA_HOME/applications
- ~/.local/share/applications
- /usr/local/share/applications
- /usr/share/applications
- $XDG_DATA_DIRS/applications

Mac OS:
- /Applications

Windows:
- %APPDATA%\Roaming\Microsoft\Windows\Start Menu\Programs
```

Font Directories

```
Unix:
- $XDG_DATA_HOME/fonts
- ~/.fonts
- ~/.local/share/fonts
- /usr/local/share/fonts
- /usr/share/fonts
- $XDG_DATA_DIRS/fonts

Mac OS:
- ~/Library/Fonts
- /Library/Fonts
- /System/Library/Fonts
- /Network/Library/Fonts

Windows:
- %windir%\Fonts
- %LOCALAPPDATA%\Microsoft\Windows\Fonts
```
