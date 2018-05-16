package account

import (
	Redis "../../redis"
	. "../../common"
	"strconv"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"sync"
)

type PersonInfo struct {
	Name     string   `json:"name"`
	HP       int      `json:"HP"`
	Atk      int      `json:"atk"`
	Def      int      `json:"def"`
	Exp      int      `json:"exp"`
	Level    int      `json:"level"`
	Gold     int      `json:"gold"`
	MaxFloor int      `json:"maxFloor"`
	RoleType string   `json:"roleType"`
	Ex       ExPerson `json:"ex"`
}

type ExPerson struct {
	Miss        int `json:"miss"`
	Critical    int `json:"critical"`
	CriticalDmg int `json:"critical_dmg"`
	Restore     int `json:"restore"`
}

type Equip struct {
	Level      int                 `json:"level"`
	Info       EquipInfo           `json:"info"`
	Rare       int                 `json:"rare"`
	ExProperty map[string]IntValue `json:"exProperty"`
}

type EquipInfo struct {
	Id int `json:"id"`
}

type IntValue struct {
	Value int `json:"value"`
}

type FightResultData struct {
	Floor int        `json:"floor"`
	P     PersonInfo `json:"p"`
	EqW   Equip      `json:"eqW"`
	EqA   Equip      `json:"eqA"`
	EqR   Equip      `json:"eqR"`
}

func CheckNameServer(ctx *gin.Context) {
	ctx.Request.ParseForm()
	if b := hasUser(ctx.Request.FormValue("name")); b {
		ctx.String(200, ToJson(BaseResp{1000, "exists"}))
	} else {
		ctx.String(200, ToJson(BaseResp{0, "ok"}))
	}
}

func AddNameServer(ctx *gin.Context) {
	name := ctx.PostForm("name")
	addUser(name)
	ctx.String(200, ToJson(BaseResp{0, "ok"}))
}

func StoreDataServer(ctx *gin.Context) {
	var data FightResultData
	ctx.BindJSON(&data)
	storeUser(GetUid(data.P.Name), ToJson(data.P),
		ToJson(data.EqW),
		ToJson(data.EqA),
		ToJson(data.EqR))
	//println("store person:" + data.P.Name)
	ctx.String(200, ToJson(BaseResp{0, "ok"}))
}

func hasUser(name string) bool {
	return Redis.HashHasField(RKEY_H_NAME, name)
}

var AccountMutex sync.Mutex
func addUser(name string) {
	AccountMutex.Lock()
	Redis.HashPutKFV(RKEY_H_NAME, name, strconv.Itoa(Redis.IncKey(RKEY_V_ID_PTR)))
	defer AccountMutex.Unlock()
}

func StoreUserByFRD(uid string, dataJson string) {
	var data FightResultData
	json.Unmarshal([]byte(dataJson), &data)
	storeUser(uid, ToJson(data.P),
		ToJson(data.EqW),
		ToJson(data.EqA),
		ToJson(data.EqR))
}

func storeUser(uid string, userInfo string, eqw string, eqa string, eqr string) {
	Redis.HashPutKFV(RKEY_H_INFO, uid, userInfo)
	Redis.HashPutKFV(RKEY_H_EQW, uid, eqw)
	Redis.HashPutKFV(RKEY_H_EQA, uid, eqa)
	Redis.HashPutKFV(RKEY_H_EQR, uid, eqr)
}

func GetUserInfo(uid string) string {
	return Redis.HashGetValue(RKEY_H_INFO, uid)
}

func GetEquips(uid string) (string, string, string) {
	return Redis.HashGetValue(RKEY_H_EQW, uid),
		Redis.HashGetValue(RKEY_H_EQA, uid),
		Redis.HashGetValue(RKEY_H_EQR, uid)
}
