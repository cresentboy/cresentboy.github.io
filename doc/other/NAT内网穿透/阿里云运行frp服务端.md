安装frps文件

进入安装目录(本文以/usr/local/为例)

```
cd /opt/software
```

 将linux系统的frp压缩文件拖至文件夹并解压

```
tar -xzvf frp_0.38.0_linux_amd64.tar.gz
```

编辑frps.ini文件

```
vim frps.ini
```

配置文件信息（这里是我需要用到的，想要全面点的可自己查找）

```
[common]
bind_addr = 0.0.0.0
bind_port = 7360
token = 123
dashboard_port = 7370
dashboard_user = *
dashboard_pwd = *
log_file = ./frps.log
log_level = info
log_max_days = 3
max_pool_count = 100
#vhost_http_port = 80
# 路由地址
#subdomain_host = www.test.com
```

启动frps，进入frps文件夹

```
nohub ./frps -c ./frps.ini &
```

输入`http://ip:7370`来查看`frps`服务状态，连不上说明启动失败

### 继续配置开机自启动

创建frps.service文件

```
sudo vim /lib/systemd/system/frps.service
```

该文件的配置内容如下

```
[Unit]
Description=frps service
After=network.target syslog.target
Wants=network.target
 
[Service]
Type=simple
#启动服务的命令（此处写你的frps的实际安装目录）
ExecStart=/opt/software/frp_0.38.0_linux_amd64/frps -c /opt/software/frp_0.38.0_linux_amd64/frps.ini 
 
[Install]
WantedBy=multi-user.target
```

再启动frps

```
sudo systemctl start frps
```

有可能报：Warning: frps.service changed on disk. Run 'systemctl daemon-reload' to reload units.

需要重新加载配置，修改systemd配置执行

```
sudo systemctl daemon-reload
```

systemd常用的命令

```
systemctl status frps.service       # 查看 frps 服务状态
systemctl cat frps.service          # 查看 frps 服务配置<br>systemctl restart frps              # 重启 frps 服务
sudo systemctl start frps.service   # 启动 frps 服务
sudo systemctl stop frps.service    # 停止 frps 服务
sudo systemctl daemon-reload        # 重新加载配置，修改 systemd 配置执行
sudo systemctl enable frps.service  # 设置开机启动，根据 install 建立软链
sudo systemctl disable frps.service # 取消开机启动，根据 install 移除软链
```

## 配置多个内网主机

#### 错误的多客户端配置

使用一台阿里云的公网服务器，我们可以配置很多内网机器的 frp 内网穿透，公网服务器上只需要按照上述的配置一次即可，但是内网机器的配置稍有不同，如果使用了一样的配置则后添加的内网机器是无法连接上公网服务器的。这里假设另一台内网机器2的 frpc.ini 配置如下，来说明会遇到的问题：

```
$ vi frpc.ini
[common]
server_addr = xxx.xxx.xxx.xxx  <==这里还是按照上面的假设，公网服务器的ip为xxx.xxx.xxx.xxx
server_port = 7000

[ssh]
type = tcp 
local_ip = 127.0.0.1
local_port = 22
remote_port = 6001     <==remote_port设置为另一个值
12345678910
```

两个内网主机的配置除了 remote_port 不一样之外，都是一样的。但是在内网机器2上运行 frpc 后，公网服务器的 nohup.out 中会记录一下的错误：

```
[W] [control.go:332] [280d36891a6ae0c7] new proxy [ssh] error: proxy name [ssh] is already in use
1
```

后来发现，frp 中是通过 [ssh] 这个名字来区分不同客户端的，所以不同的客户端要配置成不同的名字。

#### 正确的多客户端配置

内网机器1和内网机器2的配置应该区分如下：

```
内网机器1：
[ssh]                      <==不同点
type = tcp 
local_ip = 127.0.0.1
local_port = 22
remote_port = 6000         <==不同点

内网机器2：
[ssh1]                     <==不同点
type = tcp 
local_ip = 127.0.0.1
local_port = 22
remote_port = 6001         <==不同点
12345678910111213
```

在两个内网机器上分别运行 frpc 客户端程序后，一般就可以通过以下的方法 ssh 登录：

```
内网机器1：
$ ssh -p 6000 user_name1@server_addr

内网机器2：
$ ssh -p 6001 user_name2@server_addr
12345
```

> 以上参数中，server_addr是公网服务器的公网ip；user_name1、user_name2 分别是内网机器1、2的用户名，之后分别使用登录密码就可以登录。

## connection timed out 解决

但是有时候会发现按照以上的配置还是使 frp 的服务端与客户端建立连接，在客户端上会出现以下错误：

```sh
2018/09/17 22:02:23 [W] [control.go:113] login to server failed: dial tcp xxx.xxx.xxx.xxx:7000: connect: connection timed out
dial tcp xxx.xxx.xxx.xxx:7000: connect: connection timed out
```

关闭阿里云服务器上的防火墙即可：

```sh
sudo systemctl stop firewalld
sudo systemctl disable firewalld
```

