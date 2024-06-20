
output "name" {
  value       = resource.azurerm_resource_group.rsg.name
  description = "Resource group name"
}

output "id" {
  value       = resource.azurerm_resource_group.rsg.id
  description = "Resource groupd id"
}

output "raw" {
  value       = resource.azurerm_resource_group.rsg
  description = "Resource group raw response"
}