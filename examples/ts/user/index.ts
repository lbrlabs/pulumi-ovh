import * as pulumi from "@pulumi/pulumi";
import * as ovh from "@lbrlabs/ovh";

const user = new ovh.CloudProjectUser("example", {
    serviceName: "33c6e97fc8d241ff939bba3ad6e11ea7",
    description: "created by pulumi-ovh integration tests"
})