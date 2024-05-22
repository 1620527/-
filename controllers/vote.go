package controllers

import (
	"ginStudy/models"

	"github.com/gin-gonic/gin"
)

type WebVote struct {
	UserId   int `json:"userId"`
	PlayerId int `json:"playerId"`
}

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	var webVote WebVote
	if c.BindJSON(&webVote) != nil {
		ReturnError(c, 4001, "参数错误")
	}

	//TODO: 处理投票逻辑
	if webVote.UserId == 0 || webVote.PlayerId == 0 {
		ReturnError(c, 4002, "请输入正确信息")
		return
	}

	user, _ := models.GetUserInfoByUserId(webVote.UserId)
	if user.ID == 0 {
		ReturnError(c, 4003, "投票用户不存在")
		return
	}

	player, _ := models.GetPlayerInfoById(webVote.PlayerId)
	if player.ID == 0 {
		ReturnError(c, 4003, "参赛选手不存在")
		return
	}

	vote, _ := models.GetVoteInfo(webVote.UserId, webVote.PlayerId)
	if vote.Id != 0 {
		ReturnError(c, 4004, "您已经投过票了")
	}

	err := models.AddVote(webVote.UserId, webVote.PlayerId)
	if err != nil {
		ReturnError(c, 4005, "投票失败")
	}

	//TODO: 增加被投票选手得分
	err2 := models.UpdatePlayerScore(webVote.PlayerId)
	if err2 != nil {
		ReturnError(c, 4006, "更新选手得分失败")
	}

	ReturnSuccess(c, 0, "投票成功", nil, 0)

}
