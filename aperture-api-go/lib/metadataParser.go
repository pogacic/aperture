package lib

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type MetadataStruct struct {
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Exposure     string    `json:"exposure"`
	F_stop       string    `json:"f_stop"`
	Focal_length string    `json:"focal_length"`
	Lens_model   *tiff.Tag `json:"lens_model"`
	Make         *tiff.Tag `json:"make"`
	Model        *tiff.Tag `json:"model"`
	Iso          int       `json:"iso"`
}

func processFStop(fstop string) string {
	stripExp := strings.ReplaceAll(fstop, "\"", "")
	splitExp := strings.Split(stripExp, "/")

	denominator, _ := strconv.Atoi(splitExp[1])
	numerator, _ := strconv.Atoi(splitExp[0])

	s := fmt.Sprintf("%v", float64(numerator)/float64(denominator))

	return s
}

func processExpFocLen(fraction string) string {
	// For processing exposure and focal length fraction values
	stripExp := strings.ReplaceAll(fraction, "\"", "")
	splitExp := strings.Split(stripExp, "/")
	denominator, _ := strconv.Atoi(splitExp[1])

	if denominator == 1 {
		// Return numerator
		return splitExp[0]

	}
	// Return original string value without quotes
	return stripExp
}

func MetadataParse(path string) (MetadataStruct, error) {
	var md MetadataStruct

	// OPEN IMAGE AS f
	f, err := os.Open(path)
	if err != nil {
		return md, err
	}

	// IMAGE DECODE FOR HEIGHT AND WIDTH
	// This is important for the gallery plugin to display images properly
	image, _, err := image.DecodeConfig(f)

	if err != nil {
		return md, err
	}

	imWidth := image.Width
	imHeight := image.Height

	// Set file pointer back to then beginning of the file
	f.Seek(0, 0)

	// DECODE THE EXIF DATA FROM THE IMAGE
	x, _ := exif.Decode(f)

	imExposure, _ := x.Get(exif.ExposureTime)
	imFStop, _ := x.Get(exif.FNumber)
	imFocalLength, _ := x.Get(exif.FocalLength)
	imLensModel, _ := x.Get(exif.LensModel)
	imMake, _ := x.Get(exif.Make)
	imModel, _ := x.Get(exif.Model)
	imIso, _ := x.Get(exif.ISOSpeedRatings)

	md.Width = imWidth
	md.Height = imHeight
	md.Exposure = processExpFocLen(imExposure.String())
	md.F_stop = processFStop(imFStop.String())
	md.Focal_length = processExpFocLen(imFocalLength.String())
	md.Lens_model = imLensModel
	md.Make = imMake
	md.Model = imModel
	md.Iso, _ = strconv.Atoi(imIso.String())

	// Close file
	defer f.Close()

	return md, err
}
