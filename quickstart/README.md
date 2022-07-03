## 理论部分 Demo

```shell
kubectl apply -f nginx-k8s-demo/
```

## 实践部分 Go Demo

```shell

go-app 目录是要打包成镜像的 Go 程序和它的Dockerfile


kubectl apply -f go-go-app-k8s-demo/
```


##  常用命令

- kubectl apply -f  xxx.yaml 让K8s 创建在集群里按配置文件创建/更新资源对象
- kubectl get pod  | deploy | svc | ingress   查看集群中的pod、Deployment、Service、Ingress 资源的状态
- kubectl describe pod | deploy | svc | ingress  {$objectName} 查看具体资源对象当前的详细信息
- kubectl delete pod | deploy | svc | ingress  {$objectName} 删除指定对象