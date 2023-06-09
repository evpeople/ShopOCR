// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}
