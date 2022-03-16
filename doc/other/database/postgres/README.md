官方文档：https://www.postgresql.org/docs/

GitHub：https://github.com/postgres/postgres

# docker安装postgres

```
docker run -d --name postgres --restart always -e POSTGRES_USER='postgres' -e POSTGRES_PASSWORD='666888' -e ALLOW_IP_RANGE=0.0.0.0/0 -v /home/postgres/data:/var/lib/postgresql -p 5432:5432 -t postgres
```

| Options | Mean                                               |
| ------- | -------------------------------------------------- |
| -i      | 以交互模式运行容器，通常与 -t 同时使用；           |
| -t      | 为容器重新分配一个伪输入终端，通常与 -i 同时使用； |
| -d      | 后台运行容器，并返回容器ID；                       |

**进入postgres容器**

```
docker exec -it 容器ID bash
```

**更新软件源**

```
#更新软件源列表
apt-get update 

#安装vim
apt-get -y install vim 
```

**配置远程访问**

```
#切换到目录/var/lib/postgresql/data
cd /var/lib/postgresql/data
```

 **编辑postgresql.conf文件**

```
#修改：在所有IP地址上监听，从而允许远程连接到数据库服务器：
listening_address: '*'
```

**编辑pg_hba.conf文件**

```
#添加或修改：允许任意用户从任意机器上以密码方式访问数据库，把下行添加为第一条规则：
host    all             all             0.0.0.0/0               md5 
```

**修改编码格式**

```
update pg_database set encoding = pg_char_to_encoding('UTF8') where datname = 'basemap'
```

**查看pg版本**

```
show server_version;
# 或者
select version(); 
```

**尝试登录**

```
#登录数据库
psql -U postgres -W 
```

**重点：报错 psql: FATAL: Peer authentication failed for user "postgres"** 

**问题一：**

**#peer(不可信)，trust(可信)，md5(加密)**

**修改 /etc/postgresql/10/main/pg_hba.conf 文件**

**找到下面这行**

```
local   all             postgres                                peer
```

**修改成md5（加密） （或改成 trust（可信））**

```
local   all             postgres                                md5 
```

**问题二：**

**切换操作用户**

```
#切换成postgres用户
su postgres
```

**尝试登录，成功。**

**重启容器**

```
docker restart 容器name
```

# 一、PG简介



PostgreSQL 是一个免费的对象-关系数据库服务器(ORDBMS)，在灵活的BSD许可证下发行。

PostgreSQL 开发者把它念作 **post-gress-Q-L**。

PostgreSQL 的 Slogan 是 "世界上最先进的开源关系型数据库"。

“开源界的Oracle”,去O首选

PostgreSQL官网

https://www.postgresql.org/

PostgreSQL中文社区

http://www.postgres.cn/v2/home

全球数据库排行

https://db-engines.com/en/

国产数据库排行

https://www.modb.pro/dbRank

## 1、PG的历史

PostgreSQL最初设想于1986年，当时被叫做Berkley Postgres Project。

该项目一直到1994年都处于演进和修改中，直到开发人员Andrew Yu和Jolly Chen在Postgres中添加了一个SQL（Structured Query Language，结构化查询语言）翻译程序，该版本叫做Postgres95，在[开放源代码](https://baike.baidu.com/item/开放源代码/114160)社区发放。

开始以社区的形式运作。

1996年，再次对Postgres95做了较大的改动，并将其作为PostgresSQL6.0版发布。该版本的Postgres提高了后端的速度，包括增强型SQL92标准以及重要的后端特性（包括子选择、默认值、约束和触发器）。

2005年，发布8.0版本，开始支持windows系统环境

**PostgreSQL 9.0 ：**支持64位windows系统，异步流数据复制、Hot Standby；

**PostgreSQL 9.1 ：**支持数据同步复制，unlogged tabels、serializable snapshot isolation、FDW 外部表。

此版本后，PostgreSQL 开始得到中国多个行业用户的关注，开始有应用于电信、保险、制造业等边缘系统。

目前生产环境主流的版本是**PostgreSQL 12**

2021-09-30，PostgreSQL全球开发组宣布，功能最为强大的开源数据库，PostgreSQL 14版本正式发布！

## 2、PG的社区

PG为什么没有被商业公司控制？

### 纯社区

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635228905598-8cbbb2d7-d9e7-4f12-9539-a7766360840e.png)



