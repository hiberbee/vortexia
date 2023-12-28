resource "google_container_cluster" "production" {
  name = "vortexia"
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "workers"
  cluster    = google_container_cluster.production.id
  node_count = 1

  node_config {
    disk_size_gb = 50

    preemptible = true
    machine_type = "e2-medium"
  }
}

data "google_container_cluster" "production" {
  name     = google_container_cluster.production.name
  location = google_container_cluster.production.location
}

resource "local_file" "kubeconfig_file" {
  content  = <<KUBECONFIG
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ${data.google_container_cluster.production.master_auth.0.cluster_ca_certificate}
    server: https://${data.google_container_cluster.production.endpoint}
  name: gke_${data.google_container_cluster.production.name}
contexts:
- context:
    cluster: gke_${data.google_container_cluster.production.name}
    user: gke_${data.google_container_cluster.production.name}
  name: gke_${data.google_container_cluster.production.name}
current-context: gke_${data.google_container_cluster.production.name}
kind: Config
users:
- name: gke_${data.google_container_cluster.production.name}
  user:
    auth-provider:
      config:
        cmd-args: config config-helper --format=json
        cmd-path: gcloud
        expiry-key: '{.credential.token_expiry}'
        token-key: '{.credential.access_token}'
      name: gcp
KUBECONFIG
  filename = "${path.module}/kubeconfig.yaml"
}
