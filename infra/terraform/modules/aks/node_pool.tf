resource "azurerm_kubernetes_cluster_node_pool" "pool" {
  count                  = length(var.node_pool)
  name                   = var.node_pool[count.index].name
  kubernetes_cluster_id  = azurerm_kubernetes_cluster.aks.id
  vm_size                = var.node_pool[count.index].vm_size
  node_count             = var.node_pool[count.index].node_count
  enable_host_encryption = var.node_pool[count.index].enable_host_encryption
  os_disk_size_gb        = var.node_pool[count.index].os_disk_size_gb != "" ? var.node_pool[count.index].os_disk_size_gb : null
  os_disk_type           = var.node_pool[count.index].os_disk_type == "" ? "Managed" : var.node_pool[count.index].os_disk_type
  os_sku                 = var.node_pool[count.index].os_sku == "" ? "Ubuntu" : var.node_pool[count.index].os_sku
  pod_subnet_id          = var.node_pool[count.index].pod_subnet_id != "" ? var.node_pool[count.index].pod_subnet_id : null
  ultra_ssd_enabled      = var.node_pool[count.index].ultra_ssd_enabled
  zones                  = length(var.node_pool[count.index].zones) == 0 ? ["1", "2", "3"] : var.node_pool[count.index].zones
  enable_node_public_ip  = false
  scale_down_mode        = var.scale_down_mode
  upgrade_settings {
    max_surge = "10%"
  }


  depends_on = [
    azurerm_kubernetes_cluster.aks
  ]
  tags = merge(var.node_pool[count.index].tags, var.tags)
}