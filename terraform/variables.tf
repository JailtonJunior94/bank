variable "do_token" {
  type        = string
  description = "Token do provedor Digital Ocean"
  default     = ""
}

variable "region" {
  type        = string
  description = "Região onde será disponibilizado o cluster"
  default     = "nyc3"
}