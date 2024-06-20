variable "name" {
  type        = string
  description = "Name of resource group"

  validation {
    condition     = length(var.name) == 12
    error_message = "The resource group name should be a 12 character"
  }

  validation {
    condition     = can(regex("^[a-z]{1,7}[0-9]{1,5}$", var.name))
    error_message = "The resource group name should be starts with character"
  }
}

variable "tags" {
  type        = map(string)
  description = "List of tags to resource group."
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
