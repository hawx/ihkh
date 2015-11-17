package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"hawx.me/code/ihkh/flickr"
	"hawx.me/code/ihkh/views"
	"hawx.me/code/route"
	"hawx.me/code/serve"
)

var (
	port   = flag.String("port", "8080", "")
	socket = flag.String("socket", "", "")
	userId = flag.String("user-id", "", "")
	apiKey = flag.String("api-key", "", "")
)

func getCtx(client *flickr.Client, userInfo views.UserInfo, page int) (views.Ctx, error) {
	resp, err := client.PublicPhotos(userInfo.Id, 10, page)
	if err != nil {
		return views.Ctx{}, err
	}

	ctx := views.Ctx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Photos:   []views.Photo{},
		UserInfo: userInfo,
		Width:    500,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/%d", resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/%d", resp.Photos.Page+1)
	}

	for _, photo := range resp.Photos.Photo {
		ctx.Photos = append(ctx.Photos, views.Photo{
			Id:     photo.Id,
			Src:    photo.Url,
			Width:  photo.Width,
			Height: photo.Height,
		})
	}

	return ctx, nil
}

func getPhotosetCtx(client *flickr.Client, userInfo views.UserInfo, photoset string, photosetId, page int) (views.Ctx, error) {
	resp, err := client.Photoset(userInfo.Id, photosetId, 10, page)
	if err != nil {
		return views.Ctx{}, err
	}

	ctx := views.Ctx{
		Title:    fmt.Sprintf("ihkh : %s : %s", userInfo.UserName, photoset),
		Photos:   []views.Photo{},
		UserInfo: userInfo,
		Width:    500,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/set/%s/%d", photoset, resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/set/%s/%d", photoset, resp.Photos.Page+1)
	}

	for _, photo := range resp.Photos.Photo {
		ctx.Photos = append(ctx.Photos, views.Photo{
			Id:     photo.Id,
			Src:    photo.Url,
			Width:  photo.Width,
			Height: photo.Height,
		})
	}

	return ctx, nil
}

func main() {
	flag.Parse()

	client := flickr.New(*apiKey)

	user, err := client.UserInfoForId(*userId)
	if err != nil {
		log.Println(err)
		return
	}

	userInfo := views.UserInfo{
		Id:         *userId,
		PhotosUrl:  user.Person.PhotosUrl,
		ProfileUrl: user.Person.ProfileUrl,
		UserName:   user.Person.Username,
		RealName:   user.Person.Realname,
	}

	photosets, err := client.Photosets(*userId)
	if err != nil {
		log.Println(err)
		return
	}

	route.HandleFunc("/:page", func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(route.Vars(r)["page"])
		if err != nil || page < 1 {
			page = 1
		}

		ctx, err := getCtx(client, userInfo, page)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		views.Photostream(w, ctx)
	})

	photosetHandler := func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)
		name := vars["name"]

		page, err := strconv.Atoi(vars["page"])
		if err != nil || page < 1 {
			page = 1
		}

		ctx, err := getPhotosetCtx(client, userInfo, name, photosets[name], page)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		views.Photostream(w, ctx)
	}

	route.HandleFunc("/set/:name", photosetHandler)
	route.HandleFunc("/set/:name/:page", photosetHandler)

	serve.Serve(*port, *socket, route.Default)
}
