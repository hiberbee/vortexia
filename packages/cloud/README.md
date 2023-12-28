# Setup

## Gcloud Auth

[Service accounts](https://console.cloud.google.com/iam-admin/serviceaccounts?authuser=1&hl=en&project=vortexia)

```bash
terraform init
terraform plan

# vortexia is in us-east1 zone, set before kubectl
gcloud config set compute/zone us-east1

# GCE Auth plugin
gcloud components install gke-gcloud-auth-plugin # install auth plugin
export USE_GKE_GCLOUD_AUTH_PLUGIN=True # enable it
gcloud components update
gcloud container clusters get-credentials vortexia # merge config to ~/.kube/config

```

### Kubectl

Default config location: `~/.kube/config`

### Helmfile

```bash
helmfile deps # update repo index

helm plugin install https://github.com/databus23/helm-diff # install helm diff plugin
helmfile diff
```
