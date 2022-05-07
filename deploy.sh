#! bin/bash
# 外部传给内部的变量脚本
app_name="second-website"
port="8000"
docker_image_username="zcxzcxczcx/second-website"

echo "对容器存活进行判断"
if test -n "$(docker ps -a |grep $app_name)" ; then
  echo "停止并且删除容器和上版本镜像"
  docker stop $app_name
  docker rm $app_name
  docker rmi $docker_image_username
else
  echo "未检查到$app_name容器运行"
fi

echo "构建最新镜像"
docker build -t $docker_image_username .
echo "启动服务"
# 有需要添加的环境变量随时添加
docker run --name $app_name -p $port:$port -d $docker_image_username