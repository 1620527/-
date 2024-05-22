package controllers

import (
	"ginStudy/models"

	"github.com/gin-gonic/gin"
)

type WebPlayer struct {
	Aid int `json:"aid"`
}

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	// aid := c.DefaultPostForm("aid", "0")
	// aid, _ := strconv.Atoi(aidstr)
	var webPlayer WebPlayer
	if err := c.ShouldBindJSON(&webPlayer); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	aid := webPlayer.Aid

	players, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "获取成功", players, 1)
}

func (p PlayerController) GetRanking(c *gin.Context) {
	var webPlayer WebPlayer
	if err := c.ShouldBindJSON(&webPlayer); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	aid := webPlayer.Aid

	players, err := models.GetPlayers(aid, "score desc")
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "获取成功", players, 1)
}
