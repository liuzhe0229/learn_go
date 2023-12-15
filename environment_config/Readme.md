## 环境搭建

### VirtualBox
1. 只有Windows环境下需要虚拟机后续不再标明
2. 下载后直接安装即可 https://www.virtualbox.org/wiki/Downloads
3. 在VirtualBox中安装 linux 系统
#### 安装CentOS Linux 镜像
1. 搜索阿里云centOS 镜像,找到 https://mirrors.aliyun.com/centos/7/isos/x86_64/ 中的`CentOS-7-x86_64-DVD-2009.iso`版本下载，下载64位的即可。
2. 打开VirtualBox `新建` 一个操作系统，`名称`可以任意取;`Folder` 不要选到 `C`盘; `类型`选择`Linux`;`版本`选择 `Linux 64`（64位就可以）；`ISO Image`选择刚刚下好的镜像。
3. 内存一般调到4G就可以了，内存50G，一路默认就可以创建了。
4. 镜像系统安装一路默认，但要选一下 `INSTALLATION SOURCE` 与`NETWORK &HOST NAME` 配置网卡，`SOFTWARE SELECTION` 配置 `Server with GUI` 中的 `FTP Server`, `DNS NAME Server`。
5. 设置Administer用户名与密码。
6. 在虚拟机与物理中分别使用 `ifconfig`, `ipconfig`查看IP地址。
7. 在虚拟机与物理中互相 `ping IP` 确保互通。

### XShell 安装
1. 直接下载安装
2. 新建会话，填写名称（`centOS`）， 主机（CentOS中的ip地址）
3. 左边列表中刚刚新建的会话（`centOS`）,输入centOS中的用户名与密码
4. XShell就可以直接打开centOS的终端，而虚拟机中的就不用管了（需要`centOS`处于开机状况）。
5. 下载git `yum install git` 配置好用户名与邮件

### Docker install & Usage
1. Docker 的操作均通过XShell中链接的CentOS来操作
2. 下载命令 `curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun` 
3. 设置开机启动docker `systemctl enable docker` 如果不启动docker的，那么在终端中docker 的命令是都不能执行的,Mac中安装的 `Docker desktop`直接启动就可以了。
4. 如果docker被关闭 使用 `systemctl start docker`启动,重启docker`systemctl restart docker`.
5. 查看所有开启的docker 容器服务 `docker ps -a` ,查看所有容器 `docker container ls`,查看所有镜像 `docker image ls`
6. 配置阿里云镜像 https://www.yuque.com/bobby-zpcyu/bq1fxp/rkrbgo#RmCHv
7. 下载镜像`docker image pull [imageName]`, 删除镜像 `docker image rm [imageName]`
8. 开启服务`docker (container) run [imageName]` ,结束一个服务 ` docker container kill [containID]`
9. 查看日志`docker logs [containerId]`.
10. 重启之前启动过的服务，不要每次都docker run 不然每一次都会创建一个新都服务。 `docker restart [containerID]`.
#### docker-compose
1. docker是一个单容器，当我一个服务有很多的容器需要启动时就需要用到docker-compose。
2. 下载 `curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose`,Mac不需要单独安装compose.
3. 加权限 `sudo chmod +x /usr/local/bin/docker-compose` 
#### 安装 mysql
1. `docker pull mysql:5.7`
2. `docker run -p 3306:3306 --name mymysql -v $PWD/conf:/etc/mysgl/conf.d -v $PWD/logs:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7`
3. 进入 mysql容器 `docker exec -it [containerId] /bin/bash`
4. 登陆 `mysql -uroot -p123456`
5. 授权 `GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root' WITH GRANT OPTION;
   GRANT ALL PRIVILEGES ON *.* TO 'root'@'127.0.0.1' IDENTIFIED BY 'root' WITH GRANT OPTION;
   GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY 'root' WITH GRANT OPTION;
   FLUSH PRIVILEGES;`

### linux 下载node
1. `wget https://nodejs.org/dist/v16.15.0/node-v16.15.0-linux-x64.tar.xz`
2. 解压 `tar -xvf node-v16.15.0-linux-x64.tar.xz`
#### 建立软连接 类似于配置全局变量
1. `ln -s /root/node-v16.15.0-linux-x64/bin/node /usr/bin/node`
2. `ln -s /root/node-v16.15.0-linux-x64/bin/npm /usr/bin/npm`
#### 移除软连接
1. `rm  /usr/bin/npm`
2. `rm /usr/bin/node`