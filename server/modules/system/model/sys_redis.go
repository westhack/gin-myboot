package model

type RedisInfo struct {
	Key        string `json:"key" binding:"required" required_err:"缓存key不能为空"`
	Value      string `json:"value" binding:"required" required_err:"缓存value不能为空"`
	ExpireTime int64  `json:"expireTime"`
}
