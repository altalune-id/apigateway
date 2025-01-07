package httpapi

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/jsii-runtime-go"
)

func NewNoahApi(stack awscdk.Stack, stage string) awsapigatewayv2.HttpApi {
	apiName := config.NoahApp + "-api-" + config.StageName(stage)

	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("NoahApi"), &awsapigatewayv2.HttpApiProps{
		ApiName: jsii.String(apiName),
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: &[]*string{jsii.String("*")},
		},
	})

	awscdk.NewCfnOutput(stack, jsii.String("NoahApiOutputUrl"), &awscdk.CfnOutputProps{
		Value:      api.Url(),
		ExportName: jsii.String(apiName + "-url"),
	})
	awscdk.NewCfnOutput(stack, jsii.String("NoahApiOutputID"), &awscdk.CfnOutputProps{
		Value:      api.ApiId(),
		ExportName: jsii.String(apiName + "-id"),
	})

	return api
}
