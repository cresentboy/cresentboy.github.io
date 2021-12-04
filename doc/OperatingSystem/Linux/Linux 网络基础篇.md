#

# 第一章 计算机基础

## 第一节 进制转换

- **十进制数：逢十进一**
  **1 2 3 4 5 6 7 8 9 10**
- **二进制数：逢二进一**
  **0 1 10 11 100 101 110 111 1000**
  **1001 1010** **（1-10 的表示）**
  **计算机只能识别二进制**
- **十进制转化二进制**
  - **十进制数除二，把余数放到右边，直到除到得数为 0 或者 1，用最后的得数开头把最右边的余数放在一起，就是二进制数
    比如：**
    **123 的二进制数是 1111011**
    ￼
  - **二进制转化十进制**
    - **求权，然后把它加起来，从后面开始加**
      **比如：11010011
      1\*2^0+1\*2^1+0\*2^2+0\*2^3+1\*2^4+0\*2^5+1\*2^6+1\*2^7=211**
    - **装水桶法**
      **适合不能太大，258 以内，不支持小数，用在心算，算 IP
      先列个表，写二进制对位**

| 1   | 1   | 1   | 1   | 1   | 1   | 1   | 1   |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 128 | 64  | 32  | 16  | 8   | 4   | 2   | 1   |

**十进制转二进制。如：137**
**137 大于 128，减去 128 有余数则，128 下面有 1**
**1 0 0 0 0 0 0 0**
**减完之后剩 9，得挨个往后减，不能跨过，减不动就计 0，换下一位
1 0 0 0 1 0 0 0
直到碰到 8，计 1
1 0 0 0 1 0 0 1**

**十进制换二进制。如：110110
把数对应到表格下面，不够补 0
1 1 1 1 1 1 1 1
128 64 32 16 8 4 2 1
0 0 1 1 0 1 1 0
然后把有 1 的对应的数加起来
32+16+4+2=54**

**十六进制：逢十六进一**
**1-9 表示：1-9 a-f 表示：10-15**

- **十六进制转化二进制
  先转换为二进制(四个二进制位代表一个十六进制位)，直接连起来就为二进制
  再转换为十进制**
  **比如：8ea 转换十进制
  8 e(14) a(10)
  1000 1110 1010
  二进制表示：100011101010
  十进制表示：装水桶法
  2048+128+64+32+8+2=2282**
- **十进制转化十六进制
  先转化二进制
  再转化十六进制
  比如：297 转化为十六进制
  297 二进制：100101001
  装水桶法：
  0001 0010 1001 不管左侧补 0
  1 2 9
  十六进制数 123**

# 第二章 网络基础

## 第一节 互联网概述

**Internet 简介**

- **Internet 是国际互联网。**

**Internet 的发展历史**

- **ARPAnet ——> TCP/IP ——> NSFnet ——> ANSnet ——> Internet**

**Internet 的作用**

- **互联网缩短了时空的距离，大大加快了消息的传递。使得社会的各种资源得以共享**
- **互联网创造出了更多的机会，可以有效的提高传统产业的生产效率，有力的拉动消费需求，从而促进经济增长**
- **互联网为各个层次的文化交流提供了良好的平台**

**Internet 的应用领域**

- **在线招聘**
- **网上银行**
- **在线教育系统**
- **交友或参与讨论**

**互联网的典型应用**

- **WWW：万维网**
- **FTP：文件传输协议**
- **E-MAIL：电子邮件**

**WWW WWW（World Wide Web）又称全球网，环球网或万维网等，也可简称为 Web
服务器 S <——> 客户端 C C/S 架构**

**URL：统一资源定位
协议+域名或者 IP：端口+网页路径+网页名
http://www.lampbrother.net:80/index.html**

- **http 协议**
- **.net 一级域名**
- **lampbrother 二级域名**
- **www 三级域名**
- **:80 端口，默认 80**

**常见的一级域名** **（也叫顶级域名)**

| 一级域名 | 组织           |
| -------- | -------------- |
| Edu      | 教育组织       |
| Com      | 商业组织       |
| gov      | 非军事政府机构 |
| Mil      | 军事结构       |
| Org      | 其他组织       |
| net      | 网络服务机构   |

| 一级域名 | 地区     |
| -------- | -------- |
| au       | 澳大利亚 |
| cn       | 中国     |
| in       | 印度     |
| us       | 美国     |
| uk       | 英国     |

**二级域名自己申请**

**遇到钓鱼网站注意二级域名，他是唯一的，别被骗**

## 第二节 互联网接入方法

**互联网接入方法**

- **ADSL：非对称数字用户环路**
- **FTTH：光纤入户**
- **小区宽带**
- **固定 IP 光纤**

**ADSL**
**非对称数字用户环路。ADSL 可以提供最高 1Mbps 的上行速率和最高 8Mbps(1Mbps=1024Kbps 比特率除以 8 才为字节)的下行速率。最新的 ADSL2+技术可以提供最高 24Mbps 的下行速率，ADSL2+打破了 ADSL 接入方式带宽限制的瓶颈，使其应用范围更加广泛。**

- 优点：
  - **使用电话线，节省了布网成本**
  - **上网的同时可以打电话，节省了点话费**
- 缺点：
  - **铜线耗材昂贵**
  - **带宽限制**
  - **动态 IP 地址(拨号)**

**FTTH
光纤入户的带宽更高，而光纤的原材料是二氧化硅，在自然界取之不尽用之不竭。当然也是动态 IP 地址。**

**小区宽带
小区宽带是一个大局域网，所有客户都在同一个网段中。外网接口可以是 FTTH，也可以是固定 IP 的光纤。在用户高峰期太卡，所有人用一个带宽，共享带宽。最好不要选择这个，非常不安全，监听网关可以抓获所有的数据包。**

**固定 IP 的光纤
带宽自由申请，价格最贵。因为固定 IP 地址，所以可以搭建服务器。一般是公司和服务器使用。**

# 第三章 网络通信协议

## 第一节 OSI 七层模型和 TCP 四层模型

**OSI 的七层框架**
![OSI 七层框架](https://img-blog.csdnimg.cn/20200530224016933.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)**计算机 A 通过应用层到达物理层，然后通过网络层传输到计算机 B 的物理层，然后从物理层传输到应用层。**

- **物理层：设备之间的比特流的传输，物理接口，电气特性等**
- **数据链路层：成帧，用 MAC 地址访问媒介，错误检测与修正，局域网主要靠 MAC（MAC 地址）**
- **网络层：提供逻辑地址，选路(路由选择)（网络 IP）**
- **传输层：可靠与不可靠的传输，传输前的错误检测，流控 （包头端口）**
- **会话层：对应用户会话的管理，同步**
- **表示层：数据的表现形式，特定功能的实现如—加密**
- **应用层：用户接口**

**TCP/IP 协议 4 层模型**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200530224106947.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
**TCP/IP 模型与 OSI 模型的对应**

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200530224132997.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)

- **网络接口层
  网络接入层与 OSI 参考模型中的物理层和数据链路层相对应。它负责监视数据在主机和网络之间的交换。事实上，TCP/IP 本身并未定义盖层的协议，而由参与互联的各网络使用自己的物理层和数据链路层进行连接。地址解析协议(ARP 把 IP 解析为 MAC 地址)工作在此层，即 OSI 参考模型的数据链路层。**

- **网际互联层
  网际互联层对应于 OSI 参考模型的网络层，主要解决主机到主机的通信问题。它所包含的协议设计数据包在整个网络上的逻辑传输。该层有三个主要协议：网际协议(IP)，互联网组管理协议(IGMP)和互联网控制报文协议(ICMP)。**

- **传输层
  传输层对应于 OSI 参考模型的传输层，为应用层实体提供端到端的通信功能，保证了数据包的顺序传送及数据的完整性。该层定义了两个主要的协议：传输控制协议(TCP)和用户数据报协议(UDP)。**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200530224159218.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
  **TCP/IP 也叫面向连接的可靠传输协议**

  **端口是应用层与传输层的接口**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200530224225526.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
  **对于 TCP 和 UDP 可以使用同一个端口，通过分别协议不同来区分不同服务，比如：DNS
  有的服务可以使用多个端口**

- **应用层
  应用层对应于 OSI 参考模型的表示层和应用层，为用户提供所需要的各种服务，例如：FTP，Telnet，DNS，SMTP 等。**

**数据封装过程**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200530224253204.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)

**TCP/IP 模型与 OSI 模型的比较**

- 共同点：
  - **OSI 参考模型和 TCP/IP 参考模型都采用了层次结构的概念**
  - **都能提供面向连接和无连接两种通信服务机制**
- 不同点：
  - **前者是七层模型，后者是四层模型**
  - **对可靠性要求不同(后者更高)**
- **OSI 模型是在协议开发前设计的，具有通用性。TCP/IP 是先有协议集然后建立模型，不适用于非 TCP/IP 网络。**
- **实际市场应用不同(OSI 模型只是理论上的模型，并没有成熟的产品，而 TCP/IP 已经成为“实际上的国际标准”)**

# 第三章 网络通信协议

## 第一节 网络协议层与 IP 地址的划分

- **网络层协议**
- **网际协议(IP)**
- **互联网控制报文协议(ICMP)**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531225929967.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
- **IP 包头**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531225956501.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
  **注：字段后面括号中的数字是指该字段在 IP 数据包头部信息所占的位数
  IPV4 中为 32 位**
- **IP 地址的分类**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/2020053123003026.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
- **子网掩码的使用**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531230050740.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
- **变长子网掩码以及子网规划**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531230108663.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
- **如何判断子网掩码是否合法的唯一标准为：网络位是否全为连续的 1**
- **网络地址为每一个网段的第一个 IP，用来表示网段，不可以使用**
- **广播地址为每一个网段最后一个 IP，用来表示网段，不可以使用**
- **IP 与子网掩码相与可以算出来网络地址，而只要是子网掩码为连续的 0 的字段就把 IP 地址对应的字段变为 1，其他字段照着 IP 写出来，这样出来的为广播地址**

**ICMP 协议**

- **ICMP 消息通过 IP 数据报传送，被用来发送错误和控制信息。（ping 操作）**

**ICMP 定义了很多信息类型，例如：**

- **目的地不可达**
- **TTL 超时**
- **信息请求**
- **信息应答**
- **地址请求**
- **地址应答**

# 第四章 Linux 网络基础

## 第一节 Linux 的 IP 地址配置

**Linux 的 IP 地址配置地址的方法**

- **ifconfig 命令临时配置 IP 地址**
- **Setup 工具永久配置 IP 地址**
- **修改网络配置文件**
- **图形界面配置 IP 地址**

**ifconfig 命令**

- **Ifconfig 命令：查看与配置网络状态命令**
- **ifconfig eth0 192.169.0.200 netmask 255.255.255.0**
  **临时设置 eth0 网卡的 IP 地址与子网掩码**
  **eth0 表示网卡，但是这个命令看不到网关**
  **ifconfig eth0:1 192.169.0.200**
  **设置另一个 IP 也生效**
- **ifconfig eth0:1 down 取消设置**

**setup 工具**
**红帽专有图形化工具设置 setup 设置 IP 地址**
**CentOS 7 为 nmtui 命令 最小化安装需要下载**

## 第二节 Linux 网络配置文件

**网卡信息文件**
**vim /etc/sysconfig/network-scripts/ifcfg-eth0**

