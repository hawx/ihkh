package flickr

import (
	"net/url"
	"strconv"
)

type photosetsResponse struct {
	Photosets struct {
		Photoset []struct {
			Id    int    `xml:"id,attr"`
			Title string `xml:"title"`
		} `xml:"photoset"`
	} `xml:"photosets"`
}

func (client *Client) Photosets(nsid string) (sets map[string]int, err error) {
	var v photosetsResponse
	_, err = client.get("flickr.photosets.getList", url.Values{
		"user_id": {nsid},
	}, &v)

	if err != nil {
		return
	}

	sets = map[string]int{}
	for _, set := range v.Photosets.Photoset {
		sets[set.Title] = set.Id
	}

	return
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

func (client *Client) Photoset(nsid string, photosetId, perPage, page int) (PhotosetResponse, error) {
	var v PhotosetResponse
	_, err := client.get("flickr.photosets.getPhotos", url.Values{
		"user_id":     {nsid},
		"photoset_id": {strconv.Itoa(photosetId)},
		"page":        {strconv.Itoa(page)},
		"per_page":    {strconv.Itoa(perPage)},
		"extras":      {"url_l"},
	}, &v)

	return v, err
}
