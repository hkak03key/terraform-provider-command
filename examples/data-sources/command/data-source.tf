data "command" "example" {
  command = ["echo", "Hello", "world."]
}

output "example" {
  value = data.command.example
}