- **DEVICE=eth0 网卡设备名**
- **BOOTPROTO=none 是否自动获取 IP(none，static，dhcp)**
- **HWADDR=00:0c:29:17:c4:09 MAC 地址**
- **NM_CONTROLLED=yes 是否可以由 Network Manager 图形管理工具托管**
- **ONBOOT=yes 是否随网络服务启动，eth0 生效**
- **TYPE=Ethernet 类型为以太网**
- **UUID=“” 唯一识别码**
- **IPADDR=192.168.0.252 IP 地址**
- **NETMASK=255.255.255.0 子网掩码**
- **GATEWAY=192.168.0.1 网关**
- **DNS1=202.106.0.20 DNS**
- **IPV6INIT=no IPv6 没有启用**
- **USERCTL=no 不允许非 root 用户控制此网卡**

**主机名文件**
**vim /etc/sysconfig/network**

- **NETWORKING=yes**
- **HOSTNAME=localhost.localdomain**
- **hostname [主机名]**
  **查看与临时设置主机名**

**DNS 配置文件**
**vim /etc/resolv.conf**
**nameserver DNS 服务**

## 第三节 常用网络命令(1)

**ifconfig 命令**

**hostname [主机名]**

**关闭与启动网卡**

- **ifup 网卡设备名
  启用该网卡**
- **ifdown 网卡设备名
  禁用该网卡**

**查询网络状态**
**netstat [选项]**

- **-t：列出 TCP 协议端口**

- **-u：列出 UDP 协议端口**

- **-n：不使用域名与服务名，而使用 IP 地址和端口号**

- **-l：仅列出在监听状态网络服务**

- -a：列出所有的网络连接

  netstat -rn

  - **-r：列出路由列表，功能与 route 命令一致**

**route 命令**
**route -n
查看路由列表(可以看到网关)**

**域名解析命令
nslookup [主机名或 IP]
进行域名与 IP 地址解析**
**nslookup**
**>server
查看本机 DNS 服务器**

**如果没有这个命令 使用 yum -y install bind-utils 下载**

**yum provides 命令 查询该命令在哪个软件包里**

## 第四节 常用网络命令(2)

**ping 命令**
**ping [选项] ip 或者域名**
**探测指定 IP 或域名的网络状况**

- **-c 次数：指定 Ping 包的次数**

**telnet 命令
telnet [域名或 ip] [端口]
远程管理与端口探测命令
telnet 192.168.0.252 80**

**traceroute
traceroute [选项] IP 或域名
路由跟踪命令
-n 使用 IP，不使用域名，速度更快**

**wget 命令
wget 下载网址
下载命令，下载的东西的具体目录，不能为上一级目录**

## 第五节 虚拟机网络参数配置

**vim /etc/sysconfig/network-scripts/ifcfg-eth0 把 ONBOOT=“no”改为 yes**

**利用 setup 配置 IP**

**在虚拟机中配置**

- **桥接：使用真实网卡，占用和本机想通的网段，还可以和同一个网段的局域网连接**
- **NAT 连接 使用 VMnet8 只能和本机通信，不能连局域网，虚拟机可以通过主机单向网络上的其他工作站，其他工作站不能访问虚拟机。**
- **Host-only 使用 VMnet1 此时虚拟机只能与虚拟机、主机互访。也就是不能上互联网**

**选择不同的模式就把 I 虚拟机 IP 设置为对应模式相同的网段，选择真实 IP 网段，VMnet1 IP 网段，VMnet8 IP 网段**

**VMware 虚拟机克隆镜像需要改 UUID**

- **vim /etc/sysconfig/network-scripts/ifcfg-eth0
  删除 MAC 地址行**
- **rm -rf /etc/udev/rules.d/70-persistent-net.rules
  删除网卡和 MAC 地址绑定文件**
- **重启系统**

# 第五章 SSH 远程管理服务

## 第一节 SSH 简介

**ssh(安全外壳协议)
SSH 为 Secure Shell 的缩写，SSH 为建立在应用层和传输层基础上的安全协议**

**SSH 端口**

- **SSH 端口：22**
- **Linux 中守护进程：sshd**

**安装服务：OpenSSH**

- **服务端主程序：/usr/sbin/sshd**
- **客户端主程序：/usr/bin/ssh**

**相关配置文件**

- **服务端配置文件：/etc/ssh/sshd_config**
- **客户端配置文件：/etc/ssh/ssh_config**

**网络监听**
**tcpdump -I 网卡文件名 -nnX port 21**

## 第二节 SSH 原理

