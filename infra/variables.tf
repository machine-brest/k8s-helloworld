variable "aws_vpc_name" {
  type        = string
  description = "AWS VPC Name"
  default     = "vpc01"
}

variable "aws_access_key" {
  type        = string
  description = "AWS Access Key"
}

variable "aws_secret_key" {
  type        = string
  description = "AWS Secret Key"
}

variable "aws_region" {
  description = "AWS region"
}

variable "cluster_name" {
  description = "ECS Cluster Name"
}
