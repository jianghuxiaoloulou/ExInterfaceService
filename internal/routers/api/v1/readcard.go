package v1

import (
	"WowjoyProject/ExInterfaceService/global"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary 读卡信息
// @Produce  json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
// @Router /api/v1/ReadCardInfo [get]
func ReadCardInfo(c *gin.Context) {
	params := make(map[string]string)
	params["requestType"] = "7"
	params["cardType"] = "1"
	url := global.ReadCardSetting.URL

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		global.Logger.Error(err)
		return
	}
	// req.Header.Set("Authorization", token)
	// add params
	// 设置AK
	// req.Header.Set("accessKey", global.ObjectSetting.OBJECT_AK)
	que := req.URL.Query()
	if params != nil {
		for key, val := range params {
			que.Add(key, val)
		}
		req.URL.RawQuery = que.Encode()
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Info("******请求失败******")
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	global.Logger.Info(string(body))
	body_str := strings.Replace(string(body), "(", "", -1)
	body_str = strings.Replace(body_str, ")", "", -1)

	global.Logger.Info(body_str)
	var result = make(map[string]interface{})
	//调用json包的解析，解析请求的body

	err = json.Unmarshal([]byte(body_str), &result)
	if err != nil {
		global.Logger.Error(err)
	}
	global.Logger.Info("解析后的数据: ", result)

	c.JSON(http.StatusOK, result)
}