**对称加密算法
采用单钥密码系统的加密方式，同一个密钥可以同时用作信息的加密和解密，这种加密方式称为对称加密，也称为单密钥加密**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531231653803.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**非对称加密算法
非对称加密算法(asymmetric cryptographic algorithm)又名“公开密钥加密算法”，非对称加密算法需要两个密钥：公开密钥(publickey)和私有密钥(privatekey)**![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531231710964.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**SSH 安全外壳协议**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531231743252.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**保护信息安全**

## 第三节 SSH 配置文件

**服务端配置文件：/etc/ssh/sshd_config**

- **Port 22 端口**
- **ListenAddress 0.0.0.0 监听的 IP**
- **Protocol 2 SSH 版本选择**
- **HostKey /etc/ssh/ssh_host_dsa_key 私钥保存位置**
- **ServerKeyBits 1024 私钥的位数**
- **SyslogFacility AUTH 日志记录 SSH 登陆情况**
- **LogLevel NFO 日志等级**
- **GSSAPIAuthentication yes GSSAPI 认证开启**

**在 Linux 远程管理 Linux 是，如果没有搭建 DNS 服务，这个建议关闭，把客户端的 GSSAPIAuthentication yes 改为 no，默认是 yes，取消注释即可，否则登陆会异常缓慢**

- **客户端配置文件：/etc/ssh/ssh_config**
  **安全设定部分**
- **PermitRootLogin yes 允许 root 的 ssh 登陆**
- **PubkeyAuthentication yes 是否使用公钥登陆**
- **AuthorizedKeysFile .ssh/authorized_keys
  公钥的保存位置**
- **PasswordAuthentication yes 允许使用密码验证登陆**
- **PermitEmptyPasswords no 不允许空密码登陆**

## 第四节 常用 SSH 命令

- **Linux 管理其他 Linux**
  **ssh 用户名@IP**
  **可以用来苹果 Mac 笔记本终端远登陆服务器**
- **scp**
  - **下载**
    **scp root@192.168.44.2:/root/test.txt .**
  - **上传
    scp -r /root/123 root@192.168.44.2/root**
    **例子：**
    **scp -r localfile.txt username@192.168.0.1:/home/username/**
    **其中，
    １）scp 是命令，-r 是参数
    ２）localfile.txt 是文件的路径和文件名
    ３）username 是服务器账号
    ４）192.168.0.1 是要上传的服务器 ip 地址
    ５）/home/username/是要拷入的文件夹路径
    scp -r /Users/yangyangyang/Desktop/1.txt root@47.95.5.171:/tmp**
- **Sftp 文件传输**
  **sftp root@192.168.4.2**
  - **ls 查看服务器端数据**
  - **cd 切换服务器端目录**
  - **lls 查看本地数据**
  - **lcd 切换本地目录**
  - **get 下载**
  - **put 上传**

## 第五节 SSH 连接工具

- **butty**
- **secureCRT**
- **Xshel**l
- **Mac 自带终端**

## 第六节 秘钥对登陆

**密钥对验证**![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531232358384.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**即安全又方便**

**步骤 1：**

- **Client 客户端：**
  **ssh-keygen -t rsa**
- **Sever 服务端：**
  把公钥上传到服务器端
  **cat id_rsa.pub >> /root/.ssh/authorized_keys
  chmod 600 /root/.ssh/authorized_keys**

**步骤 2:**
**修改服务器端 ssh 配置文件**

- **RSAAuthentication yes 开启 RSA 验证**
- **PubkeyAuthentication yes 是否使用公钥验证**
- **AuthorizedKeysFile .ssh/aurhorized_keys
  公钥保存位置**
- **PasswordAuthentication no 禁止使用密码验证登陆**

**步骤 3:**

- **服务器端关闭 SELinux 服务**
  **vim /etc/seliunx/config**
- **重启系统 reboot**
- **服务器端重启 ssh 服务**
  **service sshd restart**

# 第六章 DHCP 服务

## 第一节 DHCP 简介与原理

**DHCP 服务作用(Dynamic Host Configuration Protocol 动态主机配置协议)：**

- **为大量客户机自动分配地址，提供集中管理**
- **减轻管理和维护成本，提供网络配置效率**

**可分配的地址信息主要包括：**

- **网卡的 IP 地址，子网掩码**
- **对应的网络地址，广播地址**
- **默认的网关服务**
- **DNS 服务器地址**
- **引导文件，TFTP 服务器地址**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531232636131.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
  **DHCP 的原理**
- **1.客户端寻找服务器**
- **2.服务器提供地址信息(仅给 IP)**
- **3.接受广播(ARP 协议确认 IP 是否有效)**
- **4.服务器确认**
- **5.客户端重新登陆**
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200531232653346.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
  **DHCP 四次握手到了第四步的时候就已经完成
  第五步申请续租 IP，分为三次询问，如果 IP 租用时间为 8 天，就在 4 天，6 天，8 天头上进行三次续租申请，如果服务器都没回应，重新开始发送广播，如果服务器回应了，续租成功，第六步。**

## 第二节 DHCP 服务器相关文件

**安装 DHCP 服务器**
**RHEL6 的 DHCP 软件包**
**dhcp-4.1.1-31.P1.el6.i686.rpm**

**对应的端口**
**端口号：**

- **ipv4 udp67，udp68**
- **ipv6 udp546，udp547**

**相关文件**

- **服务名：dhcpd**
- **主配置文件：/etc/dhcp/dhcpd.conf**
- **模版文件：/usr/share/doc/dhcp-4.1.1/dhcpd.conf.sample**

**需要用模板文件覆盖主配置文件**

## 第三节 DHCP 配置文件

**全局配置**

- **option domain-name 设定所在 DNS 域**
- **option domain-name-servers 设置 DNS 服务器地址**
- **default-lease-time 设置默认租约时间，单位为秒**
- **max-lease-time 设置最大租约时间，单位为秒**
- **log-facility 指定日志设备**
- **ddns-update-style 设定 DNS 的更新方式**
- **authoritative 标识权威服务器**
- **log-facility local7 日志发送到 local7 日志服务中**

## 第四节 配置 DHCP 服务器

**服务器端配置**

- **修改配置文件**
- **重启 DHCP 服务**
- **service dhcpd restart**

**客户端配置**
**vim /etc/sysconfig/network-scripts/ifcfg-eth0**

- **修改以下文件**
  **DEVICE=eth0**
  **BOOTPROTO=dhcp**
  **ONBOOT=yes**
  **TYPE=Ethernet**
- **重启网络服务**
  **service network restart**

**服务器端租约文件**
**vim /var/lib/dhcpd/dhcpd.leases**

# 第七章 VSFTP 服务

## 第一节 FTP 简介与原理

**FTP 简介**
**FTP(File Transfer Protocol)中文称为“文件传输协议”。用于 Internet 上的控制文件的双向传输。**

- **“下载”文件就是从远程主机拷贝文件至自己的计算机上；**
- **“上传”文件就是将文件从自己的计算机中拷贝至远程主机上。**

**主动模式**
**主动模式：服务的从 20 端口主动向客户端发起连接**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200602200052344.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
**被动模式**
**被动模式：服务端在指定范围内某个端口被动等待客户端连接**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200602200159763.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
**主动模式不好实验，客户端不在管理员的控制范围之内，FTP 连接有可能被客户端的防火墙中断。**

**端口**
**FTP 连接端口**

- **控制连接：TCP 21，用于发送 FTP 命令信息**
- **数据连接： TCP 20，用于上传，下载数据**

## 第二节 FTP 相关文件

**常见的 FTP 服务器程序**

- **WIndows：IIS，Serv-U**
- **Linux：wu-ftpd(淘汰),Proftpd，Vsftpd(Very Secure FTP Daemon)**

**安装
Vsftpd-2.2.2-11.el6.i686.rpm**

**相关文件**

- **主配置文件**
- **/etc/vsftpd/vsftpd.conf**
- 用户控制列表文件
  - **/etc/vsftpd/ftpusers(黑名单，禁止登陆用户)**
  - **/etc/vsftpd/user_list(通过配置文件修改成白名单，只有这个上面的用户可以登陆)**
    **默认禁止 root 登陆，因为密码是明文显示，防止被抓包**

**FTP 相关用户**

- **匿名用户：**
  **anonymous 或 ftp**
- **本地用户**
  **使用 Linux 系统用户和密码**
- **虚拟用户**
  **管理员自定义的模拟用户**

**注意事项**

- **关闭防火墙**
- **关闭 SElinux**

## 第三节 配置文件详解

**/etc/vsftpd/vsftpd.conf**

- **anonymous_enable=YES 允许匿名用户登陆**
- **local_enable=YES 允许本地用户登陆**
- **write_enable=YES 允许本地用户上传**
- **local_umask=022 本地用户上传的默认权限**
- **dirmessage_enable=YES 用户进入目录时，显示.message 文件中信息**
- **message_file=.message 指定信息文件(登陆界面警告信息)**
- **xferlog_enable=YES 激活记录日志**
- **connect_from_port_20=YES 主动模式数据传输接口**
- **xferlog_std_format=YES 使用标准的 ftp 日志格式**
- **#ftpd_banner=Welcome to blah FTP service.
  登陆欢迎信息**
- **listen=NO 允许被监听**
- **pam_service_name=vsftpd 设置 PAM 外挂模块提供的认证服务所使用的配置文件名，即/etc/pam.d/vsftpd 文件**
- **userlist_enable=YES 用户限制登陆**
- **tcp_wrappers=YES 是否使用 tcp_wrappers 作为主机访问控制方式**

**常用全局配置**

- **listen_address=192.168.1.1 设置监听的 IP 地址**
- **listen_port 设置监听 FTP 服务的端口号**
- **download_enable=YES 是否允许下载文件**
- **max_per_ip=0 限制同一 IP 地址的并发连接数**

**被动模式**

- **pasv_enable=YES 开启被动模式**
- **pasv_min_port=24500 被动模式最小端口**
- **pasv_max_port=24600 被动模式最多端口**
  **大于 10000，小于 65535**

**常用安全配置**

- **#accept_timeout=60 被动模式，连接超时时间**
- **#connect_timeout=60 主动模式，连接超时时间**
- **#idle_session_timeout=600 600 秒没有任何操作就断开端口连接**
- **#data_connection_timeout=120 资料传输时，超过 500 秒没有完成，就断开传输**

## 第四节 客户端使用

- **使用命令行连接**
  **ftp IP**
  - **-help 获取帮助**
  - **-get 下载**
  - **-mget 下载一批文件**
  - **-put 上传**
  - **-mput 上传一批文件**
  - **-exit 退出**
- **不支持目录下载，不支持断点续传**
- **使用 Windows 对话框
  在文件夹路径搜索页面输入 ftp://192.168.44.3/输入服务器 IP
  支持目录下载，不支持断点续传**
- **使用第三方工具
  flashFXP
  只要工具支持就可以目录下载，断点续传**

## 第五节 匿名用户访问

**anonymous_enable=YES 允许匿名用户访问**
**#anon_upload_enable=YES 允许匿名用户上传**
**#anon_mkdir_write_enable=YES 允许匿名用户建立目录**
**#anon_umask 设置上传的默认文件权限(默认是 600)**

**匿名用户如果想上传，先修改配置文件，还得修改服务器中/var/ftp 目录中 pub 目录的权限，把这个目录的所有者修改为 ftp 用户 chown ftp /var/ftp/pub，不能直接把权限修改为 777，这样有很大的安全隐患**
注意事项：

- **默认上传目录：/var/ftp/pub**
- **如果允许上传，需要服务权限和系统目录权限同时允许**
- **vsfptd 服务的伪用户是 ftp**

## 第六节 本地用户访问

### 第一讲 本地用户基本设置

**本地用户基本配置**

- **local_enable=YES 允许本地用户登陆**
- **write_enable=YES 允许本地用户上传**
- **local_umask=022 本地用户上传的默认权限**
- **local_root=/var/ftp 设置本地用户的 FTP 根目录(注意目录权限)**
- **local_max_rate=0 限制最大传输速率(字节/秒)**

**把用户限制在家目录**

**chroot_local_user=YES**
**#开启用户目录限制(只有此行，把所有用户都限制在用户目录中)**
**在 CentOS 7 中用 yum 下载的 ftp，在这一步之后会报错
需要在配置文件中加入这条语句**
**allow_writeable_chroot=YES ##验证在 vsftpd.conf 中增加该项配置，vsftpd 服务无法正常启动。**

**给一些人畅通无阻的权限的时候，需要把下面三条配置文件全打开**

- **chroot_local_user=YES**
- **chroot_list_enable=YES**
- **chroot_listl_filer=/etc/vsftpd/chroot_list**
  **写入/etc/vsftpd/chroot_list 文件中的用户可以访问任何目录，其他用户限制在用户主目录**

### 第二讲 用户访问控制

**用户控制列表文件**

- **/etc/vsftpd/ftpusers(黑名单，禁止登陆用户)**
- **/etc/vsftpd/user_list(通过修改配置文件修改白名单，只有这个上面的用户可以登陆)**

**访问控制**
**这几条配置文件没有写**

- **userlist_enable=YES 开启用户访问控制**

- **userlist_deny=YES**

- userlist_file=/etc/vsftpd/user_list 写入/etc/vsftpd/user_list 文件中的用户不能访问 ftp 服务器，没有写入的用户可以访问(默认如此)

  我们需要改为

  - **userlist_enable=YES 开启用户访问控制**
  - **userlist_deny=NO**
  - **userlist_file=/etc/vsftpd/user_list 写入/etc/vsftpd/user_list 文件中的用户可以访问 ftp 服务器，没有写入的用户不能访问(默认如此)**

**不要把限制用户主目录和用户访问限制搞混**
**chroot_local_user=YES 用于把用户禁锢在主目录**
**uselist_enable=YES 用于用户访问控制**

## 第七节 虚拟用户访问

### 第一讲 配置虚拟用户访问

**配置虚拟用户登陆步骤**

- **添加虚拟用户口令**
- **生成虚拟用户口令认证文件**
- **编辑 vsftpd 的 PAM 认证文件**
- **建立本地映射用户并设置宿主目录权限**
- **修改配置文件**
- **重启 vsftpd 服务，并测试**
- **调整虚拟用户权限**

**添加虚拟用户口令**
**vim /etc/vsftpd/vuser.txt
cangls 用户名
123 密码
bols 用户名
123 密码**

**生成虚拟用户口令认证文件**

**yum -y install db4-utils 如果没安装口令认证命令，需要安装
或者使用 wget http://rpmfind.net/linux/centos/6.10/os/i386/Packages/db4-utils-4.7.25-22.el6.i686.rpm
安装
db_load -T -t hash -f /etc/vsftpd/vuser.txt
/etc/vsftpd/vuser.db 把文本文档转变为认证的数据库**

**编辑 vsftpd 的 PAM 认证文件**

**vim /etc/pam.d/vsftpd
auth required /lib/security/pam_userdb.so db=/etc/vsftpd/vuser
account required /lib/security/pam_userdb.so db=/etc/vsftpd/vuser
注释掉其他的行，加入 这两行即可
注释掉其他行，可以禁止本地用户登陆，因为本地用户登陆时的验证依然依赖这个文件**

**建立本地映射用户并设置宿主目录权限**

- **useradd -d /home/vftproot -s /sbin/nologin vuser
  此用户不需要登陆，只是映射用户
  用户名必须和下一步配置文件中一致**
- **chmod 755 /home/vftproot**

**修改配置文件
vim /etc/vsftpd/vsftpd.conf，加入：**

- **guest_enable=YES 允许虚拟用户登陆**
- **guest_username=vuser FTP 虚拟用户对应的系统用户**
- **pam_service_name=vsftpd PAM 认证文件(默认存在)**

**重启 vsftpd 服务，并测试
systemctl restart vsftpd.service
此时虚拟用户可以登陆，查看，下载，不能上传
默认上传文件的位置是宿主用户的家目录
权限使用的是匿名用户权限进行管理**

**调整虚拟用户权限**

- **vim /etc/vsftpd/vsftpd.conf**
- **anonymous_enable=NO 关闭匿名用户登陆，更加安全(不影响虚拟用户登陆)**
- **anon_upload_enable=YES**
- **anon_mkdir_write_enable=YES**
- **anon_other_write_enable=YES 给虚拟用户设置权限，允许所有虚拟用户上传**

### 第二讲 为每个虚拟用户建立自己的配置文件，单独定义权限

**可以给每个虚拟用户单独建立目录，并建立自己的配置文件。这样方便单独配置权限，并可以单独指定上传目录**

**修改配置文件
vim /etc/vsftpd/vsftpd.conf
user_config_dir=/etc/vsftpd/vusers_dir
指定保存虚拟用户配置文件的目录**

**手工建立目录
mkdir /etc/vsftpd/vusers_dir**

**为每个虚拟用户建立配置文件
vim /etc/vsftpd/vusers_dir/cangls
anon_upload_enable=YES
anon_mkdir_write_enable=YES
anon_other_write_enable=YES
允许此用户上传
local_root=/tmp/cangls
给 cangls 指定独立的上传目录**

**建立上传目录
mkdir /tmp/vcangls
chown vuser /tmp/vcangls/**

**如果不给 bols 指定单独的配置文件，则遵守主配置文件(/etc/vsftpd/vsftpd.conf)的权限**

**配置完效果如下：**

- **禁止匿名用户登陆(配置文件修改)**
- **禁止本地系统用户登陆(pam 文件修改)**
- **允许虚拟用户登陆(配置文件修改)**
- **cangls 的上传目录是/tmp/vcangls，并且允许查看，下载，上传**
- **bols 的上传目录是虚拟用户的默认目录/home/vftproot，只能查看，下载。但是不能上传**

## 第八节 总结

**我是阿里云忠实用户！！！，我会在里面穿插基于阿里云 CentOS7 FTP 服务器搭建！！！！**

**匿名用户很简单，下载好启动服务就可以，不过非常不安全，**
**需要把匿名用户登陆修改为 anonymous_enable=YES 默认即可，anon_upload_enable=YES 把匿名上传权限的注释去掉，
local_root=/var/ftp/，为匿名用户默认上传位置
修改/var/ftp/pub 权限。允许匿名用户写入文件
，chmod o+w /var/ftp/pub/ 或者 chown ftp /var/ftp/pub，修改所有者，不能直接把权限修改为 777，这样有很大的安全隐患**
**接着 systemctl restart vsftpd.service 重启服务即可
这种模式非常不安全，最好不要使用，这里只是介绍一下！！！！**

**所以只讲本地用户和虚拟用户**
**允许本地用户登陆**

- **下载好之后，先把防火墙，selinux 关闭，**
- **创建一个新用户 useradd test1
  设置密码 passed test1
  123456**
- **建立一个用户登陆 FTP 以后的家目录** **mkdir/var/ftp/test
  修改文件所有者，以及权限 chown -R test:test /var/ftp/test**
  **进入 /etc/vsftpd/vsftpd.conf 文件**
- **主动模式：**
  - **需要把匿名用户登陆修改为 anonymous_enable=NO ，local_enable=YES，write_enable=yes ，#listen_ipv6=YES，关闭监听 IPv6 sockets，listen=YES 开启监听 IPv4 sockets，**
  - **现在需要控制 FTP 用户登陆后的范围**
    **chroot_local_user=YES 全部用户被限制在主目录**
    **chroot_list_enable=YES 启用可以登陆用户白名单**
    **chroot_list_file=/etc/vsftpd/chroot_list 指定白名单用户列表文件，列表中的用户不被锁定在主目录**
    **allow_writeable_chroot=YES 新版 vsftp 在 vsftpd.conf 中增加该项配置，如果没有 vsftpd 服务无法正常启动。**
  - **local_root=/var/ftp/test 设置登录用户的默认目录**
- **被动模式**
  - **其他和主动模式一样，只需另外加入以下语句(阿里云环境需要开启被动模式，公网共享，连接不到服务器)**
    **pasv_enable=YES 开启被动模式
    pasv_address=<阿里云实例公网 IP> 阿里云实例公网 IP
    pasv_min_port= 设置被动模式下，建立数据传输可使用的端口范围的最小值
    pasv_max_port= 设置被动模式下，建立数据传输可使用的端口范围的最大值**
    建议把端口范围设置在一段比较高的范围内，例如 50000~50010，有助于提高访问 FTP 服务器的安全性。
- **创建白名单**
  **vim /etc/vsftpd/chroot_list 这里没有例外用户时，也必须创建 chroot_list 文件，内容可为空。
  必须输入:wq 保存并退出**
- **重启服务**
  **systemctl restart vsftpd.service**
- **阿里云服务器用户需要设置安全组**
  - **主动模式：**
    ￼![在这里插入图片描述](https://img-blog.csdnimg.cn/20200602202020694.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70#pic_center)
  - **被动模式：**
    ￼![在这里插入图片描述](https://img-blog.csdnimg.cn/20200602202036203.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
    **其中这里第二个安全组的端口范围是刚才文件中设置的**
- **限制用户登陆需要用到黑名单**
  **/etc/vsftpd/ftpusers
  里面加入登陆的用户名**
- **如果需要开启用户访问控制**
  **首先写入
  userlist_enable=YES 开启用户访问控制
  接着在下面写入以下语句
  userlist_deny=NO**
  **userlist_file=/etc/vsftpd/user_list 写入/etc/vsftpd/user_list 文件中的用户可以访问 ftp 服务器，没有写入的用户不能访问(默认如此)**
  **总结：**
  - **userlist_enable 是控制 userlist_deny 文件能否生效的开工，如果是 YES，就可以生效，如果为 NO，无论 userlist_deny 怎么调都没有作用。**
  - **当 userlist_enable=YES 时，userlist_deny=YES 时：user_list 是一个黑名单，即：所有出现在名单中的用户都会被拒绝登入**
  - **而当为 userlist_deny=NO 时：user_list 是一个白名单，即：只有出现在名单中的用户才会被准许登入(user_list 之外的用户都被拒绝登入)**
  - **另外需要特别提醒的是：使用白名单后，匿名用户将无法登入！除非显式在 user_list 中加入一行：anonymous**
- **虚拟用户：**
  - **1.添加虚拟用户口令
    vim /etc/vsftpd/vuser.txt
    用户名一行，密码一行
    保存退出**
  - **2.生成虚拟用户口令认证文件**
    - **yum -y install db4-utils 如果没安装口令认证命令，需要安装
      或者
      wget http://rpmfind.net/linux/centos/6.10/os/i386/Packages/db4-utils-4.7.25-22.el6.i686.rpm
      安装
      db_load -T -t hash -f /etc/vsftpd/vuser.txt /etc/vsftpd/vuser.db 把文本文档转变为认证的数据库**
    - **编辑 vsftpd 的 PAM 认证文件
      vim /etc/pam.d/vsftpd
      auth required /lib/security/pam_userdb.so db=/etc/vsftpd/vuser /lib64/security/pam_userdb.so
      account required /lib/security/pam_userdb.so db=/etc/vsftpd/vuser
      注释掉其他的行，加入这两行即可**
  - **3.建立本地映射用户并设置宿主目录权限**
    **useradd -d /home/vftproot -s /sbin/nologin vuser
    此用户不需要登陆，只是映射用户
    用户名必须和下一步配置文件中一致
    chmod 755 /home/vftproot**
  - **4.修改配置文件
    vim /etc/vsftpd/vsftpd.conf，加入：
    guest_enable=YES 开始虚拟用户
    guest_username=vuser FTP 虚拟用户对应的系统用户
    pam_service_name=vsftpd PAM 认证文件(默认存在)**
  - **5.重启 vsftpd 服务，并测试**
  - **6.调整虚拟用户权限
    vim /etc/vsftpd/vsftpd.conf
    anonymous_enable=NO 关闭匿名用户登陆，更加安全(不影响虚拟用户登陆)
    anon_upload_enable=YES
    anon_mkdir_write_enable=YES
    anon_other_write_enable=YES 给虚拟用户设置权限，允许所有虚拟用户上传**
  - **7.可以给每个虚拟用户单独建立目录，并建立自己的配置文件。这样方便单独配置权限，并可以单独指定上传目录
    修改配置文件
    vim /etc/vsftpd/vsftpd.conf
    user_config_dir=/etc/vsftpd/vusers_dir
    指定保存虚拟用户配置文件的目录
    手工建立目录
    mkdir /etc/vsftpd/vusers_dir
    为每个虚拟用户建立配置文件
    vim /etc/vsftpd/vusers_dir/cangls
    anon_upload_enable=YES
    anon_mkdir_write_enable=YES
    anon_other_write_enable=YES
    允许此用户上传
    local_root=/tmp/vcangls
    给 cangls 指定独立的上传目录
    建立上传目录
    mkdir /tmp/vcangls
    chown vuser /tmp/vcangls/**

**OK 这个服务器就搭建好啦！！！！大家辛苦了！**

# 第七章 Samba 服务

## 第一节 samba 简介

**数据共享的方法**

**Windows 中最常用的是“网上邻居”。网上邻居使用的文件系统是 CIFS(通用互联网文件系统)协议进行数据共享
Linux 中最常用的是 NFS 服务**

**Samba 的由来**

- **在 1991 年 Andrew Tridgell 为了实现 Unix 和 Windows 之间文件共享,开发了 SMB( Server Message Block,服务消息块)文件系统**
- **由于 SMB 无法注册,就取名为 Samba,热情的桑巴舞**

**Samba 与 NetBIOS 协议**

- **IBM 开发的 NetBIOS 协议是为了局域网内少数计算机进行通信的协议**
- **Samba 基于 NetBIOS 协议开发，所以可以和 Windows 通信,但是只能在局域网通信**
- **相当于在 Linux 上搭建 Windows 的网上邻居，减少病毒传播**

**Samba 主要应用**

- **文件共享**
- **打印服务器**
- **Samba 登陆时身份验证**
- **可以进行 Windows 的主机名解析**

**常见文件服务器的比较**

| 服务名称 | 使用范围   | 服务器端       | 客户端         | 局限性                                               |
| -------- | ---------- | -------------- | -------------- | ---------------------------------------------------- |
| FTP      | 内网和公网 | Windows、Linux | Windows、Linux | 无法直接在服务器端修改数据无法直接在服务器端修改数据 |
| Samba    | 内网       | Windows、Linux | Windows、Linux | 只能在内网使用                                       |
| NFS      | 内网和公网 | Linux          | Linux          | 只能 Linux 之间使用                                  |

## 第二节 Samba 安装与端口

**安装**

- **samba: 主服务包**
- **samba-client: 客户端**
- **samba-common: 通用工具**
- **samba4-libs: 库**
- **samba-winbind: Windows 域映射**
- **samba-winbind-clients: 域映射客户端**

**samba 的守护进程**

- **smbd:提供对服务器中文件、打印资源的共享访问 139 445**
- **nmbd:提供基于 NetBIOS 主机名称的解析 137 138**

**启动 samba 服务**

- **service smb start**
- **service nmb start**

## 第三节 Samba 相关文件

**常用文件**

- **/etc/samba/smb.conf #配置文件**
- **/etc/samba/lmhosts #对应 NetBOIS 名与主机的 IP 的文件,一般 samba 会自动搜索(只对本机生效) IP 与 主机名对应**
- **etc/samba/smbpasswd #samba 密码保存文件,默认不存在**
- **/etc/samba/smbusers #用户别名,用于适用不同操作系统中用户名习惯。需要配置文件中“username map”选项支持**

**常用命令**

**testparm 检测配置文件是否正确**

## 第四节 Samba 配置文件详解

**配置文件结构 smb .conf**

**Global Settings**

- **[global]: 全局设置**

**Share Definitions**

- **[homes]: 用户目录共享设置**
- **[printers]: 打印机共享设置**
- **[myshare]:自定义名称的共享目录设置**

**Global Settings**

- **workgroup:所在工作组名称(Windows)我的电脑里面有**
- **server string:服务器描述信息**
- **log file:日志文件位置**
- **max log size:日志文件的最大容量**
- security:安全级别,可用值如下:
  - **share: 不需要密码可以访问**
  - **user: 使用系统用户,samba 密码登陆**
  - **server: 由其他服务器提供认证**
  - **domain: 由域控制器提供认证**

**Share Definitions**

- **comment: 描述信息**
- **path: 共享的路径**
- **guest ok:允许所有人访问,等同于 public**
- **read only: 所有人只读**
- **writable: 所有人可写**
- **write list: 拥有写权限的用户列表**
- **browseable: 是否浏览可见**
- **valid users: 指定可以访问的用户**

**常见的变量**

- **%v: samba 的版本号**
- **%S: 任意用户可以登陆**
- **%m: client 端的 NetBIOS 主机名**
- **%L: 服务器端的 NetBIOS 主机名**
- **%u: 当前登陆的用户名**
- **%g: 当前登陆的用户组名**

## 第五节 基本使用

### 第一讲 share 权限访问

**配置文件修改**

**[global]**

- **workgroup = MYGROUP**
- **server string = Samba Server Lamp**
- **log file = /var/log/samba/log.%m**
- **max log size = 50**
- **security = share**
- **passdb backend = tdbsam**
  **这里需要注意：samba4 较之前的 SAMBA 3 有一个重大的变化是：security 不再支持 share，参数需要做调整**
  **原来：
  security=share
  现在:
  security=user
  map to guest =Bad User**

**Share Definitions**

- [movie]
  - **comment = study material**
  - **path = /study**
  - **browseable = yes**
  - **writable = yes**
  - **guest ok = yes**

**建立共享目录**

- **mkdir /study #建立共享目录**
- **chown nobody /study/ #赋予用户访问权限**
  **映射为 Linux 中的 nobody 用户**

**测试配置文件命令**

**testparm**

**重启 Samba 服务**

- **service smb restart**
- **service nmb restart**

### 第二讲 客户端的使用

**Windows 客户端访问**

- **网上邻居访问**
  **\192.168.44.4**
- Linux 客户端访问
  - **smbclient -L 192.168.44.4 #查看主机的共享资源**
  - **smbclient //192.168.44.4/movie #访问共享目录**
  - **如果有密码：
    smbclient //192.168.44.4/movie -U 用户名
    即可访问**

### 第三讲 user 级别访问

**配置文件修改**
**[global]**

- **workgroup = MYGROUP**
- **server string = Samba Server Lamp**
- **log file = /var/log/samba/log.%m**
- **max log size = 50**
- **security = user**
- **passdb backend = tdbsam**

**Share Definitions**

- [homes]
  - **comment = Home Directories**
  - **browseable = no**
  - **writable = yes**
- [pub]
  - **path = /public**
  - **browseable = no**
  - **writable = yes**
  - **write list = sc**
  - **valid users = sc 控制访问目录权限**

**建立共享目录与测试配置文件**

- **mkdir /public**
- **testparm**
- **重启 Samba 服务**
  **service smb restart**
  **service nmb restart**
- **添加 Samba 用户**
  **smbpasswd -a 系统用户名#注意,要给用户添加 samba 密码,该用户必须已经是系统用户**

**区别：**

- **FTP:使用系统用户,系统密码登陆**

- Samba:使用系统用户,samba 登陆
  权限设定

  如果要想访问与上传:

  - **Linux 系统权限需要允许**
  - **Samba 服务权限也需要允许**

### 第四讲 添加、删除、查看用户

**添加 samba 用户**

**注意用户必须已经是系统用户,才能添加为 samba 用户**

- **pdbedit -a -u 系统用户**
- **或 smbpasswd -a 系统用户名**

**删除 samba 用户**

- **pdbedit -x -u 系统用户名**
- **或 smbpasswd -x 系统用户**

**查看 samba 用户**

- **pebedit -L**

## 第六节 Samba 访问实验

**Samba 共享目录权限**

- **系统权限要对共享目录生效**
- **Samba 服务器权限也会对共享目录生效**
- **在实际的工作中,可能的共享目录众多, 用户众多,权限交叉。建议使用系统权限控制共享目录权限,而不是使用 samba 权限**

**例子：**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200603195339458.png)
**Samba 共享目录**

