package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StageProps struct {
	awscdk.StageProps
}

func NewStage(scope constructs.Construct, id string, props *StageProps) {
	stage := awscdk.NewStage(scope, jsii.String(id), &props.StageProps)

	_ = NewStack(stage, &StackProps{
		StackProps: awscdk.StackProps{
			Env: props.Env,
		},
		Stage: id,
	})
}
