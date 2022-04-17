package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) CreatePost(ctx *gin.Context) {

}

func (h *Handler) GetAllPosts(ctx *gin.Context) {
	fmt.Sprintln("bitch")
	log.Println("bitch requested")
}

func (h *Handler) GetPostByID(ctx *gin.Context) {

}

func (h *Handler) UpdatePost(ctx *gin.Context) {

}

func (h *Handler) DeletePost(ctx *gin.Context) {

}