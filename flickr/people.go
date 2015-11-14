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
			Id             string `xml:"id,attr"`
			Secret         string `xml:"secret,attr"`
			Server         string `xml:"server,attr"`
			Farm           int    `xml:"farm,attr"`
			Title          string `xml:"title,attr"`
			License        int    `xml:"license,attr"`
			DateTaken      string `xml:"datetaken,attr"`
			DateUploaded   string `xml:"dateuploaded,attr"`
			Owner          string `xml:"owner,attr"`
			Tags           string `xml:"tags,attr"`
			OriginalSecret string `xml:"originalsecret,attr"`
			OriginalFormat string `xml:"originalformat,attr"`

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
		"extras":   {"date_taken,date_upload,original_format,tags,license,url_l"},
	}, &v)

	return v, err
}
