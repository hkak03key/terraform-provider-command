---
page_title: "command Provider"
subcategory: ""
description: |-
  
---

# command Provider


## Example Usage
```terraform
terraform {
  required_providers {
    command = {
      source = "hkak03key/command"
    }
  }
}

provider "command" {
}
```

## Why not external provider?
This provider is similar to external provider.

External provider requires json output so we should use jq (= additional dependency) or create json manually (= it is very bother).  
In this provider, we can get stdout as string.

So we can use shell command such as `find` directly.  
This is very simple.

