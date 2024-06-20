
output "name" {
  value       = resource.azurerm_log_analytics_workspace.law.name
  description = "Log analytics workspace name"
}

output "id" {
  value       = resource.azurerm_log_analytics_workspace.law.id
  description = "Log analytics workspace id"
}

output "raw" {
  value       = resource.azurerm_log_analytics_workspace.law
  description = "Log analytics workspace raw response"
}