package main

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/altalune-id/apigateway/httpapi"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StackProps struct {
	awscdk.StackProps
	Stage string
}

func NewStack(scope constructs.Construct, props *StackProps) awscdk.Stack {
	stackID := jsii.String(config.AppName)
	stack := awscdk.NewStack(scope, stackID, &props.StackProps)

	httpapi.NewNoahApi(stack, props.Stage)

	return stack
}