他们为什么要贡献核心代码？

**最终用户**

- 希望社区长久，期望可以享受免费的、可持续发展的、开源的、不被任何商业公司、不被任何国家控制的企业级数据库。去O，去DB2 ，去Sybase；
- 不靠数据库赚钱；

- PG用到的人越多，越多人背书，使用越靠谱（事实也是如此）；
- 抛砖引玉，企业投入2个研发持续贡献(一年可能一两百万)，实际上整个PG社区有数千人在贡献，对最终用户来说，简直赚到了。使用商业数据库，除了LICENSE等成本，依旧需要投管理、研发、外包资源，一年数千万甚至上亿，公司越大，越有动力去贡献社区。从趋势来看，给PG贡献代码的大客户只会越来越多；

**云厂商**

- 开源数据库与云厂商发生利益冲突，纷纷改协议；
- 数据库市场巨大；

- 自研是最佳选择，但是自研有一些问题：譬如需要培养生态，需要市场背书，需要大量研发资源，可能需要重复造轮子；
- BASE PG的好处：

1、免去自己培养生态，

2、避免重复造轮子，

3、PG的代码基础非常不错（开源界的Oracle）

4、防止其他厂商控制PG失去市场主导能力（AWS,google,IBM,微软都已成为PG社区的赞助商）



**数据库厂商**

- 推一款新的商业数据库，通常都需要背书，小厂产品，谁为你背书？

1、有技术的厂商，很难挑战已有的数据库市场格局

2、有渠道的厂商，需要抓住窗口期，快速占领市场，避免重复造轮子。

需要一款可以无法律风险，二次分发的开源数据库，唯有PG

可以贡献核心代码，社区所有的用户都可以为之背书。



**数据库服务|DaaS服务提供商**

- 开源产品的服务提供商，能力如何体现？

当然是你的架构能力，优化能力，管理能力，FIX BUG的能力

最好能贡献核心代码，有PG为你背书



### 完善的组织结构



![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635229140010-d92f02f9-617e-4e0b-b675-817033408197.png)

### 开源许可独特性

（Postgres遵守BSD许可证发行）却使开发者们得以获取源代码并进一步开发系统。



