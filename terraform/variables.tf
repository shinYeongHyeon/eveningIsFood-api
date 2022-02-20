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

variable "db_password" {
  default     = ""
  description = "Database password"
  type        = string
}