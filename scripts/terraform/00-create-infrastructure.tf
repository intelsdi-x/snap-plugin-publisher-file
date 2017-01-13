variable "deploymentName" {
  type = "string"
  description = "The desired name of your deployment."
}

variable "ami"{
  type = "string"
  description = "AMI to use for all VMs in cluster."
}

variable "SSHKey" {
  type = "string"
  description = "SSH key to use for VMs."
}

variable "instance_type" {
  type = "string"
  description = "Size of VMs to use"
}

variable "terminate_protect" {
  type = "string"
  default = "false"
}

variable "awsRegion" {
  type = "string"
}

provider "aws" {
  region = "${var.awsRegion}"
}

resource "aws_security_group" "allow_ssh" {
  name = "allow_ssh"
  description = "Allow ssh inbound traffic"

  ingress {
      from_port = 0
      to_port = 22
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
      from_port = 0
      to_port = 0
      protocol = "-1"
      cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "build" {
    count = "1"
    ami = "${var.ami}"
    instance_type = "${var.instance_type}"
    vpc_security_group_ids = ["${aws_security_group.allow_ssh.id}"]
    key_name = "${var.SSHKey}"
    disable_api_termination = "${var.terminate_protect}"
    tags {
      Name = "${var.deploymentName}-build-${count.index + 1}"
    }
    provisioner "remote-exec" {
      inline = ["# Connected!"]
      connection {
        user = "ubuntu"
      }
    }
}

output "mapped_ips" {
   value = "${map("build", aws_instance.build.public_ip)}"
}

output "mapped_ports" {
  value = "${map("build", 22)}"
}

output "mapped_users" {
  value = "${map("build", "ubuntu")}"
}
