// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consul

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccConsulServiceDefaultsConfigCEEntryTest(t *testing.T) {
	providers, _ := startTestServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipTestOnConsulEnterpriseEdition(t) },
		Providers: providers,
		Steps: []resource.TestStep{
			{
				Config: testConsulServiceDefaultsConfigEntryWithDestination,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "name", "service-defaults-test-1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "kind", "service-defaults"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "meta.key", "value"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "protocol", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "balance_inbound_connections", "exact_balance"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "mode", "test"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "transparent_proxy.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "transparent_proxy.3186228498.outbound_listener_port", "1001"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "transparent_proxy.3186228498.dialed_directly", "true"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "mutual_tls_mode", "strict"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "envoy_extensions.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "envoy_extensions.0.name", "builtin/aws/lambda"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "envoy_extensions.0.required", "false"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "envoy_extensions.0.arguments.ARN", "some-arn"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "envoy_extensions.0.consul_version", "1.16"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "envoy_extensions.0.envoy_version", "1.2"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "destination.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "destination.2840622507.addresses.0", "10.0.0.1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "destination.2840622507.addresses.1", "10.0.0.2"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "destination.2840622507.port", "1000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "max_inbound_connections", "0"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "local_connect_timeout_ms", "5000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "local_request_timeout_ms", "15"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "mesh_gateway.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "mesh_gateway.2285830552.mode", "mode"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "external_sni", "10.0.0.1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.checks", "false"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.0.listener_port", "0"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.0.local_path_port", "0"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.0.path", "/local/dir"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.0.protocol", "http"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.1.listener_port", "20"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.1.local_path_port", "10"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.1.path", "/test"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.foo", "expose.539725082.paths.1.protocol", "http"),
				),
			},
			{
				Config: testConsulServiceDefaultsConfigEntryWithUpstreamConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "name", "service-defaults-test-2"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "kind", "service-defaults"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "meta.key", "value"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "protocol", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "balance_inbound_connections", "exact_balance"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "mode", "test"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.name", "backend"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.connect_timeout_ms", "500"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.mesh_gateway.3192341522.mode", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.balance_outbound_connections", "exact_balance"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.limits.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.limits.1033039851.max_connections", "1900"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.limits.1033039851.max_pending_requests", "1900"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.limits.1033039851.max_concurrent_requests", "9399"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.passive_health_check.749601092.interval", "19"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.passive_health_check.749601092.max_failures", "8"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.passive_health_check.749601092.enforcing_consecutive_5xx", "10"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.passive_health_check.749601092.max_ejection_percent", "10"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.0.passive_health_check.749601092.base_ejection_time", "30"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.name", "frontend"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.protocol", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.connect_timeout_ms", "5000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.mesh_gateway.3192341522.mode", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.balance_outbound_connections", "exact_balance"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.limits.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.limits.1033039851.max_connections", "1900"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.limits.1033039851.max_pending_requests", "1900"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.limits.1033039851.max_concurrent_requests", "9399"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.passive_health_check.749601092.interval", "19"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.passive_health_check.749601092.max_failures", "8"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.passive_health_check.749601092.enforcing_consecutive_5xx", "10"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.passive_health_check.749601092.max_ejection_percent", "10"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.overrides.1.passive_health_check.749601092.base_ejection_time", "30"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.protocol", "http"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.connect_timeout_ms", "5000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.mesh_gateway.3192341522.mode", "tcp"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.balance_outbound_connections", "exact_balance"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.limits.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.limits.3460439324.max_connections", "1000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.limits.3460439324.max_pending_requests", "9000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.limits.3460439324.max_concurrent_requests", "2900"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.passive_health_check.2313231934.interval", "23900"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.passive_health_check.2313231934.max_failures", "29030"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.passive_health_check.2313231934.enforcing_consecutive_5xx", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.passive_health_check.2313231934.max_ejection_percent", "12"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "upstream_config.1485775592.defaults.2638836468.passive_health_check.2313231934.base_ejection_time", "30309"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "transparent_proxy.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "transparent_proxy.3186228498.outbound_listener_port", "1001"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "transparent_proxy.3186228498.dialed_directly", "true"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "mutual_tls_mode", "strict"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "envoy_extensions.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "envoy_extensions.0.name", "builtin/aws/lambda"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "envoy_extensions.0.required", "false"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "envoy_extensions.0.arguments.ARN", "some-arn"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "envoy_extensions.0.consul_version", "1.16"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "envoy_extensions.0.envoy_version", "1.2"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "max_inbound_connections", "0"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "local_connect_timeout_ms", "5000"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "local_request_timeout_ms", "15"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "mesh_gateway.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "mesh_gateway.3963046091.mode", "strict"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "external_sni", "10.0.0.1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.#", "1"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.checks", "false"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.0.listener_port", "0"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.0.local_path_port", "0"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.0.path", "/local/dir"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.0.protocol", "http"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.1.listener_port", "20"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.1.local_path_port", "10"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.1.path", "/test"),
					resource.TestCheckResourceAttr("consul_config_entry_service_defaults.bar", "expose.539725082.paths.1.protocol", "http"),
				),
			},
		},
	})
}

