# 一  安装

管理员身份运行cmd

```cmd
@”%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe” -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command “iex ((New-Object System.Net.WebClient).DownloadString(‘https://chocolatey.org/install.ps1‘))” && SET “PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin”
```

管理员运行power shell

```cmd
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

默认路径
C:\ProgramData\chocolatey 
更改软件安装路径
（1.执行“开始/运行”命令（或者WIN + R），输入“regedit”，打开注册表。

#    二  修改路径

展开注册表到下面的分支[HKEY＿LOCAL＿MACHINE\SOFTWARE\Microsoft\Windows   \CurrentVersion]，在右侧窗口中找到名为“ProgramFilesDir”的字符串，双击把数值“C:\Program    Files”修改为“D：\ProgramFiles”，确定退出后,即可更改常用软件的安装路径了。）

# 三  安装软件

先安装一个试试，看看路径是否正确 choco install putty 

查看通过chocolatey安装的软件

```cmd
choco list -local-only
```

查看本地已安装软件

```cmd
choco list -li
choco list -lai
```

升级安装包

```sh
choco upgrade all -y
```

升级自身

```cmd
choco upgrade chocolatey
```

安装常用软件

```cmd
choco install firefox googlechrome potplayer adobereader notepadplusplus.install python3 winrar 7zip git.install office365business foxitreader ccleaner putty vscode microsoft-windows-terminal procexp virtualbox golang powertoys advanced-ip-scanner sublimetext3 procmon intellijidea-community cpu-z.install telegram.install internet-download-manager listary quicklook v2rayn diskgenius fscapture xmind datagrip tim wechat typora vmwareworkstation wiznote maven jdk8
```

[list] - 列出远程或本地包

[search] - 搜索远程或本地包（列表的别名）

[info] - 检索包信息。choco搜索的简写pkgname --exact --verbose

[install]- 从各种来源安装包

[pin] - 抑制包的升级

[outdated] - 检索过时的包。类似于升级全部--noop

[upgrade] 从各种来源升级包

[uninstall] 卸载软件包

[pack] 将nuspec打包到已编译的nupkg

[push] 推送编译的nupkglx

[new] 从模板生成choco包所需的文件

[sources]查看和配置默认源（源的别名）

[source]查看和配置默认源
[config]检索并配置配置文件设置

[feature]查看和配置choco功能

[features]查看和配置choco功能（功能的别名）

[setapikey]检索或保存特定源的apikey（apikey的别名）

[apikey] 检索或保存特定源的apikey

[download]下载包 - 可选择内化所有远程资源

[synchronize]与系统安装的软件同步 - 生成缺少的包

[sync]与系统安装的软件同步 - 生成缺少的软件包

[optimize]优化安装，减少空间使用

