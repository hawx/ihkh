package flickr

import (
	"net/url"
	"strconv"
)

type PhotosetsResponse struct {
	Photosets struct {
		Photoset []struct {
			Id    string `xml:"id,attr"`
			Title string `xml:"title"`
		} `xml:"photoset"`
	} `xml:"photosets"`
}

func (client *httpClient) Photosets(nsid string) (PhotosetsResponse, error) {
	var v PhotosetsResponse
	_, err := client.get("flickr.photosets.getList", url.Values{
		"user_id": {nsid},
	}, &v)

	return v, err
}

type PhotosetResponse struct {
	Photos struct {
		Page    int `xml:"page,attr"`
		PerPage int `xml:"perpage,attr"`
		Pages   int `xml:"pages,attr"`

		Photo []struct {
			Id     string `xml:"id,attr"`
			Url    string `xml:"url_l,attr"`
			Height int    `xml:"height_l,attr"`
			Width  int    `xml:"width_l,attr"`
		} `xml:"photo"`
	} `xml:"photoset"`
}

func (client *httpClient) Photoset(nsid, photosetId string, perPage, page int) (PhotosetResponse, error) {
	var v PhotosetResponse
	_, err := client.get("flickr.photosets.getPhotos", url.Values{
		"user_id":     {nsid},
		"photoset_id": {photosetId},
		"page":        {strconv.Itoa(page)},
		"per_page":    {strconv.Itoa(perPage)},
		"extras":      {"url_l"},
	}, &v)

	return v, err
}

type PhotosetInfo struct {
	Photoset struct {
		Title string `xml:"title"`
	} `xml:"photoset"`
}

func (client *httpClient) PhotosetInfo(nsid string, photosetId string) (PhotosetInfo, error) {
	var v PhotosetInfo
	_, err := client.get("flickr.photosets.getInfo", url.Values{
		"user_id":     {nsid},
		"photoset_id": {photosetId},
	}, &v)

	return v, err
}
