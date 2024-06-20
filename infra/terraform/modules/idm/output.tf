
output "name" {
  value       = resource.azurerm_user_assigned_identity.idm.name
  description = "Identity name"
}

output "id" {
  value       = resource.azurerm_user_assigned_identity.idm.id
  description = "Identityd id"
}

output "raw" {
  value       = resource.azurerm_user_assigned_identity.idm
  description = "Identity raw response"
}