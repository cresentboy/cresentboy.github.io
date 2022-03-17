**Windows进程：指在Windows系统中正在运行的一个应用程序。**

##### 特点：

- 动态性：进程是程序的一次执行过程，是临时的、有生命周期的；
- 独立性：进程是系统进行资源分配和调度的一个独立单位；
- 并发性：多个进程可以在处理机上交替执行；
- 结构性：系统为每个进程建立一个进程控制块。

# Windows系统进程：

| 进程文件       | 进程名称                               | 描述                                                         |
| -------------- | -------------------------------------- | ------------------------------------------------------------ |
| system process | Windows内存处理系统进程                | Windows页面内存管理进程，拥有0级优先                         |
| alg.exe        | 应用层网关服务                         | 用于网络共享                                                 |
| csrss.exe      | Client/Server Runtime Server Subsystem | 客户端服务子系统，用于控制Windows图形相关子系统              |
| ddhelp.exe     | DirectDraw Helper                      | DirectX用于图形服务的一个组成部分                            |
| dllhost.exe    | DCMO DLL Host                          | 支持基于COM对象支持DLL以运行Windows程序                      |
| inetinfo.exe   | IIS Admin Services Helper              | InterInfo是Microsoft Internet Infomation Services (IIS)的一部分，用于Debug调试除错 |
| internat.exe   | 无效Locales                            | 输入控制图标，用于更改类似国家设置、键盘类型和日期格式       |
| kernel32.dll   | Windows壳进程                          | 用于管理多线程、内存和资源                                   |
| lsass.exe      | 本地安全权限服务                       | 控制Windows安全机制                                          |
| mdm.exe        | Machine Debug Manager                  | Debug除错管理用于调试应用程序和Microsoft Office中的Microsoft脚本Editor脚本编辑器 |
| mmtask.tsk     | 多媒体支持进程                         | 控制多媒体服务                                               |
| mprexe.exe     | Windows路由进程                        | 向适当的网络部分发出网络请求                                 |
| msgsrv32.exe   | Windows信使服务                        | 调用Windows驱动和程序管理再启动                              |
| mstask.exe     | Windows计划服务                        | 用于继承在特定的时间或日期备份或者运行                       |
| regsvc.exe     | 远程注册表服务                         | 用于访问在远程计算机的注册表                                 |
| rpcss.exe      | RPC Portmapper                         | Windows的RPC端口映射进程处理RPC调用（远程模块调用）然后把它们映射给指定的服务提供者 |
| services.exe   | Windows Services Controller            | 管理Windows服务                                              |
| smss.exe       | Session Manager Subsystem              | 会话管理子系统，用以初始化系统变量                           |
| snmp.exe       | Microsoft SNMP Agent                   | Windows简单的网络协议代理（SNMP）用于监听和发送请求到适当的网络部分 |
| spool32.exe    | Printer Spooler                        | Windows打印任务控制程序，用以打印机就绪                      |
| spoolsv.exe    | Printer Spooler Services               | Windows打印任务控制程序，用以打印机就绪                      |
| stisvc.exe     | Still Image Service                    | 用于控制扫描仪和数码相机连接在Windows上                      |
| svchost.exe    | Service Host Process                   | 标准的动态链接库主机处理服务                                 |
| system         | Windows System Process                 | Microsoft Windows系统进程                                    |
| taskmon.exe    | Windows Task Optimizer                 | windows任务优化器监视使用某个程序的频率来整理优化硬盘        |
| tcpsvcs.exe    | TCP/IP Services                        | TCP/IP Services Application支持透过TCP/IP连接局域网和Internet |
| winlogon.exe   | Windows Logon Process                  | Windows NT用户登录程序                                       |
| winmgmt.exe    | Windows Management Services            | 透过WMI技术处理来自应用客户端的请求                          |

# Windows程序进程：