- [教学]
  - **path = /share/jiaoxue**
  - **browseable = yes**
  - **writable = yes**
- [教务]
  - **path = /share/jiaowu**
  - **browseable = yes**
  - **writable = yes**

**Linux 系统控制权限**

- **chmod 700 /jiaoxue /jiaowu**
  **教学的权限**
- **setfacl –m u:jx:rwx /share/jiaoxue**
- **setfacl –m u:jw:rx /share/jiaoxue**
- **setfacl –m u:xz:rx /share/jiaoxue**
- **setfacl –m u:wl:rx /share/jiaoxue**
  **教务的权限**
- **setfacl –m u:jw:rwx /share/jiaowu**
- **setfacl –m u:jx:rwx /share/jiaowu**
- **setfacl –m u:xz:rx /share/jiaowu**
- **setfacl –m u:tg:rx /share/jiaowu**

## 第七节 Samba 其他使用

**客户端挂载到本地使用**

**mount -t cifs -o username=用户 原路径 目标路径
例如:
mount -t cifs -o username=user1 //192.168.44.3/pub /test**

**开机自动挂载**

**vi /etc/fstab //192.168.44.3/pub /test cifs defaults,username=user1,password=456 1 2**

**Samba 别名**

**Linux 的用户如 root、nobody 在 Windows 中可能对应的用户是 administrator、guest 用户 Samba 是跨平台的,所以用别名让他们可以通用**