**BSD许可协议**（英语：**B**erkeley **S**oftware **D**istribution license）是[自由软件](https://zh.wikipedia.org/wiki/自由軟體)中使用最广泛的[许可协议](https://zh.wikipedia.org/wiki/授權條款)之一。[BSD](https://zh.wikipedia.org/wiki/BSD)就是遵照这个许可证来发布，也因此而得名 BSD许可协议。

BSD包最初所有者是[加州大学](https://zh.wikipedia.org/wiki/加州大學)的[董事会](https://zh.wikipedia.org/wiki/董事會)，这是由于 BSD 源自[加州大学伯克利分校](https://zh.wikipedia.org/wiki/加州大学伯克利分校)。BSD开始后，BSD许可协议得以修正，使得以后许多BSD变种，都采用类似风格的条款。

跟其他条款相比，从[GNU通用公共许可证](https://zh.wikipedia.org/wiki/GNU通用公共許可證)（GPL）到限制重重的[著作权](https://zh.wikipedia.org/wiki/著作權)（Copyright），BSD许可证比较宽松，甚至跟[公有领域](https://zh.wikipedia.org/wiki/公有領域)更为接近。"Take it down to the copy center and make as many copies as you want"[[1\]](https://zh.wikipedia.org/wiki/BSD许可证#endnote_copycenter)。可以说，GPL强迫后续版本必须一样是[自由软件](https://zh.wikipedia.org/wiki/自由軟體)，BSD的后续版本可以选择要继续是BSD或其他自由软件条款或[封闭软件](https://zh.wikipedia.org/wiki/封閉軟體)等等。



众所周知，MySQL被Oracle所控制，MySQL同时使用了GPL和一种商业许可（称为双重许可）。

GPL(General Public license)是公共许可，遵循了GPL的软件是公共的。如果某软件使用了GPL软件，那么该软件也需要开源，如果不开源，就不能使用GPL软件，这和是否把该软件商用与否是没关系的。

如果无法满足GPL，就需要获得商业许可，通过与Oracle公司联系，制定解决方案，受Oracle公司约束。



同为开源软件，PostgreSQL源码使用自由友好、商业应用不受任何公司实体所控制，而MySQL则在一定程度上有所限制。



## 3、PostgreSQL与MySQL的比较

**PostgreSQL相对于MySQL的优势**

1）在SQL的标准实现上要比MySQL完善，而且功能实现比较严谨。

2）对表连接支持较完整，优化器的功能较完整，支持的索引类型很多，复杂查询能力较强。

3）PG主表采用堆表存放，MySQL采用索引组织表，能够支持比MySQL更大的数据量。 

4）PG的主备复制属于物理复制，相对于MySQL基于binlog的逻辑复制，数据的一致性更加可靠，复制性能更高，对主机性能的影响也更小。

5）PostgreSQL支持JSON和其他NoSQL功能，如本机XML支持和使用HSTORE的键值对。它还支持索引JSON数据以加快访问速度，特别是10版本JSONB更是强大。 

6）PostgreSQL完全免费，而且是BSD协议，如果你把PostgreSQL改一改，然后再拿去卖钱，也没有人管你，这一点很重要，这表明了PostgreSQL数据库不会被其它公司控制。相反，MySQL现在主要是被Oracle公司控制。

**MySQL相对于PG的优势** 

1）innodb的基于回滚段实现的MVCC机制，相对PG新老数据一起存放的基于XID的MVCC机制，是占优的。新老数据一起存放，需要定时触 发VACUUM，会带来多余的IO和数据库对象加锁开销，引起数据库整体的并发能力下降。而且VACUUM清理不及时，还可能会引发数据膨胀。 

2）MySQL采用索引组织表，这种存储方式非常适合基于主键匹配的查询、删改操作，但是对表结构设计存在约束。 

3）MySQL的优化器较简单，系统表、运算符、数据类型的实现都很精简，非常适合简单的查询操作。

4）MySQL相对于PG在国内的流行度更高，PG在国内显得就有些落寞了。 

5）MySQL的存储引擎插件化机制，使得它的应用场景更加广泛，比如除了innodb适合事务处理场景外，myisam适合静态数据的查询场景。

**总结**

从应用场景来说，PG更加适合严格的企业应用场景（比如金融、电信、ERP、CRM），但不仅仅限制于此，PostgreSQL的json，jsonb，hstore等数据格式，特别适用于一些大数据格式的分析；而MySQL更加适合业务逻辑相对简单、数据可靠性要求较低的互联网场景（比如google、facebook、alibaba），当然现在MySQL的在innodb引擎的大力发展，功能表现良好

扩展阅读：

PostgreSQL 是中国第一的开源数据库？

https://cloud.tencent.com/developer/article/1847734

解密：为什么国产数据库使用PostgreSQL而不是MySQL

https://www.iidba.com/thread-290491-1-1.html

神仙打架：PG 和 MySQL 到底哪个更好用？

https://dbaplus.cn/news-11-3235-1.html

# 二、PostgreSQL的下载安装

## 1、Windows 上安装 PostgreSQL

### (1)下载安装

访问官网下载地址

https://www.enterprisedb.com/downloads/postgres-postgresql-downloads

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635318760002-facae52c-6e92-43a6-bf4a-448912bc4bcf.png)

下载最新发布的PostgreSQL 14。

双击下载安装包，开始安装



![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319055318-57ae9772-35dc-413b-abd2-4b51efc8d7dc.png)



![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635318960879-5ae95eeb-9831-4069-b2b1-f572eb965060.png)

你可以修改安装路径

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319149228-d5100b89-48c4-4a2e-9b1e-426e70ffa0b1.png)

选择安装组件，不懂的选就是全部勾上：

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319206987-c425c4c0-a1db-48d7-91a6-ca8944bd554a.png)

