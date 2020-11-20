package config

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtConf struct {
	SignKey         string `yaml:"sign-key"`
	TokenExpireTime int    `yaml:"token-expire-time"`
}

type JwtClaims struct {
	Username string
	Id       int
	jwt.StandardClaims
}
