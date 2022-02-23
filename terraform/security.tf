module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.12.0"

  name = "vpc"
  cidr = "10.0.0.0/16"

  azs              = var.availability_zones
  public_subnets   = ["10.0.10.0/24", "10.0.20.0/24"]
  private_subnets  = ["10.0.11.0/24", "10.0.21.0/24"]
  database_subnets = ["10.0.12.0/24", "10.0.22.0/24"]

  enable_nat_gateway     = true
  single_nat_gateway     = true
  one_nat_gateway_per_az = false

  create_database_subnet_group           = true
  create_database_subnet_route_table     = true
  create_database_internet_gateway_route = true

  enable_dns_hostnames = "true"
  enable_dns_support   = "true"
}

resource "aws_security_group" "dev_web" {
  vpc_id = module.vpc.vpc_id
  name   = "dev_web"
  description = "eif-web-dev security group"

  ingress {
    from_port    = 22
    to_port      = 22
    protocol     = "tcp"
    cidr_blocks  = var.allow_cidrs
  }

  ingress {
    from_port    = 80
    to_port      = 80
    protocol     = "tcp"
    cidr_blocks  = ["0.0.0.0/0"]
  }

  ingress {
    from_port    = 443
    to_port      = 443
    protocol     = "tcp"
    cidr_blocks  = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "dev_rds" {
  vpc_id = module.vpc.vpc_id
  name   = "dev_rds"
  description = "eif-rds-dev security group"

  ingress {
    from_port    = 3306
    to_port      = 3306
    protocol     = "tcp"
    cidr_blocks  = var.allow_cidrs
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "eif_web" {
  key_name   = "eif_web"
  public_key = var.pubkey
}
