// https://github.com/qeesung/image2ascii

package guides

import (
	"github.com/qeesung/image2ascii/convert"
)

func RenderImage(path, _ string, width int, height int) string {
	convertOptions := convert.DefaultOptions

	convertOptions.FixedWidth = width
	convertOptions.FixedHeight = height - 1
	convertOptions.StretchedScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return converter.ImageFile2ASCIIString(path, &convertOptions)
}
