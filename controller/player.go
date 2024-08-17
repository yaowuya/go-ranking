package controller

import (
	"github.com/gin-gonic/gin"
	"go-ranking/cache"
	"go-ranking/models"
	"strconv"
	"time"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "")
	aid, err := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid, "id asc")
	if err == nil {
		ReturnSuccess(c, 2000, "success", rs, 1)
		return
	}
	ReturnError(c, 4004, "没有相关信息")
}

func (p PlayerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "")
	aid, _ := strconv.Atoi(aidStr)
	var redisKey string
	redisKey = "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id != 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 2000, "success-redis", players, int64(len(players)))
		return
	}
	rsDb, errDb := models.GetPlayers(aid, "score desc")
	if errDb == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score))
		}
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 2000, "success", rsDb, int64(len(rsDb)))
		return
	}
	ReturnError(c, 4004, "没有相关信息")

}
