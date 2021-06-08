### 安装MySQL

- 创建ConfigMap 把my.cnf里的配置放在ConfigMap对象里，后面作为配置文件会挂载进Pod
```
kubectl apply -f mysql-singleton/mysql-configmap.yaml
```
- 创建MySQL应用的Pod和向外提供服务的Service
```
kubectl apply -f mysql-singleton/deployment-service.yaml
```

- 从电脑链接MySQL
```
# 假设电脑上Kubernetes用的是Minikube， 这里只是例子， 也可以用K3d，用什么安装工具在电脑上安装Kubernetes不重要
# 获得集群对外提供访问的ip, 在每个人的电脑上ip不一样，下面以{minikube-ip}替代。
minikube ip  

mysql -uroot -puser -h {minikube-ip} -P 30306
```

### 安装Redis

- 创建ConfigMap
```
kubectl apply -f mysql-singleton/mysql-config.yaml
```

- 创建调度Redis应用的Deployment和向外提供服务的Service
```
kubectl apply -f redis-singleton/deployment-service.yaml
```

### 安装ETCD

创建命名空间etcd, etcd的资源都放在这个空间下:

    $ kubectl create namespace etcd

创建用到的Services:

    $ kubectl apply -f resources/services.yml -n etcd

创建etcd 集群:

    $ cat resources/etcd.yml.tmpl | resources/config.bash | kubectl apply -n etcd -f -


通过minikube ip和nodeport 访问etcd

    $ IP=$(minikube ip)
    $ PORT=$(kubectl get services -o jsonpath="{.spec.ports[].nodePort}" etcd-client -n etcd)
    $ ETCDCTL_API=3 ./etcdctl --endpoints ${IP}:${PORT}  get root --prefix
