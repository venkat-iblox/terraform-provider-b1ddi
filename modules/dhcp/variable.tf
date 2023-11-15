variable "ip_space_name" {
  default     = "tf_space_demo"
  description = "Name of the IP Space"
}

variable "range" {
  description = "Map of range attributes to be created."
  type        = map(string)
  default = {
    range_start = "192.168.3.5"
    range_end   = "192.168.3.10"
    name        = "tf_range"
  }
}

variable "dhcp_host" {
  //default     = "ZTP_am-demo-tf-wReT7_7745890282978598756"
  default = "00_Mehran-OnPrem-1"
  description = "DHCP Host name"
}

variable "address_block" {
  description = "Map of address block attributes to be created."
  type        = map(string)
  default = {
    name    = "example_block"
    address = "192.168.1.0"
    cidr    = 24
  }
}

variable "subnet" {
  description = "Map of subnet attributes to be created."
  type        = map(map(string))
  default = {
    "next_available" = {
      name  = "tf_subnet"
      cidr  = 26
      count = 1
    }
    "static" = {
      name    = "tf_subnet_static"
      address = "192.168.3.0"
      cidr    = 25
    }

  }
}

// Can make use of the next available IP in address_block
variable "fixed_address" {
  description = "Map of fixed address attributes to be created. This is the Gateway address"
  type        = map(string)
  default = {
    name    = "tf_fixed_address"
    address = "192.168.3.1"
  }
}
