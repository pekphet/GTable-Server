package app

import (
	"github.com/gin-gonic/gin"
	. "../common"
	"encoding/json"
)

type updateConf struct {
	CurrentVer  int    `json:"current_version"`
	DownloadUrl string `json:"download_url"`
	WikiUrl     string `json:"wiki_url"`
}

type updateResp struct {
	HasNew  bool   `json:"hasNew"`
	Url     string `json:"downloadUrl"`
	WikiUrl string `json:"wikiUrl"`
	BaseResp
}

var update updateConf

func init() {
	json.Unmarshal(ReadConfFile("./update_conf.json"), &update)
}

func CheckVersion(ctx *gin.Context) {
	ctx.Request.ParseForm()
	ctx.String(200, updateRespStr(update.CurrentVer > AtoI(ctx.Request.FormValue("version")), update.DownloadUrl))
}

func ChangeUpdateConf(ctx *gin.Context) {
	ctx.Request.ParseForm()
	confBlob, _ := json.Marshal(updateConf{AtoI(ctx.Request.FormValue("newVersion")),
		ctx.Request.FormValue("downloadUrl"), update.WikiUrl})

	WriteConfFile("./update_conf.json", string(confBlob))
	ctx.String(200, "ok!")
	json.Unmarshal(ReadConfFile("./update_conf.json"), &update)
}

func updateRespStr(hasNew bool, url string) string {
	str, err := json.Marshal(updateResp{hasNew, url, update.WikiUrl, BaseResp{ErrMsg: "ok"}})
	if err != nil {
		return "{\"errCode\":-1,\"errMsg\":\"" + err.Error() + "\""
	} else {
		return string(str)
	}

}
