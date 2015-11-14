package main

import (
	"flag"
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
		Title:    "ihkh : " + userInfo.UserName,
		Photos:   []views.Photo{},
		UserInfo: userInfo,
		Width:    500,
		PrevPage: -1,
		NextPage: -1,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = resp.Photos.Page - 1
		ctx.HasPrevPage = true
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = resp.Photos.Page + 1
		ctx.HasNextPage = true
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

	serve.Serve(*port, *socket, route.Default)
}
