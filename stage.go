package main

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StageProps struct {
	awscdk.StageProps
	Config *config.Config
}

func NewStage(scope constructs.Construct, props *StageProps) {
	stage := awscdk.NewStage(scope, jsii.String(props.Config.StageName), &props.StageProps)
	stackProps := &awscdk.StackProps{Env: props.Env}

	_ = NewStack(stage, stackProps, props.Config)
}
