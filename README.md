# Aperture (Photo Gallery)

Aperture is a _simple_ photo gallery webstack. Created to serve my needs but it may be useful for others.

This document should have everything you need to start.

See a working instance [here](https://pogacic.com).


## Features
- Uses folder structure to layout content into categories and albums.
- Define information such as: location, dates, description etc. within each album's `config.yml` file.
- The API serves image metadata on the fly and displays that information when a thumbnail is expanded.
- Thumbnail lazy loading and once clicked, displays the full image, with zoom and grab capabilities.
- Gallery masonary to display landscape and portait images together cleanly.
- Create "hidden" categories by prefixing a categories folder name with an underscore. Example `_my-hidden-category`. These will not be displayed on the frontend but can still be shared with other users.
- Some SEO support.

### Planned Future Features
- `.png` support and `.mp4` support.
- Cache image metadata instead of processing each image, everytime.
- Add comprehensive logging to the API.
- Dark theme.


## What tech is this using?
- Frontend
  - [Nuxt3](https://nuxt.com)
    - [DaisyUI](https://daisyui.com/)
    - [lightgalleryjs](https://www.lightgalleryjs.com)
    - [vue-masonry-wall](https://www.npmjs.com/package/@yeger/vue-masonry-wall)
    - [iconoir](https://icones.js.org/collection/iconoir)
- Backend
  - [Go](https://go.dev/)
    - [Gin](https://github.com/gin-gonic/gin)
- Containerisation
  - [Docker](https://www.docker.com/)

## Why the standalone REST API?

Not only can it be used in conjunction with the Aperture UI - you may use it for your own projects. Build your own photoframe with a Raspberry Pi? A two-way mirror with a screen behind it displaying your photos? Build your own web interface that will most likely look better than mine? Go crazy 🤙


# Installation
## Build with Docker Compose
1. `git clone git@github.com:pogacic/aperture.git`
2. `cd aperture`
3. Configure variables within `aperature-ui/nuxt.config.ts` (under `runtimeConfig->public`)
4. Configure CORS variable within `docker-compose.yml` (under the API service). This value is the URL of your UI. For example: `https://pogacic.com`.
4. `docker-compose up`

## Build and run API standalone
1. `cd aperture-api`
2. `docker build -t aperture-api .`
3. `docker run -dp 8080:8080 aperture-api`
4. http://localhost:8080/categories

# Photo folder (volume mount for API)
The default API Dockerfile will mount `/photos` to `/opt/aperture_photos`.
Look inside `docker-compose.yml` to change this volume mount under the `aperture-api` service.

You may also need to chown the user and docker group for the API to read the files within.

1. `sudo mkdir /opt/aperture_photos`
2. `sudo chown <user>:docker /opt/aperture_photos`


# Update Aperture

Just bring everything down, git pull, and docker compose everything back up!
Though for a sitemap to generate properly, bring up the API separately before starting the UI.

1. `cd aperture`
2. `git pull`
3. `docker-compose down && docker-compose up --build -d aperture-api && docker-compose build --no-cache  aperture-ui && docker-compose up -d aperture-ui`


# Aperture UI
## Nuxt configuration

It is recommended to change the values within the `nuxt.config.ts` file under the `aperture-ui/` folder before building. This is where you'll be able to change the:
- Site name
- Site description
- Site URL
- What header image will be displayed behind the `Site name`
- Where the API is being served from
- The sitemap hostname


## Theme configuration

If you wish to change the theme of the website then change line 23 of `aperture-ui/tailwind.config.js`. You can find the list of themes within the DaisyUI [documentation](https://daisyui.com/docs/themes/).

# Aperture API

## Quick API endpoints

### Data endpoints
- `/albums`
- `/categories`
- `/<category>/<album>`

### Media endpoints
- `/photos/<category>/<album>/header.jpg`
- `/photos/<category>/<album>/album.jpg`
- `/photos/<category>/<album>/large/large-1.jpg`
- `/photos/<category>/<album>/thumbnail/thumbnail-1.jpg`


## API endpoints example outputs

### `/categories` - will return categories and a random album photo.
Example:
```yaml
{
  "categories": {
    "adventures": {
      "randomAlbumCover": "photos/adventures/andorra-2022/album.jpg"
    },
    "automotive": {
      "randomAlbumCover": "photos/automotive/salone-events-podium-place-2023/album.jpg"
    }
  }
}
```

### `/albums` - will return album config data (no images).
Example:
```yaml
{
  "albums": {
    "adventures": {
      "andorra-2022": {
      "name": "Andorra 2022",
      "subtitle": "Some ski slope",
      "description": "Photos containing snow",
      "placeURL": "google.com/andorra_something_something",
      "placeName": "Andorra",
      "startDate": "2022-03-03T00:00:00Z",
      "endDate": "2022-03-08T00:00:00Z",
      "images": null
      }
    },
    "automotive": {
      "salone-events-podium-place-2023": {
      "name": "Salone Events Podium Place 2023",
      "subtitle": "Cars and coffee",
      "description": "Supercars and classics",
      "placeURL": "https://saloneevents.co.uk/buy/podium-place-car-meet-march",
      "placeName": "Newbury, UK",
      "startDate": "2023-03-04T00:00:00Z",
      "endDate": "2023-03-04T00:00:00Z",
      "images": null
      }
    }
  }
}
```

### `/automotive/salone-events-podium-place-2023` - will return a specific albums photo metadata
Example:
```yaml
{
  "name": "Salone Events Podium Place 2023",
  "subtitle": "Cars and coffee",
  "description": "Supercars and classics",
  "placeURL": "https://saloneevents.co.uk/buy/podium-place-car-meet-march",
  "placeName": "Newbury, UK",
  "startDate": "2023-03-04T00:00:00Z",
  "endDate": "2023-03-04T00:00:00Z",
  "images": {
    "large-1.jpg": {
    "large": "large-1.jpg",
    "thumbnail": "thumbnail-1.jpg",
      "metadata": {
        "width": 3000,
        "height": 2000,
        "exposure": "1/800",
        "f_stop": "2.8",
        "focal_length": "106",
        "lens_model": "EF70-200mm f/4L IS USM",
        "make": "Canon",
        "model": "Canon EOS M50",
        "iso": 200
      }
    },
    "large-10.jpg": {
    "large": "large-10.jpg",
    "thumbnail": "thumbnail-10.jpg",
      "metadata": {
        "width": 3000,
        "height": 2000,
        "exposure": "1/1000",
        "f_stop": "2.8",
        "focal_length": "97",
        "lens_model": "EF70-200mm f/4L IS USM",
        "make": "Canon",
        "model": "Canon EOS M50",
        "iso": 200
      }
    }
  }
}
...
```


## Folder structure
The API looks through the the `/photo` folder, this folder structure has been setup to allow for `categories -> albums`.

Below describes the folder/file layout that has to be adhered to:

```
photos/                         # Root image folder
├─ automotive/                  # Category
│  ├─ drift-event-2023          # Album
│  │  ├─ large                  # Large image folder
│  │  │  ├─ large-1.jpg    
│  │  │  ├─ large-2.jpg
│  │  │  ├─ large-3.jpg
│  │  ├─ thumbnail              # Thumbnail image folder
│  │  │  ├─ thumbnail-1.jpg
│  │  │  ├─ thumbnail-2.jpg
│  │  │  ├─ thumbnail-3.jpg
│  │  ├─ album.jpg              # Album JPG
│  │  ├─ config.yml             # Configuration file
```
An example of this will be within the codebase, see the `aperture-api-go/photos` folder.


## Image file naming
You may name your images however you wish, but suffix `-1`, `-2`, `-3`, `-X` at the end of the filename so full "large" image's can be linked to their respective thumbnail counterparts. (See Folder Structure to get an idea).


## Image file type
This application at the moment only supports the `JPG` file type.


## Why the large/ and thumbnail/ folders?
The `large/` folder can contain fullsized, high quality JPEG's, the `thumbnail/` folder contains those exact same images but at a smaller resolution.


## Config.yml structure
Each album can contain a `config.yml` file - this file allows the user to define data such as the name, subtitle, description of the album.

```yml
name: Andorra 2022
subtitle: Some ski slope
description: Photos containing snow
placeURL: google.com/andorra_something_something
placeName: Andorra
startDate: 2022-03-03
endDate: 2022-03-08
```
- `name`: Name of the album (string).
- `subtitle`: Subtitle of the album (string).
- `description`: Description of the album (string).
- `placeURL`: URL of the place or anything relevant (string).
- `placeName`: Location of the album `YYYY-MM-DD` format (string).
- `startDate`: Start date of album `YYYY-MM-DD` format (string).
- `endDate`: End date of album (string).


# Image Metadata

Image metadata is processed when an album is loaded (You can find metadata functions and structs within `aperture-api-go/lib/metadataParser.go`).

I use Lightroom to export my photography. You can find my `thumbnail` and `large` Lightroom export templates within the `misc/` folder.

# NGINX Configuration

Below are some NGINX configuration blocks to give you a helping hand setting up.

```sh
server {
        server_name api.pogacic.com;

        location / {
                proxy_set_header   X-Forwarded-For $remote_addr;
                proxy_set_header   Host $http_host;
                proxy_pass         "http://localhost:8080";
        }
}

server {
        server_name pogacic.com;

        location / {
                proxy_set_header   X-Forwarded-For $remote_addr;
                proxy_set_header   Host $http_host;
                proxy_pass         "http://localhost:3000";
        }
}

```