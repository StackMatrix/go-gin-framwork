package response

type Login struct {
	Jwt string `form:"jwt" json:"jwt"`
}
