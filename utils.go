package xdgicons

import (
	"os"
	"path"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func GetBaseDirs() (baseDirs []string) {
	// Let's construct this and append "icons" to each path:
	// "${XDG_DATA_HOME:-$HOME/.local/share}:${XDG_DATA_DIRS:-/usr/local/share:/usr/share}"

	xdgDataDirs := "/usr/local/share:/usr/share"
	if value := os.Getenv("XDG_DATA_DIRS"); value != "" {
		xdgDataDirs = value
	}
	dataDirs := strings.Split(xdgDataDirs, ":")

	dataHomeDir := ""
	if value := os.Getenv("XDG_DATA_HOME"); value != "" {
		dataHomeDir = value
	} else if value := os.Getenv("HOME"); value != "" {
		dataHomeDir = path.Join(value, ".local", "share")
	}

	pixmapDir := "/usr/share/pixmaps"

	if dataHomeDir != "" {
		baseDirs = append(baseDirs, path.Join(dataHomeDir, "icons"))
	}

	for _, dataDir := range dataDirs {
		baseDirs = append(baseDirs, path.Join(dataDir, "icons"))
	}

	baseDirs = append(baseDirs, pixmapDir)
	return baseDirs
}
