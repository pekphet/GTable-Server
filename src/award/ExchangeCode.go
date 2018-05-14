package award

import (
	. "../common"
	"github.com/gin-gonic/gin"
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
}