| 进程文件        | 进程名称                                        | 描述                                                         |
| --------------- | ----------------------------------------------- | ------------------------------------------------------------ |
| absr.exe        | Backdoor.Autoupder Virus                        | 这个进程是Backdoor.Autoupder后门病毒程序创建的               |
| acrobat.exe     | Adobe Acrobat                                   | Acrobat Writer用于创建PDF文档                                |
| acrord32.exe    | Acrobat Reader                                  | 一个用于阅读PDF文档的软件                                    |
| agentsvr.exe    | OLE Automation Server                           | Microsoft Agent的一部分                                      |
| alogserv.exe    | McAfee VirusScan                                | 反病毒软件，用于扫描文档和E-mail中的病毒                     |
| avconsol.exe    | McAfee VirusScan                                | 反病毒软件，用于扫描文档和E-mail中的病毒                     |
| avsynmgr.exe    | McAfee VirusScan                                | 反病毒软件，用于扫描文档和E-mail中的病毒                     |
| backweb.exe     | BackWeb Adware                                  | 广告插件，来自BackWeb Technologies                           |
| bcb.exe         | Borland C++ Builder                             | Borland C++ Builder                                          |
| calc.exe        | Calculator                                      | Microsoft Windows计算器程序                                  |
| ccapp.exe       | Symantec Common Client                          | Symantce公用应用客户端                                       |
| cdpalyer.exe    | CD Player                                       | Microsoft Windows包含的CD播放器                              |
| charmap.exe     | Windows Character Map                           | Windows字符映射表，用来帮助查找不常用的字符                  |
| idaemon.exe     | Microsoft Indexing Service                      | 在后台运行的Windows索引服务，用于帮助搜索文件在下次变得更快  |
| cisvc.exe       | Microsoft Index Service Helper                  | 监视Microsoft Indexing Services的内存占用情况                |
| cmd.exe         | Windows Command Prompt                          | Windows控制台程序                                            |
| cmesys.exe      | Gator GAIN Adware                               | Gator GAIN是一个Adware广告插件                               |
| ctfmon.exe      | Alternative User 无效 Services                  | 控制Alternative User无效Text Processor(TIP)和Microsoft Office语言条，该进程提供语音识别、手写识别、键盘、翻译和其他用户输入技术的支持 |
| ctsvccda.exe    | Create CD-ROM Services                          | 在Win9X创建CD-ROM访问服务                                    |
| cutftp.exe      | CuteFTP                                         | 一个流行的FTP客户端用于从FTP服务器上传/下载文件              |
| defwatch.exe    | Norton AntiVirus                                | 扫描文件和E-mail以检查病毒                                   |
| devldr32.exe    | Create Device Loader                            | Creative Device Loader属于Create Soundblaster驱动            |
| directcd.exe    | Adaptec DirectCD                                | 一个用于文件管理器式的界面，烧录文件到光盘的软件             |
| dreamweaver.exe | Macromedia DreamWeaver                          | 一个HTML编辑器用于创建站点和其它类别的HTML文档               |
| em_execute.exe  | Logitech Mouse Settings                         | 用于用户访问控制鼠标属性和查看Mouse Ware帮助                 |
| excel.exe       | Microsoft Excel                                 | 一个电子表格程序，包括在Microsoft Office中                   |
| findfast.exe    | Microsoft Office Indexing                       | Microsoft Office索引服务，用于提高Microsoft Office索引Office文档的速度 |
| frontpage.exe   | Microsoft FrontPage                             | 一个HTML编辑器用于创建站点和其它类别的HTML文档               |
| gmt.exe         | Gator Spyware Component                         | 一个广告插件，随Gator安装和启动                              |
| hh.exe          | Gator Windows Help                              | Windows Help程序用以打开帮助文件和文档，包括在很多Windows程序中 |
| hidserv.exe     | Microsoft Human Interface Device Audio Services | 后台服务，用来支持USB音效部件和USB多媒体键盘                 |
| iexplore.exe    | Internet Explorer                               | IE网络浏览器透过HTTP访问WWW万维网                            |
| kodakimage.exe  | Imaging                                         | 一个图片查看软件，包括在Windows，用以打开图像文件            |
| loadwc.exe      | Load WebCheck                                   | 用以定制一些Internet Explorer的设定，添加、删除或者更新用户profiles设定 |
| mad.exe         | System Attendant Service                        | Microsoft Exchange Server的后台程序，用以读取Microsoft Exchange的DLL文件，写log信息和生成离线地址簿 |
| mcshield.exe    | McAfee VirusScan                                | 一个反病毒软件，用以扫描你的文件和E-mail中的病毒             |
| mgabg.exe       | Matrox BIOS Guard                               | Matrox BIOS守护进程                                          |
| mmc.exe         | Microsoft Management Console                    | 管理控制程序集成了很多的系统控制选项。例如：设备管理或计算机的权限控制 |
| mobsync.exe     | Microsift Synchronization Manager               | Internet Explorer的一个组成部分，用以在后台同步离线查看页面  |
| mplayer.exe     | Windows Media Player                            | 一个用以打开音乐、声音和视频文件的软件                       |
| msaccess.exe    | Microsoft Access                                | 一个数据库软件，包括在Microsoft Office中                     |
| msbb.exe        | MSBB Web3000 Spyware Application                | 包括在一些adware产品中，利用注册表随Windows启动              |
| msdtc.exe       | Distributed Transaction Coordinator             | 控制多个服务器的传输，被安装在Microsoft Personal Web Server和Microsoft SQL Server上 |
| msi_execute.exe | Windows Installer Component                     | Windows Installer的一部分，用来帮助Windows Installer package files(MSI)格式的安装文件 |
| msimn.exe       | Microsoft Outlook Express                       | 一个Email和新闻组客户端，包括在Microsoft Windows中           |
| msmsgs.exe      | MSN Messenger Traybar Process                   | 一个在线聊天和即时通讯客户端                                 |
| msoobe.exe      | Windows Product Activation                      | Windows XP License的Product Activation产品激活程序           |
| mspaint.exe     | Microsoft Paint                                 | 一个图像编辑器，包括在Microsoft Windows中                    |
| mspmspsv.exe    | WMDM PMSP Service                               | Windows Media Player7需要安装的Helper Service                |
| mysqld-nt.exe   | MySQL Daemon                                    | 控制访问MySQL数据库                                          |
| navapsvc.exe    | Norton AntiVirus Auto-Protect Service           | 扫描你文件和Email中的病毒                                    |
| ndetect.exe     | ICQ Ndetect Agent                               | ICQ用来侦测网络连接的程序                                    |
| netscape.exe    | NetScape                                        | 网络浏览器通过HTTP浏览WWW万维网                              |
| notepad.exe     | Notepad                                         | 字符编辑器，用于打开文档，在Windows中附带                    |
| ntbackup.exe    | Windows Backup                                  | Windows备份工具，用于备份文件和文件夹                        |
| ntvdm.exe       | Windows 16-bit Virtual Machine                  | 为了兼容旧的16位Windows和DOS程序而设置的虚拟机               |
| nvsvc32.exe     | NVIDIA Driver Helper Service                    | 在NVIDA显卡驱动中被安装                                      |
| nwiz.exe        | NVIDIA nView Control Panel                      | NVIDIA nView控制面板在NVIDA显卡驱动中被安装，用于调整和设定  |
| osa.exe         | Office Startup Assistant                        | Microsoft Office启动助手。随Windows启动，增强启动、Office字体、命令和Outlook事务提醒等特性 |
| outlook.exe     | Microsoft Outlook                               | 一个Email客户端，包括在Microsoft Office中                    |
| photoshop.exe   | Adobe Photoshop                                 | 一个图像编辑软件，能够打开和编辑其它更多类型格式的图片       |
| point32.exe     | Microsoft Intelimouse Monitor                   | 添加一个鼠标设定图标在工具栏                                 |
| powerpnt.exe    | Microsoft PowerPoint                            | 一个演示软件，包括在Microsoft Office中                       |
| pstores.exe     | Protected Storage Service                       | 控制保密的类容密码                                           |
| qttask.exe      | Quick Time Tray Icon                            | Quick Time任务栏图标在你运行Quick Time的时候启动             |
| realplay.exe    | Real Player                                     | 一个媒体播放器，用来打开和播放音乐、声音和Real Media格式的视频文件 |
| rnaapp.exe      | Windows Modem Connection                        | 控制拨号modem连接                                            |
| rtvscan.exe     | Norton AntiVirus                                | 用以扫描你的文件和Email中的病毒                              |
| rundll32.exe    | Windows RUNDLL32 Helper                         | 支持需要调用DLLs的程序                                       |
| sndrec32.exe    | Windows Sound Recorder                          | Windows录音机，用以播放和录制声音文件（.wav）                |
| sndvol32.exe    | Windows Volume Control                          | Windows声音控制进程在任务栏驻留，用以控制音量和声卡相关选项  |
| spoolss.exe     | Printer Spooler Subsystem                       | Windows打印机控制子程序，用以调用需要打印的内容从磁盘到打印机 |
| starter.exe     | Creative Labs Ensoniq Mixer Tray Icon           | 状态栏图标在Creative Sound Mixer中被安装，为了Creative声卡（Soundblaster） |
| systray.exe     | Windows Power Management                        | Windows电源管理程序，用以控制节能和恢复启动                  |
| tapisrv.exe     | TAPI Service                                    | Windows Telephony(TAPI)的后台服务程序                        |
| userinit.exe    | UserInit Process                                | UserInit程序运行登录脚本，建立网络连接和启动Shell壳          |
| visio.exe       | Microsoft Visio                                 | 一个图形化管理软件                                           |
| vptray.exe      | Norton AntiVirus                                | 扫描你的文件和Email中的病毒                                  |
| vshwin32.exe    | McAfee VirusScan                                | 一个反病毒软件，用以扫描你的文件和Email中的病毒              |
| vsmon.exe       | True Vector Internet Monitor                    | 是ZoneAlarm个人防火墙的一部分，用以监视网络流经数据和攻击    |
| vsstat.exe      | McAfee VirusScan                                | 一个反病毒软件，用以扫描你的文件和Email中的病毒              |
| wab.exe         | Address Book                                    | 在Outlook中的地址簿，用来存放Email地址、联系信息             |
| webscanx.exe    | McAfee VirusScan                                | 一个反病毒软件，用以扫描你的文件和Email中的病毒              |
| winamp.exe      | WinAmp                                          | WinAmp Media Player是一个用来打开音乐、声音和视频文件以及用以管理MP3文件的软件 |
| winhlp32.exe    | Windows Help                                    | Windows帮助文件查看程序，用来打开帮助文档，该程序包括在很多的Windows程序中 |
| winoa386.exe    | MS-DOS Console                                  | Windows MS-DOS控制台                                         |
| winproj.exe     | Microsoft Project                               | 一个项目计划编制程序                                         |
| winroute.exe    | WinRoute                                        | 一个基于Windows的防火墙/路由/连接共享软件                    |
| winword.exe     | Microsoft Word                                  | 一个字处理程序，包括在Microsoft Office                       |
| winzip32.exe    | WinZip                                          | 一个文件压缩工具，用于创建、打开和解压zip文件                |
| wkcalrem.exe    | Microsoft Works Calender Reminder               | 工作日程提醒，在后台处理和显示弹出计划的工作日志提醒         |
| wkqkpick.exe    | WinZip Traybar Icon                             | WinZip的状态图标，被允许在WinZip启动时启动                   |
| wmplayer.exe    | Windows Media Player                            | 一个用来打开和播放音乐、声音和视频的软件                     |
| wordpad.exe     | Wordpad                                         | 一个字符编辑器，用以打开和编辑txt和rtf文档                   |
| ypager.exe      | Yahoo Messenger Helper                          | Yahoo Messenger的状态栏图标，随Yahoo Messenger运行，是其一部分 |