设置数据库路径

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319236720-df341ce1-14a9-43b5-8f40-a3d32e014cf6.png)

设置超级用户的密码

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319282822-693d8a1f-e436-4117-a4e5-8e8e088fe8de.png)

设置端口号，可以直接用默认就行

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319316525-635d93e6-bed0-4bbd-b278-223326c2ae47.png)

直接点 Next

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319344997-a2ffa5f1-279d-44ca-ab4c-4b334f2f114d.png)

点 Next

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319372134-8fc73db1-0ac3-4cd8-91b7-b13dee1897e6.png)

去掉勾选，直接点 Finish

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319395103-df55a31f-206e-44e9-838c-3298d629d495.png)

打开 pgAdmin 4

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319430982-be180641-3f4f-46ed-a47b-c792061b0e25.png)



pgAdmin 主页如下

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319456610-51aac50e-ce07-4e9c-99e3-cab37e8a11a3.png)

点击左侧的 Servers > Postgre SQL 10

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319496936-48f320d5-c3b7-4ea5-a282-456a135f8ff2.png)

输入密码，点击 OK 即可

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319519298-229e3dfa-1254-48e6-b7dd-ec9ec08f5112.png)

控制面板如下

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319556687-992ad93d-c23a-404e-9b54-11941c3dad67.png)

打开 SQL Shell(psql)

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319578088-40390393-c1ce-4952-90fb-2b96a0e2fb5c.png)

### (2)远程访问

\1. 打开postgresql安装目录的data子目录



![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635321031015-740c6c65-0aeb-40ea-90a8-1de15b62db8c.png)

\2. 修改pg_hba.conf文件：在IPV4部分添加新的一行:host all all 0.0.0.0/0 md5

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635321145389-760ad8c6-ba31-44de-89c9-6cb8b25694d4.png)

\3. 控制面板-->系统与安全-->Windows防火墙,关闭防火墙。

4.重启服务

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635321308676-fc6ad978-27a1-471b-91b1-ecb74d63fd44.png)



## 2、Linux上安装PostgreSQL

### (1)下载安装

访问官网下载地址

https://www.postgresql.org/download/

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635319868963-1e64e7c0-660f-4efc-a780-8a36a6bddcd4.png)

选择相应的版本和平台；

