package config

import (
	"os"
	"strings"
)

const (
	NoahApp = "noah"
	AppName = "apigateway"
)

var (
	AwsAccountID = os.Getenv("AWS_ACCOUNT_ID")
	AwsRegion    = os.Getenv("AWS_REGION")
)

func StageName(stage string) string {
	return strings.ToLower(stage)
}
