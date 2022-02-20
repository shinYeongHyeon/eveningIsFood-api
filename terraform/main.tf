terraform {
  cloud {
    organization = "evening-is-food"

    workspaces {
      name = "evening-is-food"
    }
  }
}

provider "aws" {
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}

