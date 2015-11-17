package views

import "html/template"

var (
	Photostream = template.Must(template.New("photostream").Parse(photostream)).Execute
	Sets        = template.Must(template.New("sets").Parse(sets)).Execute
)

type PhotosCtx struct {
	Title    string
	UserInfo UserInfo
	Width    int
	PrevPage string
	NextPage string
	Photos   []Photo
}

type SetsCtx struct {
	Title    string
	UserInfo UserInfo
	Width    int
	Sets     []Set
}

type Photo struct {
	Id     string
	Src    string
	Width  int
	Height int
}

type Set struct {
	Id    string
	Title string
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
