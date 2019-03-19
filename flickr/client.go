package flickr

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "https://api.flickr.com/services/rest/"
	defaultUserAgent = "me.hawx.ihkh"
)

type Client interface {
	UserInfo(nsid string) (UserInfoResponse, error)
	PublicPhotos(nsid string, perPage, page int) (PhotosResponse, error)

	Photosets(nsid string) (PhotosetsResponse, error)
	Photoset(nsid, photosetID string, perPage, page int) (PhotosetResponse, error)
	PhotosetInfo(nsid string, photosetID string) (PhotosetInfo, error)

	Tags(nsid string) (TagsResponse, error)
	Tag(nsid, tag string, perPage, page int) (PhotosResponse, error)
}

func New(apiKey string) Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &httpClient{
		client: http.DefaultClient,
		apiKey: apiKey,

		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
	}
}
