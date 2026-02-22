package services

import (
	"fmt"
	"manga-downloader/utils"
	"os"
	"path/filepath"
)

func DownloadManga(manga string, chapter int) (error, string) {
	chapterDir := filepath.Join("downloads", manga)
	pdfPath := filepath.Join(chapterDir, fmt.Sprintf("%s-%d.pdf", manga, chapter))

	if _, err := os.Stat(pdfPath); err == nil {
		return nil, pdfPath
	}

	url := utils.GetMangaUrl(manga, chapter)

	if err := downloadImagesFromUrl(url, chapterDir); err != nil {
		return err, ""
	}

	if err := createPdfFromImages(chapterDir, manga, chapter); err != nil {
		return err, ""
	}

	if err := deleteDownloadedImages(chapterDir); err != nil {
		return err, ""
	}

	return nil, pdfPath
}

func deleteDownloadedImages(directory string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ext := filepath.Ext(file.Name()); ext == ".pdf" {
			continue
		}

		filePath := filepath.Join(directory, file.Name())
		os.Remove(filePath)
	}

	return nil
}