**别名文件**

**vi /etc/samba/smbusers
原名 = 别名 (多个别名 空格隔开) user1 = lamp1**

**配置文件中开启别名**

**vi /etc/samba/smb.conf
加入 username map = /etc/samba/smbusers
重启,就可以使用别名访问 samba 了**

**总结**
**Vsftp:**

- **内网 外网服务器: Windows Linux**
- **客户端: Windows Linux**
- **用户: 系统用户,系统密码**
- **缺点:不能直接在服务器上修改文件**

**Samba**

- **内网使用服务器: Windows Linux**
- **客户端: Windows Linux**
- **用户: 系统用户,samba 密码**
- **缺点: 只能在内网使用**

# 第八章 NFS 服务

## 第一节 NFS 简介

**什么是 NFS**
**NFS(Network File System):**
**NFS 可以让客户端把服务器的共享目录,挂载到本机使用,就像使用本机分区一样,使用非常方便**

**RPC 服务**
**RPC(远程调用)服务:**
**NFS 是被 RPC 服务管理的,所以必须安装 RPC 的主程序 rpcbind**

**NFS 端口**

- **NFS 端口:2049**
- **RPC 端口:111**
- **NFS daemon 端口:随机**

**和其他文件服务器对比(SFU Windows 客户端搭建)**

| 服务名称 | 使用范围   | 服务器端       | 客户端          | 局限性                                        |
| -------- | ---------- | -------------- | --------------- | --------------------------------------------- |
| FTP      | 内网和公网 | Windows、Linux | Windows、Linux  | 无法直接在服务器端修改数据                    |
| Samba    | 内网       | Windows、Linux | Windows、Linux  | 只能在内网使用                                |
| NFS      | 内网和公网 | Linux          | Linux (Windows) | 客户端需要挂载使用,对普通用户有一定的技术要求 |

## 第二节 NFS 权限说明

**权限说明**

- **Linux 系统目录权限会生效**
- **NFS 服务共享权限也会生效**

**用户身份映射(系统权限)**
**NFS 没有用户登陆认证机制,所以客户端登陆到服务器之后,会把客户端的身份映射到服务器端。就会出现一下四种可能:**

- **第一种可能**
  - **client 和 server 上刚好有相同的帐号和用户组(用户名和 UID 都要相同)**
  - **client 上用户可以在 server 上按照用户权限使用文件**
- **第二种可能**
  - **当 clinet 和 server 上拥有相同的 UID,但是用户名不同,
    假设 clinet 上有用户 aa(UID:500),server 上有用户 bb(UID:500)。在 clinet 在使用 server 共享目录时,身份识别为 bb(因为 Linux 权限绑定在 UID 上)
    此种情况尽量避免出现,容易逻辑混乱**
- **第三种可能**
  - **server 上没有 clinet 的 UID**
    **clinet 用户访问 server 时,server 上没有此 UID, 则把此用户自动转变为 nfsnobody(UID: 65534)用户**
- **第四种可能**
  - **clinet 上是 root 使用共享**
    **默认把 root 也转变为 nfsnobody。服务器端可以修改配置文件,允许 root 访问 nfs 服务器**

## 第三节 NFS 服务器端设置

**服务器端安装**
**默认已经安装,如果需要手工安装**
**– NFS 主程序:nfs-utils**
**– RPC 主程序:rpcbind(旧版本 portmap)**

**相关文件**
**配置文件: /etc/exports**

**服务器端管理**
**默认已经启动,如果需要手工启动:**

- **service nfs start**
- **service rpcbind start**

**守护进程**
**ps aux | grep -E "nfs|rpc”**
**– rpc.rquotad:NFS 配额
– rpc.mountd:处理客户端挂载
– nfsd:NFS 守护进程
– rpcbind:RPC 守护进程**

**RPC 服务注册情况**
**rpcinfo -p IP 或主机名**

## 第四节 NFS 服务器端配置文件

**配置文件
/etc/exports
共享目录 客户端(权限)
多个共享，空格隔开**

**可以识别的客户端**

- **指定 IP:192.168.44.4**
- **指定网段: 192.168.44.0/24 192.168.44.0/255.255.255.0**
- **指定主机名:www.itxdl.cn**
- **所有主机:\***

**常用权限**

- **rw: 读写**
- **ro: 只读**
- **all_squash:不论登陆是谁,都压缩为匿名用户 nfsnobody**
- **root_squash:如果登陆的是 root,压缩为 nfsnobody**
- **no_root_squash:允许 root 身份登陆,不推荐 anonuid:把所有的登陆用户,不再压缩为 nfsnobody 用户,而是压缩为指定 uid 用户**
- **sync:将数据同步写入内存缓冲区与磁盘中,效率低, 但可以保证数据的一致性;**
- **async:将数据先保存在内存缓冲区中,必要时才写入磁盘**
- **举例 1**
  **任何人可以访问,允许 root 访问**
  **vi /etc/exports**
  **/home/test \*(rw,no_root_squash)**
- **举例 2**
  **同时共享多个目录**
  **vi /etc/exports**
  **/home/test 192.168.44.3(rw,no_root_squash) \*(ro)**
  **/home/soft 192.168.44.0/24(rw)**
- **举例 3**
  **匿名用户访问**
  **vi /etc/exports**
  **/home/soft 192.168.44.0/24(rw,all_squash,anonuid=600)**

**常用命令**
**exportfs 选项
选项:
-a: 按照配置文件挂载/卸载所有目录
-r: 重新挂载
-u: 卸载
-v: 显示详细信息**

**exportfs 举例**

- **exportfs -auv #全部卸载所有目录**
- **exportfs -arv #重新挂载所有目录,不用重启 NFS 服务**
  **这种是为了不用重启服务，就可以使修改的配置文件生效，而且不会把正在登陆的用户踢掉**

**showmount 命令**

- **showmount 查看共享目录**
- **showmount -e IP 或主机名**
  **选项:
  -e: 查看某个主机的共享目录**

## 第五节 NFS 客户端使用

**客户端需要启动的服务
需要启动 rpcbind,默认已经启
手工启动命令:
service rpcbind start**

**查看服务器端共享目录
showmount -e 192.168.44.3 #指定服务器端 IP 即可**

**把服务器共享目录挂载到本地**

- **mkdir /home/client #建立挂载点**
- **mount -t nfs 192.168.44.3:/home/soft /home/client #挂载到本地使用**
- **mount #查看挂载**

**卸载
umount /home/client**

**开机自动挂载
vim /etc/fstab
192.168.44.3:/home/test /home/client nfs defaults 0 0**

## 第六节 权限实验

- **第一种可能**
  **client 和 server 上刚好有相同的帐号和用户组(用户名和 UID 都要相同)
  client 上用户可以在 server 上按照用户权限使用文件**
  **服务器:
  – [root@localhost ~]# useradd -u 600 test1
  – [root@localhost ~]# passwd test1**
  **客户端:
  – 一样的用户与 UID
  – [root@localhost test]# su - test1
  – [test1@localhost ~]$ cd /home/test
  – [test1@localhost test]$ touch cde
  – [test1@localhost test]$ ll cde
  – -rw-rw-r-- 1 test1 test1 0 11 月 14 2016 cde**
- **第二种可能**
  **当 clinet 和 server 上拥有相同的 UID,但是用户名不同,
  假设 clinet 上有用户 aa(UID:500),server 上有用户 bb(UID:500)。在 clinet 在使用 server 共享目录时,身份识别为 bb(因为 Linux 权限绑定在 UID 上)
  此种情况尽量避免出现,容易逻辑混乱**
  **服务器端:
  – 用户名:test1 UID:600**
  **客户端:
  – [root@localhost ~]# useradd -u 600 user1
  – [root@localhost ~]# su - user1
  – [user1@localhost ~]$ cd /home/test/
  – [user1@localhost test]$ touch def
  – [user1@localhost test]$ ll def
  – -rw-rw-r-- 1 user1 user1 0 11 月 14 2016 def
  服务器端:
  – [root@localhost ~]# ll /tmp/def
  – -rw-rw-r-- 1 test1 test1 0 11 月 14 04:40 /tmp/def**
- **第三种可能
  server 上没有 clinet 的 UID
  clinet 用户访问 server 时,server 上没有此 UID, 则把此用户自动转变为 nfsnobody(UID: 65534)用户**
  **服务器端:
  – 没有 UID 为 700 的用户**
  **客户端:
  – [root@localhost ~]# useradd -u 700 test2
  – [test2@localhost test]$ touch qwe
  – [test2@localhost test]$ ll qwe
  – -rw-rw-r-- 1 nobody nobody 0 11 月 14 2016 qwe**
