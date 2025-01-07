package config

import (
	"os"
	"strings"
)

const NoahApp = "noah"

var (
	AwsAccountID = os.Getenv("AWS_ACCOUNT_ID")
	AwsRegion    = os.Getenv("AWS_REGION")
)

func stageName(stage string) string {
	return strings.ToLower(stage)
}

func NoahExportedLambdaARN(function string, stage string) string {
	return NoahApp + "-" + function + "-" + stageName(stage) + "-arn"
}