![img](https://cdn.nlark.com/yuque/0/2021/png/12417724/1635320024169-a9769ce4-fc34-41c7-a27a-95ec477230b3.png)

```plain
# Install the repository RPM:
sudo yum install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-7-x86_64/pgdg-redhat-repo-latest.noarch.rpm
# Install PostgreSQL:
sudo yum install -y postgresql14-server
# Optionally initialize the database and enable automatic start:
sudo /usr/pgsql-14/bin/postgresql-14-setup initdb
sudo systemctl enable postgresql-14
sudo systemctl start postgresql-14
```

导入yum源

```plain
sudo yum install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-7-x86_64/pgdg-redhat-repo-latest.noarch.rpm
```

安装PostgreSQL服务

```plain
sudo yum install -y postgresql14-server
```

初始化数据库

```plain
sudo /usr/pgsql-14/bin/postgresql-14-setup initdb
```

启动PostgreSQL服务

```plain
#设置PostgreSQL服务为开机启动
sudo systemctl enable postgresql-14
#启动PostgreSQL服务
sudo systemctl start postgresql-14
```

### (2)修改postgres账号密码

PostgreSQL安装成功之后，会默认创建一个名为postgres的Linux用户，初始化数据库后，会有名为postgres的数据库，来存储数据库的基础信息，例如用户信息等等，相当于MySQL中默认的名为mysql数据库。

postgres数据库中会初始化一名超级用户`postgres`

为了方便我们使用postgres账号进行管理，我们可以修改该账号的密码



进入PostgreSQL命令行

通过su命令切换linux用户为postgres会自动进入命令行

```plain
su postgres
```

启动SQL Shell

```plain
psql
```

修改密码

```plain
ALTER USER postgres WITH PASSWORD 'NewPassword';
```



### (3)配置远程访问

开放端口

```plain
sudo firewall-cmd --add-port=5432/tcp --permanent
sudo firewall-cmd --reload
```

修改IP绑定

```plain
#修改配置文件
vi /var/lib/pgsql/14/data/postgresql.conf

#将监听地址修改为*
#默认listen_addresses配置是注释掉的，所以可以直接在配置文件开头加入该行
listen_addresses='*'
```

允许所有IP访问

```plain
#修改配置文件
vi /var/lib/pgsql/14/data/pg_hba.conf

#在问价尾部加入
host  all  all 0.0.0.0/0 md5
```

重启PostgreSQL服务

```plain
#重启PostgreSQL服务
sudo systemctl restart postgresql-14
```



# 三、PostgreSQL的基本使用

## 登录

```plain
#psql -h 服务器 -U 用户名  -d 数据库 -p 端口地址  // -U 是大写
psql -U dbuser -d exampledb -h 127.0.0.1 -p 5432


$ psql (连接数据库，默认用户和数据库都是postgres)
#相当于系统用户postgres以同名数据库用户的身份，登录数据库，这是不用输入密码的。如果一切正常，系统提示符会变为"postgres=#"，表示这时已经进入了数据库控制台。
```

## 数据库操作

```plain
#创建数据库
CREATE DATABASE mydb;

#查看所有数据库
\l

#切换当前数据库
\c mydb



#删除数据库
drop database <dbname>
```



## 数据库表操作



创建表格时每列都必须使用数据类型。PotgreSQL中主要有三类数据类型：

- 数值数据类型
- 字符串数据类型

- 日期/时间数据类型



数值

常见数值类型包括：

| 名字     | 存储长度 | 描述                 | 范围                                         |
| -------- | -------- | -------------------- | -------------------------------------------- |
| smallint | 2 字节   | 小范围整数           | -32768 到 +32767                             |
| integer  | 4 字节   | 常用的整数           | -2147483648 到 +2147483647                   |
| bigint   | 8 字节   | 大范围整数           | -9223372036854775808 到 +9223372036854775807 |
| decimal  | 可变长   | 用户指定的精度，精确 | 小数点前 131072 位；小数点后 16383 位        |
| numeric  | 可变长   | 用户指定的精度，精确 | 小数点前 131072 位；小数点后 16383 位        |
| real     | 4 字节   | 可变精度，不精确     | 6 位十进制数字精度                           |
| double   | 8 字节   | 可变精度，不精确     | 15 位十进制数字精度                          |



字符串字符串类型包括

- char(size)，character(size)：固定长度字符串，size 规定了需存储的字符数，由右边的空格补齐；
- varchar(size)，character varying(size)：可变长度字符串，size 规定了需存储的字符数；

- text：可变长度字符串。



日期/时间

表示日期或时间的数据类型有：

- timestamp：日期和时间；
- date：日期，无时间；

- time：时间；

其他数据类型类型还有布尔值 boolean （true 或 false），货币数额 money 和 几何数据等。





```plain
#创建表
CREATE TABLE test(id int,body varchar(100));

#在表中插入数据
insert into test(id,body) values(1,'hello,postgresql');

#查看当前数据库下所有表
\d

#查看表结构，相当于desc
\d test
```

PostgreSQL 使用序列来标识字段的自增长，数据类型有 smallserial、serial 和 bigserial 。这些属性类似于 MySQL 数据库支持的 AUTO_INCREMENT 属性。

SMALLSERIAL、SERIAL 和 BIGSERIAL 范围：

| 伪类型        | 存储大小 | 范围                          |
| ------------- | -------- | ----------------------------- |
| `SMALLSERIAL` | 2字节    | 1 到 32,767                   |
| `SERIAL`      | 4字节    | 1 到 2,147,483,647            |
| `BIGSERIAL`   | 8字节    | 1 到 922,337,2036,854,775,807 |



示例

```plain
#创建表
CREATE TABLE COMPANY(
   ID  SERIAL PRIMARY KEY,
   NAME           TEXT      NOT NULL,
   AGE            INT       NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL
);

#插入数据
INSERT INTO COMPANY (NAME,AGE,ADDRESS,SALARY)
VALUES ( 'Paul', 32, 'California', 20000.00 );

INSERT INTO COMPANY (NAME,AGE,ADDRESS,SALARY)
VALUES ('Allen', 25, 'Texas', 15000.00 );
```



## Schema

PostgreSQL 模式（SCHEMA）可以看着是一个表的集合。

一个模式可以包含视图、索引、数据类型、函数和操作符等。

相同的对象名称可以被用于不同的模式中而不会出现冲突，例如 schema1 和 myschema 都可以包含名为 mytable 的表。

使用模式的优势：

- 允许多个用户使用一个数据库并且不会互相干扰。
- 将数据库对象组织成逻辑组以便更容易管理。

- 第三方应用的对象可以放在独立的模式中，这样它们就不会与其他对象的名称发生冲突。

模式类似于操作系统层的目录，但是模式不能嵌套。

```plain
#创建schema： 
create schema myschema;

create table myschema.company(
   ID   INT              NOT NULL,
   NAME VARCHAR (20)     NOT NULL,
   AGE  INT              NOT NULL,
   ADDRESS  CHAR (25),
   SALARY   DECIMAL (18, 2),
   PRIMARY KEY (ID)
);

#删除schema： 
drop schema myschema；

#删除一个模式以及其中包含的所有对象：
DROP SCHEMA myschema CASCADE;
```



## 如何备份PostgreSQL数据库

如果您在生产环境中使用[PostgreSQL](https://cloud.tencent.com/product/postgresql?from=10680)，请务必采取预防措施以确保用户的数据不会丢失。

**单数据库**

PostgreSQL提供了`pg_dump`实用程序来简化备份单个数据库的过程。 必须以对要备份的数据库具有读取权限的用户身份运行此命令。

以`postgres`用户身份登录：

```plain
sudo su - postgres
```

通过运行以下命令将数据库的内容转储到文件中。替换`dbname`为要备份的数据库的名称。

```plain
pg_dump dbname > dbname.bak
```

生成的备份文件`dbname.bak`可以使用`scp`传输到另一台主机，也可以存储在本地以供以后使用。

要演示恢复丢失的数据，请删除示例数据库并在其位置创建一个空数据库：

使用`psql`恢复数据库

```plain
 psql test < dbname.bak
```

备份格式有几种选择：

- `*.bak`：压缩二进制格式
- `*.sql`：明文转储

- `*.tar`：tarball

注意：默认情况下，PostgreSQL将忽略备份过程中发生的任何错误。这可能导致备份不完整。要防止这种情况，您可以使用`-1`选项运行`pg_dump`命令。 这会将整个备份过程视为单个事务，这将在发生错误时阻止部分备份。



**所有数据库**

由于`pg_dump`一次只创建一个数据库的备份，因此它不会存储有关数据库角色或其他群集范围配置的信息。 要存储此信息并同时备份所有数据库，可以使用`pg_dumpall`。

创建备份文件：

```plain
pg_dumpall > pg_backup.bak
```

从备份还原所有数据库：

```plain
psql -f pg_backup.bak postgres
```

示例：

```plain
#备份数据库
$ pg_dump -U postgres -f /tmp/postgres.sql postgres (导出postgres数据库保存为postgres.sql)
$ pg_dump -U postgres -f /tmp/postgres.sql -t test postgres (导出postgres数据库中表test的数据)
$ pg_dump -U postgres -F t -f /tmp/postgres.tar postgres (导出postgres数据库以tar形式压缩保存为postgres.tar)

#恢复数据库
$ psql -U postgres -f /tmp/postgres.sql bk01 (恢复postgres.sql数据到bk01数据库)
#pg_restore --  从pg_dump创建的备份文件中恢复PostgreSQL数据库,用于恢复由pg_dump 转储的任何非纯文本格式中的PostgreSQL数据库。
$ pg_restore -U postgres -d bk01 /tmp/postgres.tar  (恢复postgres.tar数据到bk01数据库)
```







## 用户操作

```plain
#创建用户并设置密码
CREATE USER 'username' WITH PASSWORD 'password';
CREATE USER test WITH PASSWORD 'test';


#修改用户密码
$ ALTER USER 'username' WITH PASSWORD 'password';


#数据库授权,赋予指定账户指定数据库所有权限
$ GRANT ALL PRIVILEGES ON DATABASE 'dbname' TO 'username';
#将数据库 mydb 权限授权于 test
GRANT ALL PRIVILEGES ON DATABASE mydb TO test;
#但此时用户还是没有读写权限，需要继续授权表
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO xxx;
#注意，该sql语句必须在所要操作的数据库里执行



#移除指定账户指定数据库所有权限
REVOKE ALL PRIVILEGES ON DATABASE mydb from test

#删除用户
drop user test



# 查看用户
\du
注意：
```

`pg_hba.conf`配置中的第一项设置的意思是：本地用户通过unix socket登陆时，使用peer方式认证。

```plain
# "local" is for Unix domain socket connections only
local   all             all                                     peer
```

[peer](https://www.postgresql.org/docs/current/static/auth-methods.html#AUTH-PEER)是用PostgreSQL所在的操作系统上的用户登陆。

peer方式中，client必须和PostgreSQL在同一台机器上。只要当前系统用户和要登陆到PostgreSQL的用户名相同，就可以登陆。

在刚部署PostgreSQL之后，切换到系统的postgres用户后，直接执行`psql`就能进入PostgreSQL就是这个原因（当前系统用户为名postgre，PostgreSQL中的用户名也是postgre)。

在PostgreSQL中创建一个没有密码的用户：

```plain
create user local_user1;
```

在PostgreSQL所在的机上，创建一个同名的用户：

```plain
useradd local_user1;
```

切换到local_user1用户后，就可以直接通过`unix_socket`登陆PostgreSQL:

```plain
# su - local_user1
[local_user1@10 ~]$ psql postgres     
psql (9.2.24)
Type "help" for help.

postgres=>
```

注意：要指定数据库名，如果不指定默认使用与用户同名的数据库。

peer不是常用的方式！最常用的方式是通过密码远程登陆。



## PostgreSQL 角色管理

在PostgreSQL 里没有区分用户和角色的概念，"CREATE USER" 为 "CREATE ROLE" 的别名，这两个命令几乎是完全相同的，唯一的区别是"CREATE USER" 命令创建的用户默认带有LOGIN属性，而"CREATE ROLE" 命令创建的用户默认不带LOGIN属性



创建david 角色和sandy 用户

```plain
postgres=# CREATE ROLE david;　　//默认不带LOGIN属性
CREATE ROLE
postgres=# CREATE USER sandy;　　//默认具有LOGIN属性
CREATE ROLE
postgres=# \du
                             List of roles
 Role name |                   Attributes                   | Member of 
-----------+------------------------------------------------+-----------
 david     | Cannot login                                   | {}
 postgres  | Superuser, Create role, Create DB, Replication | {}
 sandy     |                                                | {}

postgres=# 
postgres=# SELECT rolname from pg_roles ;
 rolname  
----------
 postgres
 david
 sandy
(3 rows)

postgres=# SELECT usename from pg_user;         //角色david 创建时没有分配login权限，所以没有创建用户
 usename  
----------
 postgres
 sandy
(2 rows)

postgres=#
```



角色属性

| 属性        | 说明                                                         |
| ----------- | ------------------------------------------------------------ |
| login       | 只有具有 LOGIN 属性的角色可以用做数据库连接的初始角色名。    |
| superuser   | 数据库超级用户                                               |
| createdb    | 创建数据库权限                                               |
| createrole  | 允许其创建或删除其他普通的用户角色(超级用户除外)             |
| replication | 做流复制的时候用到的一个用户属性，一般单独设定。             |
| password    | 在登录时要求指定密码时才会起作用，比如md5或者password模式，跟客户端的连接认证方式有关 |
| inherit     | 用户组对组员的一个继承标志，成员可以继承用户组的权限特性     |



**创建用户时赋予角色属性**



如果要在创建角色时就赋予角色一些属性，可以使用下面的方法。

首先切换到postgres 用户。

创建角色bella 并赋予其CREATEDB 的权限。

```plain
postgres=# CREATE ROLE bella CREATEDB ;
CREATE ROLE
postgres=# \du
                             List of roles
 Role name |                   Attributes                   | Member of 
-----------+------------------------------------------------+-----------
 bella     | Create DB, Cannot login                        | {}
 david     |                                                | {}
 postgres  | Superuser, Create role, Create DB, Replication | {}
 sandy     |                                                | {}

postgres=#
```

创建角色renee 并赋予其创建数据库及带有密码登录的属性。

```plain
postgres=# CREATE ROLE renee CREATEDB PASSWORD 'abc123' LOGIN;
CREATE ROLE
postgres=# \du
                             List of roles
 Role name |                   Attributes                   | Member of 
-----------+------------------------------------------------+-----------
 bella     | Create DB, Cannot login                        | {}
 david     |                                                | {}
 postgres  | Superuser, Create role, Create DB, Replication | {}
 renee     | Create DB                                      | {}
 sandy     |                                                | {}

postgres=#
```



测试renee 角色

```plain
postgres@CS-DEV:~> psql -U renee -d postgres
psql (9.1.0)
Type "help" for help.

postgres=> 
```



**给已存在用户赋予各种权限**



赋予登录权限

```plain
postgres=# ALTER ROLE bella WITH LOGIN;
ALTER ROLE
postgres=# \du
                             List of roles
 Role name |                   Attributes                   | Member of 
-----------+------------------------------------------------+-----------
 bella     | Create DB                                      | {}
 david     |                                                | {}
 postgres  | Superuser, Create role, Create DB, Replication | {}
 renee     | Create DB                                      | {}
 sandy     |                                                | {}

postgres=#
```

赋予renee 创建角色的权限

```plain
postgres=# ALTER ROLE renee WITH CREATEROLE;
ALTER ROLE
postgres=# \du
                             List of roles
 Role name |                   Attributes                   | Member of 
-----------+------------------------------------------------+-----------
 bella     | Create DB                                      | {}
 david     |                                                | {}
 postgres  | Superuser, Create role, Create DB, Replication | {}
 renee     | Create role, Create DB                         | {}
 sandy     |                                                | {}

postgres=#
```



## 控制台常用命令总结

```
\password命令（设置密码）
\q命令（退出）
\h：查看SQL命令的解释，比如\h select。
\?：查看psql命令列表。
\l：列出所有数据库。
\c [database_name]：连接其他数据库。
\d：列出当前数据库的所有表格。
\d [table_name]：列出某一张表格的结构。
\du：列出所有用户。
```

