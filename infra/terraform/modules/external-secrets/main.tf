/**
 * # Main title
 *
 * External secrets - Secrets
 *
 * Install the External Secrets
 */

resource "helm_release" "external-secrets" {
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