- **第四种可能
  clinet 上是 root 使用共享
  默认把 root 也转变为 nfsnobody。服务器端可以修改配置文件,允许 root 访问 nfs 服务器**
  **服务器端:
  – 不允许 root 访问**
  **客户端:
  – [root@localhost ~]# cd /home/test
  – [root@localhost test]# touch asd
  – [root@localhost test]# ll asd
  – -rw-r–r-- 1 nfsnobody nfsnobody 0 11 月 14 2016 asd
  服务器端:
  – 允许 root 访问
  – [root@localhost ~]# vi /etc/exports
  – /tmp 192.168.44.4(rw,no_root_squash)
  – [root@localhost ~]# exportfs -auv
  – [root@localhost ~]# exportfs -arv
  客户端:
  – [root@localhost ~]# cd /home/test
  – [root@localhost test]# touch zxc
  – [root@localhost test]# ll zxc
  – -rw-r–r-- 1 root root 0 11 月 14 2016 zxc**

## 第七节 总结

**简单写一下关于 NFS 服务器搭建中的一些问题**

**首先要知道在我们搭建这些服务器的时候，对我们最大的阻碍就是 SELinux 与防火墙，我们先从这两点开始说起
SE Linux 是美国国家安全局开发的 Linux 安全模块，它在本来已经很安全的 Linux 上，凌驾于 root 权限之上，设置了很多额外的条条框框；**
**如果你了解这些条条框框，那还好；但如果不了解，那 SELinux 可能并没有帮什么忙，却给你带来了很多不确定因素，通俗来讲， 把它关了就好，在刚开始学 Linux 这个东西对我们没有一点作用，所以我们在搭建这些 Samba，FTP，NFS 服务的时候，首先就要做到把它先关闭，**
**vim /etc/sysconfig/selinux 命令进入 SELinux 的配置文件，把 SELINUX=disabled 这一句设置好了，他就干干净净的从你的世界消失了，再也没有他的打扰，接下来第二个困扰便是防火墙，这个东西更加头疼，如果你是虚拟机搭建 ，没有太大的安全隐患，直接 iptables -F 或者 systemctl stop firewalld.service 关了就好，但是如果在比如阿里云，华为云，腾讯云这上面的服务器来说，就有一定的安全隐患了，而且直接把它关了也很令人不爽，你如果到了企业工作，不能也直接把企业的防火墙关了吧。**
**所以说，我们还是老老实实的学一下关于防火墙的一点东西吧，为什么我在刚才说需要 iptables -F 或者 systemctl stop firewalld.service 这么两个命令呢？这是因为在 Linux 中有两种防火墙软件，CentOS7.0 以上使用的是 firewall，CentOS7.0 以下使用的是 iptables，有一个就好。**
**对于防火墙，我们该如何理解呢？这样举个例子吧，有一个男孩喜欢一个女孩，那她怎么才能追到这个女孩呢？很简单，让这个女孩喜欢他就好了，这就是男追女，隔座山；女追男，隔层纱的意思，我们把服务器端(Server)比作男孩，客户端(Client)比作女孩，他们之间要想相互通信，就得有连接啊，但是很遗憾，防火墙就像女孩的父亲一样隔在你和女孩之间，如果你一味的去追女孩，因为有防火墙的存在，你是不可能建立联系的，这就好比 FTP 服务器的主动模式与被动模式，我们可以控制自己服务器的防火墙，但是却不能控制客户端的防火墙，而客户端也有可能对防火墙不太理解，你不能强迫人家说：你要想连接，去吧防火墙关了，或者开端口。这些对于非专业人士来说，是很难办到的。所以我们就有让我们自己的服务器开启相应的端口，让我们被动，我们去接客户端传来连接请求。这样女孩的父亲也管不住了，你们就可以完美的建立联系。所以，对于阿里云这些服务器，他的公网 IP 会映射到内网上，也就是你的私有 IP，阿里云的环境需要开启被动模式，公网共享，否则连接不到服务器。**
**对于端口，他就是好像在防火墙上的一扇窗，通过特定开放的端口就可以进行连接，所以你也可以使用公网 IP 加端口(47.95.1.123:21)这样来链接，我们在服务器端需要把服务的每一个端口都开启，对于阿里云用户来说，就是在阿里云控制台里面的设置安全组，把里面的所需的服务的端口都打开。这样才能建立连接，FTP 主动模式需要开 20，21 两个端口，服务从 20 端口主动向客户端发起连接，而被动模式你也需要另外在加开端口，比如 50000-50010，这是一个范围，服务端在指定范围内某个端口被动等待客户端连接。这便是被动模式。
而对于 NFS 这种服务，他是先通过 RPC 服务的注册，才能正常工作，NFS 是被 RPC 服务管理的,所以必须安装 RPC 的主程序 rpcbind，才能允许 NFS，在开启服务的时候也需要先开启 rpcbind 服务，才能开启 nfs 服务，而 NFS 主进程端口是 2049，RPC 端口是 111 ，但是 NFS 还有许多子进程，这便是通过 RPC 注册的 NFS daemon 端口，他的端口号是随机的，因为端口随机，如果关了防火墙还好说，啥都可以过，但是一旦开启呢？你总不能在使用服务之前先查一下端口号再来配置防火墙规则吧，这是不现实的，所以我们就需要把随机端口给固定了。可以更改配置文件固定 NFS 服务相关端口。这样设置固定端口以后即便重启机器也很方便挂载，如果不设置固定端口，机器或服务重启后之前添加的 iptables 规则就失效了，当然你也可以 systemctl stop iptables.service，systemctl stop firewalld.service 来临时关闭防火墙，重启之后就会失效。那如何配置 nfs 的效果端口配置文件呢？**
**vim /etc/sysconfig/nfs**
**RQUOTAD_PORT=30001
LOCKD_TCPPORT=30002
LOCKD_UDPPORT=30002
MOUNTD_PORT=30003
STATD_PORT=30004
添加这几条规则，就把端口限制了。
systemctl restart rpcbind.service
systemctl restart nfs.service
重启 rpc，nfs 服务，配置生效
rpcinfo -p 查看服务端口
你会发现，全都是你设定的端口号
接下来就需要给防火墙设置规则，开窗户了：
添加 iptables 规则**
**iptables -A INPUT -s 192.168.214.0/24 -p tcp –dport 111 -j ACCEPT
iptables -A INPUT -s 192.168.214.0/24 -p udp –dport 111 -j ACCEPT
iptables -A INPUT -s 192.168.214.0/24 -p tcp –dport 2049 -j ACCEPT
iptables -A INPUT -s 192.168.214.0/24 -p udp –dport 2049 -j ACCEPT
iptables -A INPUT -s 192.168.214.0/24 -p tcp –dport 30001:30004 -j ACCEPT
iptables -A INPUT -s 192.168.214.0/24 -p udp –dport 30001:30004 -j ACCEPT**
**保存退出后，重启防火墙**
**systemctl restart iptables.service
添加 firewall 规则
firewall-cmd --permanent --add-port=111/tcp
firewall-cmd --permanent --add-port=111/udp
firewall-cmd --permanent --add-port=2049/tcp
firewall-cmd --permanent --add-port=2049/udp
firewall-cmd --permanent --add-port=30001/tcp
firewall-cmd --permanent --add-port=30001/udp
firewall-cmd --permanent --add-port=30002/tcp
firewall-cmd --permanent --add-port=30002/udp
firewall-cmd --permanent --add-port=30003/tcp
firewall-cmd --permanent --add-port=30003/udp
firewall-cmd --permanent --add-port=30004/tcp
firewall-cmd --permanent --add-port=30004/udp
保存退出后，重启防火墙**
**systemctl restart iptables.service
重启服务器
reboot
重启之后执行"rpcinfo -p"命令
就发现全部都 OK 了
这两种选其一即可，看你电脑是开启的哪种防火墙**

**这些都弄完之后，就可以配置 NFS 的配置文件了**

**如果还有报错的话，学会使用/var/log/secure 日志文件，查看报错原因，再进行谷歌。**

# 第九章 DNS 服务

## 第一节 DNS 服务（一）

**结构**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607141830986.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**域名系统**
**DNS 系统的作用:**

- **正向解析:根据主机名称(域名)查找对应的 IP 地址**
- **反向解析:根据 IP 地址查找对应的主机域名**

**域名：**
**域名：baidu.com
主机名：www
www.baidu.com 完整的 FQDN 名称**

**IANA** **军方把 DNS 的分配权限转给 ICANN** **国际联盟组织**

**DNS 发展阶段：**

- 第一阶段，通过文件维护

  - **/etc/hosts 主机映射文件**

- 第二阶段，通过服务器维护

  - **DNS Server**
  - **比文件维护效率更高，但是服务器压力过高**

- 第三阶段，通过分布式存储服务器

  - 搭建更多的服务器分担压力

    缺点：

    - **时间长，上级服务器只需要维护直属下级服务器**
    - **管理相对麻烦，systemctl status iptables.service 上级服务器只需要维护直属下级服务器**

  - 优点：

    - **管理简单，允许重名，从而减低网络中域名的复杂程度**

**域名结构**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607141858619.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**www.baidu.com.:80/index.html
这应该是全名，浏览器帮我们简化掉了.:80/index.html 这一部分，其中的.为根域**

**比如：www.sina.com.cn.
根域；.
一级域；.cn
二级域；.com.cn
三级域；sina.com.cn
主机名；www**

**域名解析的工作模式：**

- **递归查询：压力在服务器端**
- **迭代查询：按照域名等级挨个询问服务器，服务器压力大大减小，压力在客户端**

**根 DNS 的特点：**

- **全球有 13 台根 NDS 服务器**
  **A INTERNIC.NET(美国,弗吉尼亚州) 198.41.0.4
  B 美国信息科学研究所(美国,加利弗尼亚州) 128.9.0.107
  C PSINet 公司(美国,弗吉尼亚州) 192.33.4.12
  D 马里兰大学(美国马里兰州) 128.8.10.90
  E 美国航空航天管理局[NASA](美国加利弗尼亚州) 192.203.230.10
  F 因特网软件联盟(美国加利弗尼亚州) 192.5.5.241
  G 美国国防部网络信息中心(美国弗吉尼亚州) 192.112.36.4
  H 美国陆军研究所(美国马里兰州) 128.63.2.53
  I Autonomica 公司(瑞典,斯德哥尔摩)192.36.148.17
  J VeriSign 公司(美国,弗吉尼亚州) 192.58.128.30
  K RIPE NCC(英国,伦敦) 193.0.14.129
  L IANA (美国,弗吉尼亚州) 198.32.64.12
  M 日本 WIDE 项目 202.12.27.33**
- **不支持递归查询**

**DNS 安装**
**安装：
软件 BIND
安装软件名：bind
服务开启名：named
端口：53
配置： 1.修改主配置文件 2.修改区域文件 3.修改解析数据文件**

