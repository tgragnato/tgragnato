---
title: aMule
---

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