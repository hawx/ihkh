package handler

type PhotosCtx struct {
	Title    string
	UserInfo UserInfo
	PrevPage string
	NextPage string
	Photos   []Photo
}

type SetsCtx struct {
	Title    string
	UserInfo UserInfo
	Sets     []Set
}

type TagsCtx struct {
	Title    string
	UserInfo UserInfo
	Tags     []string
}

type Photo struct {
	ID     string
	Src    string
	Width  int
	Height int
}

type Set struct {
	ID    string
	Title string
}

type UserInfo struct {
	ID         string
	PhotosURL  string
	ProfileURL string
	RealName   string
	UserName   string
}

func (info UserInfo) Name() string {
	if len(info.RealName) > 0 {
		return info.RealName
	}

	return info.UserName
}
