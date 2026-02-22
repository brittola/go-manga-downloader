package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func GetMangaUrl(manga string, chapter int) string {
	var urlByManga = map[string]string{
		"one-piece": "https://mugiwarasoficial.com/manga/manga-one-piece/capitulo-",
	}

	url := urlByManga[manga]

	return fmt.Sprintf("%s%d", url, chapter)
}

func OrderImages(files []os.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		numI, errI := strconv.Atoi(strings.TrimSuffix(files[i].Name(), filepath.Ext(files[i].Name())))
		numJ, errJ := strconv.Atoi(strings.TrimSuffix(files[j].Name(), filepath.Ext(files[j].Name())))

		if errI != nil || errJ != nil {
			return files[i].Name() < files[j].Name()
		}

		return numI < numJ
	})
}
