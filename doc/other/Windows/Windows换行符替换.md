问题描述
在编写Linux 脚本后，运行该脚本出现如下错误。

```
[root@master kafka]# ./kk.sh stop
-bash: ./kk.sh: /bin/bash^M: 坏的解释器: 没有那个文件或目录

```

原因
我使用Windows 10系统上的Sublime Text3编写脚本，然后上传到Linux服务器执行。

Windows系统中的换行符是**\n\r**，Linux系统中的换行符是**\n**，因此需要将**\r**替换为空白。

解决方案
Linux sed命令可自动编辑一个或多个文件，执行以下命令即可。

```
sed -i 's/\r$//' xxx.sh
```

