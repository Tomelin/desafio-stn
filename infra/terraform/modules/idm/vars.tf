variable "name" {
  type        = string
  description = "Name of identity"

  validation {
    condition     = length(var.name) == 12
    error_message = "The identity name should be a 12 character"
  }
  validation {
    condition     = can(regex("^[a-z]{1,7}[0-9]{1,5}$", var.name))
    error_message = "The identity name should be starts with character"
  }
}

variable "tags" {
  type        = map(string)
  description = "List of tags to identity."
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