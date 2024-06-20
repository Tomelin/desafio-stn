output "vault" {
  value       = helm_release.vault
  sensitive   = true
  description = "Result of installation a Hashicopr Vault"
}
