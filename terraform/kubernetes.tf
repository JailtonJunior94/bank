resource "digitalocean_kubernetes_cluster" "k8s" {
  name    = "bank-k8s"
  region  = var.region
  version = "1.21.2-do.2"

  node_pool {
    name       = "bank-worker-pool"
    size       = "s-1vcpu-2gb"
    node_count = 1
  }
}
