/**
 * # Main title
 *
 * ArgoProj - ArgoCD
 *
 * Module to deploy install in Kubernetes the ArgoCD
 */

resource "helm_release" "argocd" {
  name       = var.name
  repository = "https://argoproj.github.io/argo-helm"
  chart      = var.chart
  version    = var.chart_version
  namespace  = var.namespace
    create_namespace     = var.create_namespace


    set {
    name  = "installCRDs"
    value = true
  }
}