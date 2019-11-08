# Cloud computing

## Virtualisation

Virtualisation can be offered on different hardware and software layers, like Central Processing Unit (CPU), disk, memory, filesystems, etc. In this chapter, we will look at some examples of creating virtual machines (VMs) after emulating the different kinds of hardware to install a guest OS on them.

**Virtual machines** are created on top of a hypervisor, which runs on top of the host machine's operating system. With hypervisors, we emulate hardware like CPU, disk, network, memory, etc., and install guest machines on it.

An hypervisor will allow a user to create multiple guest machines with different OSs. For example, we can take a Linux machine running on **bare metal**, install an hypervisor, and create different guest machines that run Linux or Windows. Examples of hypervisors:

- KVM
- VMWare
- VirtualBox

### KVM

- Kernel based Virtual Machine.
- Supports nested guests (a vm running inside a vm).
- Supports **overcommitting**, ie allocating more RAM per vm that the one available in the system. It does that by "borrowing" resources from other guests that are not using those resources.
- Highly scalable.
- Secure since it uses SELinux by default.

### VirtualBox

- By Oracle for Win, Linux and Mac.

### Vagrant

- Cross platform.
- Useful to orchestrate multiple VMs usage.
- It supports Docker.

Installation is configured by a ruby file, `Vagrantfile` that has instructions to handle multiple VMs.
It describes software requirements, packages, operating system configuration, users, and more. An example [here](https://github.com/patrickdlee/vagrant-examples/blob/master/example2/Vagrantfile).

It provides synced folders so that we can share a directory between the VM and the host system. For ex:

```rb
config.vm.synced_folder "../data", "vagrant_data"
```

Will share the local `/data` dir the the `vagrant_data` present in the VM.
**Provisioners** take care of to automatically install software or execute commands after the VM is booted.
For ex. we can tell Vagrant to use the Linux shell to install vim on our guest system:

```rb
config.vm.provision "shell", inline: <<-SHELL
              yum install vim -y
   SHELL
```
