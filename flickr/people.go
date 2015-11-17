package flickr

import (
	"net/url"
	"strconv"
)

type PeopleGetInfoResponse struct {
	Person struct {
		Username   string `xml:"username"`
		Realname   string `xml:"realname"`
		PhotosUrl  string `xml:"photosurl"`
		ProfileUrl string `xml:"profileurl"`
	} `xml:"person"`
}

func (client *Client) UserInfoForId(nsid string) (PeopleGetInfoResponse, error) {
	var v PeopleGetInfoResponse
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
			Id     string `xml:"id,attr"`
			Url    string `xml:"url_l,attr"`
			Height int    `xml:"height_l,attr"`
			Width  int    `xml:"width_l,attr"`
		} `xml:"photo"`
	} `xml:"photos"`
}

func (client *Client) PublicPhotos(nsid string, perPage, page int) (PhotosResponse, error) {
	var v PhotosResponse
	_, err := client.get("flickr.people.getPublicPhotos", url.Values{
		"user_id":  {nsid},
		"page":     {strconv.Itoa(page)},
		"per_page": {strconv.Itoa(perPage)},
		"extras":   {"url_l"},
	}, &v)

	return v, err
}
