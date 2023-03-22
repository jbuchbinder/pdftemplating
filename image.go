package pdftemplating

import (
	"image"

	"github.com/karmdip-mi/go-fitz"
)

// PdfPageToImage creates an image.Image object from the specified PDF page
func PdfPageToImage(fn string, pageNo int) (image.Image, error) {
	doc, err := fitz.New(fn)
	if err != nil {
		return nil, err
	}

	defer doc.Close()

	img, err := doc.Image(pageNo - 1)
	return img, err
}

// PdfPageToImages creates an array of image.Image objects from the
// specified PDF file
func PdfPageToImages(fn string) ([]image.Image, error) {
	doc, err := fitz.New(fn)
	if err != nil {
		return nil, err
	}

	defer doc.Close()

	out := make([]image.Image, 0)
	for i := 0; i < doc.NumPage(); i++ {
		img, err := doc.Image(i)
		if err != nil {
			return out, err
		}
		out = append(out, img)
	}
	return out, err
}

// PdfPageCount returns the number of pages in a PDF document
func PdfPageCount(fn string) (int, error) {
	doc, err := fitz.New(fn)
	if err != nil {
		return 0, err
	}

	defer doc.Close()

	return doc.NumPage(), nil
}
