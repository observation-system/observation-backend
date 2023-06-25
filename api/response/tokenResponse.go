package response

// 認証サーバーのレスポンス auth: user_check
type UserCheckToken struct {
	UserKey string `json:"user_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUserCheckToken(userKey string, username string, email string) *UserCheckToken {
	return &UserCheckToken{
		UserKey:  userKey,
		Username:  username,
		Email:    email,
	}
}
