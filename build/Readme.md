#加载镜像
gunzip -c api_1.0.0.1.tar.gz | docker load

#开启容器
./start.sh

#关闭删除容器
./stop.sh


