// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Vrack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Utilities;
import com.pulumi.ovh.Vrack.VrackArgs;
import com.pulumi.ovh.Vrack.inputs.VrackState;
import com.pulumi.ovh.Vrack.outputs.VrackOrder;
import com.pulumi.ovh.Vrack.outputs.VrackPlan;
import com.pulumi.ovh.Vrack.outputs.VrackPlanOption;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 *             .description(&#34;my vrack order cart&#34;)
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
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -&gt; getCartResult.ovhSubsidiary()))
 *             .description(&#34;my vrack&#34;)
 *             .plan(VrackPlanArgs.builder()
 *                 .duration(vrackCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(vrackCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.planCode()))
 *                 .pricingMode(vrackCartProductPlan.applyValue(getCartProductPlanResult -&gt; getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:Vrack/vrack:Vrack")
public class Vrack extends com.pulumi.resources.CustomResource {
    /**
     * yourvrackdescription
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return yourvrackdescription
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * yourvrackname
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return yourvrackname
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Details about an Order
     * 
     */
    @Export(name="orders", refs={List.class,VrackOrder.class}, tree="[0,1]")
    private Output<List<VrackOrder>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Output<List<VrackOrder>> orders() {
        return this.orders;
    }
    /**
     * OVHcloud Subsidiary
     * 
     */
    @Export(name="ovhSubsidiary", refs={String.class}, tree="[0]")
    private Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    @Export(name="paymentMean", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentMean;

    /**
     * @return Ovh payment mode
     * 
     */
    public Output<Optional<String>> paymentMean() {
        return Codegen.optional(this.paymentMean);
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="plan", refs={VrackPlan.class}, tree="[0]")
    private Output<VrackPlan> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<VrackPlan> plan() {
        return this.plan;
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="planOptions", refs={List.class,VrackPlanOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<VrackPlanOption>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<Optional<List<VrackPlanOption>>> planOptions() {
        return Codegen.optional(this.planOptions);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vrack(String name) {
        this(name, VrackArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vrack(String name, VrackArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vrack(String name, VrackArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/vrack:Vrack", name, args == null ? VrackArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vrack(String name, Output<String> id, @Nullable VrackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/vrack:Vrack", name, state, makeResourceOptions(options, id));
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
    public static Vrack get(String name, Output<String> id, @Nullable VrackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vrack(name, id, state, options);
    }
}
