package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	
	"go-cdk-test/stacks/lambda_api"
)

func main() {
	app := awscdk.NewApp(nil)

	// CloudFormation stack name
	lambda_api.NewAppServerlessCdkGoStack(app, "CDKSVRLESSGOSTACK", &lambda_api.AppServerlessCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}