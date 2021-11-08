package model

type RepData struct {
	CaptchaID           interface{} `json:"captchaId"`
	ProjectCode         interface{} `json:"projectCode"`
	CaptchaType         interface{} `json:"captchaType"`
	CaptchaOriginalPath interface{} `json:"captchaOriginalPath"`
	CaptchaFontType     interface{} `json:"captchaFontType"`
	CaptchaFontSize     interface{} `json:"captchaFontSize"`
	SecretKey           string      `json:"secretKey"`
	OriginalImageBase64 string      `json:"originalImageBase64"`
	Point               Point       `json:"point"`
	JigsawImageBase64   string      `json:"jigsawImageBase64"`
	WordList            []string    `json:"wordList"`
	PointList           interface{} `json:"pointList"`
	PointJSON           interface{} `json:"pointJson"`
	Token               string      `json:"token"`
	Result              bool        `json:"result"`
	CaptchaVerification interface{} `json:"captchaVerification"`
}
type CaptchaInfo struct {
	RepCode string  `json:"repCode"`
	RepMsg  string  `json:"repMsg"`
	RepData RepData `json:"repData"`
	Success bool    `json:"success"`
	Error   bool    `json:"error"`
}

type CaptchaCheckRequest struct {
	CaptchaType string `json:"captchaType"`
	PointJSON   string `json:"pointJson"`
	Token       string `json:"token"`
}

type CaptchaRequest struct {
	CaptchaType string `json:"captchaType"`
	ClientUid   string `json:"clientUid"`
	Ts          int64  `json:"ts"`
}

type CaptchaVerificationRequest struct {
	CaptchaVerification string `json:"captchaVerification"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Config struct {
	Offset float32
}

type BlockPuzzleCheckInfo struct {
	Point     Point  `json:"point"`
	SecretKey string `json:"secretKey"`
}

type ClickWordCheckInfo struct {
	Points    []Point `json:"points"`
	SecretKey string  `json:"secretKey"`
}
