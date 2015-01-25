# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "hashicorp/precise64"
  config.vm.box_check_update = false


  config.vm.provider "virtualbox" do |v|
    v.gui = false
    v.memory = 1024
    v.cpus = 2
    v.customize ["modifyvm", :id, "--cpuexecutioncap", "70"]
  end


  config.vm.define "vm_with_dockers" do |vdocker|

    # Install the latest version of Docker and nsenter
    vdocker.vm.provision "shell", inline: <<SH
      # If you'd like to try the latest version of Docker:
      # First, check that your APT system can deal with https URLs:
      # the file /usr/lib/apt/methods/https should exist.
      # If it doesn't, you need to install the package apt-transport-https.
      [ -e /usr/lib/apt/methods/https ] || {
      apt-get -y update
      apt-get -y install apt-transport-https
      }
      # Then, add the Docker repository key to your local keychain.
      sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 36A1D7869245C8950F966E92D8576A8BA88D21E9
      # Add the Docker repository to your apt sources list,
      # update and install the lxc-docker package.
      # You may receive a warning that the package isn't trusted.
      # Answer yes to continue installation.
      sh -c "echo deb https://get.docker.io/ubuntu docker main > /etc/apt/sources.list.d/docker.list"
      apt-get -y update
      apt-get -y install lxc-docker
      # Now install nsenter
      docker run --rm -v /usr/local/bin:/target jpetazzo/nsenter
SH

    # Install GIT and Fig
    vdocker.vm.provision "shell", inline: <<SH
      # It is easiest to install Git on Linux using the preferred
      # package manager of your Linux distribution.
      # Debian/Ubuntu
      # $ apt-get install git
      apt-get -y install git
      # Automatically chdir to vagrant directory upon “vagrant ssh”
      echo "\n\ncd /vagrant\n" >> /home/vagrant/.bashrc
      # Installing Fig
      sudo apt-get -y install curl
      curl -L https://github.com/docker/fig/releases/download/1.0.1/fig-`uname -s`-`uname -m` > /usr/local/bin/fig 
      chmod +x /usr/local/bin/fig
      # Installing node
      curl -sL https://deb.nodesource.com/setup | sudo bash -
      sudo apt-get install -y nodejs
      sudo npm install grunt-cli -g
SH

    vdocker.vm.network "forwarded_port", guest: 9000, host: 8000
    vdocker.vm.network "forwarded_port", guest: 27017, host: 27018
    vdocker.vm.network :private_network, ip: "192.168.33.10"

  end

end
