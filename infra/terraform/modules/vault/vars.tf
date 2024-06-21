
variable "name" {
  type        = string
  default     = "vault"
  description = "Hashicorp vault name"
}

variable "repository" {
  type        = string
  default     = "https://helm.releases.hashicorp.com"
  description = "Repository helm to install Hashicorp vault"
}

variable "chart" {
  type        = string
  default     = "vault"
  description = "Chart of helm"
}

variable "chart_version" {
  type        = string
  default     = "0.28.0"
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
