#!/bin/sh

for node in 'node1' 'node2' 'node3'; do
    scp assets/auth/kubeconfig core@$node.example.com:/home/core/kubeconfig;
    ssh core@$node.example.com 'sudo modprobe dm_thin_pool'
    ssh core@$node.example.com 'sudo mv kubeconfig /etc/kubernetes/kubeconfig'
done
scp -r assets core@node1.example.com:/home/core
ssh core@node1.example.com 'sudo mv assets /opt/bootkube/assets && sudo systemctl start bootkube'

