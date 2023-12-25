package auth

type LoginResponseData struct {
	Jwt string `form:"jwt" json:"jwt"`
}
