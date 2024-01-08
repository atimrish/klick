package PostController

import (
	"backend/controllers/PostController/types"
	"backend/database/models/post"
	"backend/helpers"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
)

func GetPosts(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	offset, err := strconv.Atoi(c.Query("offset"))
	helpers.HandleError(err)

	posts, err := post.GetPosts(int64(limit), int64(offset))
	helpers.HandleError(err)

	c.JSON(200, gin.H{
		"data": *posts,
	})
}

func GetPostById(c *gin.Context) {
	postId := c.Param("post_id")

	fmt.Println("post_id ", postId)

	data, err := post.GetPostById(postId)

	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(404, gin.H{
			"message": "ничего не найдено",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
	return
}


func AddPost(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	helpers.HandleError(err)

	payload, err := helpers.GetPayloadJWT(accessToken)
	if err != nil {
		helpers.TokenExpiredResponse(c)
		return
	}
	UserId := payload.UserId

	var form types.AddPostForm
	err = c.Bind(&form)

	if err != nil {
		helpers.ErrorResponse(c, 422, err)
		return
	}

	requestForm, err := c.MultipartForm()
	helpers.HandleError(err)

	var photos []string
	var videos []post.Video

	for _, photo := range requestForm.File["photos"] {
		photos = append(photos, helpers.SaveFile(c, photo, "storage/post/photos/"))
	}

	for i, video := range requestForm.File["videos"] {
		newVideo := post.Video{
			Link: helpers.SaveFile(c, video, "storage/post/videos/"),
		}

		if requestForm.File["previews"][i] != nil {
			newVideo.Preview = helpers.SaveFile(
				c,
				requestForm.File["previews"][i],
				"storage/post/previews/",
			)
		}

		videos = append(videos, newVideo)
	}

	helpers.HandleError(err)

	categoriesString := strings.Split(form.Categories, ",")
	var categories []int64

	for _, val := range categoriesString {
		num, err := strconv.Atoi(val)
		helpers.HandleError(err)

		categories = append(categories, int64(num))
	}

	newPost := post.Post{
		Text:       form.Text,
		Photos:     photos,
		UserId:     UserId,
		Videos:     videos,
		Categories: categories,
	}

	go newPost.Insert()

	c.JSON(201, gin.H{
		"message": "пост добавлен",
	})
	return
}

func DeleteById(c *gin.Context) {
	postId := c.Param("post_id")



	err := post.DeletePostById(postId)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "неизвестная ошибка",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "пост удален",
	})
	return
}
