variable "do_token" {
  type        = string
  description = "Token do provedor Digital Ocean"
  default     = "034bdfb3309f2f5d39e67ca941d4566b3734648c381de0bd79369ea569836cb6"
}

variable "region" {
  type        = string
  description = "Região onde será disponibilizado o cluster"
  default     = "nyc3"
}