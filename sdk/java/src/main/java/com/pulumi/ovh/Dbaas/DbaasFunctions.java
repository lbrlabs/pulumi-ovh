// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ovh.Dbaas.inputs.GetLogsClusterArgs;
import com.pulumi.ovh.Dbaas.inputs.GetLogsClusterPlainArgs;
import com.pulumi.ovh.Dbaas.inputs.GetLogsInputEngineArgs;
import com.pulumi.ovh.Dbaas.inputs.GetLogsInputEnginePlainArgs;
import com.pulumi.ovh.Dbaas.inputs.GetLogsOutputGraylogStreamArgs;
import com.pulumi.ovh.Dbaas.inputs.GetLogsOutputGraylogStreamPlainArgs;
import com.pulumi.ovh.Dbaas.outputs.GetLogsClusterResult;
import com.pulumi.ovh.Dbaas.outputs.GetLogsInputEngineResult;
import com.pulumi.ovh.Dbaas.outputs.GetLogsOutputGraylogStreamResult;
import com.pulumi.ovh.Utilities;
import java.util.concurrent.CompletableFuture;

public final class DbaasFunctions {
    /**
     * Use this data source to retrieve informations about a DBaas logs cluster tenant.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsClusterArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsCluster(GetLogsClusterArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLogsClusterResult> getLogsCluster(GetLogsClusterArgs args) {
        return getLogsCluster(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve informations about a DBaas logs cluster tenant.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsClusterArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsCluster(GetLogsClusterArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLogsClusterResult> getLogsClusterPlain(GetLogsClusterPlainArgs args) {
        return getLogsClusterPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve informations about a DBaas logs cluster tenant.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsClusterArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsCluster(GetLogsClusterArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLogsClusterResult> getLogsCluster(GetLogsClusterArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Dbaas/getLogsCluster:getLogsCluster", TypeShape.of(GetLogsClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve informations about a DBaas logs cluster tenant.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsClusterArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsCluster(GetLogsClusterArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLogsClusterResult> getLogsClusterPlain(GetLogsClusterPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Dbaas/getLogsCluster:getLogsCluster", TypeShape.of(GetLogsClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about a DBaas logs input engine.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsInputEngineArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsInputEngine(GetLogsInputEngineArgs.builder()
     *             .isDeprecated(true)
     *             .name(&#34;logstash&#34;)
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .version(&#34;6.8&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLogsInputEngineResult> getLogsInputEngine(GetLogsInputEngineArgs args) {
        return getLogsInputEngine(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a DBaas logs input engine.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsInputEngineArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsInputEngine(GetLogsInputEngineArgs.builder()
     *             .isDeprecated(true)
     *             .name(&#34;logstash&#34;)
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .version(&#34;6.8&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLogsInputEngineResult> getLogsInputEnginePlain(GetLogsInputEnginePlainArgs args) {
        return getLogsInputEnginePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a DBaas logs input engine.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsInputEngineArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsInputEngine(GetLogsInputEngineArgs.builder()
     *             .isDeprecated(true)
     *             .name(&#34;logstash&#34;)
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .version(&#34;6.8&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLogsInputEngineResult> getLogsInputEngine(GetLogsInputEngineArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Dbaas/getLogsInputEngine:getLogsInputEngine", TypeShape.of(GetLogsInputEngineResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about a DBaas logs input engine.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsInputEngineArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var logstash = DbaasFunctions.getLogsInputEngine(GetLogsInputEngineArgs.builder()
     *             .isDeprecated(true)
     *             .name(&#34;logstash&#34;)
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .version(&#34;6.8&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLogsInputEngineResult> getLogsInputEnginePlain(GetLogsInputEnginePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Dbaas/getLogsInputEngine:getLogsInputEngine", TypeShape.of(GetLogsInputEngineResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about a DBaas logs output graylog stream.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsOutputGraylogStreamArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var stream = DbaasFunctions.getLogsOutputGraylogStream(GetLogsOutputGraylogStreamArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .title(&#34;my stream&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLogsOutputGraylogStreamResult> getLogsOutputGraylogStream(GetLogsOutputGraylogStreamArgs args) {
        return getLogsOutputGraylogStream(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a DBaas logs output graylog stream.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsOutputGraylogStreamArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var stream = DbaasFunctions.getLogsOutputGraylogStream(GetLogsOutputGraylogStreamArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .title(&#34;my stream&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLogsOutputGraylogStreamResult> getLogsOutputGraylogStreamPlain(GetLogsOutputGraylogStreamPlainArgs args) {
        return getLogsOutputGraylogStreamPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a DBaas logs output graylog stream.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsOutputGraylogStreamArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var stream = DbaasFunctions.getLogsOutputGraylogStream(GetLogsOutputGraylogStreamArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .title(&#34;my stream&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLogsOutputGraylogStreamResult> getLogsOutputGraylogStream(GetLogsOutputGraylogStreamArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Dbaas/getLogsOutputGraylogStream:getLogsOutputGraylogStream", TypeShape.of(GetLogsOutputGraylogStreamResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about a DBaas logs output graylog stream.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Dbaas.DbaasFunctions;
     * import com.pulumi.ovh.Dbaas.inputs.GetLogsOutputGraylogStreamArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var stream = DbaasFunctions.getLogsOutputGraylogStream(GetLogsOutputGraylogStreamArgs.builder()
     *             .serviceName(&#34;ldp-xx-xxxxx&#34;)
     *             .title(&#34;my stream&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLogsOutputGraylogStreamResult> getLogsOutputGraylogStreamPlain(GetLogsOutputGraylogStreamPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Dbaas/getLogsOutputGraylogStream:getLogsOutputGraylogStream", TypeShape.of(GetLogsOutputGraylogStreamResult.class), args, Utilities.withVersion(options));
    }
}
