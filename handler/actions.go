package handler

import (
	"fmt"

	"hawx.me/code/ihkh/flickr"
)

func getIndex(client flickr.Client, userInfo UserInfo, _ string, page, pageSize int) (PhotosCtx, error) {
	resp, err := client.PublicPhotos(userInfo.ID, pageSize, page)
	if err != nil {
		return PhotosCtx{}, err
	}

	ctx := PhotosCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Photos:   []Photo{},
		UserInfo: userInfo,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/%d", resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/%d", resp.Photos.Page+1)
	}

	for _, photo := range resp.Photos.Photo {
		ctx.Photos = append(ctx.Photos, Photo{
			ID:     photo.ID,
			Src:    photo.URL,
			Width:  photo.Width,
			Height: photo.Height,
		})
	}

	return ctx, nil
}

func getAllSets(client flickr.Client, userInfo UserInfo) (interface{}, error) {
	resp, err := client.Photosets(userInfo.ID)
	if err != nil {
		return nil, err
	}

	ctx := SetsCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Sets:     []Set{},
		UserInfo: userInfo,
	}

	for _, set := range resp.Photosets.Photoset {
		ctx.Sets = append(ctx.Sets, Set{
			ID:    set.ID,
			Title: set.Title,
		})
	}

	return ctx, nil
}

func getSet(client flickr.Client, userInfo UserInfo, photosetID string, page, pageSize int) (PhotosCtx, error) {
	info, err := client.PhotosetInfo(userInfo.ID, photosetID)
	if err != nil {
		return PhotosCtx{}, err
	}

	resp, err := client.Photoset(userInfo.ID, photosetID, pageSize, page)
	if err != nil {
		return PhotosCtx{}, err
	}

	ctx := PhotosCtx{
		Title:    fmt.Sprintf("ihkh : %s : %s", userInfo.UserName, info.Photoset.Title),
		Photos:   []Photo{},
		UserInfo: userInfo,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/sets/%s/%d", photosetID, resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/sets/%s/%d", photosetID, resp.Photos.Page+1)
	}

	for _, photo := range resp.Photos.Photo {
		ctx.Photos = append(ctx.Photos, Photo{
			ID:     photo.ID,
			Src:    photo.URL,
			Width:  photo.Width,
			Height: photo.Height,
		})
	}

	return ctx, nil
}

func getAllTags(client flickr.Client, userInfo UserInfo) (interface{}, error) {
	resp, err := client.Tags(userInfo.ID)
	if err != nil {
		return nil, err
	}

	ctx := TagsCtx{
		Title:    fmt.Sprintf("ihkh : %s", userInfo.UserName),
		Tags:     []string{},
		UserInfo: userInfo,
	}

	for _, tag := range resp.Tags.Tag {
		ctx.Tags = append(ctx.Tags, tag)
	}

	return ctx, nil
}

func getTag(client flickr.Client, userInfo UserInfo, tag string, page, pageSize int) (PhotosCtx, error) {
	resp, err := client.Tag(userInfo.ID, tag, pageSize, page)
	if err != nil {
		return PhotosCtx{}, err
	}

	ctx := PhotosCtx{
		Title:    fmt.Sprintf("ihkh : %s : %s", userInfo.UserName, tag),
		Photos:   []Photo{},
		UserInfo: userInfo,
	}

	if resp.Photos.Page > 1 {
		ctx.PrevPage = fmt.Sprintf("/tags/%s/%d", tag, resp.Photos.Page-1)
	}
	if resp.Photos.Page != resp.Photos.Pages {
		ctx.NextPage = fmt.Sprintf("/tags/%s/%d", tag, resp.Photos.Page+1)
	}

	for _, photo := range resp.Photos.Photo {
		ctx.Photos = append(ctx.Photos, Photo{
			ID:     photo.ID,
			Src:    photo.URL,
			Width:  photo.Width,
			Height: photo.Height,
		})
	}

	return ctx, nil
}
