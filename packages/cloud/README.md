# Setup

## Gcloud Auth


[Service accounts](https://console.cloud.google.com/iam-admin/serviceaccounts?authuser=1&hl=en&project=vortexia)

```bash
terraform init
```


### Helmfile

```bash
helmfile deps # update repo index
helmfile diff
```

### Kubectl

Default config location: `~/.kube/config`


```bash

- GCE Auth plugin

```bash
gcloud components install gke-gcloud-auth-plugin** # install auth plugin 
export USE_GKE_GCLOUD_AUTH_PLUGIN=True # enable it
gcloud container clusters get-credentials # merge config to ~/.kube/config
```
