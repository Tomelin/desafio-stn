
variable "name" {
  type        = string
  default     = "external-secrets"
  description = "External Secrets name"
}

variable "repository" {
  type        = string
  default     = "https://charts.external-secrets.io"
  description = "Repository helm to install External secrets"
}

variable "chart" {
  type        = string
  default     = "secrets"
  description = "Chart of helm"
}

variable "chart_version" {
  type        = string
  default     = "0.9.19"
  description = "Chart version"
}

variable "namespace" {
  type        = string
  default     = "vault"
  description = "Namespace"
}

variable "installCRDs" {
  type        = bool
  default     = true
  description = "Enabled  install installCRDs"
}

variable "create_namespace" {
  type        = bool
  default     = true
  description = "Enable create a namespace"
}
