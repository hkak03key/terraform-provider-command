package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"bytes"
	"os/exec"
	"strings"
)

func dataSourceCommand() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Execute shell command.",

		ReadContext: dataSourceCommandRead,

		Schema: map[string]*schema.Schema{
			"command": {
				Description: "execute command and args.  \nNOTE: These will be 2nd or later args of exec.CommandContext().",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"stdin": {
				Description: "stdin.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"trim_space": {
				Description: "remove stdout beginning/ending whitespaces.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"stdout": {
				Description: "stdout.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceCommandRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	var command []string
	if v, ok := d.GetOk("command"); ok {
		commandRaw := v.([]interface{})
		command = make([]string, len(commandRaw))
		for i, v := range commandRaw {
			command[i] = v.(string)
		}
	}

	stdin := d.Get("stdin").(string)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return diag.FromErr(err)
	}

	if d.Get("trim_space").(bool) {
		d.Set("stdout", strings.TrimSpace(stdout.String()))
	} else {
		d.Set("stdout", stdout.String())
	}

	d.SetId("dummy")

	return nil
}
