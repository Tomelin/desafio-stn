/**
 * # Main title
 *
 * Hashicorp Vault - Vault
 *
 * Install the Hashicorp vault
 */

resource "helm_release" "vault" {
  name             = "vault"
  chart            = "vault"
  repository       = "https://helm.releases.hashicorp.com"
  version          = var.chart_version
  namespace        = var.namespace
  create_namespace = var.create_namespace


  set {
    name  = "installCRDs"
    value = true
  }

  set {
    name  = "namespace"
    value = "vault"
  }
}