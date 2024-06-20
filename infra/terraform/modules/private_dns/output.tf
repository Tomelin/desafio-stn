
output "name" {
  value       = resource.azurerm_private_dns_zone.private.name
  description = "Private DNS name"
}

output "id" {
  value       = resource.azurerm_private_dns_zone.private.id
  description = "Private DNS id"
}

output "raw" {
  value       = resource.azurerm_private_dns_zone.private
  description = "Private DNS raw response"
}