// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Vrack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Utilities;
import com.pulumi.ovh.Vrack.IpAddressArgs;
import com.pulumi.ovh.Vrack.inputs.IpAddressState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Attach an IP block to a VRack.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Order.OrderFunctions;
 * import com.pulumi.ovh.Order.inputs.GetCartArgs;
 * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
 * import com.pulumi.ovh.Vrack.Vrack;
 * import com.pulumi.ovh.Vrack.VrackArgs;
 * import com.pulumi.ovh.Vrack.inputs.VrackPlanArgs;
 * import com.pulumi.ovh.Ip.IpService;
 * import com.pulumi.ovh.Ip.IpServiceArgs;
 * import com.pulumi.ovh.Ip.inputs.IpServicePlanArgs;
 * import com.pulumi.ovh.Vrack.IpAddress;
 * import com.pulumi.ovh.Vrack.IpAddressArgs;
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
 *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
 *             .ovhSubsidiary(&#34;fr&#34;)
 *             .description(&#34;my cart&#34;)
 *             .build());
 * 
 *         final var vrackCartProductPlan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
 *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
 *             .priceCapacity(&#34;renew&#34;)
 *             .product(&#34;vrack&#34;)
 *             .planCode(&#34;vrack&#34;)
 *             .build());
 * 
 *         var vrackVrack = new Vrack(&#34;vrackVrack&#34;, VrackArgs.builder()        
 *             .description(mycart.applyValue(getCartResult -&gt; getCartResult.description()))
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -&gt; getCartResult.ovhSubsidiary()))
 *             .plan(VrackPlanArgs.builder()
 *                 .duration(vrackCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(vrackCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.planCode()))
 *                 .pricingMode(vrackCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .build())
 *             .build());
 * 
 *         final var ipblockCartProductPlan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
 *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
 *             .priceCapacity(&#34;renew&#34;)
 *             .product(&#34;ip&#34;)
 *             .planCode(&#34;ip-v4-s30-ripe&#34;)
 *             .build());
 * 
 *         var ipblockIpService = new IpService(&#34;ipblockIpService&#34;, IpServiceArgs.builder()        
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -&gt; getCartResult.ovhSubsidiary()))
 *             .description(mycart.applyValue(getCartResult -&gt; getCartResult.description()))
 *             .plan(IpServicePlanArgs.builder()
 *                 .duration(ipblockCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(ipblockCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.planCode()))
 *                 .pricingMode(ipblockCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .configurations(IpServicePlanConfigurationArgs.builder()
 *                     .label(&#34;country&#34;)
 *                     .value(&#34;FR&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var vrackblock = new IpAddress(&#34;vrackblock&#34;, IpAddressArgs.builder()        
 *             .serviceName(vrackVrack.serviceName())
 *             .block(ipblockIpService.ip())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:Vrack/ipAddress:IpAddress")
public class IpAddress extends com.pulumi.resources.CustomResource {
    /**
     * Your IP block.
     * 
     */
    @Export(name="block", refs={String.class}, tree="[0]")
    private Output<String> block;

    /**
     * @return Your IP block.
     * 
     */
    public Output<String> block() {
        return this.block;
    }
    /**
     * Your gateway
     * 
     */
    @Export(name="gateway", refs={String.class}, tree="[0]")
    private Output<String> gateway;

    /**
     * @return Your gateway
     * 
     */
    public Output<String> gateway() {
        return this.gateway;
    }
    /**
     * Your IP block
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return Your IP block
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * The internal name of your vrack
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your vrack
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Where you want your block announced on the network
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return Where you want your block announced on the network
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpAddress(String name) {
        this(name, IpAddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpAddress(String name, IpAddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpAddress(String name, IpAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/ipAddress:IpAddress", name, args == null ? IpAddressArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpAddress(String name, Output<String> id, @Nullable IpAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/ipAddress:IpAddress", name, state, makeResourceOptions(options, id));
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
    public static IpAddress get(String name, Output<String> id, @Nullable IpAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpAddress(name, id, state, options);
    }
}
