package services

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"manga-downloader/constants"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly/v2"
)

func downloadImagesFromUrl(pageURL string, targetDir string) error {
	c := colly.NewCollector()

	var imageURLs []string
	c.OnHTML("img.wp-manga-chapter-img", func(e *colly.HTMLElement) {
		imgURL := e.Attr("src")
		if imgURL != "" {
			imageURLs = append(imageURLs, e.Request.AbsoluteURL(imgURL))
		}
	})

	c.Visit(pageURL)

	if err := os.MkdirAll(targetDir, constants.READ_AND_ACCESS_PERMISSION); err != nil {
		return err
	}

	for i, imgURL := range imageURLs {
		imgURL = strings.TrimSpace(imgURL)

		resp, err := http.Get(imgURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		parsedURL, err := url.Parse(imgURL)
		if err != nil {
			continue
		}

		imageExtension := filepath.Ext(parsedURL.Path)
		if imageExtension == "" {
			imageExtension = ".jpg"
		}

		filename := filepath.Join(targetDir, fmt.Sprintf("%d%s", i+1, imageExtension))
		file, err := os.Create(filename)
		if err != nil {
			continue
		}
		defer file.Close()

		io.Copy(file, resp.Body)
	}

	return nil
}

func convertWebpToJpeg(webpPath string) (string, error) {
	file, err := os.Open(webpPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	jpegPath := strings.TrimSuffix(webpPath, filepath.Ext(webpPath)) + ".jpg"
	outFile, err := os.Create(jpegPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return "", err
	}

	return jpegPath, nil
}
