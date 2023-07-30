package utils

import "github.com/go-playground/validator"

const endpointProfile string = "https://www.googleapis.com/oauth2/v2/userinfo"

var validate = validator.New()
