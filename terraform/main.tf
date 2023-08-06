terraform {
  required_providers {
    linode = {
      source  = "linode/linode"
      version = "2.3.0"
    }
  }
}

provider "linode" {
  token = var.token
}

resource "linode_instance" "okr-generator" {
  label           = "ubuntu"
  type            = "g6-nanode-1"
  region          = "ap-south"
  image           = "linode/ubuntu20.04"
  authorized_keys = [var.authorized_keys]

  connection {
    type        = "ssh"
    user        = "root"
    private_key = file("~/.ssh/id_github")
    host        = self.ip_address
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt update",
      "sudo apt install -y apt-transport-https ca-certificates curl software-properties-common"
    ]
  }
}
