package main

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewStage(app, "Local", &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
	})

	NewStage(app, "Dev-ID", &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
	})

	NewStage(app, "Dev-SG", &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
	})

	NewStage(app, "Prod-ID", &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
	})

	NewStage(app, "Prod-SG", &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	account := "000000000000"
	region := "us-east-1"

	if config.AwsAccountID != "" {
		account = config.AwsAccountID
	}

	if config.AwsRegion != "" {
		region = config.AwsRegion
	}

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
