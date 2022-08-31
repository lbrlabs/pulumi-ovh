using System.Collections.Generic;
using Pulumi;
using Pulumiverse.Ovh;

return await Deployment.RunAsync(() =>
{

   var user = new CloudProjectUser("example", new CloudProjectUserArgs{
      ServiceName = "33c6e97fc8d241ff939bba3ad6e11ea7",
      Description = "created by pulumi-ovh integration tests for DotNet"
   });

});
