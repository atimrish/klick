package types

import "mime/multipart"

type UpdatePostForm struct {
	Text       string                  `form:"text"`
	Photos     *[]multipart.FileHeader `form:"photo"`
	Videos     *[]multipart.FileHeader `form:"videos"`
	Previews   *[]multipart.FileHeader `form:"previews"`
	Categories string                  `form:"categories"`
}
