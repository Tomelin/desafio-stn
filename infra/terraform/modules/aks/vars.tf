variable "name" {
  type        = string
  description = "Name of Azure kubernetes service"

  validation {
    condition     = length(var.name) == 12
    error_message = "The Azure kubernetes service name should be a 12 character"
  }

  validation {
    condition     = can(regex("^[a-z]{1,7}[0-9]{1,5}$", var.name))
    error_message = "The Azure kubernetes service name should be starts with character"
  }
}

variable "tags" {
  type        = map(string)
  description = "List of tags to Azure kubernetes service."
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

variable "dns_prefix" {
  type        = string
  default     = "internal"
  description = "DNS prefix to deploy with kubernetes"
}

variable "identity" {
  type = object({
    type         = string
    identity_ids = list(string)
  })
  description = "Identity of Kubernetes"
}


variable "default_node_pool" {
  type = object({
    name                   = string
    node_count             = number
    enable_host_encryption = optional(string)
    os_disk_size_gb        = optional(string)
    os_disk_type           = optional(string)
    os_sku                 = optional(string)
    ultra_ssd_enabled      = optional(bool)
    vm_size                = optional(string)
    pod_subnet_id          = optional(string)
    zones                  = list(string)
  })

  description = "AKS default node pool"
}

variable "node_pool" {
  type = list(object({
    name                   = string
    node_count             = number
    enable_host_encryption = optional(string)
    os_disk_size_gb        = optional(string)
    os_disk_type           = optional(string)
    os_sku                 = optional(string)
    ultra_ssd_enabled      = optional(bool)
    vm_size                = optional(string)
    pod_subnet_id          = optional(string)
    zones                  = list(string)
    tags                   = optional(map(string))
  }))
  default     = []
  description = "AKS node pool"
}

variable "scale_down_mode" {
  type        = string
  default     = "Deallocate"
  description = "Specifies the autoscaling behaviour of the Kubernetes Cluster"
}

variable "ultra_ssd_enabled" {
  type        = bool
  default     = false
  description = "Used to specify whether the UltraSSD is enabled in the Default Node Pool."
}

variable "enable_auto_scaling" {
  type        = bool
  default     = false
  description = "Should the Kubernetes Auto Scaler be enabled for this Node Pool"
}

variable "network_profile" {
  type = object({
    load_balancer_sku = optional(string)
    network_plugin    = optional(string)
    network_policy    = optional(string)
    outbound_type     = optional(string)
    pod_cidr          = optional(string)
    service_cidr      = optional(string)
    dns_service_ip    = optional(string)
  })
  default     = {}
  description = "Network profile kubernetes"
}

variable "oms_agent" {
  type = object({
    log_analytics_workspace_id = string
  })
  description = "OMS agent to linked with log analytics workspace"
}

variable "ingress_application_gateway" {
  type = object({
    subnet_cidr = string
  })
  default = {
    subnet_cidr = "10.224.0.0/12"
  }
  description = "Ingress application gateway"
}
