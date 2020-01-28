package provider

import (
	"bitbucket.org/accezz-io/terraform-provider-symcsc/service"
	"bitbucket.org/accezz-io/terraform-provider-symcsc/service/dto"
	"bitbucket.org/accezz-io/terraform-provider-symcsc/utils"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
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
		Schema: sshSchema,
		Create: resourceCreateSSHApplication,
		Read:   resourceReadSSHApplication,
		Update: resourceUpdateSSHApplication,
		Delete: resourceDeleteSSHApplication,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateSSHApplication(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] LUMINATE CREATE APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	newApp := extractSSHApplicationFields(d)

	app, err := client.Applications.CreateApplication(newApp)
	if err != nil {
		return err
	}

	err = client.Applications.BindApplicationToSite(app, newApp.SiteID)
	if err != nil {
		return err
	}

	d.SetId(app.ID)
	setSHHApplicationFields(d, app)

	return resourceReadSSHApplication(d, m)
}

func resourceReadSSHApplication(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] LUMINATE READ APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	app, err := client.Applications.GetApplicationById(d.Id())
	if err != nil {
		return err
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

func resourceUpdateSSHApplication(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] LUMINATE UPDATE APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	app := extractSSHApplicationFields(d)

	updApp, err := client.Applications.UpdateApplication(app)
	if err != nil {
		return err
	}

	err = client.Applications.BindApplicationToSite(app, app.SiteID)
	if err != nil {
		return err
	}

	updApp.SiteID = app.SiteID
	setSHHApplicationFields(d, updApp)

	return resourceReadSSHApplication(d, m)
}

func resourceDeleteSSHApplication(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] LUMINATE DELETE APP")

	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	err := client.Applications.DeleteApplication(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")

	return resourceReadSSHApplication(d, m)
}

func setSHHApplicationFields(d *schema.ResourceData, application *dto.Application) {
	d.Set("name", application.Name)
	d.Set("icon", application.Icon)
	d.Set("type", application.Type)
	d.Set("visible", application.Visible)
	d.Set("site_id", application.SiteID)
	d.Set("notification_enabled", application.NotificationsEnabled)
	d.Set("internal_address", application.InternalAddress)
	d.Set("external_address", application.ExternalAddress)
	d.Set("subdomain", application.Subdomain)
	d.Set("custom_external_address", application.CustomExternalAddress)
	d.Set("luminate_address", application.LuminateAddress)
}

func extractSSHApplicationFields(d *schema.ResourceData) *dto.Application {
	return &dto.Application{
		ID:                    d.Id(),
		Name:                  d.Get("name").(string),
		Icon:                  d.Get("icon").(string),
		SiteID:                d.Get("site_id").(string),
		Type:                  "ssh",
		Visible:               d.Get("visible").(bool),
		NotificationsEnabled:  d.Get("notification_enabled").(bool),
		InternalAddress:       d.Get("internal_address").(string),
		ExternalAddress:       d.Get("external_address").(string),
		Subdomain:             d.Get("subdomain").(string),
		CustomExternalAddress: d.Get("custom_external_address").(string),
	}
}