// Destination and Upstream Config are exclusive.
const testConsulServiceDefaultsConfigEntryWithDestination = `
resource "consul_config_entry_service_defaults" "foo" {
	name      = "service-defaults-test-1"
	meta      = {
		key = "value"
	}
	protocol                    = "tcp"
	balance_inbound_connections = "exact_balance"
	mode                        = "test"
	transparent_proxy {
		outbound_listener_port = 1001
		dialed_directly        = true
	}
	mutual_tls_mode = "strict" # only supported when services are in transparent proxy mode
	envoy_extensions {
		name           = "builtin/aws/lambda"
		required       = false
		arguments      = {
			"ARN" = "some-arn"
		}
		consul_version = "1.16"
		envoy_version  = "1.2"
	}
	destination {
		addresses = [
			"10.0.0.1",
			"10.0.0.2"
		]
		port = 1000
	}
	max_inbound_connections  = 0
	local_connect_timeout_ms = 5000
	local_request_timeout_ms = 15
	mesh_gateway {
		mode = "mode"
	}
	external_sni = "10.0.0.1"

	expose {
		checks = false
		paths {
			path            = "/local/dir"
			local_path_port = 0
			listener_port   = 0
			protocol        = "http"
		}
		paths {
			path            = "/test"
			local_path_port = 10
			listener_port   = 20
			protocol        = "http"
		}
	}
}
`

const testConsulServiceDefaultsConfigEntryWithUpstreamConfig = `
resource "consul_config_entry_service_defaults" "bar" {
	name      = "service-defaults-test-2"
	meta      = {
		key = "value"
	}
	protocol                    = "tcp"
	balance_inbound_connections = "exact_balance"
	mode                        = "test"
	upstream_config {
		overrides {
			name               = "backend"
			protocol           = "tcp"
			connect_timeout_ms = 500
			mesh_gateway {
				mode = "tcp"
			}
			balance_outbound_connections = "exact_balance"
			limits {
				max_connections         = 1900
				max_pending_requests    = 1900
				max_concurrent_requests = 9399
			}
			passive_health_check {
				interval                  = 19
				max_failures              = 8
				enforcing_consecutive_5xx = 10
				max_ejection_percent      = 10
				base_ejection_time        = 30
			}
		}
		overrides {
			name               = "frontend"
			protocol           = "tcp"
			connect_timeout_ms = 5000
			mesh_gateway {
				mode = "tcp"
			}
			balance_outbound_connections = "exact_balance"
			limits {
				max_connections         = 1900
				max_pending_requests    = 1900
				max_concurrent_requests = 9399
			}
			passive_health_check {
				interval                  = 19
				max_failures              = 8
				enforcing_consecutive_5xx = 10
				max_ejection_percent      = 10
				base_ejection_time        = 30
			}
		}
		defaults {
			protocol           = "http"
			connect_timeout_ms = 5000
			mesh_gateway {
				mode = "tcp"
			}
			balance_outbound_connections = "exact_balance"
			limits {
				max_connections         = 1000
				max_pending_requests    = 9000
				max_concurrent_requests = 2900
			}
			passive_health_check {
				interval                  = 23900
				max_failures              = 29030
				enforcing_consecutive_5xx = 1
				max_ejection_percent      = 12
				base_ejection_time        = 30309
			}
		}
	}
	transparent_proxy {
		outbound_listener_port = 1001
		dialed_directly        = true
	}
	mutual_tls_mode = "strict" # only supported when services are in transparent proxy mode
	envoy_extensions {
		name           = "builtin/aws/lambda"
		required       = false
		arguments      = {
			"ARN" = "some-arn"
		}
		consul_version = "1.16"
		envoy_version  = "1.2"
	}
	max_inbound_connections  = 0
	local_connect_timeout_ms = 5000
	local_request_timeout_ms = 15
	mesh_gateway {
		mode = "strict"
	}
	external_sni = "10.0.0.1"

	expose {
		checks = false
		paths {
			path            = "/local/dir"
			local_path_port = 0
			listener_port   = 0
			protocol        = "http"
		}
		paths {
			path            = "/test"
			local_path_port = 10
			listener_port   = 20
			protocol        = "http"
		}
	}
}
`
