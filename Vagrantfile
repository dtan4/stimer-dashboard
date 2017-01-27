# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"

  config.vm.synced_folder Dir.pwd, "/home/ubuntu/src/github.com/dtan4/stimer-dashboard"

  config.vm.network :private_network, ip: "172.17.8.101"

  config.vm.provision "shell", inline: <<-SHELL
    wget -q https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
    tar zxf go1.7.4.linux-amd64.tar.gz -C /usr/local
    rm go1.7.4.linux-amd64.tar.gz
    echo 'export GOROOT=/usr/local/go' >> /home/ubuntu/.bashrc
    echo 'export GOPATH=$HOME' >> /home/ubuntu/.bashrc
    echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /home/ubuntu/.bashrc
    mkdir -p /home/ubuntu/bin && chown -R ubuntu:ubuntu /home/ubuntu/bin
    mkdir -p /home/ubuntu/src && chown -R ubuntu:ubuntu /home/ubuntu/src
    add-apt-repository ppa:masterminds/glide
    apt-get update
    apt-get install -y cmake glide
    echo 'cd /home/ubuntu/src/github.com/dtan4/stimer-dashboard' >> /home/ubuntu/.bashrc
  SHELL
end
