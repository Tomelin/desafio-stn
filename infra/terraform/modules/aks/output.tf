
output "name" {
  value       = resource.azurerm_kubernetes_cluster.aks.name
  description = "Kubernetes service name"
}

output "id" {
  value       = resource.azurerm_kubernetes_cluster.aks.id
  description = "Kubernetes service id"
}

output "raw" {
  value       = resource.azurerm_kubernetes_cluster.aks
  description = "Kubernetes service raw response"
}