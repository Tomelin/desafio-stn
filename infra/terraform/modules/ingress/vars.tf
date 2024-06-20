variable "name" {
  type        = string
  default     = "nginx-ingress-controller"
  description = "Name of nginx controller"
}

variable "repository" {
  type        = string
  default     = "https://charts.bitnami.com/bitnami"
  description = "Repository ghelm to install nginx controller"
}

variable "chart" {
  type        = string
  default     = "nginx-ingress-controller"
  description = "Chart of helm"
}

variable "labels" {
  type        = map(string)
  default     = {}
  description = "Labels of ingress"
}

variable "namespace" {
  type        = string
  default     = "ingress-controller"
  description = "Namespace name"
}

variable "create_namespace" {
  type        = bool
  default     = true
  description = "Enable create a namespace"
}