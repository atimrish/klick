package types

import "mime/multipart"

type AddPostForm struct {
	Text       string `form:"text"`
	Userid     int64
	Photos     *[]multipart.FileHeader `form:"photo"`
	Videos     *[]multipart.FileHeader `form:"videos"`
	Previews   *[]multipart.FileHeader `form:"previews"`
	Categories string                  `form:"categories"`
}
