output "created" {
  value       = local.created_at
  description = "Created time"
}

output "tags" {
  value       = local.tags
  description = "Tags to created"
}

output "resource_group" {
  value       = module.resource_group.name
  description = "Created a resource group"
}

output "identity" {
  value       = module.identity.name
  description = "Created a identity"
}

output "private_dns" {
  value       = module.private_dns.name
  description = "Created a private dns"
}

output "log_analytics" {
  value       = module.log_analytics.name
  description = "Created a log analytics"
}

output "aks" {
  value       = module.aks[*].name
  description = "Created a kubernetes service"
}


output "ingress" {
  value       = module.ingress
  description = "Values about the ingress controller"
  sensitive   = true
}