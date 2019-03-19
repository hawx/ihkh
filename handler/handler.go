package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"hawx.me/code/ihkh/flickr"
	"hawx.me/code/route"
)

func Index(client flickr.Client, userInfo UserInfo, pageSize int, templates *template.Template) *Handler {
	return &Handler{
		client:    client,
		userInfo:  userInfo,
		pageSize:  pageSize,
		templates: templates,
		showView:  "photostream.gotmpl",
		get:       getIndex,
	}
}

func Sets(client flickr.Client, userInfo UserInfo, pageSize int, templates *template.Template) *Handler {
	return &Handler{
		client:    client,
		userInfo:  userInfo,
		pageSize:  pageSize,
		templates: templates,
		listView:  "sets.gotmpl",
		showView:  "photostream.gotmpl",
		getAll:    getAllSets,
		get:       getSet,
	}
}

func Tags(client flickr.Client, userInfo UserInfo, pageSize int, templates *template.Template) *Handler {
	return &Handler{
		client:    client,
		userInfo:  userInfo,
		pageSize:  pageSize,
		templates: templates,
		listView:  "tags.gotmpl",
		showView:  "photostream.gotmpl",
		getAll:    getAllTags,
		get:       getTag,
	}
}

type Handler struct {
	client   flickr.Client
	userInfo UserInfo
	pageSize int

	templates *template.Template
	listView  string
	showView  string

	getAll func(flickr.Client, UserInfo) (interface{}, error)
	get    func(flickr.Client, UserInfo, string, int, int) (PhotosCtx, error)
}

func (h *Handler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := h.getAll(h.client, h.userInfo)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := h.templates.ExecuteTemplate(w, h.listView, ctx); err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)
		param := vars["param"]

		page, err := strconv.Atoi(vars["page"])
		if err != nil || page < 1 {
			page = 1
		}

		ctx, err := h.get(h.client, h.userInfo, param, page, h.pageSize)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		if err := h.templates.ExecuteTemplate(w, h.showView, ctx); err != nil {
			log.Println(err)
		}
	}
}
