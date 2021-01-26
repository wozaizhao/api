package qqmap

import (
	"encoding/json"
	"fmt"

	"wozaizhao.com/api/common"
	"wozaizhao.com/api/util"
)

// ResGeocoder 地址解析返回
type ResGeocoder struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Title    string `json:"title"`
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
	} `json:"result"`
}

// {
// 	"status": 0,
// 	"message": "query ok",
// 	"result": {
// 		"title": "华漕社区卫生服务中心",
// 		"location": {
// 			"lng": 121.273338,
// 			"lat": 31.242786
// 		},
// 			"ad_info": {
// 			"adcode": "310112"
// 		},
// 		"address_components": {
// 			"province": "上海市",
// 			"city": "上海市",
// 			"district": "闵行区",
// 			"street": "",
// 			"street_number": ""
// 		},
// 			"similarity": 0.8,
// 			"deviation": 1000,
// 			"reliability": 7,
// 			"level": 11
// 	}
// }

const geocoderURL = "https://apis.map.qq.com/ws/geocoder/v1/?address=%s&region=%s&key=%s"

// Geocoder 地址解析
func Geocoder(address, region string) (result ResGeocoder, err error) {
	key := common.GetQQMAPKey()
	urlStr := fmt.Sprintf(geocoderURL, address, region, key)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Status != 0 {
		err = fmt.Errorf("Geocoder error : errmsg=%v", result.Message)
		return
	}
	return
}
