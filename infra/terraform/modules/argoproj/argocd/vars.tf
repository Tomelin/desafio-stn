
variable name {
  type        = string
  default     = "argocd"
  description = "ArgoCD name"
}

variable "repository" {
  type        = string
  default     = "https://argoproj.github.io/argo-helm"
  description = "Repository helm to install Argo CD"
}

variable "chart" {
  type        = string
  default     = "argo-cd"
  description = "Chart of helm"
}

variable "chart_version" {
  type        = string
  default     = "5.46.8"
  description = "Chart version"
}

variable namespace {
  type        = string
  default     = "argoproject"
  description = "Namespace"
}

variable installCRDs {
  type        = bool
  default     = true
  description = "Enabled  install installCRDs"
}

variable "create_namespace" {
  type        = bool
  default     = true
  description = "Enable create a namespace"
}