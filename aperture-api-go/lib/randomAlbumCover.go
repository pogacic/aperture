package lib

import (
	"log"
	"math/rand"
	"os"
)

var PHOTOS_PATH = "photos/"

type albumsStruct struct {
	Albums map[string]map[string]AlbumStruct `json:"albums"`
}

type categoriesStruct struct {
	Categories map[string]RandomAlbumStruct `json:"categories"`
}

type RandomAlbumStruct struct {
	RAlbumCover string `json:"randomAlbumCover"`
}

func isDirectory(fullpath string) bool {
	checkEntry, err := os.Stat(fullpath)
	if err != nil {
		return false
	}
	if checkEntry.IsDir() {
		return true
	}
	return false
}

func getAlbums(category string) map[string]AlbumStruct {
	data := make(map[string]AlbumStruct)
	path := "photos/"

	entries, err := os.ReadDir(path + category)
	if err != nil {
		return data
	}
	for _, e := range entries {
		fpath := path + category + "/" + e.Name()
		if isDirectory(fpath) {
			yamlData, err := UnmarshalAlbumConfig(fpath + "/config.yml")
			if err != nil {
				return data
			}

			var album AlbumStruct

			album.Name = yamlData.Name
			album.Subtitle = yamlData.Subtitle
			album.Description = yamlData.Description
			album.PlaceURL = yamlData.PlaceURL
			album.PlaceName = yamlData.PlaceName
			album.StartDate = yamlData.StartDate
			album.EndDate = yamlData.EndDate
			data[e.Name()] = album
		}
	}
	return data
}

func ReturnOnlyAlbums() (albumsStruct, error) {
	var data albumsStruct

	entries, err := os.ReadDir(PHOTOS_PATH)
	if err != nil {
		return data, err
	}

	albumMap := make(map[string]map[string]AlbumStruct)

	for _, e := range entries {
		category := e.Name()
		if isDirectory(PHOTOS_PATH + category) {
			albumMap[category] = getAlbums(category)
		}
	}

	data.Albums = albumMap

	return data, err
}

func ReturnOnlyCategories() (categoriesStruct, error) {
	var data categoriesStruct

	entries, err := os.ReadDir(PHOTOS_PATH)
	if err != nil {
		return data, err
	}

	categoryMap := make(map[string]RandomAlbumStruct)

	for _, e := range entries {
		category := e.Name()
		if isDirectory(PHOTOS_PATH + category) {
			categoryMap[category] = getRandomAlbum(category)
		}
	}

	data.Categories = categoryMap

	return data, err
}

func getRandomAlbum(category string) RandomAlbumStruct {
	var x RandomAlbumStruct
	path := "photos/" + category + "/"

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var albumSlice []string
	for _, e := range entries {
		if isDirectory(path + e.Name()) {
			albumSlice = append(albumSlice, e.Name())
		}
	}

	x.RAlbumCover = path + albumSlice[rand.Intn(len(albumSlice))] + "/album.jpg"

	return x
}
