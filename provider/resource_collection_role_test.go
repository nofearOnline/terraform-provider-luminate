package provider

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const testCollectionRole = `
	resource "luminate_collection" "collection" {
		name = "collectionToBeAssign"
	} 
	resource "luminate_collection_role" "collection-role" {
		role_type = "PolicyOwner"
		identity_provider_id =  "local"
		entity_id = "a8a48219-835f-4183-a2a9-bbba8cad8eb8"
		entity_type = "User"
		collection_id = "${luminate_collection.collection.id}"
	}
`

func TestAccLuminateCollectionRole(t *testing.T) {
	resourceName := "luminate_collection_role.collection-role"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testCollectionRole,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "role_type", "PolicyOwner")),
			},
		},
	})
}
