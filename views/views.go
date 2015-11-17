package views

import "html/template"

var Photostream = template.Must(template.New("photostream").Parse(photostream)).Execute

type Ctx struct {
	Title       string
	Photos      []Photo
	UserInfo    UserInfo
	Width       int
	PrevPage    string
	HasPrevPage bool
	NextPage    string
	HasNextPage bool
}

type Photo struct {
	Id     string
	Src    string
	Width  int
	Height int
}

type UserInfo struct {
	Id         string
	PhotosUrl  string
	ProfileUrl string
	RealName   string
	UserName   string
}

func (info UserInfo) Name() string {
	if len(info.RealName) > 0 {
		return info.RealName
	}

	return info.UserName
}
