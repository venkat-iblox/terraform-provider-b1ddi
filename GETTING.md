# Getting the provider plugin

## Building a Binary from the Source Code and Using it

Complete the following steps to build the binary:
* Install and set up Golang version 1.16 or later from
  https://golang.org/doc/install
* Install Terraform CLI from
  https://www.terraform.io/downloads.html
* Clone the repo and build it as follows:
  ```
  $ cd `go env GOPATH`/src
  $ mkdir -p github.com/infobloxopen
  $ cd github.com/infobloxopen
  $ git clone https://github.com/infobloxopen/terraform-provider-b1ddi
  $ cd terraform-provider-b1ddi
  $ make build
  ```  

* To install the resulting binary as a plugin, follow the [Terraform instruction](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers).
  
