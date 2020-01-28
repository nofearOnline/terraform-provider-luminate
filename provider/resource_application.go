package provider

import (
	"bitbucket.org/accezz-io/terraform-provider-symcsc/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func CommonApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "Site ID to which the application will be bound",
			ValidateFunc: utils.ValidateUuid,
		},
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "Name of the application",
			ValidateFunc: utils.ValidateApplicationName,
		},
		"icon": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Base64 representation of 40x40 icon",
			ValidateFunc: utils.ValidateString,
		},
		"visible": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      true,
			ValidateFunc: utils.ValidateBool,
			Description:  "Indicates whether to show this application in the applications portal.",
		},
		"notification_enabled": {
			Type:         schema.TypeBool,
			Optional:     true,
			Default:      true,
			ValidateFunc: utils.ValidateBool,
			Description:  "Indicates whether notifications are enabled for this application.",
		},
		"subdomain": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The application DNS subdomain.",
		},
		"custom_external_address": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: utils.ValidateString,
			Description:  "The application custom DNS address that exposes the application.",
		},
		"external_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"luminate_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
