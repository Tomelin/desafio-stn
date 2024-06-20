/**
 * # Main title
 *
 * Nginx Ingress Controller
 *
 * Module to deploy install in Kubernetes the Ingress Controller
 */

resource "helm_release" "nginx_ingress" {
  name             = var.name
  repository       = var.repository
  chart            = var.chart
  create_namespace = var.create_namespace
  namespace        = var.namespace

  set {
    name  = "service.type"
    value = "ClusterIP"
  }
}
