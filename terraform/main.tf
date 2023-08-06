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
      "sudo apt install -y apt-transport-https ca-certificates curl software-properties-common",
      "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg",
      "echo 'deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu focal stable' | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null",
      "sudo apt update",
      "sudo apt install -y docker-ce docker-ce-cli containerd.io",
      "docker version"
    ]
  }
}
