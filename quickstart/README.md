## 理论部分 Demo

```shell
kubectl apply -f nginx-k8s-demo/
```

## 实践部分 Go Demo

```shell

# go-app 目录是要打包成镜像的 Go 程序和它的Dockerfile

cd ./go-app

docker build -t registry.cn-hangzhou.aliyuncs.com/docker-study-lab/simple-app-go:v0.2 .

docker push registry.cn-hangzhou.aliyuncs.com/docker-study-lab/simple-app-go:v0.2

# go-app-k8s-demo 里是所有 K8s 资源的 YAML 定义

kubectl apply -f go-app-k8s-demo/
```


##  常用命令

- kubectl apply -f  xxx.yaml 让K8s 创建在集群里按配置文件创建/更新资源对象
- kubectl get pod  | deploy | svc | ingress   查看集群中的pod、Deployment、Service、Ingress 资源的状态
- kubectl describe pod | deploy | svc | ingress  {$objectName} 查看具体资源对象当前的详细信息
- kubectl delete pod | deploy | svc | ingress  {$objectName} 删除指定对象

**关于本教程使用的任何问题，可关注公众号「网管叨bi叨」进行咨询，另外我还准备了版 PDF 教程资料，可以发送私信「Kubernetes学习资料」领取**。
![#公众号：网管叨bi叨](https://cdn.learnku.com/uploads/images/202109/24/6964/ZXgD1fAlOU.png!large)
