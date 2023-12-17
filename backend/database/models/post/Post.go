package post

const tableName = "posts"

type Post struct {
	id         uint16
	text       string
	photos     []string
	userId     uint16
	video      Video
	categories []uint16
}

type Video struct {
	link    string
	preview string
}
