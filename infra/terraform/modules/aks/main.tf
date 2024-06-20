/**
 * # Main title
 *
 * Azure Kubernetes service
 *
 * Module to deploy Azure kubernetes service 
 */

resource "azurerm_kubernetes_cluster" "aks" {
  name                = "${var.name}-aks"
  location            = var.location
  resource_group_name = var.resource_group
  dns_prefix          = var.dns_prefix

  public_network_access_enabled = false

  key_vault_secrets_provider {
    secret_rotation_enabled = true
  }

  default_node_pool {
    name                   = var.default_node_pool.name
    node_count             = var.default_node_pool.node_count < 1 ? 2 : var.default_node_pool.node_count
    enable_host_encryption = var.default_node_pool.enable_host_encryption
    os_disk_size_gb        = var.default_node_pool.os_disk_size_gb != "" ? var.default_node_pool.os_disk_size_gb : null
    os_disk_type           = var.default_node_pool.os_disk_type == "" ? "Managed" : var.default_node_pool.os_disk_type
    os_sku                 = var.default_node_pool.os_sku == "" ? "Ubuntu" : var.default_node_pool.os_sku
    pod_subnet_id          = var.default_node_pool.pod_subnet_id != "" ? var.default_node_pool.pod_subnet_id : null
    ultra_ssd_enabled      = var.default_node_pool.ultra_ssd_enabled
    vm_size                = var.default_node_pool.vm_size == null ? "Standard_DS2_v2" : var.default_node_pool.vm_size
    zones                  = length(var.default_node_pool.zones) == 0 ? ["1", "2", "3"] : var.default_node_pool.zones
    enable_node_public_ip  = false
    scale_down_mode        = var.scale_down_mode
    upgrade_settings {
      max_surge = "10%"
    }
  }

  network_profile {
    load_balancer_sku = var.network_profile.load_balancer_sku == null ? "standard" : var.network_profile.load_balancer_sku
    network_plugin    = var.network_profile.network_plugin == null ? "azure" : var.network_profile.network_plugin
    network_policy    = var.network_profile.network_policy == null ? "azure" : var.network_profile.network_policy
    outbound_type     = var.network_profile.outbound_type == null ? "loadBalancer" : var.network_profile.outbound_type
    # pod_cidr          = var.network_profile.pod_cidr == null ? var.network_profile.network_plugin == "azure" ? null : "10.10.0.0/16" : var.network_profile.pod_cidr
    pod_cidr       = var.network_profile.pod_cidr == null ? null : var.network_profile.pod_cidr
    service_cidr   = var.network_profile.service_cidr == null ? "172.18.0.0/16" : var.network_profile.service_cidr
    dns_service_ip = var.network_profile.dns_service_ip == null ? "172.18.0.10" : var.network_profile.dns_service_ip
  }

  oms_agent {
    log_analytics_workspace_id = var.oms_agent.log_analytics_workspace_id
  }

  identity {
    type         = var.identity.type
    identity_ids = var.identity.identity_ids
  }

  tags = var.tags

  lifecycle {
    ignore_changes = [
      tags["created_at"]
    ]
  }
}
