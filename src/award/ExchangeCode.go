package award

import (
	. "../common"
	"github.com/gin-gonic/gin"
	Redis "../redis"
)

type ReqExCode struct {
	Code string `json:"code"`
}

type RespExCode struct {
	AwardType string `json:"awardType"`
	Value     int    `json:"value"`
	Eid       int    `json:"eid"`
	Rare      int    `json:"rare"`
	BaseResp
}

func ExchangeCodeServer(ctx *gin.Context) {
	ctx.Request.ParseForm()
	mCode := ctx.Request.FormValue("code")
	mResult := Redis.GetString(mCode)
}
