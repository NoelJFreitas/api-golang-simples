package controllers

import (
	"net/http"

	"github.com/NoelJFreitas/api-golang/api/entities"
	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "falha ao converter o JSON",
		})
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:i], t.tweets[i+1:]...)
			ctx.JSON(http.StatusOK, t.tweets)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "nao encontrado",
	})
}
