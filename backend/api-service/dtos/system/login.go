package system

type LoginRequest struct {
	LoginName  string `json:"login_name"`
	Password   string `json:"password"`
	Captcha    string `json:"captcha"`
	CaptchaKey string `json:"captcha_key"`
}
