package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccRDPApplication_minimal = `
resource "luminate_site" "new-site" {
	name = "tfAccSite"
}
resource "luminate_rdp_application" "new-rdp-application" {
	site_id = "${luminate_site.new-site.id}"
	name = "tfAccRDP"
	internal_address = "127.0.0.2"
}
`
const testAccRDPApplication_changeInternalAddress = `
	resource "luminate_site" "new-site" {
		name = "tfAccSite"
	}
	resource "luminate_rdp_application" "new-rdp-application" {
		site_id = "${luminate_site.new-site.id}"
		name = "tfAccRDP"
		internal_address = "tcp://127.0.0.2:33"
	}
`

const testAccRDPApplication_changeInternalAddress_2 = `
	resource "luminate_site" "new-site" {
		name = "tfAccSite"
	}
	resource "luminate_rdp_application" "new-rdp-application" {
		site_id = "${luminate_site.new-site.id}"
		name = "tfAccRDP"
		internal_address = "tcp://127.0.0.2"
	}
`

const testAccRDPApplication_changeInternalAddress_3 = `
	resource "luminate_site" "new-site" {
		name = "tfAccSite"
	}
	resource "luminate_rdp_application" "new-rdp-application" {
		site_id = "${luminate_site.new-site.id}"
		name = "tfAccRDP"
		internal_address = "tcp://127.0.0.3:3389"
	}
`

const testAccRDPApplication_options = `
resource "luminate_site" "new-site" {
	name = "tfAccSite"
}

resource "luminate_rdp_application" "new-rdp-application" {
	site_id = "${luminate_site.new-site.id}"
	name = "tfAccRDPUpd"
	internal_address = "tcp://127.0.0.5:126"
}
`
const testAccRDPApplication_collection = `
resource "luminate_site" "new-site-collection" {
	name = "tfAccSiteCollection"
}
resource "luminate_collection" "new-collection" {
	name = "tfAccCollectionForApp"
}
resource "luminate_collection_site_link" "new-collection-site-link" {
	site_id = "${luminate_site.new-site-collection.id}"
	collection_ids = sort(["${luminate_collection.new-collection.id}"])
}
resource "luminate_rdp_application" "new-rdp-application-collection" {
	site_id = "${luminate_site.new-site-collection.id}"
	collection_id = "${luminate_collection.new-collection.id}"
	name = "tfAccRDPWithCollection"
	internal_address = "tcp://127.0.0.2"
    depends_on = [luminate_collection_site_link.new-collection-site-link]
}
`

func TestAccLuminateRDPApplication(t *testing.T) {
	resourceName := "luminate_rdp_application.new-rdp-application"
	resourceNameCollection := "luminate_rdp_application.new-rdp-application-collection"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newTestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccRDPApplication_minimal,
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "tfAccRDP"),
					resource.TestCheckResourceAttr(resourceName, "visible", "true"),
					resource.TestCheckResourceAttr(resourceName, "notification_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "internal_address", "tcp://127.0.0.2:3389"),
					resource.TestCheckResourceAttr(resourceName, "external_address", fmt.Sprintf("tfaccrdp.rdp.%s", testAccDomain)),
					resource.TestCheckResourceAttr(resourceName, "luminate_address", fmt.Sprintf("tfaccrdp.rdp.%s", testAccDomain)),
				),
			},
			{
				Config:  testAccRDPApplication_changeInternalAddress,
				Destroy: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "internal_address", "tcp://127.0.0.2:33"),
				),
			},
			{
				Config:  testAccRDPApplication_changeInternalAddress_2,
				Destroy: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "internal_address", "tcp://127.0.0.2:3389"),
				),
			},
			{
				Config:  testAccRDPApplication_changeInternalAddress_3,
				Destroy: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "internal_address", "tcp://127.0.0.3:3389"),
				),
			},
			{
				Config: testAccRDPApplication_options,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "tfAccRDPUpd"),
					resource.TestCheckResourceAttr(resourceName, "internal_address", "tcp://127.0.0.5:126"),
					resource.TestCheckResourceAttr(resourceName, "external_address", fmt.Sprintf("tfaccrdp.rdp.%s", testAccDomain)),
					resource.TestCheckResourceAttr(resourceName, "luminate_address", fmt.Sprintf("tfaccrdp.rdp.%s", testAccDomain)),
				),
			},
			{
				Config: testAccRDPApplication_collection,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceNameCollection, "name", "tfAccRDPWithCollection")),
			},
		},
	})
}
