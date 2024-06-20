output "argocd" {
  value       = helm_release.argocd
  sensitive   = true
  description = "Result of installation a ArgoCD"
}
