package flickr

import (
	"net/url"
	"strconv"
)

type TagsResponse struct {
	Tags struct {
		Tag []string `xml:"tag"`
	} `xml:"who>tags"`
}

func (client *httpClient) Tags(nsid string) (TagsResponse, error) {
	var v TagsResponse
	_, err := client.get("flickr.tags.getListUser", url.Values{"user_id": {nsid}}, &v)

	return v, err
}

func (client *httpClient) Tag(nsid, tag string, perPage, page int) (PhotosResponse, error) {
	var v PhotosResponse
	_, err := client.get("flickr.photos.search", url.Values{
		"user_id":  {nsid},
		"tags":     {tag},
		"page":     {strconv.Itoa(page)},
		"per_page": {strconv.Itoa(perPage)},
		"extras":   {"url_l"},
	}, &v)

	return v, err
}
