package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	TokenDetails struct {
		AccessToken  string
		RefreshToken string
		AccessUuid   string
		RefreshUuid  string
		AtExpires    int64
		RtExpires    int64
	}

	Claims struct {
		jwt.StandardClaims
		UserID      string `json:"user_id,omitempty"`
		Authorized  bool   `json:"authorized,omitempty"`
		AccessUUID  string `json:"access_uuid,omitempty"`
		RefreshUUID string `json:"refresh_uuid,omitempty"`
		Exp         int64  `json:"exp,omitempty"`
	}
)
