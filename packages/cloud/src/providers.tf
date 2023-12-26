provider "google" {
  credentials = file("account.json")
  project     = "vortexia"
  region      = "us-east1"
}
