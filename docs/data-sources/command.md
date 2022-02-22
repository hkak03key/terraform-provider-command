---
page_title: "command Data Source - terraform-provider-command"
subcategory: ""
description: |-
  Execute shell command.
---
# command (Data Source)
Execute shell command.

## Example Usage
```terraform
data "command" "example" {
  command = ["echo", "Hello", "world."]
}

output "example" {
  value = data.command.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **command** (List of String) execute command and args.
NOTE: This will 2nd or later args of exec.CommandContext().

### Optional

- **stdin** (String) stdin.
- **trim_space** (Boolean) remove stdout beginning/ending whitespaces.

### Read-Only

- **stdout** (String) stdout.
- **id** (String) dummy.  <!-- edit manually -->
