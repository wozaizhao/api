package user

import (
	"net/http"

	"wozaizhao.com/api/common"

	"github.com/gin-gonic/gin"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	log "github.com/sirupsen/logrus"
)

var officialAccount *officialaccount.OfficialAccount

var (
	appID     string
	appSecret string
	serverURL string
)

func init() {
	id := common.GetWxID()
	appID = id["appID"]
	appSecret = id["appSecret"]
	serverURL = id["serverURL"]
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，�memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     appID,
		AppSecret: appSecret,
		Token:     "xxx",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	officialAccount = wc.GetOfficialAccount(cfg)
}

// WxGetConfig 获取js-sdk配置
func WxGetConfig(c *gin.Context) {
	js := officialAccount.GetJs()
	config, err := js.GetConfig(serverURL)
	if err != nil {
		log.Warn(err)
		c.Abort()
	}
	c.JSON(http.StatusOK, config)
}
