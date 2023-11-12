package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"gin-api/lib"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set access control headers for CORS
		c.Header("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))
		c.Header("Access-Control-Allow-Methods", "GET")

		c.Next()
	}
}

func getAlbumData(c *gin.Context) {
	// Create vars for URL route parameters
	category := c.Param("category")
	album := c.Param("album")

	// Load given album config.yml
	yamlData, err := lib.UnmarshalAlbumConfig("photos/" + category + "/" + album + "/config.yml")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	data := new(lib.AlbumStruct)

	// Populate struct
	data.Name = yamlData.Name
	data.Subtitle = yamlData.Subtitle
	data.Description = yamlData.Description
	data.PlaceURL = yamlData.PlaceURL
	data.PlaceName = yamlData.PlaceName
	data.StartDate = yamlData.StartDate
	data.EndDate = yamlData.EndDate

	path := "photos/" + category + "/" + album + "/large/"

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Could not read /photos folder")
	}

	imagesMap := make(map[string]lib.ImageStruct)

	for _, img := range entries {

		metadata, err := lib.MetadataParse(path + img.Name())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "could not parse metadata - check logs"})
			return
		}

		thumbnail := strings.Replace(img.Name(), "large", "thumbnail", -1)

		p := lib.ImageStruct{
			Large:     img.Name(),
			Thumbnail: thumbnail,
			Metadata:  metadata,
		}

		imagesMap[img.Name()] = p
	}

	data.Images = imagesMap
	c.JSON(http.StatusOK, data)
}

func getCategories(c *gin.Context) {
	data, err := lib.ReturnOnlyCategories()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "could not generate albums"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func getAlbums(c *gin.Context) {
	data, err := lib.ReturnOnlyAlbums()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "could not generate albums"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func main() {
	router := gin.New()

	// Don't explode on panic
	router.Use(gin.Recovery())

	// Set /photo local dir to /photos endpoint to serve static files
	router.Static("/photos", "./photos")

	// Set routes
	router.GET("/categories", CORSMiddleware(), getCategories)
	router.GET("/albums", CORSMiddleware(), getAlbums)
	router.GET("/:category/:album", CORSMiddleware(), getAlbumData)

	fmt.Println("Starting Aperture :)")

	// Run Gin on port 8080
	router.Run("0.0.0.0:8080")
}
