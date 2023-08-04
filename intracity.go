package weapp

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const (
	apiCreateStore  = "/cgi-bin/express/intracity/createstore"
	apiQueryStore   = "/cgi-bin/express/intracity/querystore"
	apiUpdateStore  = "/cgi-bin/express/intracity/updatestore"
	apiStoreCharge  = "/cgi-bin/express/intracity/storecharge"
	apiStoreRefund  = "/cgi-bin/express/intracity/storerefund"
	apiQueryFlow    = "/cgi-bin/express/intracity/queryflow"
	apiBalanceQuery = "/cgi-bin/express/intracity/balancequery"
	apiAddOrder     = "/cgi-bin/express/intracity/addorder"
	apiQueryOrder   = "/cgi-bin/express/intracity/queryorder"
	apiCancelOrder  = "/cgi-bin/express/intracity/cancelorder"
)

// ServiceTransPreferType 运力ID类型
type ServiceTransPreferType = string

// 所有绑定运力ID类型
const (
	Dada = "DADA" // 达达
	Sftc = "SFTC" // 顺丰同城
)

// ExpressStore 门店
type ExpressStore struct {
	OutStoreId         string                 `json:"out_store_id"`         // 自定义门店编号
	StoreName          string                 `json:"store_name"`           // 门店名称
	OrderPattern       uint32                 `json:"order_pattern"`        // 运力偏好
	ServiceTransPrefer ServiceTransPreferType `json:"service_trans_prefer"` // 优先使用的运力ID
	AddressInfo        struct {
		Province string  `json:"province"`
		City     string  `json:"city"`
		Area     string  `json:"area"`
		Street   string  `json:"street"`
		House    string  `json:"house"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
		Phone    string  `json:"phone"`
	} `json:"address_info"`
}

func (cli *Client) CreateStore(es *ExpressStore) (*AddStoreResult, error) {
	api := baseURL + apiCreateStore

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.createStore(api, token, es)
}

func (cli *Client) createStore(api, token string, es *ExpressStore) (*AddStoreResult, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(AddStoreResult)
	if err := cli.request.Post(url, es, res); err != nil {
		return nil, err
	}

	return res, nil
}

type AddStoreResult struct {
	request.CommonError
	WxStoreID  string `json:"wx_store_id"` //微信门店编号
	Appid      string `json:"appid"`
	OutStoreID string `json:"out_store_id"` //自定义门店编号
}
