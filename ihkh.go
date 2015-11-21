package main

import (
	"flag"
	"fmt"
	"io"
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

type handler struct {
	client   flickr.Client
	userInfo views.UserInfo

	listView func(io.Writer, interface{}) error
	showView func(io.Writer, interface{}) error

	getAll func(flickr.Client, views.UserInfo) (interface{}, error)
	get    func(flickr.Client, views.UserInfo, string, int) (views.PhotosCtx, error)
}

func (h *handler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := h.getAll(h.client, h.userInfo)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		h.listView(w, ctx)
	}
}

func (h *handler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)
		param := vars["param"]

		page, err := strconv.Atoi(vars["page"])
		if err != nil || page < 1 {
			page = 1
		}

		ctx, err := h.get(h.client, h.userInfo, param, page)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		h.showView(w, ctx)
	}
}

func getIndex(client flickr.Client, userInfo views.UserInfo, _ string, page int) (views.PhotosCtx, error) {
	resp, err := client.PublicPhotos(userInfo.Id, 10, page)
	if err != nil {
		return views.PhotosCtx{}, err
	}

	ctx := views.PhotosCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Photos:   []views.Photo{},
		UserInfo: userInfo,
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

func getAllSets(client flickr.Client, userInfo views.UserInfo) (interface{}, error) {
	resp, err := client.Photosets(userInfo.Id)
	if err != nil {
		return nil, err
	}

	ctx := views.SetsCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Sets:     []views.Set{},
		UserInfo: userInfo,
	}

	for _, set := range resp.Photosets.Photoset {
		ctx.Sets = append(ctx.Sets, views.Set{
			Id:    set.Id,
			Title: set.Title,
		})
	}

	return ctx, nil
}

func getSet(client flickr.Client, userInfo views.UserInfo, photosetId string, page int) (views.PhotosCtx, error) {
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
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/sets/%s/%d", photosetId, resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/sets/%s/%d", photosetId, resp.Photos.Page+1)
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

func getAllTags(client flickr.Client, userInfo views.UserInfo) (interface{}, error) {
	resp, err := client.Tags(userInfo.Id)
	if err != nil {
		return nil, err
	}

	ctx := views.TagsCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Tags:     []string{},
		UserInfo: userInfo,
	}

	for _, tag := range resp.Tags.Tag {
		ctx.Tags = append(ctx.Tags, tag)
	}

	return ctx, nil
}

func getTag(client flickr.Client, userInfo views.UserInfo, tag string, page int) (views.PhotosCtx, error) {
	resp, err := client.Tag(userInfo.Id, tag, 10, page)
	if err != nil {
		return views.PhotosCtx{}, err
	}

	ctx := views.PhotosCtx{
		Title:    fmt.Sprintf("ihkh : %s : %s", userInfo.UserName, tag),
		Photos:   []views.Photo{},
		UserInfo: userInfo,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/tags/%s/%d", tag, resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/tags/%s/%d", tag, resp.Photos.Page+1)
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

	user, err := client.UserInfo(*userId)
	if err != nil {
		log.Println(err)
		return
	}

	userInfo := views.UserInfo{
		Id:         *userId,
		PhotosUrl:  user.PhotosUrl,
		ProfileUrl: user.ProfileUrl,
		UserName:   user.Username,
		RealName:   user.Realname,
	}

	index := &handler{
		client:   client,
		userInfo: userInfo,
		showView: views.Photostream,
		get:      getIndex,
	}

	route.HandleFunc("/:page", index.Show())

	sets := &handler{
		client:   client,
		userInfo: userInfo,
		listView: views.Sets,
		showView: views.Photostream,
		getAll:   getAllSets,
		get:      getSet,
	}

	route.HandleFunc("/sets", sets.List())
	route.HandleFunc("/sets/:param", sets.Show())
	route.HandleFunc("/sets/:param/:page", sets.Show())

	tags := &handler{
		client:   client,
		userInfo: userInfo,
		listView: views.Tags,
		showView: views.Photostream,
		getAll:   getAllTags,
		get:      getTag,
	}

	route.HandleFunc("/tags", tags.List())
	route.HandleFunc("/tags/:param", tags.Show())
	route.HandleFunc("/tags/:param/:page", tags.Show())

	serve.Serve(*port, *socket, route.Default)
}
