package request

type GetRedisKey struct {
	Key string `json:"key" binding:"required" required_err:"缓存key不能为空"`
}

type GetRedisKeys struct {
	Key []string `json:"key" binding:"required" required_err:"缓存key不能为空"`
}
