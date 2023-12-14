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