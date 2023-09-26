// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consul

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccConsulServiceSplitterConfigCEEntryTest(t *testing.T) {
	providers, _ := startTestServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipTestOnConsulEnterpriseEdition(t) },
		Providers: providers,
		Steps: []resource.TestStep{
			{
				Config: testConsulServiceSplitterConfigEntryCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "name", "web"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "meta.key", "value"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.#", "2"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.0.weight", "90"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.0.service", "web"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.0.service_subset", "v1"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.1.weight", "10"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.1.service_subset", "v2"),
					resource.TestCheckResourceAttr("consul_service_splitter_config_entry.foo", "splits.1.service", "web"),
				),
			},
		},
	})
}

const testConsulServiceSplitterConfigEntryCE = `
resource "consul_config_entry" "web" {
	name = "web"
	kind = "service-defaults"

	config_json = jsonencode({
		Protocol         = "http"
		Expose           = {}
		MeshGateway      = {}
		TransparentProxy = {}
	})
}

resource "consul_config_entry" "service_resolver" {
	kind = "service-resolver"
	name = consul_config_entry.web.name

	config_json = jsonencode({
		DefaultSubset = "v1"

		Subsets = {
			"v1" = {
				Filter = "Service.Meta.version == v1"
			}
			"v2" = {
				Filter = "Service.Meta.version == v2"
			}
		}
	})
}

resource "consul_service_splitter_config_entry" "foo" {
	name = consul_config_entry.service_resolver.name
	meta = {
		key = "value"
	}
	splits {
		weight  = 90                   
		service_subset  = "v1"                
		service = "web"
	}
	splits {
		weight  = 10
		service = "web"
		service_subset  = "v2"
	}
}
`
