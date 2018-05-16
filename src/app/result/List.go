package result

import (
	"github.com/gin-gonic/gin"
	. "../../common"
	Redis "../../redis"
	Account "../account"
	"encoding/json"
	"fmt"
	"sync"
)

type ResultListResp struct {
	ResultList []Account.FightResultData `json:"resultList"`
	BaseResp
}

func ReceiveResultListServer(ctx *gin.Context) {
	mFloor := AtoI(ctx.PostForm("floor"))
	mRoleType := ctx.PostForm("roleType")
	mName := ctx.PostForm("name")
	mScore := Redis.ZSetCount(mRoleType, mFloor*100, mFloor*100+99) + mFloor*100
	storeTopList(mRoleType, mScore, mName)
	Account.StoreUserByFRD(GetUid(mName), ctx.PostForm("info"))
	ctx.String(200, ToJson(BaseResp{0, "ok"}))
}

var ListMutex sync.Mutex

func storeTopList(roleType string, score int, name string) {
	ListMutex.Lock()
	typeKey := getKeyByRoleType(roleType)
	uid := GetUid(name)
	currRank := Redis.ZSetRank(typeKey, GetUid(name))
	if currRank >= 20 { //NOT in using top list
		if Redis.ZSetScore(typeKey, uid) >= score {
			defer ListMutex.Unlock()
			return
		}
	}
	Redis.ZSetPut(typeKey, score, uid)
	defer ListMutex.Unlock()
}

func RetResultListServer(ctx *gin.Context) {
	ctx.Request.ParseForm()
	typeKey := getKeyByRoleType(ctx.Request.FormValue("type"))
	str, err := json.Marshal(ResultListResp{getInfosByKey(typeKey), BaseResp{0, "ok"}})
	if err != nil {
		fmt.Println("ERR:", err)
	}
	ctx.String(200, string(str))
}

func getInfosByKey(typeKey string) []Account.FightResultData {
	uids := Redis.ZSetGet(typeKey, 20)
	retArray := make([]Account.FightResultData, len(uids))
	for i, uid := range uids {
		var p Account.PersonInfo
		var eqw Account.Equip
		var eqa Account.Equip
		var eqr Account.Equip
		json.Unmarshal([]byte(Redis.HashGetValue(RKEY_H_INFO, uid)), &p)
		json.Unmarshal([]byte(Redis.HashGetValue(RKEY_H_EQW, uid)), &eqw)
		json.Unmarshal([]byte(Redis.HashGetValue(RKEY_H_EQA, uid)), &eqa)
		json.Unmarshal([]byte(Redis.HashGetValue(RKEY_H_EQR, uid)), &eqr)
		retArray[i] = Account.FightResultData{
			Floor: Redis.ZSetScore(typeKey, uid) / 100, P: p, EqW: eqw, EqA: eqa, EqR: eqr}
	}
	return retArray
}

func getKeyByRoleType(roleType string) (typeKey string) {
	switch roleType {
	case "ROGUE":
		typeKey = RKEY_Z_TOP_ROG
	case "NEC":
		typeKey = RKEY_Z_TOP_NEC
	case "KNIGHT":
		typeKey = RKEY_Z_TOP_KNT
	case "FIGHTER":
		typeKey = RKEY_Z_TOP_KNT
	default:
		typeKey = RKEY_Z_TOP_FHT
	}
	return
}
