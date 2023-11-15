variable "ip_space_name" {
  default     = "tf_space_demo"
  description = "Name of the IP Space"
}
/*variable "address_block_name" {
    default = "example_block"
    description = "Name of the Address Block"
}
variable "address_block_address" {
    default = "192.168.1.0"
    description="Address value"
}
variable "address_block_cidr" {
    type=number
    default=24
    description = "CIDR for the address"
}*/
/*variable "subnet_name"{
    default = "example_subnet"
    description = "Name of the Subnet"
}
variable "subnet_address" {
    default = "192.168.3.0"
    description = "Address of the subnet"
}
variable "subnet_CIDR" {
    type=number
    default = 24
    description = "CIDR of the subnet"
}*/
/*variable "range_start" {
    default = "192.168.3.5"
    description = "Start address of the range"
}
variable "range_end" {
    default = "192.168.3.10"
    description = "End address of the range"
}
variable "range_name" {
    default = "example_range"
    description = "Name of the range"
}*/

variable "range" {
  description = "Map of range attributes to be created."
  type        = map(string)
  default = {
    range_start = "192.168.3.5"
    range_end   = "192.168.3.10"
    name        = "tf_range"
  }

}
/*variable "fixed_address_name" {
    default = "FA"
    description = "Name of the fixed address block"  
}
variable "fixed_address" {
    default = "192.168.3.15"
    description = "Address for the FA"
}*/

/*variable "next_available_subnet_name" {
    default = "subnet-nas"
    description = "Name of the subnet to be created using NAS"
}

variable "next_available_subnet_cidr" {
    default = 27
    description = "subnet CIDR to be created using NAS"
}

variable "next_available_subnet_count" {
    default = 3
    description = "count of subnet to be created using NAS"
}*/

variable "dhcp_host" {
  default     = "Venkat-01"
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

variable "fixed_address" {
  description = "Map of fixed address attributes to be created."
  type        = map(string)
  default = {
    name    = "tf_fixed_address"
    address = "192.168.3.15"
  }
}