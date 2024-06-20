variable "name" {
  type        = string
  description = "Name of log analytics workspace"

  validation {
    condition     = length(var.name) == 12
    error_message = "The log analytics workspace name should be a 12 character"
  }
  validation {
    condition     = can(regex("^[a-z]{1,7}[0-9]{1,5}$", var.name))
    error_message = "The log analytics workspace name should be starts with character"
  }
}

variable "tags" {
  type        = map(string)
  description = "List of tags to log analytics workspace."
}

variable "location" {
  type        = string
  default     = "East US"
  description = "The Azure region to deploy resources"

  validation {
    condition = contains(["East US", "West US", "Brazil South", "Brazil"], var.location)

    error_message = "The location available are: 'East US', 'West US', 'Brazil South', and 'Brazil'"
  }
}

variable "resource_group" {
  type        = string
  description = "Name of resource group to deploy a resource"

  validation {
    condition     = length(var.resource_group) == 16
    error_message = "The resource group name should be a 16 character"
  }
}

variable "identity" {
  type = object({
    type         = string
    identity_ids = list(string)
  })
  description = "Identity of Kubernetes"
}


variable "retention_in_days" {
  type        = number
  default     = 30
  description = "Number of days to retention the logs"
}

variable "sku" {
  type        = string
  default     = "PerGB2018"
  description = "Specifies the SKU of the Log Analytics Workspace"
}

