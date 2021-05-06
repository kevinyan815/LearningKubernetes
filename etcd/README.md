minikube-etcd
=============

Create a etcd cluster on minikube, for use with local development against
etcd.

Prerequisites
-------------

Update hosts file with:

    $(minikube ip) etcd-0.etcd
    $(minikube ip) etcd-1.etcd
    $(minikube ip) etcd-2.etcd

Usage
-----

Start minikube:

    $ minikube start --vm-driver=xhyve

Create a namespace:

    $ kubectl create namespace etcd

Create the services:

    $ kubectl create -f resources/services.yml -n etcd

Create the cluster:

    $ cat resources/etcd.yml.tmpl | resources/config.bash | kubectl create -n etcd -f -

Verify the cluster's health:

    $ kubectl exec -it etcd-0 -n etcd etcdctl cluster-health

The cluster is exposed through minikube's IP:

    $ IP=$(minikube ip)
    $ PORT=$(kubectl get services -o jsonpath="{.spec.ports[].nodePort}" etcd-client -n etcd)
    $ etcdctl --endpoints http://${IP}:${PORT} get foo

Destroy the services:

    $ kubectl delete services,statefulsets --all -n etcd

License
-------

MIT
