package httpapi

import (
	"github.com/altalune-id/apigateway/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func NewNoahApi(stack awscdk.Stack, cfg *config.Config) awsapigatewayv2.HttpApi {
	apiName := config.NoahApp + "-api-" + cfg.StageLowerCase()

	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("NoahApi"), &awsapigatewayv2.HttpApiProps{
		ApiName: jsii.String(apiName),
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: &[]*string{jsii.String("*")},
		},
	})

	lambdaName := config.NoahApp + "-restapi-" + cfg.StageLowerCase()
	lambdaExportName := lambdaName + "-arn"
	lambdaArn := awscdk.Fn_ImportValue(jsii.String(lambdaExportName))
	lambdaFunction := awslambda.Function_FromFunctionArn(
		stack,
		jsii.String("RestapiLambda"),
		lambdaArn,
	)

	principal := awsiam.NewServicePrincipal(jsii.String("apigateway.amazonaws.com"), &awsiam.ServicePrincipalOpts{})
	lambdaFunction.AddPermission(jsii.String("ApiGatewayInvoke"), &awslambda.Permission{
		Action:    jsii.String("lambda:InvokeFunction"),
		Principal: principal,
		SourceArn: jsii.String("arn:aws:execute-api:" + *stack.Region() + ":" + *stack.Account() + ":" + *api.ApiId() + "/*"),
	})

	lambdaIntegration := jsii.String(config.NoahApp + "-lambda-integration")
	integ := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		lambdaIntegration,
		lambdaFunction,
		&awsapigatewayv2integrations.HttpLambdaIntegrationProps{
			Timeout: awscdk.Duration_Seconds(jsii.Number(15)),
		},
	)

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_ANY},
		Path:        jsii.String("/{proxy+}"),
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
