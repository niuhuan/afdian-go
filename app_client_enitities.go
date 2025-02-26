package afdian

// auth_token
type AuthToken struct {
	AuthToken string `json:"auth_token"`
}

// my/account
// "oauth":{} not implemented
type MyAccount struct {
	Login MyAccountLogin `json:"login"`
	UserPrivateId string `json:"user_private_id"`
}

type MyAccountLogin struct {
	Email         string `json:"email"`
	Phone         string `json:"phone"`
}