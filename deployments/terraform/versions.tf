terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "blackhorseya"

    workspaces {
      name = "todo-app"
    }
  }

  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
  }
}
