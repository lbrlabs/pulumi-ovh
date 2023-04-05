// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Me.InstallationTemplatePartitionSchemeHardwareRaidArgs;
import com.pulumi.ovh.Me.inputs.InstallationTemplatePartitionSchemeHardwareRaidState;
import com.pulumi.ovh.Utilities;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Use this resource to create a hardware raid group in the partition scheme of a custom installation template available for dedicated servers.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Me.InstallationTemplate;
 * import com.pulumi.ovh.Me.InstallationTemplateArgs;
 * import com.pulumi.ovh.Me.InstallationTemplatePartitionScheme;
 * import com.pulumi.ovh.Me.InstallationTemplatePartitionSchemeArgs;
 * import com.pulumi.ovh.Me.InstallationTemplatePartitionSchemeHardwareRaid;
 * import com.pulumi.ovh.Me.InstallationTemplatePartitionSchemeHardwareRaidArgs;
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
 *         var mytemplate = new InstallationTemplate(&#34;mytemplate&#34;, InstallationTemplateArgs.builder()        
 *             .baseTemplateName(&#34;centos7_64&#34;)
 *             .templateName(&#34;mytemplate&#34;)
 *             .defaultLanguage(&#34;fr&#34;)
 *             .build());
 * 
 *         var scheme = new InstallationTemplatePartitionScheme(&#34;scheme&#34;, InstallationTemplatePartitionSchemeArgs.builder()        
 *             .templateName(mytemplate.templateName())
 *             .priority(1)
 *             .build());
 * 
 *         var group1 = new InstallationTemplatePartitionSchemeHardwareRaid(&#34;group1&#34;, InstallationTemplatePartitionSchemeHardwareRaidArgs.builder()        
 *             .templateName(scheme.templateName())
 *             .schemeName(scheme.name())
 *             .disks(            
 *                 &#34;[c1:d1,c1:d2,c1:d3]&#34;,
 *                 &#34;[c1:d10,c1:d20,c1:d30]&#34;)
 *             .mode(&#34;raid50&#34;)
 *             .step(1)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * The resource can be imported using the `template_name`, `scheme_name`, `name` of the cluster, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid group1 template_name/scheme_name/name
 * ```
 * 
 */
@ResourceType(type="ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid")
public class InstallationTemplatePartitionSchemeHardwareRaid extends com.pulumi.resources.CustomResource {
    /**
     * Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
     * 
     */
    @Export(name="disks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> disks;

    /**
     * @return Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
     * 
     */
    public Output<List<String>> disks() {
        return this.disks;
    }
    /**
     * RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output<String> mode;

    /**
     * @return RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * Hardware RAID name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Hardware RAID name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The partition scheme name.
     * 
     */
    @Export(name="schemeName", refs={String.class}, tree="[0]")
    private Output<String> schemeName;

    /**
     * @return The partition scheme name.
     * 
     */
    public Output<String> schemeName() {
        return this.schemeName;
    }
    /**
     * Specifies the creation order of the hardware RAID.
     * 
     */
    @Export(name="step", refs={Integer.class}, tree="[0]")
    private Output<Integer> step;

    /**
     * @return Specifies the creation order of the hardware RAID.
     * 
     */
    public Output<Integer> step() {
        return this.step;
    }
    /**
     * The template name of the partition scheme.
     * 
     */
    @Export(name="templateName", refs={String.class}, tree="[0]")
    private Output<String> templateName;

    /**
     * @return The template name of the partition scheme.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstallationTemplatePartitionSchemeHardwareRaid(String name) {
        this(name, InstallationTemplatePartitionSchemeHardwareRaidArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstallationTemplatePartitionSchemeHardwareRaid(String name, InstallationTemplatePartitionSchemeHardwareRaidArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstallationTemplatePartitionSchemeHardwareRaid(String name, InstallationTemplatePartitionSchemeHardwareRaidArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid", name, args == null ? InstallationTemplatePartitionSchemeHardwareRaidArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstallationTemplatePartitionSchemeHardwareRaid(String name, Output<String> id, @Nullable InstallationTemplatePartitionSchemeHardwareRaidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static InstallationTemplatePartitionSchemeHardwareRaid get(String name, Output<String> id, @Nullable InstallationTemplatePartitionSchemeHardwareRaidState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstallationTemplatePartitionSchemeHardwareRaid(name, id, state, options);
    }
}
