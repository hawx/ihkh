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

func getCtx(client *flickr.Client, userInfo views.UserInfo, page int) (views.PhotosCtx, error) {
	resp, err := client.PublicPhotos(userInfo.Id, 10, page)
	if err != nil {
		return views.PhotosCtx{}, err
	}

	ctx := views.PhotosCtx{
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

func getPhotosetsCtx(client *flickr.Client, userInfo views.UserInfo) (views.SetsCtx, error) {
	resp, err := client.Photosets(userInfo.Id)
	if err != nil {
		return views.SetsCtx{}, err
	}

	ctx := views.SetsCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Sets:     []views.Set{},
		UserInfo: userInfo,
		Width:    500,
	}

	for _, set := range resp.Photosets.Photoset {
		ctx.Sets = append(ctx.Sets, views.Set{
			Id:    set.Id,
			Title: set.Title,
		})
	}

	return ctx, nil
}

func getPhotosetCtx(client *flickr.Client, userInfo views.UserInfo, photosetId string, page int) (views.PhotosCtx, error) {
	info, err := client.PhotosetInfo(userInfo.Id, photosetId)
	if err != nil {
		return views.PhotosCtx{}, err
	}

	resp, err := client.Photoset(userInfo.Id, photosetId, 10, page)
	if err != nil {
		return views.PhotosCtx{}, err
	}

	ctx := views.PhotosCtx{
		Title:    fmt.Sprintf("ihkh : %s : %s", userInfo.UserName, info.Photoset.Title),
		Photos:   []views.Photo{},
		UserInfo: userInfo,
		Width:    500,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/set/%s/%d", photosetId, resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/set/%s/%d", photosetId, resp.Photos.Page+1)
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

	route.HandleFunc("/sets", func(w http.ResponseWriter, r *http.Request) {
		ctx, err := getPhotosetsCtx(client, userInfo)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		views.Sets(w, ctx)
	})

	photosetHandler := func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)
		setId := vars["set"]

		page, err := strconv.Atoi(vars["page"])
		if err != nil || page < 1 {
			page = 1
		}

		ctx, err := getPhotosetCtx(client, userInfo, setId, page)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		views.Photostream(w, ctx)
	}

	route.HandleFunc("/sets/:set", photosetHandler)
	route.HandleFunc("/sets/:set/:page", photosetHandler)

	serve.Serve(*port, *socket, route.Default)
}
