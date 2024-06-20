output "external-secrets" {
  value       = helm_release.external-secrets
  sensitive   = true
  description = "Result of installation a External Secrets"
}
