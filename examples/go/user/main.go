package main

import (
	"fmt"

	ovh "github.com/lbrlabs/pulumi-ovh/sdk/go/ovh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		user, err := ovh.NewCloudProjectUser(ctx, "example", &ovh.CloudProjectUserArgs{
			ServiceName: pulumi.String("33c6e97fc8d241ff939bba3ad6e11ea7"),
			Description: pulumi.String("created by pulumi-ovh integration tests for Go"),
		})

		if err != nil {
			return fmt.Errorf("error creating Ovh user: %v", err)
		}

		ctx.Export("user", user.ID())

		return nil

	})
}
