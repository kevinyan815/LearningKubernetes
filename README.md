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
# 获得集群对外提供访问的ip, 在每个人的电脑上ip不一样。
# MySQL的Service对外暴露的端口为30306
IP=$(minikube ip)  

mysql -uroot -puser -h ${IP} -P 30306
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

- 从电脑上访问Redis

```
# Redis Service对外暴露端口为：31379
IP=$(minikube ip)  

redis-cli -h ${IP} -p 31379

```

### 安装ETCD

- 创建命名空间etcd, etcd的资源都放在这个空间下:
```
kubectl create namespace etcd
```

- 创建用到的Services:
```
kubectl apply -f resources/services.yml -n etcd
```

- 创建etcd 集群:
```
cat resources/etcd.yml.tmpl | resources/config.bash | kubectl apply -n etcd -f -
```


- 从电脑上访问etcd
```
IP=$(minikube ip)

PORT=$(kubectl get services -o jsonpath="{.spec.ports[].nodePort}" etcd-client -n etcd)
ETCDCTL_API=3 ./etcdctl --endpoints ${IP}:${PORT}  get root --prefix
```

### 安装MongoDB

安装的是Standalone（单点）类型的MongoDB

- 创建root用户的用户名和密码的secret对象
```
kubectl apply -f mongodb-singleton/mongo-secret.yaml
```

- 创建Deployment和Service
```
kubectl apply -f mongodb-singleton/deployment-service.yaml
```

- 客户端访问，推荐使用MongoDB Compass
```
IP=$(minikube ip) # 假设$IP是 192.168.64.4

mongodb://username:password@192.168.64.4:30017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false
```
- 在客户端里为数据库创建用户
```
> use my-database

> db.createUser(
  {
    user: "my-user",
    pwd: "passw0rd",
    roles: [ { role: "readWrite", db: "my-database" } ]
  }
)
```
