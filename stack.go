package main

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/altalune-id/apigateway/httpapi"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewStack(scope constructs.Construct, props *awscdk.StackProps, cfg *config.Config) awscdk.Stack {
	stackID := jsii.String(config.AppName)
	stack := awscdk.NewStack(scope, stackID, props)

	_ = httpapi.NewNoahApi(stack, cfg)

	return stack
}