**BIND(Berkeley Internet Name Daemon)
伯克利 Internet 域名服务
官方站点:https://www.isc.org/
软件包:bind-9.3.3-7.el5.i386.rpm
服务名:named
端口号:53
主配置文件:/etc/named.conf
保存 DNS 解析记录的数据文件位于: /var/named/**

- **主域名服务器: 特定 DNS 区域的官方服务器,具有唯一性负责维护该区域内所有域名->IP 地址的映射记录**
- **从域名服务器: 也称为 辅助域名服务器其维护的域名->IP 地址记录 来源于主域名服务器**

**DNS 服务器**

**主 DNS 服务器**

**配置主配置文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607141920151.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**配置区域文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607141934583.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**配置正向数据文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607141945858.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**配置反向数据文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/202006071420160.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**从 DNS 服务器** \* **减轻主服务器的压力** \* **数据从主服务器复制**

**修改主 DNS 配置文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142033497.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**配置从 DNS 配置文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142047772.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**配置从 DNS 区域文件**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142055893.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)

## 第一节 DNS 服务（二）

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142106799.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**外网访问内网服务器的全过程：**
**S：200.200.200.2
D：200.200.200.200**
**通过端口映射 NET 策略**
**S：200.200.200.2
D：1.1.1.2
回信：
S：1.1.1.2
D：200.200.200.2
S：200.200.200.200
D：200.200.200.2**

**内网访问内网服务器的全过程：
S：1.1.1.1
D：200.200.200.200**
**S：1.1.1.1
D：1.1.1.2
S：1.1.1.2
D：1.1.1.1**
**原地址与目标地址不匹配，基于 TCP 的协议无法建立连接。**
**而解决这种回流现象的问题至关重要，这便是分离解析**

**分离解析 DNS 服务器
作用: 将相同的域名解析为不同的 IP 地址**
**实验环境:**
**三台虚拟机:
1、第一个网段测试机
2、网关、分离解析 DNS
3、第二个网段测试机**

**代码**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142121553.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142128196.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607142134462.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/2020060714214223.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)

# 第十章 Postfix 服务器配置

**本章主要说这几点：**

- **Postfix 概念与原理**
- **Postfix 配置文件解析**
- **邮件服务器端配置与客户端使用**
- **dovecot 的使用**
- **extmail 和 extman**

**邮件服务器概述:**

**邮件服务器概念: 电子邮件服务器是处理邮件交换的软硬件设施的总称,包括电子邮件程序、电子邮件箱等。它是为用户提供基于 E-mail 服务的电子邮件系统,人们通过访问服务器实现邮件的交换。**
**常见的邮件服务器:**

- **Sendmail、Qmail、Postfix**
- **Exchange、 Notes/Domino**
- **Coremail**

**在初期没有邮件服务器诞生：**

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143105745.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**邮件服务器出现之后：**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143118187.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**邮件系统角色:**

- **MUA(邮件用户代理)**
- **MTA(邮件传输代理)**
- **MDA(邮件分发代理)**
- **MRA (邮件检索代理)**

**邮件应用协议:**

- **SMTP,简单邮件传输协议,TCP 25 端口,加密时使用 TCP 465 端口(发送邮件)**
- **POP3,第 3 版邮局协议,TCP 110 端口 , 加密时使用 995 端口(收取邮件)**
- **IMAP4,第 4 版互联网消息访问协议,TCP 143 端口 ,加密时使用 993 端口(收取邮件)**

**邮件服务器的原理:**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143129369.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143136164.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**实验搭建邮件服务器**

- **第一步搭建邮件域
  使用一台虚拟机搭建 DNS 服务器，解析邮件域名**
- **第二步下载邮件服务器
  默认安装，没有的话**
  **yum -y install postfix**
- 第三部配置配置文件
  Vim main.cf
  - **:set nu 设置行号
    在第 75 行 myhostname = host.domain.tld
    需要修改，改为你 DNS 解析的域名**
    **第 83 行 #mydomain = domain.tld
    需要修改，设置邮件域
    比如
    12345@yangyang.com 其中 yangyang.com 就为邮件域**
    **在第 98 行 #myorigin = $myhostname
    第 99 行 #myorigin = $mydomain
    用户在发送时自动补全邮件域名称**
    **113 行去掉注释
    116 行加上注释**
    **164 行用来设置接收方信息，设置接收哪类邮件
    比如**
    **a@mail.yangyang.com
    a@localhost.yangyang.com
    a@localhost
    a@yangyang.com
    只接受这四类邮件，其他一律拒绝掉**
    **221 行 unknown_local_recipient_reject_code = 550
    用来设置拒绝不存在的本地账号，或者使用者不明的账号，拒绝掉之后返回 550 错误代码
    386 行 alias_maps = hash:/etc/aliases
    用来定义用户别名 用来用户转发，需要数据库支持，数据库文件存放在 397 行
    419 行 home_mailbox = Maildir/ 邮箱位置需要启动**
  - **一切准备好之后。在本机下载 telnet，进行邮箱连接**
    **yum -y install telent**
    **首先关闭防火墙**
    **接着输入
    telnet mail.yangyang.com 25 进行连接
    mail.yangyang.com 是域名，25 为 portfix 对应开启的端口号**
  - **链接好之后，先声明邮件服务器地址
    helo mail.yangyang.com
    然后声明发送方
    mail from:pp@mail.yangyang.com
    声明接收放
    rcpt to:gg@mail.yangyang.com
    声明邮件正文
    data
    写正文。。。。
    如果不想写了
    另起一行写. 接着回车就好
    quit 退出**
  - **接着就可以在收件人的家目录里发现一个 Maildir/目录 进去里面，有 new 目录，再进去就会发现有发的邮件啦**

**dovecot 是一个 MRA 是一种收邮件的软件
可以下载用 telnet 连接,也可以进行接收邮件
下载：yum -y install dovecot dovecot-deve1
telnet mail.yangyang.com 110
进行连接
连接好之后
输入
user 用名 回车
pass 密码 回车
就可以登陆啦
然后输入 list 就可以看到里面的内容了
查看内容 使用
retr 邮件编号
即可查看
quit 退出**

**企业级邮件服务器搭建**

- **第一步：
  先把 DNS 里添加解析 extmail.org 的语句**
- **第二步：
  安装数据库
  yum -y install mysql mysql-server**
  **启动
  systemctl start mysql.service**
- **第三步：
  创建 /var/www/extsuite 文件夹
  下载 extmail 与 extman
  https://474b.com/file/4095383-385096033
  https://474b.com/file/4095383-385096042
  解压到/var/www/extsuite/ 文件夹**
- **第四步：
  **进入 extman-1.1 文件夹的 docs/文件夹里面
  把 init.sql 的管理员用户密码修改了，在最后一行 root，修改为：\*\*
  **INSERT INTO `manager` VALUES (‘root@extmail.org’,‘123456’,‘admin’,‘root’,‘Super User’,‘my question’,‘my answer’,‘0’,‘2007-02-14 15:10:04’,‘2010-11-08’,1);**
  **即可，保存退出后接着输入
  mysql < extmail.sql**
  **这一步会报错：
  ERROR 1064 (42000) at line 50: You have an error in your SQL syntax; check the manual that corresponds to your MariaDB server version for the right syntax to use near ‘TYPE=MyISAM COMMENT=‘Ext/Webman - Admin Accounts’’ at line 15**
  **这是因为 TYPE 是老版本的参数了，mysql5.5 版本之后移除了 TYPE 选项并推荐使用 ENNIGE 代替
  所以需要将里面所有的 TYPE=MyISAM 改为 ENGINE=MyISAM
  你肯定会发现还有错误：**
  **ERROR 1364 (HY000) at line 31: Field ‘ssl_cipher’ doesn’t have a default value
  出现错误的原因是 mysql 默认配置严格模式，该模式禁止通过 insert 的方式直接修改 mysql 库中的 user 表进行添加新用户。
  进入/etc/my.cnf 文件修改：sql-mode=STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
  为：
  sql-mode=NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
  就可以了**
  如果还有报错：
  **ERROR 1007 (HY000) at line 46: Can’t create database ‘extmail’; database exists
  进入：/var/lib/mysql
  把 extmail 目录删除**
  **回到/var/www/extsuite/extman/docs/重新输入 mysql < extmail.sql
  就 OK 了**
  **接着输入：
  mysql < init.sql 把密码信息导入数据库**
  **遇到错误：
  ERROR 1062 (23000) at line 5: Duplicate entry ‘support@extmail.org’ for key ‘PRIMARY’
  使用： vim init.sql
  把 INSERT 改为 INSERT IGNORE**
- **第五步：
  把模版拷贝到主目录下
  cp mysql_virtual_domains_maps.cf mysql_virtual_alias_maps.cf mysql_virtual_mailbox_maps.cf /etc/postfix/**
  **创建虚拟用户
  useradd -u 600 vmail
  接着
  vim /etc/postfix/main.cf
  在文件最后加入：
  virtual_mailbox_base = /home/vmail**
  **virtual_uid_maps = static:600
  virtule_gid_maps = static:600
  virtual_alias_maps = mysql:/etc/postfix/mysql_virtual_alias_maps.cf
  virtual_mailbox_domains = mysql:/etc/postfix/mysql_virtual_domains_maps.cf
  virtual_mailbox_maps = mysql:/etc/postfix/mysql_virtual_mailbox_maps.cf
  保存退出**
  **重启邮件服务**
  **systemctl restart postfix.service**

**咱们来做个小测试吧**
**输入：
echo “hi” | mail -s test support@extmail.org
你会在刚刚配置文件里加入的收件目录里：
cd /home/vmail
进去之后你会惊喜的发现有一个 extmail.org 的文件夹，哈哈哈哈哈哈哈！这就代表你成功啦，再进去会发现：postmaster，再进去
就是 Maildir 了，再进去就可以看到我们熟悉的 new 目录啦，存放邮件的地方，再进去就可以查看邮件啦！**

- **第六步：
  接下来就需要我们配置 dovecot 模块了，把整个服务架起来
  yum -y install dovecot dovecot-mysql
  下载好之后，进入：
  cd /etc/dovecot/conf.d/
  第一个我们来修改
  10-mail.conf 这个文件,用来设置邮件的配置
  找到 mail_location = maildir:这一行，把它修改：
  mail_location = maildir:/home/vmail/%d/%n/Maildir
  这是用来设置我们接收邮件的地址
  %d 的意思配置文件上一行有解释
  接着把 first_valid_uid =设置为 600
  保存退出
  修改第二个文件：
  10-auth.conf 这个是设置认证方式的
  找到这一行!include auth-sql.conf.ext 把注释去掉，通过数据库认证
  保存退出
  修改第三个文件：
  先返回上一级**
  **cd …
  执行
  cp /usr/share/doc/dovecot-2.2.36/example-config/dovecot-sql.conf.ext .
  把这个文件拷贝到当前目录**
  **打开这个文件修改：**
  - **1.#driver = 这句话的注释去了后面加 mysql
    driver = mysql**
  - **2.找到 Examples:这一项，下面的第二行，把值修改为一下内容：
    connect = host=localhost dbname=extmail user=extmail password=extmail**
  - **3.#default_pass_scheme = MD5，去掉注释**
  - **4.找到这一项：#password_query =
    去掉三行注释，改为这样
    password_query =
    SELECT username, domain, password
    FROM mailbox WHERE username = ‘%u’ AND domain = '%d’**
  - **5.找到这一行：
    \# user_query = SELECT home, 501 AS uid, 501 AS gid FROM users WHERE userid = ‘%u’
    改为：
    user_query = SELECT maildir, 600 AS uid, 600 AS gid FROM mailbox WHERE username = ‘%u’
    这是用户的查询方法**
  - **6:保存退出**
- **第七步：
  启动 dovecot
  systemctl start dovecot.service
  测试
  telnet mail.extmail.org 110
  输入用户名：
  user postmaster@extmail.org
  输入密码：
  pass extmail
  查看：
  list
  retr 1
  便可以看到刚刚收到的邮件了
  quit 退出**

# 第十一章 RSYNC 文件同步

**备份服务器数据**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143209422.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**Samba 只不过是映射点，并不能保存，数据还在服务器
所以我们需要一个文件同步的工具**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143220484.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**关于 RSYNC**

- **一款快速增量备份工具**
- **Remote Sync,远程同步**
- **支持本地复制,或者与其他 SSH、rsync 主机同步**
- **官方网站:http://rsync.samba.org/**

**Rsync(remote synchronize)是一个远程数据同步工具,可通过 LAN/WAN 快速同步多台主机间的文件,也可以使用 Rsync 同步本地硬盘中的不同目录。**
**Rsync 是用于取代 rcp 的一个工具,Rsync 使用所谓的 “Rsync 算法” 来使本地和远程两个主机之间的文件达到同步,这个算法只传送两个文件的不同部分(校验和),而不是每次都整份传送,因此速度相当快。您可以参考 How Rsync Works A Practical Overview 进一步了解 rsync 的运作机制。**
**Rsync 支持大多数的类 Unix 系统,无论是 Linux、Solaris 还是 BSD 上都经过了良好的测试。此外,它在 windows 平台下也有相应的版本,比较知名的有 cwRsync 和 Sync2NAS**

**特点:**

- **能更新整个目录和树和文件系统;**
- **有选择性的保持符号链链、硬链接、文件属于、权限、设备以及时间等;**
- **对于安装来说,无任何特殊权限要求;**
- **能用 rsh、ssh 或直接端口做为传输入端口;**
- **支持匿名 rsync 同步文件,是理想的镜像工具**

**同步源与发起端**

**rsync 同步源: 指备份操作的远程服务器,也称为备份源主要包括两种:rsync 源、SSH 源**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143234222.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**搭建同步源**

- **1.基于 ssh 的同步源**
- **2.基于 RSYNC 的同步源**

**1.基于 ssh 的同步源
在/var/www/html 文件夹下面创建 vim 网页文件
vim /var/www/html/index.html
写入网站内容
保存退出
在另一台虚拟机上运行命令
curl 服务器 IP
就可以显示网页所写的内容啦
如果需要备份网站内容，使用 ssh 同步源备份到另一台服务器
如果两台服务器需要使用 rsync 同步，两台服务器全需要下载 yum -y install rsync
同步命令：rsync -avz 同步源的一个账户@同步源 IP:同步源目录 发起端目录**

**rsync 命令的用法:
基本格式:rsync [选项] 原始位置 目标位置
常用选项:
-a:归档模式,递归并保留对象属性,等同于 -rlptgoD
-v:显示同步过程的详细(verbose)信息
-z:在传输文件时进行压缩(compress)
-H:保留硬连接文件
-A:保留 ACL 属性信息
–delete:删除目标位置有而原始位置没有的文件
-r:递归模式,包含目录及子目录中所有文件
-l:对于符号链接文件仍然复制为符号链接文件
-p:保留文件的权限标记-t:保留文件的时间标记
-g:保留文件的属组标记(仅超级用户使用)
-o:保留文件的属主标记(仅超级用户使用)
-D:保留设备文件及其他特殊文件**

- **rsync -avz 同步源的一个账户@同步源 IP:同步源目录 发起端目录 下行同步**
- **rsync -avz 发起端目录 同步源的一个账户@同步源 IP:同步源目录 上行同步(发送到同步源)**
  **这一步需要在同步源设置接收文件夹的权限 最好使用 ACL**

**2.基于 RSYNC 的同步源**

**数据下行同步**

- **生成配置文件**
  **vim /etc/rsyncd.conf
  写入工作方式**
  **全局配置部分：**
  **address = IP
  port = 873
  pid file = /var/run/rsyncd.pid
  log file = /var/log/rsyncd.log**
  **共享配置部分**
  **[share]**
  **comment = soft**
  **path = /server/rsync**
  **read only = yes**
  **dont compress = \*.gz \*.bz2 \*.zip**
  **auth users = wang**
  **secrets file = /etc/rsyncd_users.db**
  **接着创建文件**
  **vim /etc/rsyncd_users.db
  wang:123456**
  **rsync 不支持特别复杂的密码设定
  因为他是通过超级守护进程 xinetd
  来管理的，需要输入
  rsync --daemon
  来进行启动**
  **创建目录：
  mkdir -p /server/rsync
  创建测试文件
  touch rsync.txt
  客户端创建目录：
  mkdir -p /client/rsync
  输入命令：
  rsync -avz wang@192.168.1.14::share /client/rsync/
  这时会报错
  @ERROR: auth failed on module share
  rsync error: error starting client-server protocol (code 5) at main.c(1649) [Receiver=3.1.2]
  因为 rsync 的密码文件的权限必须为 600
  服务器执行
  chmod 600 /etc/rsyncd_users.db**

**数据上行同步**

- **客户端创建测试文件
  touch client.txt
  修改配置文件
  vim /etc/rsyncd.conf
  read only = no
  保存退出
  杀死进程
  pkill rsync
  重启
  rsync --daemon
  修改上传目录的写入权限**
  **setfacl -m u:nobody:rwx /server/rsync
  这个时候客户端执行：
  rsync -avz /client/rsync/\* wang@192.168.1.14::share
  会发现有以下信息：
  sending incremental file list
  client.txt
  rsync: chgrp “/.client.txt.ce66QW” (in share) failed: Operation not permitted (1)**
  **sent 111 bytes received 121 bytes 92.80 bytes/sec
  total size is 0 speedup is 0.00
  rsync error: some files/attrs were not transferred (see previous errors) (code 23) at main.c(1179) [sender=3.1.2]**
  **去服务器端查看发现文件已经传输过去了
  这是因为新版本需要在配置文件 share 模块中加入
  fake super = yes #无需 rsync 以 root 身份运行，允许接受数据
  修改配置文件**
  **杀死进程
  rsync --daemon
  重启就可以啦**
  需要主要 rsync -avz /client/rsync/\* wang@192.168.1.14/share
  这种书写格式也可以上传

**命令格式**
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200607143250412.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5MTUwMTIy,size_16,color_FFFFFF,t_70)
**免密码验证**

- **基于 ssh 的免密码同步
  生成密钥对**
  **ssh-keygen -t rsa
  公钥上传到服务器
  ssh-copy-id server@服务器 IP
  此时同步不需要密码了**
- **基于 rsync 的免密码同步**
  **系统为 rsync 准备了变量
  RSYNC_PASSWORD
  我们只需要为这个变量赋值
  echo $RSYNC_PASSWORD
  查看这个变量，此时为空
  给他赋值，为我们刚开始设置的 rsync 的密码
  export RSYNC_PASSWORD=123456
  这回下行同步就不需要密码了**

**同步的优缺点**

- **定期同步的不足：
  执行备份的时间固定,延迟明显、实时性差
  当同步源长期不变化时,密集的定期任务是不必要的**
- **实时同步的优点：
  一旦同步源出现变化,立即启动备份
  只要同步源无变化,则不执行备份**

**inotify 实现实时同步**
**安装 gcc
yum -y install gcc\*
安装 inotifu-tools
wget http://github.com/downloads/rvoicilas/inotify-tools/inotify-tools-3.14.tar.gz
解压 tar -zxvf inotify-tools-3.14.tar.gz
进入解压好的包
./configure && make && make install**
**如何你希望使用效果高，可以调整 inotify 内核参数
调整 inotify 内核参数
max_queue_events:监控队列大小 max_user_instances:最多监控实例数
max_user_watches:每个实例最多监控文件数
你需要在官方网站查看不同版本的系统不同的内核参数
到/etc/sysctl.conf 里面添加参数**
**安装 inotify-tools 辅助工具
inotifywait:用于持续监控,实时输出结果 inotifywatch:用于短期监控,任务完成后再出结果
[root@localhost ~]# inotifywait -mrq -e modify,create,move,delete /var/www/html
Setting up watches. Beware: since -r was given, this may take a while! Watches established.
/var/www/html/ CREATE index.php /var/www/html/ MODIFY index.php /var/www/html/ MOVED_FROM index.php /var/www/html/ MOVED_TO test.php**
**………**

**实时监控**
**inotifywait -mrq -e modify,create,move,delete /var/www/html**
**选项：**
**-m:表示持续监控
-r:表示递归监控
-q:表示输出数据简化
-e:指定你要监控的哪些数据，比如 create,move,delete，多个命令用逗号隔开**

**这条命令输入后，界面会被锁死，需要换一个终端连接，咋检测的目录下创建删除文件，这个锁死的终端就会显示信息
ctrl+c 退出界面**

**实时同步实现**

- **inotify+rsync
  基于 ssh 的实时同步
  利用 inotify 有输出这一特性可以进行与脚本的配合，实时监控数据的增删，进行同步**
  \*\*先进行免密认证
  然后在服务器端写脚本
  cd
  vim 1.sh

```shell
#!/bin/bash**
*a=“/usr/local/bin/inotifywait -mrq -e create /var/www/html/”
b=“/usr/bin/rsync -avz /var/www/html/* 192.168.1.13:/client/ssh”
$a | while read directory event file
do
	$b
done
1234567
```

**保存退出**
**执行**
**bash 1.sh &**

- **inotify+unsion**
  **服务器建立目录
  mkdir /server1
  客户端建立目录
  mkdir /server2**
  **服务器安装 gcc
  yum -y install gcc\*
  服务器与客户端直接相互生成密钥对
  服务器：
  ssh-keygen -t rsa
  ssh-copy-id 客户端 IP
  客户端：
  ssh-keygen -t rsa
  ssh-copy-id 服务器端 IP
  服务器：
  安装 ocaml-3.10.1tar.gz
  wget http://caml.inria.fr/pub/distrib/ocaml-3.10/ocaml-3.10.1.tar.gz
  解压
  tar -zxvf ocaml-3.10.1.tar.gz
  cd ocaml-3.10.1
  ./configure
  make world opt
  make install
  安装 unison
  wget https://www.seas.upenn.edu/~bcpierce/unison/download/releases/unison-2.13.16/unison-2.13.16.tar.gz
  解压
  tar -zxvf unison-2.13.16.tar.gz
  cd unison-2.13.16
  make UISTYLE=text THREADS=true STATIC=true
  cp unison /usr/local/bin/**
  **客户端：
  安装 gcc
  yum -y install gcc\*
  安装 inotifu-tools
  wget http://github.com/downloads/rvoicilas/inotify-tools/inotify-tools-3.14.tar.gz
  解压 tar -zxvf inotify-tools-3.14.tar.gz
  进入解压好的包
  ./configure && make && make install
  安装 ocaml-3.10.1tar.gz
  wget http://caml.inria.fr/pub/distrib/ocaml-3.10/ocaml-3.10.1.tar.gz
  解压
  tar -zxvf ocaml-3.10.1.tar.gz
  cd ocaml-3.10.1
  ./configure
  make world opt
  make install
  安装 unison
  wget https://www.seas.upenn.edu/~bcpierce/unison/download/releases/unison-2.13.16/unison-2.13.16.tar.gz
  解压
  tar -zxvf unison-2.13.16.tar.gz
  cd unison-2.13.16
  make UISTYLE=text THREADS=true STATIC=true
  cp unison /usr/local/bin/**
  **全部安装完成后**
  **服务器：
  写脚本：**
  **vim 2.sh**

```shell
#!/bin/bash
a=“/usr/local/bin/inotifywait -mrq -e create,delete /server1”
b=“/usr/local/bin/unison -batch /server1 ssh://客户端IP//server2”
$a | while read directory event file
do
	$b
done
1234567
```

**scp 2.sh 客户端 IP:/root**
**客户端：**
**vim 2.sh**

```shell
#!/bin/bash
a="/usr/local/bin/inotifywait -mrq -e create,delete /server2"
b="/usr/local/bin/unison -batch /server1 ssh://服务器IP//server1"

$a | while read directory event file
do
        $b
done
12345678
```

**服务器：
bash 2.sh &**

**这回就全部完成啦
可以自己在服务器/server1 中的文件夹里面创建文件，观察客户端/server2 中的变化**
