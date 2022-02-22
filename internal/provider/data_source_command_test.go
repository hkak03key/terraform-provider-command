package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

//-----------------------------------------
const testAccDataSourceCommand_basic = `
data "command" "test" {
  command = ["echo", "Hello", "world."]
}
`

func TestAccDataSourceCommand_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCommand_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(
						"data.command.test", "stdin"),
					resource.TestCheckResourceAttr(
						"data.command.test", "trim_space", "true"),
					resource.TestCheckResourceAttr(
						"data.command.test", "stdout", "Hello world."),
				),
			},
		},
	})
}

//-----------------------------------------
const testAccDataSourceCommand_multiple = `
data "command" "test1" {
  command = ["echo", "1"]
}

data "command" "test2" {
  command = ["echo", "2"]
}
`

func TestAccDataSourceCommand_multiple(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCommand_multiple,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.command.test1", "stdout", "1"),
					resource.TestCheckResourceAttr(
						"data.command.test2", "stdout", "2"),
				),
			},
		},
	})
}

//-----------------------------------------
const testAccDataSourceCommand_stdin = `
data "command" "test" {
  command = ["sh", "-c", "echo $(tee)"]
  stdin   = "Test for stdin."
}
`

func TestAccDataSourceCommand_stdin(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCommand_stdin,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.command.test", "stdout", "Test for stdin."),
				),
			},
		},
	})
}

//-----------------------------------------
const testAccDataSourceCommand_unknownCommand = `
data "command" "test" {
  command = ["this-is-unknown-command"]
}
`

func TestAccDataSourceCommand_unknownCommand(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceCommand_unknownCommand,
				ExpectError: regexp.MustCompile("."),
			},
		},
	})
}

//-----------------------------------------
const testAccDataSourceCommand_trimSpace = `
data "command" "test" {
  command    = ["echo", "\n	Test for trim_space.\r \t"]
  trim_space = %s
}
`

func TestAccDataSourceCommand_trimSpace(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceCommand_trimSpace, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.command.test", "stdout", "Test for trim_space."),
				),
			},
			{
				Config: fmt.Sprintf(testAccDataSourceCommand_trimSpace, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.command.test", "stdout", regexp.MustCompile("^\\s+Test for trim_space\\.\\s+$")),
				),
			},
		},
	})
}
