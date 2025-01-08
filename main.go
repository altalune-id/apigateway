package main

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewStage(app, &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
		Config: &config.Config{
			StageName: "Local",
		},
	})

	NewStage(app, &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
		Config: &config.Config{
			StageName: "Dev-ID",
		},
	})

	NewStage(app, &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
		Config: &config.Config{
			StageName: "Dev-SG",
		},
	})

	NewStage(app, &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
		Config: &config.Config{
			StageName: "Prod-ID",
		},
	})

	NewStage(app, &StageProps{
		StageProps: awscdk.StageProps{Env: env()},
		Config: &config.Config{
			StageName: "Prod-SG",
		},
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
