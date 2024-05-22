package models

import (
	"ginStudy/dao"
	"time"
)

type Vote struct {
	Id       int   `json:"id"`
	UserId   int   `json:"userId"`
	PlayerId int   `json:"playerId"`
	AddTime  int64 `json:"addTime"`
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId, playerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id =? AND player_id =?", userId, playerId).First(&vote).Error
	return vote, err
}

func AddVote(userId, playerId int) error {
	vote := Vote{
		UserId:   userId,
		PlayerId: playerId,
		AddTime:  time.Now().Unix(),
	}
	error := dao.Db.Create(&vote).Error
	return error
}
