

1. 首先我们打开在运行输入框等方式打开cmd窗口后，在窗口顶部右击选择属性，选中选项后会看到默认编码为gbk

![如何修改cmd控制台默认编码为utf-8](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image001.jpg)

2. 然后我们在默认窗口路径内，输入chcp命令后回车，会输出图中的结果，936就表示gbk编码

![如何修改cmd控制台默认编码为utf-8](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image002.jpg)

3.  然后在窗口中输入chcp 65001（65001代表utf-8编码），然后回车，即可看到窗口默认编码为utf-8编码了

![如何修改cmd控制台默认编码为utf-8](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image003.jpg)

4. 上面的方法每次都要重新设置，接下来的方法是自动修改，首先win+R打开运行窗口后输入regedit，点击确定按钮

![如何修改cmd控制台默认编码为utf-8](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image004.jpg)

5.  接着会打开注册表窗口，按照图中的路径打开command processor项，

![如何修改cmd控制台默认编码为utf-8](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image005.jpg)

6. 然后会没有autorun字符串值，新建字符串值，设置完名称后右击打开窗口后输入数值数据chcp 65001，点击确定保存，即可完成设置了

![如何修改cmd控制台默认编码为utf-8](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image006.jpg)

 

 

# Powershell修改

第一种：临时性修改编码

使用 chcp 命令，例如 chcp 65001 ，这回将当前代码页变为 utf-8编码，不过这种方式在关闭 cmd 之后会自动失效。

 

常用的编码及对应的码值(10进制)：

 

十进制码值  对应编码名称

950 繁体中文

65001   UTF-8代码页

936 简体中文默认的GBK

437 MS-DOS 美国英语

 

第二种：永久性修改

永久性修改就是通过修改注册注册表达到。

 

打开注册表方法：略

 

定位到：HKEY_CURRENT_USERConsole%SystemRoot%_system32_cmd.exe

 

PS：近期看到下面评论反应很多人在这里找不到 cmd 的选项，由于最近换了新电脑，新安装的 window 10 1903 专业版确实默认找不到（旧版升级过来的可能还保留有），考虑到自己手动添加也比较麻烦，建议大家直接转投 powershell 吧，毕竟 powshell 更强大，也是微软主推的未来的趋势（未来 cmd 就会被 powshell替代掉），而且 powshell 也借鉴了一些 Linux bash 的使用习惯，而且 powshell 是跨平台的哟，你也可以在 Linux 下安装使用，替代 Linux 默认的 bash 。更多详情参考：powshell 官方文档

 

对于没有 CodePage 的，可以鼠标右键-> 新建 -> DWORD(32位)值，创建即可。

![img](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image008.png)

 之后通过 “shift + 鼠标右键” 打开的 powershell 默认代码页就是 65001 即 UTF-8。

![img](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image010.png)

![img](file:///C:/Users/FORMAL~1/AppData/Local/Temp/msohtmlclip1/01/clip_image012.png)

同理你也可以修改：powershell 的默认编码。如果有人因为权限问题无法修改的，可以右键左边选中的红框，选择权限，赋给自己完全控制的权限就OK了！

 