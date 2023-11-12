package lib

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ImageStruct struct {
	Large     string         `json:"large"`
	Thumbnail string         `json:"thumbnail"`
	Metadata  MetadataStruct `json:"metadata"`
}

type AlbumStruct struct {
	Name        string                 `json:"name"`
	Subtitle    string                 `json:"subtitle"`
	Description string                 `json:"description"`
	PlaceURL    string                 `json:"placeURL"`
	PlaceName   string                 `json:"placeName"`
	StartDate   time.Time              `json:"startDate"`
	EndDate     time.Time              `json:"endDate"`
	Images      map[string]ImageStruct `json:"images"`
}

func UnmarshalAlbumConfig(path string) (AlbumStruct, error) {
	var albumconfig AlbumStruct

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return albumconfig, err
	}

	err = yaml.Unmarshal(yamlFile, &albumconfig)
	if err != nil {
		return albumconfig, err
	}

	return albumconfig, err
}
