/**
 * # Main title
 *
 * Hashicorp Vault - Vault
 *
 * Install the Hashicorp vault
 */

resource "helm_release" "vault" {
  name             = var.name
  repository       = var.repository
  chart            = var.chart
  version          = var.chart_version
  namespace        = var.namespace
  create_namespace = var.create_namespace

  set {
    name  = "installCRDs"
    value = true
  }
}
