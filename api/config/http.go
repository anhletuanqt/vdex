package config

type LoginReq struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
}

type IDParams struct {
	ID int64 `params:"id"`
}
