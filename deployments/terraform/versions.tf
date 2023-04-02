terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "blackhorseya"

    workspaces {
      name = "ekko"
    }
  }

  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
  }
}
