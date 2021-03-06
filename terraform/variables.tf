variable "access_key" {
  default     = ""
  description = "AWS access key ID"
  type        = string
}

variable "secret_key" {
  default = ""
  description = "AWS secret key"
  type = string
}

variable "region" {
  default     = "ap-northeast-2"
  description = "The AWS region to use."
  type        = string
}

variable "db_username" {
  default     = ""
  description = "Database password provided by tfc"
  type        = string
}

variable "db_password" {
  default     = ""
  description = "Database password provided by tfc"
  type        = string
}

variable "pubkey" {
  default     = ""
  description = "public key provided by tfc"
  type        = string
}

variable "availability_zones" {
  default = [
    "ap-northeast-2a",
    "ap-northeast-2c"
  ]
  description = "A list of availability zones"
  type        = list(string)
}

variable "allow_cidrs" {
  default = [
    "125.178.212.149/32", // Den Home
    "49.166.222.94/32", // Steve Home
    "192.168.200.140/32", // Ash Home
    "58.72.1.130/32" // TeamO2
  ]
  description = "A list of CIDR to access as internal"
  type        = list(string)
}
