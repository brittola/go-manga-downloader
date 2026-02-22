package services

import (
	"fmt"
	"image"
	"manga-downloader/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/signintech/gopdf"
	_ "golang.org/x/image/webp"
)

func createPdfFromImages(directory string, manga string, chapter int) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	utils.OrderImages(files)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) == ".pdf" {
			continue
		}

		imagePath := filepath.Join(directory, file.Name())

		convertedPath := imagePath
		if strings.ToLower(filepath.Ext(file.Name())) == ".webp" {
			jpegPath, err := convertWebpToJpeg(imagePath)
			if err != nil {
				continue
			}
			convertedPath = jpegPath
		}

		orientation, err := getImageOrientation(convertedPath)
		if err != nil {
			continue
		}

		if orientation == "L" {
			pdf.AddPageWithOption(gopdf.PageOption{
				PageSize: &gopdf.Rect{W: gopdf.PageSizeA4.H, H: gopdf.PageSizeA4.W},
			})
			pdf.Image(convertedPath, 0, 0, &gopdf.Rect{W: gopdf.PageSizeA4.H, H: gopdf.PageSizeA4.W})
		} else {
			pdf.AddPage()
			pdf.Image(convertedPath, 0, 0, &gopdf.Rect{W: gopdf.PageSizeA4.W, H: gopdf.PageSizeA4.H})
		}
	}

	pdfFilename := filepath.Join(directory, fmt.Sprintf("%s-%d.pdf", manga, chapter))
	if err := pdf.WritePdf(pdfFilename); err != nil {
		return err
	}

	return nil
}

func getImageOrientation(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return "", err
	}

	if img.Width > img.Height {
		return "L", nil
	}
	return "P", nil
}
