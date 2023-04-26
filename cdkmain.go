package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdanodejs" // for node handler
)

type AppServerlessCdkGoStackProps struct {
	awscdk.StackProps
}

func NewAppServerlessCdkGoStack(scope constructs.Construct, id string, props *AppServerlessCdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create Lambda function
	getHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("cdk-go-lambds"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("./lambdasfn/go-hello"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
	})

	// node lambda function
	nodeHandler := awslambdanodejs.NewNodejsFunction(stack, jsii.String("cdk-node-lambds"), &awslambdanodejs.NodejsFunctionProps{
		Runtime: awslambda.Runtime_NODEJS_14_X(),
		// insert specific js/ts file for entry
		Entry:   jsii.String("./lambdasfn/node-hello/index.js"),
		// need dependency lock file
		DepsLockFilePath: jsii.String("./lambdasfn/node-hello/package-lock.json"),
	})

	// API Gateway name
	restApi := awsapigateway.NewRestApi(stack, jsii.String("cdkgoapi"), &awsapigateway.RestApiProps{
		RestApiName:    jsii.String("cdkgoapi"),
		CloudWatchRole: jsii.Bool(false),
	})

	// Create APIGateway RestAPI: go handler
	restApi.Root().AddResource(jsii.String("go-hello"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getHandler, &awsapigateway.LambdaIntegrationOptions{}),
		restApi.Root().DefaultMethodOptions(),
	)

	// Create APIGateway RestAPI: node handler
	restApi.Root().AddResource(jsii.String("node-hello"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(nodeHandler, &awsapigateway.LambdaIntegrationOptions{}),
		restApi.Root().DefaultMethodOptions(),
	)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	// CloudFormation stack name
	NewAppServerlessCdkGoStack(app, "CDKSVRLESSGOSTACK", &AppServerlessCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}