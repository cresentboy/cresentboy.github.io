**MacOS 常用终端命令大全**

**初识终端**

格式：“电脑用户名:当前路径 账户名$ 一条命令 ”

命令的构成：Command Name、Options、Arguments、Extras 四个部分，很多情况下后面三部分都是可省略的。

Options 部分用-作为前导符。其中许多命令的 Options 部分只包含单个字母，这时可以合并。例如:ls -lA和ls -l -A是等效的。

Arguments 部分用来细化这个命令或指定这个命令具体的实施对象。

Extras 部分则用来进一步实现其他功能。

例子：删除 QQ 这个程序。

some-pc:~ mac$ rm -R /Applications/QQ.app



**为什么要使用命令行****/****如何开启命令行？**

许多功能在图形界面不提供，只有通过命令行来实现。

Finder会隐藏许多你不太会需要的文件，然而 command line 会允许你访问所有文件。

通过 command line 可以远程访问你的 Mac（利用 SSH）。

administrators 用户可以通过 sudo命令获得 root 用户权限。

通过 command-line script 可以使工作更高效。

Terminal（终端）程序可以在“实用工具”里找到。

如果你开启手动输入用户名登录模式，登陆时在用户名处输入 >console 可以直接进入命令行界面。随后你仍然需要登录到一个账户。



**关于** **man** **命令**

不管是mac还是linux都有很多命令，不可能熟练掌握所有命令，即使忘记了使用Google也能查到。mac最强大的一个命令应该算 man xxx ，Mac有上千条命令，每条命令还有许多可选参数和具体的使用方式，但是你却不需要记住这些命令。你只需要记住一个：man，查看具体的命令说明，想要推出直接键入q即可。

大多数命令都会包含一个使用指南，会告诉你任何你需要知道的关于这个命令的所有细节，在命令行中输入 man command-name即可获取。例如，你想知道ls这个命令怎么使用，输入man ls即可进入使用指南页面。使用指南往往很长，所以你可以使用▲（上箭头）或▼（下箭头）来上下移动，使用　来翻页，输入/和关键字来按照关键字搜索，按Q来退出使用指南页面。

那么——如果你连命令名称都不知道怎么办呢？输入man -k和关键字来对整个使用指南数据库进行搜索。



**MacOS** **常用终端命令大全：**

**目录操作**

命令——功能描述——示例

mkdir——创建一个目录——mkdir dirname

rmdir——删除一个目录——rmdir dirname

mvdir——移动或重命名一个目录——mvdir dir1 dir2  (macos使用**mv**代替mvdir, macbook air 2017, 13inch macOS  11.1实测)

cd——改变当前目录——cd dirname

pwd——显示当前目录的路径名——pwd

ls——显示当前目录的内容——ls -la

dircmp——比较两个目录的内容——dircmp dir1 dir2  (改为**diff**)



**文件操作**

命令——功能描述——示例

cat——显示或连接文件————cat filename

pg分页格式化显示文件内容——pg filename

more——分屏显示文件内容——more filename

od——显示非文本文件的内容——od -c filename

cp——复制文件或目录——cp file1 file2

rm——删除文件或目录——rm filename

mv——改变文件名或所在目录——mv file1 file2

ln——联接文件——ln -s file1 file2

find——使用匹配表达式查找文件——find . -name “*.c” -print

file——显示文件类型——file filenamed

open——使用默认的程序打开文件——open filename（open . 打开当前目录）



**选择操作**

命令——功能描述——示例

head——显示文件的最初几行——head -20 filename

tail——显示文件的最后几行——tail -15 filename

cut——显示文件每行中的某些域——cut -f1,7 -d: /etc/passwd

colrm——从标准输入中删除若干列——colrm 8 20 file2

paste——横向连接文件——paste file1 file2

diff——比较并显示两个文件的差异——diff file1 file2

sed————非交互方式流编辑器——sed “s/red/green/g” filename

grep——在文件中按模式查找——grep “^[a-zA-Z]” filename

awk——在文件中查找并处理模式——awk ‘{print 111}’ filename

sort——排序或归并文件——sort -d -f -u file1

uniq——去掉文件中的重复行——uniq file1 file2

comm——显示两有序文件的公共和非公共行——comm file1 file2

wc——统计文件的字符数、词数和行数——wc filename

nl——给文件加上行号——nl file1 >file2



**安全操作**

命令——功能描述——示例

passwd——修改用户密码——passwd

chmod——改变文件或目录的权限——chmod ug+x filename

umask————定义创建文件的权限掩码——umask 027

chown——改变文件或目录的属主——chown newowner filename

chgrp——改变文件或目录的所属组——chgrp staff filename

xlock——给终端上锁——xlock -remote



**编程操作**

命令——功能描述——示例

make——维护可执行程序的最新版本——make

touch——更新文件的访问和修改时间——touch -m 05202400 filename

dbx——命令行界面调试工具——dbx a.out

xde——图形用户界面调试工具——xde a.out



**进程操作**

命令——功能描述——示例

ps——显示进程当前状态——ps u

kill——终止进程——kill -9 30142

nice——改变待执行命令的优先级——nice cc -c *.c

renice——改变已运行进程的优先级——renice +20 32768



**时间操作**

命令——功能描述——示例

date——显示系统的当前日期和时间——date

cal——显示日历——cal 8 1996

time——统计程序的执行时间——time a.out



**网络与通信操作**

命令——功能描述——示例

telnet——远程登录——telnet hpc.sp.net.edu.cn

rlogin——远程登录——rlogin hostname -l username

rsh——在远程主机执行指定命令——rsh f01n03 date

ftp——在本地主机与远程主机之间传输文件——ftp ftp.sp.net.edu.cn

rcp——在本地主机与远程主机 之间复制文件——rcp file1 host1:file2

ping——给一个网络主机发送 回应请求——ping hpc.sp.net.edu.cn

mail——阅读和发送电子邮件——mail

write——给另一用户发送报文——write username pts/1

mesg——允许或拒绝接收报文——mesg n



**Korn Shell** **命令**

命令——功能描述——示例

history——列出最近执行过的 几条命令及编号——history

r——重复执行最近执行过的 某条命令——r -2

alias——给某个命令定义别名——alias del=rm -i

unalias——取消对某个别名的定义——unalias del



**其它命令**

命令——功能描述——示例

uname——显示操作系统的有关信息——uname -a

clear——清除屏幕或窗口内容——clear

env——显示当前所有设置过的环境变量——env

who——列出当前登录的所有用户——who

whoami——显示当前正进行操作的用户名——whoami

tty——显示终端或伪终端的名称——tty

stty——显示或重置控制键定义——stty -a

du——查询磁盘使用情况——du -k subdir

df——显示文件系统的总空间和可用空间——df /tmp

w——显示当前系统活动的总信息——w



**一些常用技巧**

所以你可以使用▲（上箭头）或▼（下箭头）来上下移动，

使用　空格键 来翻页，输入/和关键字来按照关键字搜索

按Q来退出使用指南页面

tab按键自动补全唯一路径

中止一个错误的或者发疯的命令，可以使用组合键control + C。

你可以在执行前编辑命令，只需要使用箭头和键盘上的其他字母。

没有输入任何命令时，你可以用▲和▼来浏览历史命令。同样可以编辑和再次执行。

你也可以使用history命令查看历史记录。

你可以使用组合键control + L清屏。



**一、说明**

1. MAC系统采用Unix文件系统，所有文件都挂在根目录下面，没有Windows系统的盘符概念，根目录用斜杠(/)表示；
2. 根目录(/)不是可有可无，/System表示根目录下的System文件，System表示当前目录下的System文件；
3. 在 Unix系统中区别字符大小写，A.txt 不等于 a.txt；
4. 关键的标点符号：点(.)表示当前目录；两个点(..)表示上一级目录；星号(*)匹配任意字符任意次数；问号(?)匹配任意字符仅一次；
5. 获得权限：为了防止误操作破坏系统，再用户状态下没有权限操作重要的系统文件，先要获取root权限，语法：sudo -s，然后会提示输入密码，输入密码时没有任何回显，连星号都没有，输入完密码按回车键即可；
6. 编辑文件：vim directory/file_name，若指定路径的文件不存在，则新建空文件，输入字母i或o进入编辑模式，编辑好内容，点击【esc】键后，输入:w进行保存；输入:wq进行保存并退出；输入:q!进行不保存强行退出；
7. table键，单击可以实现自动补全，双击可以列出指定路径下的所有内容，类似ls命令；
8. 常用位置
    驱动所在位置： /Systme/Library/Extensions
    用户文件夹位置：/User/用户名，可以用波浪号(~)表示
    桌面位置：/User/用户名/Desktop，可以用~/Desktop表示
9. 清理系统
    按天进行清理：sudo periodic daily
    按每周进行清理：sudo periodic weekly
    按每月进行清理：sudo periodic monthly
    按上面3种情况进行清理：sudo periodic daily weekly monthly

**二、目录和文件操作**

| **命令名**        | **功能描述**                 | **举例或备注**                   |
| ----------------- | ---------------------------- | -------------------------------- |
| cd                | 进入指定文件夹路径           | cd ~/Desktop                     |
| pwd               | 显示当前的目录路径           | /Users/xz/Desktop                |
| ls                | 显示当前目录下的内容         |                                  |
| ls -la            | 显示当前目录下的详细内容     |                                  |
| ls -A             | 显示当前目录下的内容         | 含点(.)开头的文件                |
| mkdir             | 创建目录                     | mkdir dir_name                   |
| touch file.format | 创建指定格式的文件           |                                  |
| mvdir             | 移动目录                     | mvdir dir1 dir2                  |
| mv                | 移动/重命名---文件/文件夹    | mv dir1 dir2 MAC没有重命名的命令 |
| rm                | 删除文件 或 **空**目录       |                                  |
| rm -rf dir        | 删除一个 **非空** 目录       | rm -rf dir                       |
| rmdir             | 删除 **空** 目录             | 平时用得少                       |
| cp                | 复制文件或目录               | cp file1 file2                   |
| file              | 显示文件类型                 | file file_name                   |
| find              | 使用匹配表达式查找文件       | find *.file_format               |
| open              | 使用默认的程序打开文件       | open file_name                   |
| cat               | 显示或连接文件内容           | cat file                         |
| ln                | 为文件创建联接               | ln -s file1 file2 s 表示软联接   |
| head              | 显示文件的最初几行           | head -20 file_name               |
| tail              | 显示文件的最后几行           | tail -10 file_name               |
| paste             | 横向拼接文件内容             | paste file1 file2                |
| diff              | 比较并显示两个文件的内容差异 | diff file1 file2                 |
| wc                | 统计文件的字符数、词数和行数 | wc file_name                     |
| uniq              | 去掉文件中的重复行           | uniq file_name                   |
| grep              | 通过简单正则表达式搜索文件   |                                  |

**三、文件属性**

- Linux系统：一切设备都可以看成是文件。如：目录、磁盘文件、管道、网络Socket、外接U盘和SD卡等；
- 文件属性：用户组、读、写、执行权限；
- 查看文件属性

XZ:ts xz$ ls -l

total 82488

-rw-r--r--@ 1 xz staff 42233727 7 19 16:30 PowerBi.pbix

| **语法**     | **属性** | **含义说明**                                |
| ------------ | -------- | ------------------------------------------- |
| -            | 文件类型 | 横杠表示普通文件，若为d表示文件目录         |
| rw-r--r--    | 访问权限 | 分3组：用户、群组和其他用户的文件访问权限； |
| 1            | 文件数量 | 本例中仅1个文件                             |
| xz           | 所在用户 | 本例中用户名为xz                            |
| staff        | 所在群组 | 本例中用户群组为staff                       |
| 42233727     | 文件大小 | 本例中文件的字节数                          |
| 7 19 16:30   | 修改日期 | 本例中为7-19 16:30                          |
| PowerBi.pbix | 文件名称 | 本例中为PowerBi.pbix                        |

- 修改访问权限
   **语法**：chmod 用户 操作 权限 文件
   **用户**：u表示用户(user)、g表示群组(group)、o表示其他用户(other)、
     a表示全部用户。缺失的情况下默认为所有用户；
   **操作**：+表示增加权限、-表示取消权限、=表示赋值权限；
   **权限**：r表示可读(read)、w表示可写(write)、x表示可执行(execute)；
   **文件**：不指定文件名时，操作对象为当前目录下的所有文件。
- 示例：为user用户增加执行的权限

XZ:ts xz$ chmod u+x PowerBi.pbix 

XZ:ts xz$ ls -l

total 82488

-rwxr--r--@ 1 xz staff 42233727 7 19 16:30 PowerBi.pbix

**四、常用操作**

| **命令名**     | **功能描述**                 | **举例或备注**                                               |
| -------------- | ---------------------------- | ------------------------------------------------------------ |
| sudo           | 获取root权限                 | sudo -s                                                      |
| Ctr + D / exit | 退出root权限                 |                                                              |
| clear          | 清除屏幕或窗口内容           |                                                              |
| ping           | 给网络主机发送回应请求       | ping [www.baidu.com](https://links.jianshu.com/go?to=http%3A%2F%2Fwww.baidu.com) |
| man            | 查看命令说明                 | man ls                                                       |
| q              | 退出查看的命令说明           |                                                              |
| which          | 查看指定程序的路径           | which python                                                 |
| history        | 列出最近执行过的命令及编号   |                                                              |
| hostname       | 电脑在网络中的名称           |                                                              |
| env            | 显示当前所有设置过的环境变量 |                                                              |
| passwd         | 修改用户密码                 |                                                              |
| date           | 显示系统的当前日期和时间     | date                                                         |
| cal            | 显示日历                     | cal                                                          |
| time           | 统计程序的执行时间           | time                                                         |

**五、快捷键**

CTRL+A：移动光标至行首
 CTRL+E：移动光标至行尾
 CTRL+X：按住CTRL，双击 X 可以进行当前位置与行首位置，进行切换
 ESC+B：光标向左移动一个单词
 ESC+F：光标向右移动一个单词
 CTRL+U：删除光标前所有字符
 CTRL+K：删除光标后所有字符
 CTRL+W：删除光标前一个单词（根据空格识别单词分隔）
 CTRL+Y：粘贴之前（CTRL+U/K/W）删除的内容
 CTRL+C：中断操作


