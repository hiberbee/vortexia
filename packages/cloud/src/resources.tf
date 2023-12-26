resource "google_container_cluster" "production" {
  name = "vortexia"
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "workers"
  cluster    = google_container_cluster.production.id
  node_count = 2

  node_config {
    machine_type = "e2-medium"
  }
}
