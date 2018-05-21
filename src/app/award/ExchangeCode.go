package award

import (
	. "../../common"
	"github.com/gin-gonic/gin"
	Redis "../../redis"
)

type ReqExCode struct {
	Code string `json:"code"`
}

type RespExCode struct {
	AwardType string `json:"awardType"`
	Value     int    `json:"value"`
	Eid       int    `json:"eid"`
	Rare      int    `json:"rare"`
	Ex        string `json:"ex"`
	BaseResp
}

func ExchangeCodeServer(ctx *gin.Context) {
	var result string
	ctx.Request.ParseForm()
	mCode := ctx.Request.FormValue("code")
	name := ctx.Request.FormValue("name")
	result = Redis.HashGetValue(RKEY_H_CODE, mCode)
	if result == "" {
		ctx.String(200, ToJson(BaseResp{2001, "兑换码不存在"}))
		return
	}
	//if uid := GetUid(name); Redis.SetHasValue(mCode, uid) {
	//	ctx.String(200, ToJson(BaseResp{2002, "已经兑换过了"}))
	//} else {
	//	Redis.SetPut(mCode, uid)
	//	ctx.String(200, result)
	//}
	uid := GetUid(name)
	Redis.SetPut(mCode, uid)
	ctx.String(200, result)
}
