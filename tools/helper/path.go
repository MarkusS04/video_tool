package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindLastOccurrence(str string, char rune) string {
	for i := len(str) - 1; i >= 0; i-- {
		if rune(str[i]) == char {
			return str[i+1:]
		}
	}
	return ""
}

func GetLinkPath(name string, count int, prefix string) string {
	characters := []string{"<", ">", ":", "\"", "/", "\\", "|", "?", "*"}
	for _, ch := range characters {
		name = strings.ReplaceAll(name, ch, "_")
	}
	return filepath.FromSlash(fmt.Sprintf("%s/%s%d_%s.mp4", Config.FS.Destination, prefix, count, name))
}

// cleanup deletes destination folder and recreates it
func Cleanup() {
	folder := Config.FS.Destination
	os.RemoveAll(folder)
	fmt.Printf("removed %s\n\n", folder)
	os.Mkdir(folder, os.ModePerm)
}
