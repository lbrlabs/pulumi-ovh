// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getEndpoint(t *testing.T) string {
	name := os.Getenv("OVH_ENDPOINT")
	if name == "" {
		t.Skipf("Skipping test due to missing OVH_ENDPOINT environment variable")
	}

	return name
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	endpoint := getEndpoint(t)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"endpoint": endpoint,
		},
	}
}