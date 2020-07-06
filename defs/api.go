package defs

type ResUserLogin struct {
	Token   string `json:"token"`
	Expired int    `json:"expired"`
	UserId  int64  `json:"user_id"`
}
