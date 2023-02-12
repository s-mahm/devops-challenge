# frozen_string_literal: true

vagrantfile_api_version = '2'

clusters = [
  {
    name: 'master',
    ip: '192.168.60.10',
    hostname: 'master',
    ports: [80, 443],
    memory: 4096,
    cpus: 2
  },
  {
    name: 'nfs',
    ip: '192.168.60.11',
    hostname: 'nfs',
    ports: [80, 443],
    memory: 1024,
    cpus: 1
  },
  {
    name: 'gitlab',
    ip: '192.168.60.12',
    hostname: 'gitlab',
    ports: [80, 443],
    memory: 6144,
    cpus: 2
  },
  {
    name: 'rocketchat',
    ip: '192.168.60.13',
    hostname: 'rocketchat',
    ports: [80, 443],
    memory: 4096,
    cpus: 2
  },
]

Vagrant.configure(vagrantfile_api_version) do |config|
  config.vm.box = 'debian/bullseye64'
  config.vm.box_check_update = true
  config.ssh.insert_key = false
  clusters.each_with_index do |cluster, index|
    config.vm.define (cluster[:name]).to_s do |cluster_config|
      cluster_config.vm.hostname = (cluster[:hostname]).to_s
      cluster_config.vm.provider 'virtualbox' do |vbox|
        vbox.memory = cluster[:memory]
        vbox.cpus = cluster[:cpus]
        vbox.customize ['modifyvm', :id, '--uartmode1', 'file',
                        File.join(Dir.pwd, 'vagrant-logs/debian-bullseyes-11-cloudimg-console.log')]
      end
      cluster_config.vm.network 'private_network', ip: cluster[:ip]
      if index == clusters.size - 1
        cluster_config.vm.provision 'ansible' do |ansible|
          ansible.playbook = 'kubernetes-setup/setup.yaml'
          ansible.limit = 'all'
          # ansible.groups = {
          #   'master' => ['master'],
          #   'workers' => %w[worker-nfs worker-gitlab worker-webapp]
          # }
          ansible.inventory_path = 'kubernetes-setup/hosts.yaml'
        end
      end
    end
  end
end
