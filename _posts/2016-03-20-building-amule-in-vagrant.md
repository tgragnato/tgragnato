---
title: Building aMule in a Disposable Vagrant VM
title_marker: Vagrant
description: An archived Debian and VirtualBox recipe for configuring and compiling aMule from a mounted source tree
layout: default
lang: en
---

This Vagrantfile describes a disposable Debian build environment for aMule. It installs the native build dependencies, runs the project's `autogen.sh` and `configure` scripts against the source tree mounted at `/vagrant`, then invokes `make` there.

The recipe records the build options and dependency set used at the time; it does not pin package versions, and the old `debian/testing64` box and package names may no longer resolve. Treat it as a historical build note, not a current installation guide. Rebuilding old software is best done in an isolated VM, with the source revision and build dependencies recorded explicitly.

```
Vagrant.configure("2") do |config|
  config.vm.box = "debian/testing64"
  config.vm.box_check_update = false
  config.vm.provider "virtualbox" do |vb|
     vb.gui = false
     vb.memory = "2048"
  end
  config.vm.provision "shell", inline: <<-SHELL
     apt-get update
     apt-get install -y build-essential autoconf automake binutils-dev
     apt-get install -y gettext autopoint
     apt-get install -y libcrypto++-dev libgeoip-dev libupnp-dev zlib1g-dev
     apt-get install -y libwxbase3.0-dev libwxgtk3.0-dev
     /vagrant/autogen.sh
     /vagrant/configure \
      --enable-debug \
      --enable-xas \
      --enable-fileview \
      --enable-plasmamule \
      --enable-mmap \
      --enable-optimize \
      --enable-upnp \
      --without-boost \
      --enable-geoip \
      --enable-webserver \
      --enable-monolithic \
      --enable-amule-daemon \
      --enable-amulecmd \
      --enable-cas \
      --enable-alcc \
      --enable-amule-gui \
      --enable-alc \
      --enable-wxcas
     make -C /vagrant
  SHELL
end
```