package result

import (
	"github.com/gin-gonic/gin"
	. "../common"
	Redis "../redis"
	"encoding/json"
	"fmt"
)

type ResultListResp struct {
	ResultList []string `json:"resultList"`
	BaseResp
}

func ReceiveResultListServer(ctx *gin.Context) {
	mFloor := AtoI(ctx.PostForm("floor"))
	mRoleType := ctx.PostForm("roleType")
	mName := ctx.PostForm("name")
	mScore := Redis.GetSortedSetCount(mRoleType, mFloor*100, mFloor*100+99) + mFloor*100
	Redis.PutSortedSet(mRoleType, mScore, mName)
	Redis.SetString(mName, ctx.PostForm("info"))
}

func RetResultListServer(ctx *gin.Context) {
	ctx.Request.ParseForm()
	str, err := json.Marshal(ResultListResp{Redis.GetSortedSet(ctx.Request.FormValue("type"), 20), BaseResp{ErrMsg: "ok"}})
	if err != nil {
		fmt.Errorf("json marshal", err)
	}
	ctx.String(200, string(str))

}
