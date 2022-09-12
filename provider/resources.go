// Copyright 2016-2022, lbrlabs.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ovh

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/lbrlabs/pulumi-ovh/provider/pkg/version"
	ovh "github.com/ovh/terraform-provider-ovh/ovh"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	ovhPkg = "ovh"
	// modules:
	ovhMod = "index" // the y module

)

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// ovhMember manufactures a type token for the Ovh package and the given module and type.
func ovhMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(ovhPkg + ":" + mod + ":" + mem)
}

// ovhType manufactures a type token for the Ovh package and the given module and type.
func ovhType(mod string, typ string) tokens.Type {
	return tokens.Type(ovhMember(mod, typ))
}

// ovhDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Ovh package and names the file by simply lower casing the data
// source's first character.
func ovhDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ovhMember(mod+"/"+fn, res)
}

// ovhResource manufactures a standard resource token given a module and resource name.
// It automatically uses the ovh package and names the file by simply lower casing the resource's
// first character.
func ovhResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ovhType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(ovh.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "ovh",
		DisplayName:       "OVH",
		Publisher:         "lbrlabs",
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		Description:       "A Pulumi package for creating and managing OVH cloud resources.",
		Keywords:          []string{"pulumi", "ovh", "category/cloud"},
		License:           "Apache-2.0",
		LogoURL:           "https://raw.githubusercontent.com/lbrlabs/pulumi-ovh/main/assets/ovh.svg",
		Homepage:          "https://www.pulumi.com",
		Repository:        "https://github.com/lbrlabs/pulumi-ovh",
		GitHubOrg:         "ovh",
		Config: map[string]*tfbridge.SchemaInfo{
			"endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_ENDPOINT"},
				},
			},
			"application_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_APPLICATION_KEY"},
				},
			},
			"application_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_APPLICATION_SECRET"},
				},
			},
			"consumer_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_CONSUMER_KEY"},
				},
				Secret: boolRef(true),
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"ovh_cloud_project": {Tok: ovhResource(ovhMod, "CloudProject")},
			"ovh_cloud_project_containerregistry": {
				Tok: ovhResource(ovhMod, "CloudProjectContainerRegistry"),
			},
			"ovh_cloud_project_containerregistry_user": {
				Tok: ovhResource(ovhMod, "CloudProjectContainerRegistryUser"),
			},
			"ovh_cloud_project_database": {
				Tok: ovhResource(ovhMod, "CloudProjectDatabase"),
			},
			"ovh_cloud_project_database_ip_restriction": {
				Tok: ovhResource(ovhMod, "CloudProjectDatabaseIpRestriction"),
			},
			"ovh_cloud_project_database_postgresql_user": {
				Tok: ovhResource(ovhMod, "CloudProjectDatabasePostgresSqlUser"),
			},
			"ovh_cloud_project_database_user": {
				Tok: ovhResource(ovhMod, "CloudProjectDatabaseUser"),
			},
			"ovh_cloud_project_failover_ip_attach": {
				Tok: ovhResource(ovhMod, "CloudProjectFailoverIpAttach"),
			},
			"ovh_cloud_project_kube": {
				Tok: ovhResource(ovhMod, "CloudProjectKube"),
			},
			"ovh_cloud_project_kube_iprestrictions": {
				Tok: ovhResource(ovhMod, "CloudProjectKubeIpRestrictions"),
			},
			"ovh_cloud_project_kube_nodepool": {
				Tok: ovhResource(ovhMod, "CloudProjectKubeNodePool"),
			},
			"ovh_cloud_project_kube_oidc": {
				Tok: ovhResource(ovhMod, "CloudProjectKubeOidc"),
			},
			"ovh_cloud_project_network_private": {
				Tok: ovhResource(ovhMod, "CloudProjectNetworkPrivate"),
			},
			"ovh_cloud_project_network_private_subnet": {
				Tok: ovhResource(ovhMod, "CloudProjectNetworkPrivateSubnet"),
			},
			"ovh_cloud_project_user": {
				Tok: ovhResource(ovhMod, "CloudProjectUser"),
			},
			"ovh_dbaas_logs_input": {
				Tok: ovhResource(ovhMod, "DbaasLogsInput"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhResource(ovhMod, "DbaasLogsOutputGraylogStream"),
			},
			"ovh_dedicated_ceph_acl": {
				Tok: ovhResource(ovhMod, "DedicatedCephAcl"),
			},
			"ovh_dedicated_server_install_task": {
				Tok: ovhResource(ovhMod, "DedicatedServiceInstallTask"),
			},
			"ovh_dedicated_server_reboot_task": {
				Tok: ovhResource(ovhMod, "DedicatedServerRebootTask"),
			},
			"ovh_dedicated_server_update": {
				Tok: ovhResource(ovhMod, "DedicatedServerUpdate"),
			},
			"ovh_domain_zone": {
				Tok: ovhResource(ovhMod, "DomainZone"),
			},
			"ovh_domain_zone_record": {
				Tok: ovhResource(ovhMod, "DomainZoneRecord"),
			},
			"ovh_domain_zone_redirection": {
				Tok: ovhResource(ovhMod, "DomainZoneRedirection"),
			},
			"ovh_ip_reverse": {
				Tok: ovhResource(ovhMod, "IpReverse"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ip_reverse": {
						CSharpName: "IpReverseAddress",
					},
				},
			},
			"ovh_ip_service": {
				Tok: ovhResource(ovhMod, "IpService"),
			},
			"ovh_iploadbalancing": {
				Tok: ovhResource(ovhMod, "IpLoadBalancing"),
			},
			"ovh_iploadbalancing_http_farm": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingHttpFarm"),
			},
			"ovh_iploadbalancing_http_farm_server": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingHttpFarmServer"),
			},
			"ovh_iploadbalancing_http_frontend": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingHttpFrontend"),
			},
			"ovh_iploadbalancing_http_route": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingHttpRoute"),
			},
			"ovh_iploadbalancing_http_route_rule": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingHttpRouteRule"),
			},
			"ovh_iploadbalancing_refresh": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingRefresh"),
			},
			"ovh_iploadbalancing_tcp_farm": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingTcpFarm"),
			},
			"ovh_iploadbalancing_tcp_farm_server": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingTcpFarmServer"),
			},
			"ovh_iploadbalancing_tcp_frontend": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingTcpFrontend"),
			},
			"ovh_iploadbalancing_tcp_route": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingTcpRoute"),
			},
			"ovh_iploadbalancing_tcp_route_rule": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingTcpRouteRule"),
			},
			"ovh_iploadbalancing_vrack_network": {
				Tok: ovhResource(ovhMod, "IpLoadBalancingVrackNetwork"),
			},
			"ovh_me_identity_user": {
				Tok: ovhResource(ovhMod, "MeIdentityUser"),
			},
			"ovh_me_installation_template_partition_scheme": {
				Tok: ovhResource(ovhMod, "MeInstallationTemplatePartitionScheme"),
			},
			"ovh_me_installation_template_partition_scheme_hardware_raid": {
				Tok: ovhResource(ovhMod, "MeInstallationTemplatePartitionSchemeHardwareRaid"),
			},
			"ovh_me_installation_template_partition_scheme_partition": {
				Tok: ovhResource(ovhMod, "MeInstallationTemplatePartitionSchemePartition"),
			},
			"ovh_me_ipxe_script": {
				Tok: ovhResource(ovhMod, "MeIpxeScript"),
			},
			"ovh_me_ssh_key": {
				Tok: ovhResource(ovhMod, "MeSshKey"),
			},
			"ovh_vrack": {
				Tok: ovhResource(ovhMod, "Vrack"),
			},
			"ovh_vrack_cloudproject": {
				Tok: ovhResource(ovhMod, "VrackCloudProject"),
			},
			"ovh_vrack_dedicated_server": {
				Tok: ovhResource(ovhMod, "VrackDedicatedServer"),
			},
			"ovh_vrack_dedicated_server_interface": {
				Tok: ovhResource(ovhMod, "VrackDedicatedServerInterface"),
			},
			"ovh_vrack_ip": {
				Tok: ovhResource(ovhMod, "VrackIp"),
			},
			"ovh_vrack_iploadbalancing": {
				Tok: ovhResource(ovhMod, "VrackIpLoadbalancing"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"ovh_cloud_project_capabilities_containerregistry": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectCapabilitiesContainerRegistry"),
			},
			"ovh_cloud_project_capabilities_containerregistry_filter": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectCapabilitiesContainerFilter"),
			},
			"ovh_cloud_project_containerregistries": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectContainerRegistries"),
			},
			"ovh_cloud_project_containerregistry": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectContainerRegistry"),
			},
			"ovh_cloud_project_containerregistry_users": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectContainerRegistryUsers"),
			},
			"ovh_cloud_project_database": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectDatabase"),
			},
			"ovh_cloud_project_database_ip_restrictions": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectDatabaseIpRestrictions"),
			},
			"ovh_cloud_project_database_postgresql_user": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectDatabasePostgresSqlUser"),
			},
			"ovh_cloud_project_database_user": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectDatabaseUser"),
			},
			"ovh_cloud_project_database_users": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectDatabaseUsers"),
			},
			"ovh_cloud_project_databases": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectDatabases"),
			},
			"ovh_cloud_project_failover_ip_attach": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectFailoverIpAttach"),
			},
			"ovh_cloud_project_kube": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectKube"),
			},
			"ovh_cloud_project_kube_iprestrictions": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectKubeIpRestrictions"),
			},
			"ovh_cloud_project_kube_nodepool": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectKubeIpNodePool"),
			},
			"ovh_cloud_project_region": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectRegion"),
			},
			"ovh_cloud_project_regions": {
				Tok: ovhDataSource(ovhMod, "getCloudProjectRegions"),
			},
			"ovh_dbaas_logs_input_engine": {
				Tok: ovhDataSource(ovhMod, "getDbaasLogsInputEngine"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhDataSource(ovhMod, "getDbaasLogsOutputGraylogStream"),
			},
			"ovh_dedicated_ceph": {
				Tok: ovhDataSource(ovhMod, "getDedicatedCeph"),
			},
			"ovh_dedicated_installation_templates": {
				Tok: ovhDataSource(ovhMod, "getDedicatedInstallationTemplates"),
			},
			"ovh_dedicated_server": {
				Tok: ovhDataSource(ovhMod, "getDedicatedServer"),
			},
			"ovh_dedicated_servers": {
				Tok: ovhDataSource(ovhMod, "getDedicatedServers"),
			},
			"ovh_domain_zone": {
				Tok: ovhDataSource(ovhMod, "getDomainZone"),
			},
			"ovh_ip_service": {
				Tok: ovhDataSource(ovhMod, "getIpService"),
			},
			"ovh_iploadbalancing_vrack_network": {
				Tok: ovhDataSource(ovhMod, "getIpLoadbalancingVrackNetwork"),
			},
			"ovh_iploadbalancing_vrack_networks": {
				Tok: ovhDataSource(ovhMod, "getIpLoadbalancingVrackNetworks"),
			},
			"ovh_me": {
				Tok: ovhDataSource(ovhMod, "getMe"),
			},
			"ovh_me_identity_user": {
				Tok: ovhDataSource(ovhMod, "getMeIdentityUser"),
			},
			"ovh_me_identity_users": {
				Tok: ovhDataSource(ovhMod, "getMeIdentityUsers"),
			},
			"ovh_me_installation_template": {
				Tok: ovhDataSource(ovhMod, "getMeInstallationTemplate"),
			},
			"ovh_me_installation_templates": {
				Tok: ovhDataSource(ovhMod, "getMeInstallationTemplates"),
			},
			"ovh_me_ipxe_script": {
				Tok: ovhDataSource(ovhMod, "getMeIpxeScript"),
			},
			"ovh_me_ipxe_scripts": {
				Tok: ovhDataSource(ovhMod, "getMeIpxeScripts"),
			},
			"ovh_me_paymentmean_bankaccount": {
				Tok: ovhDataSource(ovhMod, "getMePaymentmeanBankAccount"),
			},
			"ovh_me_paymentmean_creditcard": {
				Tok: ovhDataSource(ovhMod, "getMePaymentmeanCreditCard"),
			},
			"ovh_me_ssh_key": {
				Tok: ovhDataSource(ovhMod, "getMeSshKey"),
			},
			"ovh_order_cart": {
				Tok: ovhDataSource(ovhMod, "getOrderCart"),
			},
			"ovh_order_cart_product": {
				Tok: ovhDataSource(ovhMod, "getOrderCartProduct"),
			},
			"ovh_order_cart_product_options": {
				Tok: ovhDataSource(ovhMod, "getOrderCartProductOptions"),
			},
			"ovh_order_cart_product_options_plan": {
				Tok: ovhDataSource(ovhMod, "getOrderCartProductOptionsPlan"),
			},
			"ovh_order_cart_product_plan": {
				Tok: ovhDataSource(ovhMod, "getOrderCartProductPlan"),
			},
			"ovh_vps": {
				Tok: ovhDataSource(ovhMod, "getVps"),
			},
			"ovh_vracks": {
				Tok: ovhDataSource(ovhMod, "getVracks"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@lbrlabs/pulumi-ovh",
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "lbrlabs_pulumi_ovh",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", ovhPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				ovhPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Lbrlabs.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
