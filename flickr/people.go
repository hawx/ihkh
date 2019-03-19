package flickr

import (
	"net/url"
	"strconv"
)

type UserInfoResponse struct {
	Username   string `xml:"person>username"`
	Realname   string `xml:"person>realname"`
	PhotosURL  string `xml:"person>photosurl"`
	ProfileURL string `xml:"person>profileurl"`
}

func (client *httpClient) UserInfo(nsid string) (UserInfoResponse, error) {
	var v UserInfoResponse
	_, err := client.get("flickr.people.getInfo", url.Values{"user_id": {nsid}}, &v)

	return v, err
}

type PhotosResponse struct {
	Photos struct {
		Page    int `xml:"page,attr"`
		Pages   int `xml:"pages,attr"`
		PerPage int `xml:"perpage,attr"`
		Total   int `xml:"total,attr"`

		Photo []struct {
			ID     string `xml:"id,attr"`
			URL    string `xml:"url_l,attr"`
			Height int    `xml:"height_l,attr"`
			Width  int    `xml:"width_l,attr"`
		} `xml:"photo"`
	} `xml:"photos"`
}

func (client *httpClient) PublicPhotos(nsid string, perPage, page int) (PhotosResponse, error) {
	var v PhotosResponse
	_, err := client.get("flickr.people.getPublicPhotos", url.Values{
		"user_id":  {nsid},
		"page":     {strconv.Itoa(page)},
		"per_page": {strconv.Itoa(perPage)},
		"extras":   {"url_l"},
	}, &v)

	return v, err
}
