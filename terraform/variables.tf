variable "access_key" {
  default     = ""
  description = "AWS access key ID"
  type        = string
}

variable "secret_key" {
  default     = ""
  description = "AWS secret key ID"
  type        = string
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

variable "db_username" {
  default     = ""
  description = "Database username"
  type        = string
}
