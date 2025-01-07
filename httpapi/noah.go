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

func NewNoahApi(stack awscdk.Stack, stage string) {
	apigwID := jsii.String(config.NoahApp + "-api")
	apigwEndpoint := jsii.String(config.NoahApp + "-api-endpoint")

	api := awsapigatewayv2.NewHttpApi(stack, apigwID, &awsapigatewayv2.HttpApiProps{
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: &[]*string{jsii.String("*")},
		},
	})

	lambdaID := jsii.String(config.NoahApp + "-lambda")
	lambdaIntegration := jsii.String(config.NoahApp + "-lambda-integration")
	lambdaArn := awscdk.Fn_ImportValue(jsii.String(config.NoahExportedLambdaARN("restapi", stage)))
	lambdaFunction := awslambda.Function_FromFunctionArn(stack, lambdaID, lambdaArn)
	principal := awsiam.NewServicePrincipal(jsii.String("apigateway.amazonaws.com"), &awsiam.ServicePrincipalOpts{})

	lambdaFunction.AddPermission(jsii.String("ApiGatewayInvoke"), &awslambda.Permission{
		Action:    jsii.String("lambda:InvokeFunction"),
		Principal: principal,
		SourceArn: jsii.String("arn:aws:execute-api:" + *stack.Region() + ":" + *stack.Account() + ":" + *api.ApiId() + "/*"),
	})

	integ := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		lambdaIntegration,
		lambdaFunction,
		&awsapigatewayv2integrations.HttpLambdaIntegrationProps{},
	)

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_ANY},
		Path:        jsii.String("/{proxy+}"),
	})

	awscdk.NewCfnOutput(stack, apigwEndpoint, &awscdk.CfnOutputProps{
		Value: api.Url(),
	})
}
