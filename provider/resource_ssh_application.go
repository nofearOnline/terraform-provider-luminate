package provider

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"

	"github.com/Broadcom/terraform-provider-luminate/service"
	"github.com/Broadcom/terraform-provider-luminate/service/dto"
	"github.com/Broadcom/terraform-provider-luminate/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func LuminateSSHApplication() *schema.Resource {
	sshSchema := CommonApplicationSchema()

	sshSchema["internal_address"] = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: utils.ValidateString,
		Description:  "Internal address of the application, accessible by connector",
	}

	return &schema.Resource{
		Schema:        sshSchema,
		CreateContext: resourceCreateSSHApplication,
		ReadContext:   resourceReadSSHApplication,
		UpdateContext: resourceUpdateSSHApplication,
		DeleteContext: resourceDeleteSSHApplication,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateSSHApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] LUMINATE CREATE APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return diag.FromErr(errors.New("unable to cast Luminate service"))
	}

	newApp := extractSSHApplicationFields(d)

	app, err := client.Applications.CreateApplication(newApp)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.Applications.BindApplicationToSite(app, newApp.SiteID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(app.ID)
	setSHHApplicationFields(d, app)

	return resourceReadSSHApplication(ctx, d, m)
}

func resourceReadSSHApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] LUMINATE READ APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return diag.FromErr(errors.New("unable to cast Luminate service"))
	}

	app, err := client.Applications.GetApplicationById(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if app == nil {
		d.SetId("")
		return nil
	}

	d.SetId(app.ID)

	app.SiteID = d.Get("site_id").(string)
	setSHHApplicationFields(d, app)

	return nil
}

func resourceUpdateSSHApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] LUMINATE UPDATE APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return diag.FromErr(errors.New("unable to cast Luminate service"))
	}

	app := extractSSHApplicationFields(d)

	updApp, err := client.Applications.UpdateApplication(app)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.Applications.BindApplicationToSite(app, app.SiteID)
	if err != nil {
		return diag.FromErr(err)
	}

	updApp.SiteID = app.SiteID
	setSHHApplicationFields(d, updApp)

	return resourceReadSSHApplication(ctx, d, m)
}

func resourceDeleteSSHApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] LUMINATE DELETE APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return diag.FromErr(errors.New("unable to cast Luminate service"))
	}

	err := client.Applications.DeleteApplication(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceReadSSHApplication(ctx, d, m)
}

func setSHHApplicationFields(d *schema.ResourceData, application *dto.Application) {
	d.Set("name", application.Name)
	d.Set("collection_id", application.CollectionID)
	d.Set("icon", application.Icon)
	d.Set("type", application.Type)
	d.Set("visible", application.Visible)
	d.Set("site_id", application.SiteID)
	d.Set("notification_enabled", application.NotificationsEnabled)
	d.Set("internal_address", application.InternalAddress)
	d.Set("external_address", application.ExternalAddress)
	d.Set("subdomain", application.Subdomain)
	d.Set("luminate_address", application.LuminateAddress)
}

func extractSSHApplicationFields(d *schema.ResourceData) *dto.Application {
	return &dto.Application{
		ID:                   d.Id(),
		CollectionID:         d.Get("collection_id").(string),
		Name:                 d.Get("name").(string),
		Icon:                 d.Get("icon").(string),
		SiteID:               d.Get("site_id").(string),
		Type:                 "ssh",
		Visible:              d.Get("visible").(bool),
		NotificationsEnabled: d.Get("notification_enabled").(bool),
		InternalAddress:      d.Get("internal_address").(string),
		ExternalAddress:      d.Get("external_address").(string),
		Subdomain:            d.Get("subdomain").(string),
	}
}
