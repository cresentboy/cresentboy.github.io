# CentOS7使用systemctl添加自定义服务

## 一、简介

Centos7开机第一个程序从init完全换成了systemd这种启动方式，同centos 5 6已经是实质差别。systemd是靠管理unit的方式来控制开机服务，开机级别等功能。
在/usr/lib/systemd/system目录下包含了各种unit文件，有service后缀的服务unit，有target后缀的开机级别unit等，这里介绍关于service后缀的文件。因为systemd在开机要想执行自启动，都是通过这些*.service 的unit控制的，服务又分为系统服务（system）和用户服务（user）。

- 系统服务：开机不登陆就能运行的程序（常用于开机自启）。
- 用户服务：需要登陆以后才能运行的程序。

## 二、配置文件说明：

- [Unit] 区块：启动顺序与依赖关系
  Description字段：给出当前服务的简单描述。
  Documentation字段：给出文档位置。
  After字段：如果network.target或sshd-keygen.service需要启动，那么sshd.service应该在它们之后启动。
  Before字段：定义sshd.service应该在哪些服务之前启动。
  注：After和Before字段只涉及启动顺序，不涉及依赖关系。

  > 举例来说，某 Web 应用需要 postgresql 数据库储存数据。在配置文件中，它只定义要在 postgresql 之后启动，而没有定义依赖 postgresql 。上线后，由于某种原因，postgresql 需要重新启动，在停止服务期间，该 Web 应用就会无法建立数据库连接。
  > 设置依赖关系，需要使用Wants字段和Requires字段。
  > Wants字段：表示sshd.service与sshd-keygen.service之间存在"弱依赖"关系，即如果"sshd-keygen.service"启动失败或停止运行，不影响sshd.service继续执行。
  > Requires字段则表示"强依赖"关系，即如果该服务启动失败或异常退出，那么sshd.service也必须退出。
  > 注意，Wants字段与Requires字段只涉及依赖关系，与启动顺序无关，默认情况下是同时启动的。

- [Service] 区块：启动行为

  - 启动命令
    ExecStart字段：定义启动进程时执行的命令
    ExecReload字段：重启服务时执行的命令
    ExecStop字段：停止服务时执行的命令
    ExecStartPre字段：启动服务之前执行的命令
    ExecStartPost字段：启动服务之后执行的命令
    ExecStopPost字段：停止服务之后执行的命令

    注：所有的启动设置之前，都可以加上一个连词号（-），表示"抑制错误"，即发生错误的时候，不影响其他命令的执行。比如`EnvironmentFile=-/etc/sysconfig/sshd`（注意等号后面的那个连词号），就表示即使`/etc/sysconfig/sshd`文件不存在，也不会抛出错误。
    **注意：[Service]中的启动、重启、停止命令全部要求使用绝对路径！**

  - 启动类型
    Type字段定义启动类型。它可以设置的值如下：
    simple（默认值）：ExecStart字段启动的进程为主进程
    forking：ExecStart字段将以fork()方式启动，此时父进程将会退出，子进程将成为主进程（后台运行）
    oneshot：类似于simple，但只执行一次，Systemd 会等它执行完，才启动其他服务
    dbus：类似于simple，但会等待 D-Bus 信号后启动
    notify：类似于simple，启动结束后会发出通知信号，然后 Systemd 再启动其他服务
    idle：类似于simple，但是要等到其他任务都执行完，才会启动该服务。一种使用场合是为让该服务的输出，不与其他服务的输出相混合

  - 重启行为
    Service区块有一些字段，定义了重启行为：
    **KillMode字段：定义 Systemd 如何停止 sshd 服务：**
    control-group（默认值）：当前控制组里面的所有子进程，都会被杀掉
    process：只杀主进程
    mixed：主进程将收到 SIGTERM 信号，子进程收到 SIGKILL 信号
    none：没有进程会被杀掉，只是执行服务的 stop 命令。
    **Restart字段：定义了 sshd 退出后，Systemd 的重启方式**
    上面的例子中，Restart设为on-failure，表示任何意外的失败，就将重启sshd。如果 sshd 正常停止（比如执行systemctl stop命令），它就不会重启。
    Restart字段可以设置的值如下。
    no（默认值）：退出后不会重启
    on-success：只有正常退出时（退出状态码为0），才会重启
    on-failure：非正常退出时（退出状态码非0），包括被信号终止和超时，才会重启
    on-abnormal：只有被信号终止和超时，才会重启
    on-abort：只有在收到没有捕捉到的信号终止时，才会重启
    on-watchdog：超时退出，才会重启
    always：不管是什么退出原因，总是重启
    注：对于守护进程，推荐设为on-failure。对于那些允许发生错误退出的服务，可以设为on-abnormal。
    **RestartSec字段：表示 Systemd 重启服务之前，需要等待的秒数。**

- [Install] 区块

  Install区块，定义如何安装这个配置文件，即怎样做到开机启动。

  WantedBy字段：表示该服务所在的 Target。

  Target的含义是服务组，表示一组服务。

  WantedBy=multi-user.target指的是：sshd 所在的 Target 是multi-user.target。

  这个设置非常重要，因为执行systemctl enable sshd.service命令时，sshd.service的一个符号链接，就会放在/etc/systemd/system目录下面的multi-user.target.wants子目录之中。

  Systemd 有默认的启动 Target。

  

  ```csharp
  systemctl get-default
  #输出multi-user.target
  ```

  上面的结果表示，默认的启动 Target 是multi-user.target。在这个组里的所有服务，都将开机启动。这就是为什么systemctl enable命令能设置开机启动的原因。

  使用 Target 的时候，systemctl list-dependencies命令和systemctl isolate命令也很有用。

  

  ```css
  #查看 multi-user.target 包含的所有服务
  systemctl list-dependencies multi-user.target
  
  #切换到另一个 target
  #shutdown.target 就是关机状态
  systemctl isolate shutdown.target
  ```

  一般来说，常用的 Target 有两个：

  multi-user.target：表示多用户命令行状态；

  graphical.target：表示图形用户状态，它依赖于multi-user.target。

## 三、注册服务实例

- 配置文件目录
  systemctl脚本目录：`/usr/lib/systemd/`
  系统服务目录：`/usr/lib/systemd/system/`
  用户服务目录：`/usr/lib/systemd/user/`

- 在/usr/lib/systemd/system目录下新建service-name.service文件：

  

  ```csharp
  [UNIT]
  #服务描述
  Description=Media wanager Service
  #指定了在systemd在执行完那些target之后再启动该服务
  After=network.target
  
  [Service]
  #定义Service的运行类型
  Type=simple
  
  #定义systemctl start|stop|reload *.service 的执行方法（具体命令需要写绝对路径）
  #注：ExecStartPre为启动前执行的命令
  ExecStartPre=/usr/bin/test "x${NETWORKMANAGER}" = xyes
  ExecStart=/home/mobileoa/apps/shMediaManager.sh -start
  ExecReload=
  ExecStop=
  
  #创建私有的内存临时空间
  PrivateTmp=True
  
  [Install]
  #多用户
  WantedBy=multi-user.target
  ```

  重载系统服务：`systemctl daemon-reload`
  设置开机启动：`systemctl enable *.service`
  启动服务：`systemctl start *.service`
  停止服务：`systemctl stop *.service`
  重启服务：`systemctl restart *.service`

注：修改完配置文件要重载配置文件。