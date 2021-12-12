# Oracle数据库对比MySQL

# 基本

Oracle默认端口：1521 默认用户：system
MySQL默认端口：3306 默认用户：root

连接MySQL：

```bash
mysql -u root -p
-- 输入密码

-- 查询所有数据库
show databases;
-- 切换到 "test" 这个数据库
use test;
-- 查询该数据库所有表
show tables;
```

连接Oracle：

```bash
sqlplus
-- 输入用户名
-- 输入密码

-- 查询该用户的表
select TABLE_NAME from user_tables;
```

注意：Oracle 登录需要授予登录用户 session权限，建表需要分配限额

# 常用字段类型

Oracle
数值 number number(10) number(10,2)
字符串 varchar2 varchar2(20)
日期 date

MySQL
数值 tinyint smallint mediumint int bigint decimal
字符串 varchar(10) 必须指定
日期 date time datetime timestamp year

# DML

## Oracle:

```sql
create table t_student(
    sid int primary key ,
    sname varchar2(10) not null ,
    enterdate date,
    gender char(2),
    mail unique,
    age number check (age>19 and age<30)
)
insert into t_student values(stuseq.nextval,'Test',to_date('1990-3-4','YYYY-MM-DD'),'男','1@outlook.com',20);
commit;
```

## MySQL

```sql
create table t_student(
    sid int primary key auto_increment,
    sname varchar(1) not null ,
    enterdate date,
    gender char(1),
    age int,
    mail varchar(10) UNIQUE
)

insert into t_student values(null,'Test','1990-3-4','男',30,'2@outlook.com')
```

MySQL插入日期使用now() 或 sysdate()，可以插入多条，使用逗号隔开
删表数据：Oracle可以省略from：delete from t_student; (删除所有数据)

外键约束：Oracle是constraints,MySQL是constraint

级联操作：

- Oracle：on delete set null 或者on delete cascade
- MySQL: on delete set null on update CASCADE

# 多表操作

Oracle：92语法：可以内连接，外连接99语法：可以内连接，外连接，全外连接(full join)

```sql
-- SQL92 左外连接（保留左边, 注意(+)要放在右边，记忆：左外，右边会出现空行要+补齐） 
where e.department_id = d.department_id(+)
-- 
```

MySQL：只支持内连接、外连接，并且只能用类似oracle中99语法的格式写，MySQL不完全符合SQL-92规范

# SQL 语句

## MySQL

大小写不敏感(关键字和字段名都不区分)

阿里巴巴Java开发手册，在MySQL建表规约里有：
【强制】表名、字段名必须使用小写字母或数字 ， 禁止出现数字开头，禁止两个下划线中间只出现数字。数据库字段名的修改代价很大，因为无法进行预发布，所以字段名称需要慎重考虑

**Windows 大小写不敏感，文件名同名大小写不同会覆盖**

MySQL 在 Windows 下不区分大小写，但在 Linux 下默认是区分大小写。因此，数据库名、 表名、字段名，都不允许出现任何大写字母，避免节外生枝
MySQL 的字段 大小写都可以查到

## Oracle

是Oracle大小写不敏感的前提条件是在没有使用双引号 "" 的前提下（表名、字段名）

CREATE TABLE "TableName"("id" number); // 如果创建表的时候是这样写的，那么就必须严格区分大小写

SELECT * FROM "TableName"; // 不仅要区分大小写而且要加双引号，以便和上面的第三种查询方式区分开
Oracle默认是大写，对字段的具体值是敏感的

# 分页

Oracle：

```sql
-- 利用rownum 
-- rownum从0开始
select * from
(select rownum rr,stu.* from (select * from t_student order by sid desc) stu )
where rr>=1 and rr<=5;
```

MySQL：

```sql
-- 记录从0开始
-- 从第0条开始，取5条数据
select * from test2 order by sid desc  limit 0,5
```

# 时间日期

## Oracle

Java中常用的 "yyyy-MM-dd mm:HH;ss" -> "2020-02-03 16:25:48"
在 Oracle 中的表示方式：'yyyy-mm-dd hh24:mi:ss'

## MySQL

```sql
-- 获取当前时间戳 
select unix_timestamp(); 
-- 1612340981

-- 获取当前日期时间
select now();
2020-02-03 16:30:22

-- 获取当前日期
select date(now());
-- 2020-02-03

-- timestamp -> datetime
select FROM_UNIXTIME(1612340981);
-- 2020-02-03 16:29:41

-- datetime -> varchar  (time与之类似：time_format(time,format))
select  DATE_FORMAT('2008-08-08 22:23:01','%Y %m %d %H %i %s');
-- 2008 08 08 22 23 01

-- varchar -> date   str_to_date(str, format)
select str_to_date('08.09.2008 08:09:30', '%m.%d.%Y %h:%i:%s'); 
-- 2008-08-09 08:09:30
```

# Oracle

Oracle DML 需要手动提交或回滚事务
DML(Data Manipulation Language): 数据操纵语言 针对表数据的增删改查
Oracle select 查询必须有from 所以可以用from dual（这是一张神奇的表）

## 类型转换

### date <--> varchar2 <--> number

date --> varchar2 : to_char(sysdate,'yyyy-mm-dd')
varchar2 --> date : to_date('2020-02-02','yyyy-mm-dd')

number --> varchar2: to_char(1111111.11,'999,999,999') -- 输出：1,111,111 使用'999,999,999'去匹配数字
varchar2 --> number :to_number('￥001,111,111','L000,000,000') from dual; -- 输出：1111111

L表示：当地的货币符号 字符串在运算时会自动隐式转换，含有非数字字符会报错：无效数字

# Oracle服务

![image-20211211091322017](oacle.assets/image-20211211091322017.png)

1、OracleService+服务名（ORCL）:

该服务是Oracle数据库的基础，只有启动该服务才能正常使用Oracle数据库。

2、OracleOraDb11g_home1TNSlistener ：

该服务为Oracle客户端提供监听程序的服务，只有启动该服务，本地的客户端程序才能通过监听连接到数据库，和数据库进行交互。

3、Oracle ORCL VSS Writer Service：

Oracle卷映射拷贝写入服务，VSS(Volume Shadow Copy Service)能够让存储基础设备(比如磁盘，阵列等)创建高保真的时间点映像，即映射拷贝(shadow copy)。它可以在多卷或者单个卷上创建映射拷贝，同时不会影响到系统的性能。(非必须启动)

4、OracleMTSRecoveryService：

服务端控制。该服务允许数据库充当一个微软事务服务器MTS、COM/COM+对象和分布式环境下的事务的资源管理器。(非必须启动)

5、 OracleOraDb11g_home1ClrAgent：

Oracle数据库 .NET扩展服务的一部分。 (非必须启动)

6、 OracleJobSchedulerORCL：

Oracle作业调度(定时器)服务，ORCL是Oracle实例标识。(非必须启动)

## 用户权限

Oracle数据库用户权限分为：系统权限和对象权限两种。

系统权限：比如：create session可以和数据库进行连接权限、create table、create view 等具有创建数据库对象权限。

对象权限：比如：对表中数据进行增删改查操作，拥有数据库对象权限的用户可以对所拥有的对象进行相应的操作。

## 数据库角色

oracle数据库角色是若干系统权限的集合，给Oracle用户进行授数据库角色，就是等于赋予该用户若干数据库系统权限。常用的数据库角色如下：

CONNECT角色：connect角色是Oracle用户的基本角色，connect权限代表着用户可以和Oracle服务器进行连接，建立session（会 话）。

RESOURCE角色：resouce角色是开发过程中常用的角色。 RESOURCE给用户提供了可以创建自己的对象，包括：表、视图、序列、过程、触发器、索引、包、类型等。

DBA角色：DBA角色是管理数据库管理员该有的角色。它拥有系统所有权限，和给其他用户授权的权限。SYSTEM用户就具有DBA权限。

因此，在实际开发过程当中可以根据需求，把某个角色或系统权限赋予某个用户。授权语句如下：

语法：授权

```sql
--GRANT 对象权限 on 对象 TO 用户    
grant select, insert, update, delete on JSQUSER to STUDENT;

--GRANT 系统权限 to 用户
grant select any table to STUDENT;

--GRANT 角色 TO 用户
grant connect to STUDENT;--授权connect角色
grant resource to STUDENT;--授予resource角色
```

语法：取消用户权限

```sql
-- Revoke 对象权限 on 对象 from 用户 
revoke select, insert, update, delete on JSQUSER from STUDENT;

-- Revoke 系统权限  from 用户
revoke SELECT ANY TABLE from STUDENT;

-- Revoke 角色（role） from 用户
revoke RESOURCE from STUDENT;
```

语法：Oracle用户的其他操作

```sql
--修改用户信息
alter user STUDENT
  identified by ******  --修改密码
  account lock;--修改用户处于锁定状态或者解锁状态 （LOCK|UNLOCK ）
```

## Oracle字段数据类型

常用的Oracle列字段的数据类型如下：

| **数据类型 **    | **类型解释 **                                                |
| ---------------- | ------------------------------------------------------------ |
| VARCHAR2(length) | 字符串类型：存储可变的长度的字符串，length:是字符串的最大长度，默认不填的时候是1，最大长度不超过4000。 |
| CHAR(length)     | 字符串类型：存储固定长度的字符串，length:字符串的固定长度大小，默认是1，最大长度不超过2000。 |
| NUMBER(a,b)      | 数值类型：存储数值类型，可以存整数，也可以存浮点型。a代表数值的最大位数：包含小数位和小数点，b代表小数的位数。例子：number(6,2)，输入123.12345，实际存入：123.12 。number(4,2)，输入12312.345，实际春如：提示不能存入，超过存储的指定的精度。 |
| DATA             | 时间类型：存储的是日期和时间，包括年、月、日、时、分、秒。例子：内置函数sysdate获取的就是DATA类型 |
| TIMESTAMP        | 时间类型：存储的不仅是日期和时间，还包含了时区。例子：内置函数systimestamp获取的就是timestamp类型 |
| CLOB             | 大字段类型：存储的是大的文本，比如：非结构化的txt文本，字段大于4000长度的字符串。 |
| BLOB             | 二进制类型：存储的是二进制对象，比如图片、视频、声音等转换过来的二进制对象 |

## 备份查询到的数据：

```
create table 表名 as select 语句
```

## insert插入一个select的结果集

```
 INSERT INTO 表 SELECT 子句，
```



truncate和delete都能删除表中的数据，他们的区别：

- 1、TRUNCATE 是 DDL 命令，命令执行完就提交，删除的数据不能恢复； DELETE 命令是 DML 命令，命令执行完需提交后才能生效，删除后的数据可以通过日志文件恢复。
- 2、如果表中的数据量较大，TRUNCATE的速度比DELETE速度快很多。
- 3、truncate删除将重新设置表索引的初始大小，而delete不能。
- 4、delete能够触发表上相关的delete触发器，而truncate则不会触发。
- 5、delete删除的原理是一次一条从表中删除数据，并将删除操作当做事物记录在数据库的日志当中，以便进行数据回滚。而truncate是一次性进行数据页的删除，因此执行速度快，但是不能回滚。

总结：truncate命令是属于DDL命令，一次性删除表中所有数据，并且数据不能恢复，在实际开发过程当中truncate命令慎用。

# Oracle伪列

Oracle的伪列是Oracle表在存储的过程中或查询的过程中，表会有一些附加列，称为伪列。伪列就像表中的字段一样，但是表中并不存储。伪列只能查询，不能增删改。Oracle的伪列有：rowid、rownum。

------

## ORACLE ROWID

Oracle表中的每一行在数据文件中都有一个物理地址， ROWID 伪列返回的就是该行的物理地址。使用 ROWID 可以快速的定位表中的某一行。 ROWID 值可以唯一的标识表中的一行。通过Oracle select 查询出来的ROWID，返回的就是该行数据的物理地址。

------

ROWID：

```sql
select t.*,t.rowid from stuinfo t ;
select t.*,t.rowid from stuinfo t where t.rowid='AAAShjAAEAAAAEFAAD';
```

## ORACLE ROWNUM

ORACLE ROWNUM表示的Oracle查询结果集的顺序，ROWNUM为每个查询结果集的行标识一个行号，第一行返回1，第二行返回2，依次顺序递增。

ROWNUM 与 ROWID 不同， ROWID 是插入记录时生成， ROWNUM 是查询数据时生成。ROWID 标识的是行的物理地址。 ROWNUM 标识的是查询结果中的行的次序。

------

ROWNUM：

```sql
select t.stuid,t.stuname,t.sex,t.classno,t.stuaddress ,rownum  from stuinfo t ;
```

ROWNUM经常用来限制查询的结果返回的行数，求前几行或前几名的数据。

#  Oracle 函数

Oracle SQL语句中经常使用到Oracle自带的函数，这些函数丰富了SQL的语言功能，为Oracle SQL提供了更多的操作性。Oracle函数可以接受零个或者多个输入参数，并返回一个输出结果。 Oracle 数据库中主要使用两种类型的函数：

1、单行函数：对每一个函数应用在表的记录中时，只能输入一行中的列值作为输入参数(或常数)，并且返回一个结果。

例如1：MOD(X,Y) 是求余函数，返回的X除以Y的余数，其中X和Y可以是列值，也可以是常数。

例如2：TO_CHAR(X,'YYYYMMDD')是时间类型转字符串的函数，其中X可以是行中某一时间类型（date）的列，也可以是一个时间类型的常数。

常用的单行函数大致以下几类：

1. 字符串函数：对字符串进行操作，例如：TO_CHAR()、SUBSTR()、DECODE()等等。
2. 数值函数：对数值进行计算或操作，返回一个数字。例如：ABS()、MOD()、ROUND()等等。
3. 转换函数：将一种数据类型转换成另外一种类型：例如：TO_CHAR()、TO_NUMBER()、TO_DATE()等等。
4. 日期函数：对时间和日期进行操作的函数。例如：TRUNC()、SYSDATE()、ADD_MONTHS()等等。

2、聚合函数：聚合函数同时可以对多行数据进行操作，并返回一个结果。比如 SUM(x)返回结果集中 x 列的总合。

# Oracle字符型函数

Oracle字符型函数是单行函数当中的一种，是用来处理字符串类型的函数，通过接收字符串参数，然后经过操作返回字符串结果的函数。

常用的函数如下表：

| **函数 **                | **说明 **                                                    | **案例 **                                                    | **结果 **      |
| ------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | -------------- |
| ASCII(X)                 | 求字符X的ASCII码                                             | select ASCII('A') FROM DUAL;                                 | 65             |
| CHR(X)                   | 求ASCII码对应的字符                                          | select CHR(65) FROM DUAL;                                    | 'A'            |
| LENGTH(X)                | 求字符串X的长度                                              | select LENGTH('ORACLE技术圈')from DUAL;                      | 9              |
| CONCATA(X,Y)             | 返回连接两个字符串X和Y的结果                                 | select CONCAT('ORACLE','技术圈') from DUAL;                  | ORACLE技术圈   |
| INSTR(X,Y[,START])       | 查找字符串X中字符串Y的位置，可以指定从Start位置开始搜索，不填默认从头开始 | SELECT INSTR('ORACLE技术圈','技术') FROM DUAL;               | 7              |
| LOWER(X)                 | 把字符串X中大写字母转换为小写                                | SELECT LOWER('ORACLE技术圈') FROM DUAL;                      | oracle技术圈   |
| UPPER(X)                 | 把字符串X中小写字母转换为大写                                | SELECT UPPER('Oracle技术圈') FROM DUAL;                      | ORACLE技术圈   |
| INITCAP(X)               | 把字符串X中所有单词首字母转换为大写，其余小写。              | SELECT INITCAP('ORACLE is good ') FROM DUAL;                 | Oracle Is Good |
| LTRIM(X[,Y])             | 去掉字符串X左边的Y字符串，Y不填时，默认的是字符串X左边去空格 | SELECT LTRIM('--ORACLE技术圈','-') FROM DUAL;                | ORACLE技术圈   |
| RTRIM(X[,Y])             | 去掉字符串X右边的Y字符串，Y不填时，默认的是字符串X右边去空格 | SELECT RTRIM('ORACLE技术圈--','-') FROM DUAL;                | ORACLE技术圈   |
| TRIM(X[,Y])              | 去掉字符串X两边的Y字符串，Y不填时，默认的是字符串X左右去空格 | SELECT TRIM('--ORACLE技术圈--','-') FROM DUAL;               | ORACLE技术圈   |
| REPLACE(X,old,new）      | 查找字符串X中old字符，并利用new字符替换                      | SELECT REPLACE('ORACLE技术圈','技术圈','技术交流') FROM DUAL; | ORACLE技术交流 |
| SUBSTR(X,start[,length]) | 截取字符串X，从start位置（其中start是从1开始）开始截取长度为length的字符串，length不填默认为截取到字符串X末尾 | SELECT SUBSTR('ORACLE技术圈',1,6) FROM DUAL;                 | ORACLE         |
| RPAD(X,length[,Y])       | 对字符串X进行右补字符Y使字符串长度达到length长度             | SELECT RPAD('ORACLE',9,'-') from DUAL;                       | ORACLE---      |
| LPAD(X,length[,Y])       | 对字符串X进行左补字符Y使字符串长度达到length长度             | SELECT LPAD('ORACLE',9,'-') from DUAL;                       | ---ORACLE      |

# Oracle日期型函数

Oracle日期类型函数是操作日期、时间类型的相关数据，返回日期时间类型或数字类型结果，常用的函数有：SYSDATE()、ADD_MONTHS（）、LAST_DAY（）、TRUNC()、ROUND()等等。

## 系统日期、时间函数：

**SYSDATE函数：**该函数没有参数，可以得到系统的当前时间。

案例代码：

```
select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') from dual;
```

**SYSTIMESTAMP函数：**该函数没有参数，可以得到系统的当前时间，该时间包含时区信息，精确到微秒。

案例代码：

```
select systimestamp from dual;
```

## 数据库时区函数：

**DBTIMEZONE函数：**该函数没有输入参数，返回数据库时区。

案例代码：

```
select dbtimezone from dual;
```

## 给日期加上指定的月份函数：

**ADD_MONTHS（r,n）函数：**该函数返回在指定日期r上加上一个月份数n后的日期。其中

r：指定的日期。

n：要增加的月份数，如果N为负数，则表示减去的月份数。

案例代码：

```
select to_char(add_months(to_date('2018-11-12','yyyy-mm-dd'),1),'yyyy-mm-dd'),
       to_char(add_months(to_date('2018-10-31','yyyy-mm-dd'),1),'yyyy-mm-dd'),
       to_char(add_months(to_date('2018-09-30','yyyy-mm-dd'),1),'yyyy-mm-dd')        
  from dual;
```

（如果指定的日期是月份的最后一天，返回的也是新的月份的最后一天，如果新的月份比指定的月份日期少，将会自动调回有效日期）

## 月份最后一天函数:

**LAST_DAY(r)函数：**返回指定r日期的当前月份的最后一天日期。

案例代码：

```
 select last_day(sysdate) from dual;
```

## 指定日期后一周的日期函数:

**NEXT_DAY(r,c)函数：**返回指定R日期的后一周的与r日期字符（c：表示星期几）对应的日期。

案例代码：

```
 select next_day(to_date('2018-11-12','yyyy-mm-dd'),'星期四') from dual;
```

## 返回指定日期中特定部分的函数：

**EXTRACT（time）函数：**返回指定time时间当中的年、月、日、分等日期部分。

案例代码：

```
select  extract( year from timestamp '2018-11-12 15:36:01') as year,
        extract( month from timestamp '2018-11-12 15:36:01') as month,        
        extract( day from timestamp '2018-11-12 15:36:01') as day,  
        extract( minute from timestamp '2018-11-12 15:36:01') as minute,
        extract( second from timestamp '2018-11-12 15:36:01') as second
 from dual;
```

## 返回两个日期间的月份数：

**MONTHS_BETWEEN(r1,r2)函数：**该函数返回r1日期和r2日期直接的月份。当r1>r2时，返回的是正数，假如r1和r2是不同月的同一天，则返回的是整数，否则返回的小数。当r1<r2时，返回的是负数。

案例代码：

```
select months_between(to_date('2018-11-12', 'yyyy-mm-dd'),
                      to_date('2017-11-12', 'yyyy-mm-dd')) as zs, --整数
       months_between(to_date('2018-11-12', 'yyyy-mm-dd'),
                      to_date('2017-10-11', 'yyyy-mm-dd')) as xs, --小数
       months_between(to_date('2017-11-12', 'yyyy-mm-dd'),
                      to_date('2018-10-12', 'yyyy-mm-dd')) as fs --负数
  from dual;
```

## 日期截取函数：

**ROUND（r[,f]）函数：**将日期r按f的格式进行四舍五入。如果f不填，则四舍五入到最近的一天。

案例代码：

```
select sysdate, --当前时间
       round(sysdate, 'yyyy') as year, --按年
       round(sysdate, 'mm') as month, --按月
       round(sysdate, 'dd') as day, --按天
       round(sysdate) as mr_day, --默认不填按天
       round(sysdate, 'hh24') as hour --按小时
  from dual;
```

**TRUNC（r[,f]）函数：**将日期r按f的格式进行截取。如果f不填，则截取到当前的日期。

案例代码：

```
select sysdate, --当前时间
       trunc(sysdate, 'yyyy') as year, --按年
       trunc(sysdate, 'mm') as month, --按月
       trunc(sysdate, 'dd') as day, --按天
       trunc(sysdate) as mr_day, --默认不填按天
       trunc(sysdate, 'hh24') as hour --按小时
  from dual;
```

# Oracle数值型函数

Oracle数值型函数可以是输入一个数值，并返回一个数值的函数，我们经常用到函数如下表：

 

| **函数 **    | **解释 **                                                    | **案例 **                                                    | **结果 **        |
| ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ---------------- |
| ABS(X)       | 求数值X的绝对值                                              | select abs(-9) from dual;                                    | 9                |
| COS(X)       | 求数值X的余弦                                                | select cos(1) from dual;                                     | 0.54030230586814 |
| ACOS(X)      | 求数值X的反余弦                                              | select acos(1) from dual;                                    | 0                |
| CEIL(X)      | 求大于或等于数值X的最小值                                    | select ceil(7.8) from dual;                                  | 8                |
| FLOOR(X)     | 求小于或等于数值X的最大值                                    | select floor(7.8) from dual;                                 | 7                |
| log(x,y)     | 求x为底y的对数                                               | select log(2,8) from dual;                                   | 3                |
| mod(x,y)     | 求x除以y的余数                                               | select mod(13,4) from dual;                                  | 1                |
| power(x,y)   | 求x的y次幂                                                   | select power(2,4) from dual;                                 | 16               |
| sqrt(x)      | 求x的平方根                                                  | select sqrt(16) from dual;                                   | 4                |
| round(x[,y]) | 求数值x在y位进行四舍五入。y不填时，默认为y=0;当y>0时，是四舍五入到小数点右边y位。当y<0时，是四舍五入到小数点左边\|y\|位。 | select round(7.816, 2), round(7.816), round(76.816, -1)  from dual; | 7.82 / 8 / 80    |
| trunc(x[,y]) | 求数值x在y位进行直接截取y不填时，默认为y=0;当y>0时，是截取到小数点右边y位。当y<0时，是截取到小数点左边\|y\|位。 | select trunc(7.816, 2), trunc(7.816), trunc(76.816, -1)  from dual; | 7.81 / 7 / 70    |

# Oracle转换函数

Oracle转换函数是进行不同数据类型转换的函数，是我们平常数据库开发过程当中用的最多的内置函数。常用的函数有to_char()、to_number()、to_date()等等。详细分析如下表：

 

| **函数 **                    | **解释 **                                                    | **案例 **                                                    | **结果 **                               |
| ---------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | --------------------------------------- |
| asciistr(x)                  | 把字符串x转换为数据库字符集对应的ASCII值                     | select asciistr('Oracle技术圈')  from dual;                  | Oracle\6280\672F\5708                   |
| bin_to_num(x1[x2...])        | 把二进制数值转换为对应的十进制数值                           | select bin_to_num(1,0,0) from dual;                          | 4                                       |
| cast(x as type)              | 数据类型转换函数，该函数可以把x转换为对应的type的数据类型，基本上用于数字，字符，时间类型安装数据库规则进行互转， | select cast('123' as number) num,cast(123 as varchar2(3)) as ch,cast(to_date('20181112','yyyymmdd') as varchar2(12)) as time  from dual; | 123/'123'/12-11月-18(三列值，用"/"隔开) |
| convert(x,d_chset[,r_chset]) | 字符串在字符集间的转换函数，对字符串x按照原字符集r_chset转换为目标字符集d_chset，当r_chset不填时，默认选择数据库服务器字符集。 | select CONVERT('oracle技术圈','US7ASCII','ZHS16GBK') from dual; | oracle???                               |
| to_char(x[,f])               | 把字符串或时间类型x按格式f进行格式化转换为字符串。           | select to_char(123.46,'999.9') from dual; select to_char(sysdate,'yyyy-mm-dd') from dual; | 123.52018-11-13                         |
| to_date(x[,f])               | 可以把字符串x按照格式f进行格式化转换为时间类型结果。         | select to_date('2018-11-13','yyyy-mm-dd') from dual;         | 2018/11/13                              |
| to_number(x[,f])             | 可以把字符串x按照格式f进行格式化转换为数值类型结果。         | select to_number('123.74','999.99') from dual                | 123.74                                  |

**提醒：**其中数值的格式f可以参考下表：

| **参数 ** | **示例 ** | **说明 **                |
| --------- | --------- | ------------------------ |
| 9         | 999       | 指定位置返回数字         |
| .         | 99.9      | 指定小数点的位置         |
| ，        | 99,9      | 指定位置返回一个逗号     |
| $         | $99.9     | 指定开头返回一个美元符号 |
| EEEE      | 9.99EEEE  | 指定科学计数法           |

# Oracle聚合函数 

Oracle聚合函数同时可以对多行数据进行操作，并返回一个结果。比如经常用来计算一些分组数据的总和和平均值之类，常用函数有AVG()、SUM()、MIN()、MAX()等等。

# Oracle序列

Oracle序列Sequence是用来生成连续的整数数据的对象，它经常用来作为业务中无规则的主键。Oracle序列可以是升序列也可以是降序列。

**创建Oracle序列的语法结构如下：**

```
CREATE SEQUENCE sequence_name
[MAXVALUE num|NOMAXVALUE]
[MINVALUE num|NOMINVALUE]
[START WITH num]
[INCREMENT BY increment]
[CYCLE|NOCYCLE]
[CACHE num|NOCACHE]
```

**语法解析：**

- 1、MAXVALUE/MINVALUE：指定的是序列的最大值和最小值。
- 2、NOMAXVALUE/NOMINVALUE：不指定序列的最大值和最小值，使用系统的默认选项，升序的最大值：10^27次方，降序是-1。升序最小值：1，降序的最小值：-10^26。
- 3、START WITH：指定从某一个整数开始，升序默认是1，降序默认是-1。
- 4、CYCLE | NOCYCLE:表示序列达到最大值或者最小值的时候，是否重新开始。CYCLE：重新开始，NOCYCLE：不重新开始。
- 5、CACHE：使用 CACHE 选项时，该序列会根据序列规则预生成一组序列号。保留在内存中，当使用下一个序列号时，可以更快的响应。当内存中的序列号用完时，系统再生成一组新的序列号，并保存在缓存中，这样可以提高生成序列号的效率 。
- 6、NOCACHE：不预先在内存中生成序列号。

# Oracle视图

oracle视图可以理解为数据库中一张虚拟的表，他是通过一张或者多张基表进行关联查询后组成一个虚拟的逻辑表。查询视图，本质上是对表进行关联查询。

视图的本身是不包含任何数据，只是一个查询结果，当基表的数据发生变化时，视图里面的数据也会跟着发生变化。我们经常在实际开发过程中遇到的视图可以大概分为三种：单表视图、多表关联视图、视图中含有子视图。

## 视图的作用和优势

既然视图在实际开发过程当中被广泛使用到，它到底有哪些作用和优势呢？

**1、使数据简单化**：可以将复杂的查询创建成视图，提供给他人使用，他人就不需要去理解其中复杂性的业务关系或逻辑关系。这样对视图的使用人员来说，就简化了数据的，屏蔽了数据的复杂性。

**2、表结构设计的补充：**系统刚刚开始设计时，大部分程序是直接访问表结构的数据的，但是随着业务的变化、系统的更新等，造成了某些表结构的不适用，这时候去修改表结构对系统的影响太大，开发成本较高，这个时候可以创建视图来对表结构的设计进行补充，降低开发成本。程序可以直接通过查询视图得到想要的数据。

**3、增加安全性：**视图可以把表中指定的字段展示给用户，而不必把表中所有字段一起展示给用户。在实际开发中，视图经常作为数据的提供方式，设置为只读权限提供给第三方人员进行查询使用。

## 创建视图

创建视图的语法结构如下：

```
CREATE [OR REPLACE]  VIEW view_name
AS
SELECT查询
[WITH READ ONLY CONSTRAINT]
```

解释：

1、OR REPLACE：如果视图已经存在，则替换旧视图。

2、WITH READ ONLY：默认不填的，用户是可以通过视图对基表执行增删改操作，但是有很多在基表上的限制（比如：基表中某列不能为空，但是该列没有出现在视图中，则不能通过视图执行 insert 操作，或者基表设置了某些约束，这时候插入视图或者修改视图的值，有可能会报错）， WITH READ ONLY 说明视图是只读视图，不能通过该视图进行增删改操作。但是在现实开发中，基本上不通过视图对表中的数据进行增删改操作。 

# Oracle索引

Oracle索引（index）最大的作用是用来优化数据库查询的效率，提升数据库的查询性能。就好比书的目录一样，可以通过目录来直接定位所需内容存在的页数，大大提高检索效率。

Oracle数据库中如果某列出现在查询的条件中，而该列的数据是无序的，查询时只能从第一行开始一行一行的匹配。创建索引就是对某些特定列中的数据进行排序或归类，生成独立的索引表。在某列上创建索引后，如果该列出现在查询条件中，Oracle 会自动的引用该索引，先从索引表中查询出符合条件记录的 ROWID，由于 ROWID 是记录的物理地址，因此可以根据 ROWID 快速的定位到具体的记录，当表中的数据非常多时，引用索引带来的查询效率非常可观 。

## 何时建立索引：

既然我们都知道建立索引有利于查询速率的提升，那是不是所有字段都可以加上索引。这是万万不行的，建立索引不仅仅要浪费空间来存储索引表，当数据量较少时，直接查询数据比经过查询索引表再定位到表数据的速度更快。索引可以提高查询的效率，但是在数据增删改时需要更新索引，因此索引对增删改时会有负面影响。所以要根据实际情况， 考虑好再建立索引。

那何时建立索引，下面大概介绍几点，其余的得在实际应用和开发过程中，酌情考虑：

1、Oracle 数据库会为表的主键和包含唯一约束的列自动创建索引，所以在建立唯一约束时，可以考虑该列是否必要建立。是否经常要作为查询条件。

2、如果某个表的数据量较大（十几二十万以上），某列经常作为where的查询条件，并且检索的出来的行数经常是小于总表的5%，那该列可以考虑建立索引。

3、对于两表连接的字段，应该考虑建立索引。如果经常在某表的一个字段进行Order By 则也经过进行索引。

4、不应该在小表上建立索引。上面也说过，小表之间查询的数据会比建立索引的查询速度更快，但是在某些字段，如性别：只有男、女和未知三种数据时，可以考虑位图索引，可以增加查询效率。

5、经常进行DML操作，即经常进行增删改的操作的表，创建表索引时就要权衡一下，因为建索引会导致进行DML操作时速度变慢。所以可以根据实际情况，选择某些字段建立索引，而不能盲目乱建。

## 索引的类别：

适当的使用索引可以提高数据检索速度，那Oracle有哪些类型的索引呢？

**1、b-tree索引：**Oracle数据中最常见的索引，就是b-tree索引，create index创建的normal就是b-tree索引，没有特殊的必须应用在哪些数据上。

**2、bitmap位图索引：**位图索引经常应用于列数据只有几个枚举值的情况，比如上面说到过的性别字段，或者我们经常开发中应用的代码字段。这个时候使用bitmap位图索引，查询效率将会最快。

**3、函数索引：**比如经常对某个字段做查询的时候经常是带函数操作的，那么此时建一个函数索引就有价值了。例如：trim（列名）或者substr(列名)等等字符串操作函数，这个时候可以建立函数索引来提升这种查询效率。

**4、hash索引：**hash索引可能是访问数据库中数据的最快方法，但它也有自身的缺点。创建hash索引必须使用hash集群，相当于定义了一个hash集群键，通过这个集群键来告诉oracle来存储表。因此，需要在创建HASH集群的时候指定这个值。存储数据时，所有相关集群键的行都存储在一个数据块当中，所以只要定位到hash键，就能快速定位查询到数据的物理位置。

**5、reverse反向索引：**这个索引不经常使用到，但是在特定的情况下，是使用该索引可以达到意想不到的效果。如：某一列的值为{10000,10001,10021,10121,11000,....}，假如通过b-tree索引，大部分都密集分布在某一个叶子节点上，但是通过反向处理后的值将变成{00001,10001,12001,12101,00011,...}，很明显的发现他们的值变得比较随机，可以比较平均的分部在各个叶子节点上，而不是之前全部集中在某一个叶子节点上，这样子就可大大提高检索的效率。

**6、分区索引和分区表的全局索引：**这两个索引是应用在分区表上面的，前者的分区索引是对分区表内的单个分区进行数据索引，后者是对分区表的全表进行全局索引。分区表的介绍，可以后期再做单独详解，这里就不累述了。

## 索引的创建

**语法结构：**

```
create[unique]|[bitmap] index index_name --UNIQUE表示唯一索引、BITMAP位图索引
on table_name(column1,column2...|[express])--express表示函数索引
[tablespace tab_name] --tablespace表示索引存储的表空间
[pctfree n1]    --索引块的空闲空间n1
[storage         --存储块的空间
 (
    initial 64K  --初始64k
    next 1M
    minextents 1
    maxextents unlimited

)];
```

**语法解析：**

1、UNIQUE:指定索引列上的值必须是唯一的。称为唯一索引，BITMAP表示位图索引。

2、index_name：指定索引名。

3、tabl_name：指定要为哪个表创建索引。

4、column_name：指定要对哪个列创建索引。我们也可以对多列创建索引，这种索引称为组合索引。也可以是函数表达式，这种就是函数索引。

------

**修改索引：**

1、重命名索引：

```
alter index index_old rename to index_new;--重新命名索引
```

2、合并索引、重新构造索引：我们索引建好后，经过很长一段时间的使用，索引表中存储的空间会产生一些碎片，导致索引的查询效率会有所下降，这个时候可以合并索引，原理是按照索引规则重新分类存储一下，或者也可以选择删除索引重新构造索引。

```
alter index index_name coalesce;--合并索引
alter index index_name rebuild;--重新构造
```

------

**删除索引：**

```
drop index index_name;
```

------

**查看索引：**

```
select t.INDEX_NAME,--索引名字
       t.index_type,--索引类型
       t.TABLESPACE_NAME,--表空间
       t.status,--状态
       t.UNIQUENESS--是否唯一索引
  from all_indexes T 
  where  t.INDEX_NAME='index_name';
```

# Oracle分区详解和创建

Oracle在实际业务生产环境中，经常会遇到随着业务量的逐渐增加，表中的数据行数的增多，Oracle对表的管理和性能的影响也随之增大。对表中数据的查询、表的备份的时间将大大提高，以及遇到特定情况下，要对表中数据进行恢复，也随之数据量的增大而花费更多的时间。这个时候，Oracle数据库提供了分区这个机制，通过把一个表中的行进行划分，归为几部分。可以减少大数据量表的管理和性能问题。利用这种分区方式把表数据进行划分的机制称为表分区，各个分区称为分区表。

------

Oracle分区对于大型表（大数据量）非常有用，分区的作用主要有：

1、改善大型表的查询性能，因为可以通过查询对应分区表中对应的数据，而不需要查询整个表。

2、表更容易管理，因为分区表的数据存储在各个分区中，所以可以按照分区建，来管理对应分区当中的数据，可以按照分区加载和删除其中的数据，比在不分区情况下，更容易管理数据。以及在特定的事故情况下，通过备份好的分区，可以快速恢复对应分区当中的数据，也不需要对全表数据进行恢复。

# oracle merge into命令

Oracle merge into命令，顾名思义就是“有则更新，无则插入”，这个也是merge into 命令的核心思想，在实际开发过程中，我们会经常遇到这种通过两表互相关联匹配更新其中一个表的某些字段的业务，有时还要处理不匹配的情况下的业务。这个时候你会发现随着表的数据量增加，类似这种业务场景的执行效率会比较慢，那是因为你需要多次重复查询两表中的数据，而通过merge into命令，只需要一次关联即可完成“有则更新，无则插入”的业务场景，大大提高语句的执行效率。

merge into命令的语法结构如下：

```
merge into A 
using B 
on (A.id = B.id)
when matched then
  update set A.col=B.col
when not matched then
  insert 语句;
```

语法解析：利用B表通过A.id=B.id的条件来匹配A表，当满足条件时，可以对A表进行更新，当不满足条件时，可以利用inert语句插入相关数据。

# oracle物化视图

Oracle物化视图可以理解为用来存储数据表的副本或者聚集，物化视图可以用来复制一个表的全部数据或者部分数据，也可以复制多表关联查询的结果集。然后按照指定的时间间隔内自动更新副本数据的刷新。本文将介绍物化视图常用的场景和创建方法，以及刷新策略。

物化视图是基于select查询结果的数据副本，简单来讲就是数据的复制表，但是它可以按照特定的刷新策略刷新复制表中的数据。常用来：

1、物化视图经常用来当做远程数据表的本地化，可以用来做数据同步用。

2、也经常用来做本地单表或多表数据的汇总，这样子，可以定时把汇总数据更新到复制表，用来做决策分析用。

3、也可以用来做复杂视图的物化作用，把这种数据不经常实时更新的视图，物化成物理表，然后在物化表上建相应的索引，大大提高查询效率。

物化视图在创建时，可以指定刷新（refresh interval）的间隔来执行定时刷新物化表的数据。也可以利用基于事务的刷新，通过主表变化的行数实时来更新物化表。

## 创建物化视图语法：

**语法：**

```
create materialized view view_name
refresh [fast|complete|force]
[
on [commit|demand] |
start with (start_time) next (next_time)
]
AS select 查询语句;
```

**语法解析：**

1、create materialized view是创建物化视图的命令关键字。view_name指的是要创建物化视图的名字。

2、refresh 是指定物化视图数据刷新的模式：

force：这个是默认的刷新方式，Oracle会自动判断该如何进行数据刷新，如果满足快速刷新的条件，就进行fast刷新，不满足就进行全量刷新。

complete:这个是全量刷新数据，当要刷新数据时，会删除物化表中的所有数据，然后物化视图中的查询语句重新生成数据。

fast:采用增量的方式进行数据的刷新。通过主表的视图日志和上次刷新数据进行比对，然后增量更新物化后的数据表。

3、on[comit|demand]，指的是选择物化视图数据刷新的模式。

on commit指的是基表的数据有提交的时候触发刷新物化视图，使得物化视图中的数据和基本的数据是一致性的。

on demand指定是物化视图数据刷新的另外一种方式，仅在需要更新数据时才刷新物化视图中的数据。一般配合 start with 参数按照特定的时间计划，按计划更新数据，比如每天晚上12点或每周更新一次物化视图中的数据。

# oracle分析函数_开窗函数详解

Oracle分析函数是Oracle系统自带函数中的一种，是Oracle专门用来解决具有复杂统计需求的函数，它可以对数据进行分组然后基于组中数据进行分析统计，最后在每组数据集中的每一行中返回这个统计值。

Oracle分析函数不同于分组统计（group by），group by只能按照分组字段返回一个固定的统计值，但是不能在原来的数据行上带上这个统计值，而Oracle分析函数正是Oracle专门解决这类统计需求所开发出来的函数。

Oracle分析函数都会带上一个开窗函数over（）。

**Oracle分析函数的语法结构：**

```
select table.column, 
Analysis_function()OVER(
[partition by 字段]
[order by 字段 [windos]]
) as 统计值
from table
```

语法解析：

  1、Analysis_function：指定分析函数名，常用的分析函数有sum、max、first_value、last_value、rank、row_number等等。

  2、over()：开窗函数名，partition by指定进行数据分组的字段，order by指定进行排序的字段，windos指定数据窗口（即指定分析函数要操作的行数），使用的语法形式大概如下：

```
over(partition by xxx order by yyy rows between zzz)
```

# Oracle行转列（PIVOT）

Oracle11g之后提供了自带函数PIVOT可以完美解决这个行转列的需求，具体语法结构如下：

```
SELECT * FROM （数据查询集）
PIVOT 
(
 SUM(Score/*行转列后 列的值*/) FOR 
 coursename/*需要行转列的列*/ IN (转换后列的值)
)
```

# Oracle列转行_unpivot

利用Oracle自带的列转行函数unpivot也可以完美解决该问题，具体语法结构如下：

```
select 字段 from 数据集
unpivot（自定义列名/*列的值*/ for 自定义列名 in（列名））
```

# Oracle游标

Oracle游标是通过关键字CURSOR的来定义一组Oracle查询出来的数据集，类似数组一样，把查询的数据集存储在内存当中，然后通过游标指向其中一条记录，通过循环游标达到循环数据集的目的。

------

## **游标的种类**

Oracle游标可以分为显式游标和隐式游标两种之分。

**显式游标：**指的是游标使用之前必须得先声明定义，一般是对查询语句的结果事进行定义游标，然后通过打开游标循环获取结果集内的记录，或者可以根据业务需求跳出循环结束游标的获取。循环完成后，可以通过关闭游标，结果集就不能再获取了。全部操作完全由开发者自己编写完成，自己控制。

**隐式游标：**指的是PL/SQL自己管理的游标，开发者不能自己控制操作，只能获得它的属性信息。

------

## 显式游标

显式游标在实际开发中经常使用到，可以丰富PL/SQL的开发程序的编写，实现一些循环类的复杂业务。游标的使用步骤如下：

**1、声明游标：**

声明游标是给游标命名并给游标关联一个查询结果集，具体声明语法如下：

```
declare cursor cursor_name(游标名)
is select＿statement(查询语句);
```

**2、打开游标：**

游标声明完，可以通过打开游标打开命令，初始化游标指针，游标一旦打开后，游标对应的结果集就是静态不会再变了，不管查询的表的基础数据发生了变化。打开游标的命令如下：

```
open cursor_name;
```

**3、读取游标中数据：**

读取游标中的数据是通过fetch into语句完成，把当前游标指针指向的数据行读取到对应的变量中（record 变量）。游标读取一般和循环LOOP一起使用，用于循环获取数据集中的记录。**
**

```
fetch cursor_name into record变量
```

**4、关闭游标：**

游标使用完，一定要关闭游标释放资源。关闭后，该游标关联的结果集就释放了，不能够再操作了，命令如下：

```
close cursor_name;
```

# Oracle触发器

Oracle触发器是使用者对Oracle数据库的对象做特定的操作时，触发的一段[PL/SQL程序代码](http://www.oraclejsq.com/plsql/010200490.html)，叫做触发器。触发的事件包括对表的DML操作，用户的DDL操作以及数据库事件等。

------

## 触发器的作用

Oracle触发器可以根据不同的数据库事件进行特定的调用触发器程序块，因此，它可以帮助开发者完成一些PL/SQL存储过程完成不了的问题，比如操作日志的记录、防止一些无效的操作、校验数据的正确性、限制一些对数据库对象的操作、提供数据同步的可行性。但是不推荐在触发器当中写业务逻辑程序，因为这样对后期数据的维护将大大提高成本。

------

## 触发器的类型

触发器按照用户具体的操作事件的类型，可以分为5种触发器。大致如下：

- **1、[数据操作（DML）触发器](http://www.oraclejsq.com/plsql/010200566.html)**：此触发器是定义在Oracle表上的，当对表执行insert、update、delete操作时可以触发该触发器。如果按照对表中行级数据进行触发或语句级触发，又可以分为行级（row）触发器，语句级触发器，按照修改数据的前后触发触发器，又可以分为 after 触发器和before触发器之分。
- **2、[数据定义操作（DDL）触发器](http://www.oraclejsq.com/plsql/010200598.html)**：当对数据库对象进行create、alter、drop操作时，触发触发器进行一些操作记录保存、或者限定操作。
- **3、用户和系统事件触发器：**该类型的触发器是作用在Oracle数据库系统上，当进行数据库事件时，触发触发器，一般用来记录登录的相关信息。
- **4、INSTEAD OF 触发器**：此类型的触发器是作用在视图上，当用户对视图进行操作时，触发该触发器把相关的操作转换为对表进行操作。
- **5、复合触发器**：指的是对数据操作(DML)触发器当中的多种类型触发器进行复合，比如；一个触发器当中包含着after（或before）的行级触发器和after（或before）的语句级触发器，来完成一些更为复杂的操作。

# Oracle存储过程

Oracle存储过程在实际数据库开发过程当中会经常使用到，作为一个数据库开发者必备的技能，它有着SQL语句不可替代的作用。所谓存储过程，就是一段存储在数据库中执行某块业务功能的程序模块。它是由一段或者多段的PL/SQL代码块或者SQL语句组成的一系列代码块。

**创建Oracle存储过程语法：**

```
create [or replace] procedure 过程名
( p1 in|out datatype,
  p2 in|out datatype,
  ...
  pn in|out datatype
    
) is 
    
    ....--声明部分

    begin
    
    ....--过程体

    end;
```

**语法解析：**

1、procedure 关键字是创建存储过程的命令。

2、create [or replace] :如果存储过程已经存在则覆盖替代原有的过程。

3、in|out ：存储过程具有入参和出参两种参数选择，in表示的是入参，out表示的是出参，在使用过程的时候，入参必须得有对应的变量传入，出参得有对应的变量接收。

4、datatype表示出入参变量对应的数据类型。

5、is后面跟着的是过程当中使用到的声明变量。

6、begin...end 中间编写的就是存储过程的具体操作。

# Oracle锁

Oracle锁是用于数据共享的情景当中，它是一种Oracle的访问机制，在访问同一个资源时，防止不同事物操作同一个数据时，出现数据问题。利用Oracle锁机制，多个会话操作同一个数据时，优先的会话会锁定该数据，其它会话只能等待。

Oracle锁就是[事物](http://www.oraclejsq.com/plsql/010200600.html)的隔离性，当前的事物不能影响其它事物。Oracle锁是Oracle自动管理的，它锁定的最小粒度是数据行记录，锁定时，其它事物不能操作该数据记录，但是可以读取锁定时的记录值。因此，Oracle锁的作用是保证数据的完整性。

------

## Oracle锁的类型

Oracle锁可以分为排它锁和共享锁两种分类：

**1、排它锁**：可以理解为写锁，这种锁是防止资源的共享，用于修改数据。如果一个事物给某个数据加了排它锁，其它事物就不能对它再加任何锁，直到事物完结，排它锁释放。

**2、共享锁**：可以理解为读锁，加了共享锁的数据，只能共享读，不能再给它加排它锁进行写的操作。

------

Oracle正是使用锁机制来实现系统的高并发，利用不同类型的排它锁或者共享锁来管理并发会话对数据的操作，Oracle锁的类型有：

**1、DML锁**：该类型的锁称为数据锁，控制DML数据操作。

**2、DDL锁**：该类型的锁用来保护数据库对象结构的。

其中，DML锁又可以细分为行级锁和表级锁两种：

**1、行级锁**：行级锁是粒度最细的DML锁，主要用来控制数据行的修改、删除操作。当对表中的某行数据进行修改时，需要对其加上行级排他锁，防止其他事物也对其进行修改，等数据修改完，事物提交时，自动释放。

**2、表级锁**：主要作用是在修改表数据时，会对表结构加一个表级锁，这个时候不应许其它事物对表进行修改或者删除（DDL操作）。

# OEM介绍

OEM(Oracle enterprise manager)是Oracle数据库提供的一个企业管理的web界面，用来管理和监控Oracle数据库运行情况的组件。

OEM叫做企业管理器，它为企业开发者管理Oracle数据库提供了一种快捷、简便的管理方式。同时也为学习者提供了一个了解Oracle相关性能和特性的平台。

OEM一般在我们安装完Oracle数据库就会自带安装上的一组组件，假如没有安装成功，也可以通过重新独立安装OEM进行重新安装，

OEM主要的功能菜单有：

1、主目录：显示数据库当前的状态，主要显示内容有监听程序、CPU负载、连接会话数、SQL响应时间、空间使用率等是否正常。

2、[性能](http://www.oraclejsq.com/oraclegl/010300644.html)：主要以图表的形式显示数据库的CPU、I/O、进程数等运行情况。查看Oracle数据库是否有过载情况，可以查看具体原因，进而进行调整数据库的软硬件。

3、[可用性](http://www.oraclejsq.com/oraclegl/010300646.html)：主要提供数据库的备份和恢复的一些功能。

4、[服务器](http://www.oraclejsq.com/oraclegl/010300648.html)：提供一些表空间、数据文件、归档日志的信息管理。

5、[方案](http://www.oraclejsq.com/oraclegl/010300650.html)：主要是对数据库对象进行一些日常管理，包括常用的[表](http://www.oraclejsq.com/article/010100139.html)、[视图](http://www.oraclejsq.com/article/010100440.html)、[过程](http://www.oraclejsq.com/plsql/010200560.html)、[函数](http://www.oraclejsq.com/plsql/010200558.html)、包体、[自定义类型](http://www.oraclejsq.com/plsql/010200492.html)等数据库对象的信息管理。

6、[数据移动](http://www.oraclejsq.com/oraclegl/010300652.html)：主要是针对DMP文件的导入导出和移动数据文件等功能。

7、软件和支持：主要用于配置一些Oracle的主机信息、进行Oracle部署，包括RAC的部署。

## 独立安装OEM

原理上安装完Oracle数据库，会自动安装上OEM管理器，但是由于各种原因，会经常出现OEM安装不成功的情况。因此可以通过独立安装OEM来实现。步骤如下：

1、登录Oracle数据库解锁并重新设置sysman、DBSNMP用户的密码，代码如下：

```
alter user sysman account unlock identified by ***;
alter user dbsnmp account unlock identified by ***;
```

2、在cmd命令框当中通过命令建立OEM，然后根据命令提示输入数据库的SID，和对应的用户密码，命令如下：

```
emca -repos create
```

3、配置OEM，通过命令进行初始化OEM数据仓库的数据，然后按照命令提示输入数据库的SID、对应的安装目录和用户名密码，命令如下：

```
emca -config dbcontrol db
```

4、启动、关闭OEM的命令：

```
emctl start/stop dbconsole
```

通过上面步骤重新安装完OEM后，就会在数据库的安装目录中生成OEM的启动地址链接。

# Oracle控制文件

Oracle控制文件是Oracle数据库存储信息的重要文件，它是一个二进制文件，控制文件主要用来存放数据库名字、数据文件位置等信息的文件。Oracle控制文件是至关重要的，没有了它，数据库就不能启动。

每一个数据库都有一个Oracle控制文件，而且每一个控制文件只属于一个数据库，不能拿来共用。控制文件在数据库创建时跟着一起创建，控制文件不能手动修改，Oracle数据库自己独立管理。

那控制文件的状态和存放位置是如何的呢？我们可以根据数据字典V$controlfile进行查询：

```
select * from v$controlfile;
```

通过查询结果，可以看出控制文件的扩展名是.ctl文件。每一个控制文件都记录着Oracle数据库的创建时间、名称、数据文件的名字、数据文件的位置、日志文件的名字及位置、表空间、备份、最近检查点等信息。因此在对数据库进行相应的操作时，比如增加数据文件时，就会更新对应的控制文件信息，而不是手动进行修改。

## 控制文件多路复用 

既然控制文件这么重要，我们该如何对它进行保护呢？Oracle数据库提供了多路复用机制对控制文件进行保护。多路复用就是把控制文件进行复制创建在不同磁盘上，这样子可以防止一个磁盘在损坏的情况下，可以从其它磁盘上进行恢复。保证Oracle数据库的安全性。

可以使用init.ora文件对控制文件进行多路复用，init.ora是Oracle数据库初始化文件，它也是Oracle创建时，就自动创建的一个文件，它里面包含了控制文件的位置信息，init.ora文件在Oracle数据库安装目录下dbs文件下。

在修改init.ora之前可以对Oracle控制文件进行备份，然后再修改init.ora文件中control_files参数。步骤如下：

1、查看参数文件spfile位置，把参数文件转换为可以编辑文件pfile

```
--查看参数文件spfile位置
show parameter spfile;
--创建pfile
create pfile from spfile;
```

2、关闭数据库。

3、对控制文件进行备份，然后对pfile文件进行修改，把contol_files参数添加上备份的控制文件。

4、再创建spfile文件，然后重启数据库。

```
create spfile from pfile;
```

![image-20211211141724437](oacle.assets/image-20211211141724437.png)

通过查询控制文件V$controlfile可以查询正好有两个控制文件，和init.ora文件中两个控制文件一一对应。因此可以通过备份控制文件，然后通过init.ora文件对控制文件进行多路复用，从而起到保护Oracle数据库的作用。

# Oracle日志文件

Oracle日志文件是Oracle数据库存储信息的重要文件，主要用来存储数据库变化的操作信息。

Oracle日志文件可以分为两种：重做日志文件（redo log file）、归档日志文件，其中重做日志文件主要记录了数据库的操作过程，可以在进行数据库恢复时，将重做日志文件在还原的数据库上进行执行，以达到数据库的最新状态。

Oracle数据库在运行时，可以选择是否开启归档日志，在非归档日志的情况下，Oracle所有的操作日志都写在重做日志当中，当所有重做日志文件写满时（Oracle重做日志是分组的，默认是分为三组），那么就把前面的日志文件覆盖继续写入。而在开启归档日志模式情况下，当重做日志都写满时，继续要写入日志时，会把要覆盖的日志文件写入归档日志当中，然后再对重做日志进行覆盖，因此使用归档日志利于后期进行数据恢复。

------

那要怎么查看数据库是否开启归档日志？可以根据以下命令进行查询：

```
select t.NAME,t.LOG_MODE from v$database t;
```

**代码解析**：V$database数据字典主要存储数据库创建后的一些配置信息，其中LOG_MODE字段记录的就是是否开启归档日志文件，NOARCHIVELOG：表示未开启归档日志文件（Oracle在安装时默认选择的是不开启归档日志模式的）。

------

那重做日志文件的位置和属性信息如何查询，我们可以根据数据字典V$logfile进行查询，代码如下：

```sql
select *from v$logfile;
```

# Oracle添加日志文件组

Oracle通过添加日志文件组达到扩充日志文件组的目的，然后再进行对其添加特定的日志文件，就达到扩充日志文件数的目的，解决了日志文件过少的性能问题。

Oracle增加日志文件组可以使用[OEM企业管理器](https://www.oraclejsq.com/oraclegl/010300640.html)、SQL命令两种方式进行添加。这里我们采用SQL命令方式进行讲解，OEM企业管理器方式读者门可以自行试验。

**创建日志文件组语法结构：**

```
alter database
add logfile group n
filename size m
```

**语法解析：**

1、group n：表示创建日志文件组的组号，在Oracle当中日志文件组的组号是唯一的。

2、filename：表示日志文件组存储的位置。

3、size m：表示日志文件组的大小，默认是50M大小。

------

下面我们增加一个新的日志文件组4，大小100M，代码如下：

```
alter database
add logfile group 4
'E:\APP\ADMIN\ORADATA\ORCL\NEWREDO04.LOG' size 100M;
```

# Oracle添加日志文件进文件组

Oracle可以通过添加日志文件进特定的文件组以达到扩充日志文件数的目的，解决了日志文件过少的性能问题。

Oracle数据库中一个日志文件组可以包含多个日志文件，但是必须包含一个文件。因此，我们可以对建好的日志文件组进行扩充，语法结构如下：

```
alter database
add logfile member
filename to group n;
```

**语法解析：**

1、filename：表示日志文件存储的位置。

2、group n：表示添加到那个日志文件组。

------

利用上面的命令，我们对日志文件组4再添加一个日志文件，命令如下：

```
 alter database
add logfile member
'E:\APP\ADMIN\ORADATA\ORCL\NEWREDO05.LOG' to group 4;
```

执行完代码，再查下一下日志文件，发现多了一个新添加的日志文件。

# Oracle删除日志文件

不仅仅可以对日志文件组中进行添加日志文件，也可以把不必要的日志文件进行删除。

Oracle删除日志文件的语法结构如下：

```
alter database
drop logfile member
filename
```

**语法解析：**

1、filename：日志文件存储的位置。

2、当日志文件组只有一个日志文件时，是不能进行删除日志文件的，必须把整个日志文件组进行删除。

------

利用删除日志文件的命令，把刚刚添加的日志文件NEWREDO05进行删除，代码命令如下：

```
 alter database
drop logfile member
'E:\APP\ADMIN\ORADATA\ORCL\NEWREDO05.LOG'
```

执行完代码，在查下日志文件，发现日志文件NEWREDO05已经删除了

# Oracle删除日志文件组

我们不仅仅可以增加日志文件组，也可以对于过多的日志文件组进行整组删除。

Oracle删除日志文件组的语法结构：

```
alter database
drop logfile
group n
```

**语法解析：**

1、group n：表示要删除的文件组组号。

2、删除文件组，会对应的里面的日志文件一并全部删除。

------

利用删除日志文件组命令，把文件组4整组删除，代码如下：

```
alter database
drop logfile
group 4;
```

执行完代码，再查看下日志文件，发现日志文件组4，已经全部删除。

# Oracle表空间

Oracle表空间是Oracle数据对象和数据存储的容器，Oracle表空间经常和数据文件成对出现，一个表空间可以对应多个数据文件，而一个数据文件只能在一个表空间当中。我们在创建表空间时，就会默认创建一个数据文件，同理，我们创建数据文件时，必须指定一个表空间。

Oracle数据库存储数据是有一个个表空间组成的，一个表空间当中存储着多个数据文件，Oracle的数据（表、索引等数据）存储在数据文件当中，在表空间当中的逻辑单位是段（segment）,例如：

我们创建一个索引时，会在指定表空间的创建一个以索引名字命名的索引段，然后在索引段当中创建一个或者多个区（extent），用来存储索引数据，一个区段只能存在于一个数据文件当中。再细分，一个区段当中，可以分为多个区块（block）。区块是Oracle数据库当中最小的空间分配单位。

一个文件在磁盘空间当中存储一般都不是连续的，因此，表空间当中的段是由不同数据文件当中的区段组成的。

------

## 默认表空间

Oracle安装完后（笔者采用的是Oracle11g），会有五个个默认的表空间，分别是：

**SYSAUX：**安装Oracle11g示例的数据库表空间。

**SYSTEM：**存储sys/system用户表、存储过程、视图等数据库对象。

**UNDOTBS1：**用于存储撤销（用于回滚）的信息。

**TEMP：**临时表空间，用于存储SQL语句处理的表和索引信息。

**USERS：**存储数据库用户创建的数据库对象信息。

------

## 查看表空间

想查看数据库所有默认表空间，可以通过数据字典dba_tablespaces进行查询:

```sql
select * from dba_tablespaces;
```



## 查看指定用户的默认表空间

如果想查看指定用户的默认表空间可以通过数据字典DBA_USERS进行查询，下面我利用查询语句，分别查询出SYS、SYSTEM、STUDENT(笔者创建的)用户分别对应的表空间是什么，查询代码如下：

```
select T.username,--用户名
       T.account_status,--用户状态
       T.default_tablespace,--默认表空间
       T.temporary_tablespace,--临时表空间
       T.created--创建时间
  from dba_users t
 where t.username in ('SYS', 'SYSTEM', 'STUDENT')
```

# Oracle创建表空间

Oracle创建[表空间](https://www.oraclejsq.com/oraclegl/010300758.html)是数据库管理员经常要做的事情，在实际当中，一般独立的业务系统会有一个独立的用户进行独立开发管理，附带的会独立创建一个自己的表空间进行存储。

Oracle创建表空间可以通过OEM企业管理器、SQL命令两种方式进行创建，笔者这里采用SQL命令方式进行讲解，OEM方式读者可以自行登录OEM后台自行试验。

------

**Oracle创建表空间语法结构如下：**

```
create tablespace tab_name
datafile 'filename'
size n
[autoextend on next n1 maxsize m /of]
[permanent] 
[extent management local/dictionary];
```

**语法解析：**

**create tablespace：**创建表空间的关键字。

**tab_name：**创建后表空间的名字。

**datafile：**指定数据文件的路径为filename。

**size n：**指定数据文件的大小。

**[autoextend on next n1 maxsize m /of ]：**表示表空间是否是自动扩展的，on 为自动扩展，of为不扩展，当自动扩展时，next n1表示自动扩展的大小,max size m 表示数据文件最大扩展到m大小。

**[permanent] ：**表示创建的表空间的类型，permanent表示永久表空间，不填都是默认永久表空间。

**[extent management local/dictionary]：**表示表空间管理的方式，local表示本地的管理模式，dictionary表示数据字典管理模式，默认都是本地管理方式。

------

**案例1、**根据表空间创建语法，创建一个100M大小数据文件（student.dbf）的表空间student，代码如下：

```sql
create tablespace student
datafile 'E:\APP\ADMIN\ORADATA\ORCL\student.DBF'
size 100m
autoextend on next 10m maxsize 500m
permanent
extent management local;
```

**案例解析：**

创建一个student表空间，指定了数据文件为“E:\APP\ADMIN\ORADATA\ORCL\student.DBF”，表空间是自动扩展的，每次自动扩展大小为10M，最大扩展到500M，创建的是永久表空间，用来存储student用户的数据库对象和数据，管理模式为本地管理。

我们查看数据字典dba_data_files和dba_tablespaces对创建好后的student表空间进行查询，查询代码如下：

```sql
select t.TABLESPACE_NAME, --表空间名
       t.FILE_NAME, --文件名
       t.AUTOEXTENSIBLE, --是否自动扩展
       t.BYTES / 1024 / 1024, --表空间初始大小
       t.MAXBYTES / 1024 / 1024, --表空间最大扩展到多少
       b.CONTENTS, --表空间类型
       b.EXTENT_MANAGEMENT --表空间管理模式
  from dba_data_files t, dba_tablespaces b
 where t.TABLESPACE_NAME = b.TABLESPACE_NAME
```

# Oracle删除表空间

Oracle删除表空间的操作经常发生在数据库部分业务拆分的情况下，会把不必要的[表空间](https://www.oraclejsq.com/oraclegl/010300758.html)和对应的数据文件删除，释放当前的数据库的硬件空间。

删除表空间可以通过OEM企业管理、SQL命令两种方式进行直接删除。Oracle删除表空间的时候不需要先删除数据文件，再删除表空间，可以选择删除表空间时，把数据文件一并删除。

------

**Oracle删除表空间语法结构：**

```sql
drop tablespace tab_name [including contents][cascade constraints]
```

**语法解析：**

**drop tablespace：**删除表空间的关键字，tab_name表示表空间名字。

**[including contents]：**表示在删除表空间的时候把表空间中的数据文件一并删除。

**[cascade constraints]：**表示在删除表空间的时候把表空间的完整性也一并删除。比如表的外键，和触发器等就是表的完整性约束。

------

**案例1、**删除student表空间，并删除表空间的数据文件和完整性，代码如下：

```sql
drop tablespace student 
including contents
cascade constraints;
```

再查下一下表空间，发现student表空间不存在了

# Oracle临时表空间

表空间是Oracle数据库存储数据和对象的逻辑容器，那临时表空间呢？

Oracle临时表空间主要是存储数据库的排序操作、临时表、中间排序结果等临时对象。例如，我们进行大数量级的排序操作时，当数据库内存不够时，就会写入临时表空间，当操作完成后，临时表空间就会自动清空释放。Oracle经常使用到临时表空间的操作有：create index（创建索引）、group by(分组查询)、order by(排序时)、集合运算时（union、minus、intersect）、多表连接查询时，当数据库内存不足时，会用到临时表空间。

------

## 创建临时表空间

Oracle数据库在安装完后就会创建一个默认的临时表空间temp。Oracle创建临时表空间的语法结构和创建持久化表空间一样，只是多了关键字**temporary**进行创建临时表空间。

**创建临时表空间语法：**

```sql
create temporary  tablespace tempname
tempfile 'filename'
size m;
```

**语法解析：**

1、create temporary tablespace：表示创建临时表空间，tempname表示创建临时表空间的名字。

2、filename：指定临时表空间数据文件的位置。

3、size m：表示临时表空间的大小。

------

案例1、创建临时表空间temp1，代码如下：

```sql
create temporary  tablespace temp1
tempfile 'E:\APP\ADMIN\ORADATA\ORCL\temp1.DBF'
size 50m;
```

创建好临时表空间temp1，我们可以通过数据字典dba_temp_files进行查询临时表空间的信息，查询代码如下：

```sql
select t.TABLESPACE_NAME, --表空间名
       t.FILE_NAME, --文件名
       t.AUTOEXTENSIBLE, --是否自动扩展
       t.BYTES / 1024 / 1024 as tsize, --表空间初始大小
       t.MAXBYTES / 1024 / 1024 msize, --表空间最大扩展到多少
       b.CONTENTS, --表空间类型
       b.EXTENT_MANAGEMENT --表空间管理模式
  from dba_temp_files t, dba_tablespaces b
 where t.TABLESPACE_NAME = b.TABLESPACE_NAME
```

# oracle临时表空间组

从[Oracle临时表空间](https://www.oraclejsq.com/oraclegl/010300764.html)章节，我们了解到临时表空间主要存储数据库进行数据操作时的中间临时数据的存储，实际应用当中，经常多用户进行大数据量关联查询时或排序查询时，我们的临时表空间经常会不够用，导致查询中断。

Oracle10g之前，我们只能通过扩充临时表空间，因为，Oracle10g之前每个用户只能指定一个临时表空间。但是在Oracle11g中Oracle数据提供了Oracle临时表空间组的概念。应许把多个临时表空间组成一个组，然后把用户指定到这个临时表空间组，从而达到一个用户可以同时使用多个临时表空间的目的。

------

Oracle临时表空间组中至少得有一个临时表空间，并且同组下的表空间不能有重名。Oracle临时表空间组不用显式的创建，在创建临时表空间的时候为他指定一下临时表空间组即可，实际上创建临时表空间组就是为表空间设定组。创建的方式有如下两种：

1、创建临时表空间时指定临时表空间组：这种创建的方式和创建表空间的语法很相似，语法如下：

```sql
create temporary  tablespace tempname
tempfile 'filename'
size m
tablespace group groupname;
```

下面，我们就在创建临时表空间temp2的同时，为它指定一个临时表空间组tempgroup。代码如下：

```sql
create temporary  tablespace temp2
tempfile 'E:\APP\ADMIN\ORADATA\ORCL\temp2.DBF'
size 50m
tablespace group tempgroup;
```

然后，我们通过数据字典dba_tablespace_groups临时表空间组的创建情况：

```sql
select * from dba_tablespace_groups;
```

2、把原有的临时表空间转移到创建好的临时表空间组当中，下面把临时表空间temp1转移到tempgroup组中，代码如下：

```
alter tablespace temp1
tablespace group tempgroup;
```

然后，我们通过数据字典dba_tablespace_groups查询下临时表空间组的创建情况

创建好临时表空间组，可以把数据库的默认临时表空间设置为表空间组，也可以把对应的用户的临时表空间替换成临时表空间组，从而达到优化临时表空间的目的，代码如下：

```
--修改数据库默认临时表空间
alter database default temporary tablespace tempgroup;
--修改用户默认临时表空间
alter user student temporary tablespace tempgroup;
```

# Oracle用户管理

Oracle用户管理是数据库管理员的必备技能，只有数据库管理员才具有创建、修改、删除用户的权限。对用户的管理主要涉及用户权限的管理，根据每个用户的需求不同，分配给每个用户的权限也不一样，如果数据库管理员对于Oracle用户权限分配的合理，会大大提高数据库的安全性，如果对权限的分配不合理，可能给数据库带来安全隐患。Oracle用户管理的主要内容包括：

1、进行Oracle用户的创建、修改、删除。

2、Oracle权限管理：对用户或角色进行授权和撤销授权。

3、Oracle角色管理：对用户设定特定的角色、创建角色并授予角色特定权限、对特定角色进行修改、删除。

4、介绍Oracle概要文件（PROFILE），对特定用户设置特定的概要文件，配置用户对系统资源的访问限制和口令管理。

# Oracle创建用户

Oracle用户创建是Oracle数据库管理员才具有的权限，同理利用具有DBA权限的用户，如SYS、SYSTEM用户具有创建用户的权限。利用create user关键字命令进行用户的创建，具体的语法如下：

```
create user username
  identified by "password"
  [default tablespace ts_name]
  [temporary tablespace tempname|tempgroupname]
  [quota n size|unlimited on ts_name]
  [profile DEFAULT/profilename]
  [account lock|unlock]
  [password expire];
```

**语法解析：**

**1、username**：指定创建用户的名字。

**2、identified by "password"**：指定用户的密码为password。

**3、default tablespace ts_name：**指定用户存储的表空间为TS_NAME。

**4、temporary tablespace tempname|tempgroupname：**指定用户的临时表空间为tempname，当存在临时表空间组的时候，指定用户的临时表空间组为tempgroupname。

**5、quota n size|unlimited on ts_name：**指定用户使用表空间的最大值为n，unlimited则表示对表空间使用不限制。

**6、profile DEFAULT/profilename：**表示指定用户的概要文件。

**7、account lock|unlock ：**指定用户的锁定状态，lock：锁定状态，unlock：解锁状态。

**8、password expire：**设置当前用户的密码为过期状态，使用户不能登录，要登录必须得重新修改密码。

# Oracle修改删除用户

## Oracle修改用户

Oracle用户创建好后，有时候会随着业务系统的改变，需要对用户的一些信息进行修改，那Oracle数据库是利用关键 字alter user对用户信息进行修改

**Oracle修改用户的语法结构：**

```
alter user username
  identified by "password"
  [default tablespace ts_name]
  [temporary tablespace tempname|tempgroupname]
  [quota n size|unlimited on ts_name]
  [profile DEFAULT/profilename]
  [account lock|unlock]
  [password expire];
```

**语法解析：**

通过上面修改用户的语法会发现和Oracle创建用户语法相似，这里就不一一解释了，可以参考Oracle创建用户的语法。

------

**案例1**、修改用户teacher的密码，代码如下：

```
alter user teacher 
identified by "234567";
```

------

**案例2**、通过修改teacher用户的临时表空间，把临时表空间指向临时表空间组tempgroup，代码如下：

```
alter user teacher 
temporary tablespace tempgroup;
```

然后通过查询数据字典dba_users查看下结果

## oracle删除用户

数据库经常会存在某些废弃的用户，这个时候可以通过drop user关键字对用户进行删除操作。Oracle删除用户时会一并把该用户下的所有对象进行删除，因此再进行删除操作时，应当谨慎考虑。

**Oracle删除用户的语法结构：**

```
drop user username cascade;
```

**语法解析：**

Oracle删除用户的操作很简单，就是通过drop user 关键字和用户名即可删除

# Oracle权限授权管理

Oracle的权限主要有系统权限和对象权限两种。系统权限主要是连接权限（session）、user权限等，系统权限主要是对数据库具有系统级的操作。对象权限，指的是对数据库对象具有特定操作的权限。

Oracle授予权限的对象可以是[用户](https://www.oraclejsq.com/oraclegl/010300768.html)，也可以是角色。授予权限的操作包括授予系统权限或对象权限给用户（或角色）。

我们在进行授予权限时候，要注意，系统权限的授予必须具有DBA权限的用户进行授权。如：sys、system用户。对象权限的授予，可以使用对象所有者的用户进行授权，或者使用具有DBA用户的进行授权

------

## 授予系统权限

**授予系统权限的语法结构：**

```
grant system_privilege|all privileges to user|role
[with admin option]
```

**语法解析：**

1、grant system_privilege|all privileges :指定授予什么权限。system_privilege:授予的权限名称，all privileges：指授予所有系统权限。

2、user|role：授予权限给用户还是角色。

3、[with admin option]：指的是当前被授权的用户具有授权给其它用户系统权限的权利。

------

**例1**、给用户teacher授予系统的create session权限，代码如下：

```
grant create session to TEACHER;
```

通过查询数据字典dba_sys_privs可以查询一下TEACHER的系统权限。

## 授予对象权限

**授予对象权限的语法结构如下：**

```
grant obj_privilege|all 
on obj_name  to user|role
[with grant option]
```

**语法解析：**

1、grant obj_privilege|all：给指定用户授予对象权限，all：指的是授予全部对象权限。

2、on obj_name to user|role：指的是把对象obj_name的权限授予给用户user或者角色role。

3、with grant option：指定是当前被授权的用户具有授权给其它用户该对象权限的权利。这里with grant option对应对象的授权是级联，而对于上面系统权限的with admin option的授权不是级联的，回收该用户的权限，不会把授予给别人的权限收回，而级联授权的就会把对应的权限都收回。

------

**例2**、把student用户下的学生信息表（stuinfo）授权给teacher用户，代码如下：

```
grant select on student.stuinfo to teacher with grant option;
```

然后通过查询数据字典dba_tab_privs查看下teacher用户具有的对象权限

**注意：**给用户授权时，要注意自己是不能为自己授权的，要么利用第三方具有该对象授权权限的用户进行授权，或者利用具有DBA权限的用户授权。

# Oracle撤销权限

Oracle撤销权限的概念就是对[用户](https://www.oraclejsq.com/oraclegl/010300768.html)或角色的权限的回收，也就是对用户删除某个系统权限或者删除某个对象的操作权限。

------

## Oracle撤销系统权限

Oracle撤销系统权限只有具有DBA权限的用户才能进行撤销（revoke）操作。撤销系统权限的语法如下：

```
revoke system_privilege from user|role;
```

语法解析：

1、system_privilege：指的是撤销的系统权限的名称，撤销系统权限时，必须是该用户具有了该系统权限，假如不存在系统权限是不能进行撤销的。

2、user|role：指的是撤销权限的对象是用户还是角色。

------

例1、利用revoke命令撤销用户teacher的create session权限，代码如下：

```
revoke create session from teacher;
```

然后通过数据字典dba_sys_privs可以查询一下TEACHER的系统权限，发现create session已经不存在了。

## Oracle撤销对象权限

Oracle撤销对象权限也是利用revoke命令进行的，语法结构如下：

```
revoke obj_privilege|all 
on object 
from  user|role;
```

语法解析：

1、obj_privilege|all：指的是对应的数据库对象的操作权限，all表示把所有的操作权限都撤销。

------

例2、利用revoke命令撤销用户teacher对学生信息表（stuinfo）的delete（删除）的操作权限，代码如下：

```
revoke delete  on student.stuinfo from  teacher ;
```

然后，通过数据字典dba_tab_privs查询一下stuinfo的操作权限，发现已经删除，只剩下查询权限。

注意：在进行撤销权限时，系统权限的撤销和对象权限的撤销是不一样的，通过上一章Oracle权限授权管理中知道，系统权限的授权不是级联的，对象权限的授权是级联的。在撤销权限时也是一样的，在撤销对象权限时，那么该用户授予给其它用户的对象权限也要跟着收回。而撤销系统权限就不会收回其它用户的权限。

# Oracle角色

Oracle角色其实就是一组权限的集合，比如我们经常用的DBA、connect、resource等角色，都是Oracle系统为我们定义好的一些常用的角色，一般利用这些角色就可以控制好不同需求用户的权限。

Oracle中角色可以授予给多个[用户](https://www.oraclejsq.com/article/010100133.html)，并且一个用户可以设置多个角色。因此，Oracle数据库也提供了自定义角色的功能，方便不同用户的权限授予。例如，可以自定义个一个角色，可以[查询](https://www.oraclejsq.com/article/010100196.html)、[更新](https://www.oraclejsq.com/article/010100202.html)特定几个关联表的权限。那么以后但凡有用户想用到这些表的操作权限时，可以直接给该用户设置定义好的角色。

------

## Oracle创建角色

Oracle创建角色分为两个步骤：创建角色、给角色授权：

**1、创建角色的语法结构：**

```
create role role_name ;
```

**2、授予角色权限的语法结构：**

授予角色权限的语法结构和[授予用户权限](https://www.oraclejsq.com/oraclegl/010300774.html)的语法结构一样，只是授予对象是角色而不是对象而已，具体语法如下：

```
--授予系统权限
grant system_privilege|all privileges to role_name
[with admin option]
--授予对象权限
grant obj_privilege|all 
on obj_name  to role_name
[with grant option]
```

------

**例1**、利用上面创建角色的语法结构，创建一个具有操作学生信息表（stuinfo）和班级表（class）权限的角色student_role，具体代码如下：

```
--创建角色 student_role
create role student_role;

--给student_role授于操作stuinfo、class的操作权限
grant all on student.stuinfo to student_role;
grant all on student.class to student_role;
```

然后，通过数据字典role_tab_privs(角色对应的对象操作权限)、role_sys_privs(角色对应的系统权限)查询对应的权限，可以发现student_role已经具有了表stuinfo和表class的所有操作权限。

## 设置角色

角色创建好后，并不是直接生效，而是要把角色设置给对应的用户后才能生效的。具体语法如下：

```
grant role_name to user;
```

例2、把角色student_role授权给teacher用户，使得teacher用户可以直接操作stuinfo表和class表的权限。代码如下：

```
grant student_role to teacher;
```

然后，可以通过数据字典dba_role_privs查看下，用户teacher设定好的角色。发现已经存在student_role角色

# Oracle概要文件profile

在我们创建用户时，都会为用户指定一个概要文件。概要文件主要是设置用户限制使用数据库资源（主要CPU资源）和Oracle口令参数管理的信息。经常利用Oracle概要文件对数据库用户进行口令管理、指定口令有效期、用户连接时间以及最大空闲等待时间等。

在Oracle系统中，在不给用户指定特定的概要文件时，系统会自动使用默认的概要文件DEFAULT。Oracle系统管理员可以利用不同的概要文件对不同的用户组设置不同的限制权限，可以起到合理分配系统的目的。

------

## Oracle创建概要文件

**Oracle创建概要文件的语法结构：**

```
create profile pro_name
limit
[数据库资源参数|口令参数]
```

**语法解析：**

1、create profile：Oracle创建概要文件的关键字，指定概要文件的名字为pro_name。

2、limit [数据库资源参数|口令参数]：指定要限制的参数信息。

**数据库资源参数（KERNEL），常用的有：**

1、CPU_PER_SESSION：限制会话使用的CPU时间，单位是百分之一秒。

2、SESSIONS_PER_USER：限制用户所允许建立的最大并发会话数。

3、CONNECT_TIME： 限制每个会话能连接到数据库的最长时间， 超过这个时间会话将自动断开，单位是分钟。

4、IDLE_TIME：限制每个会话所允许的最长连续空闲时间，超过这个时间会话将自动断开，单位是分钟。

5、LOGICAL_READS_PER_SESSION：限制每个会话所能读取的数据块数目。

6、PRIVATE_SGA：每个会话分配的私有SGA 区大小（以字节为单位）。

7、CPU_PER_CALL：用于指定每条SQL 语句可占用的最大CPU时间，单位是百分之一秒。

8、LOGICAL_READS_PER_CALL：用于指定每条SQL 语句最多所能读取的数据块数目。

**口令参数(PASSWORD)，常用的有：**

1、FAILED_LOGIN_ATTEMPTS： 该参数指定允许的输入错误口令的次数， 超过该次数后用户帐户被自动锁定。

2、PASSWORD_LOCK_TIME：用于指定指定账户被锁定的天数，单位为天。

3、PASSWORD_LIFE_TIME： 指定口令的有效期，单位为天。 如果在达到有效期前用户还没有更换口令，它的口令将失效。这个时候必须重新设置新密码才能登陆。

4、PASSWORD_GRACE_TIME：用于指定口令失效的宽限期，单位为天

5、PASSWORD_REUSE_TIME： 指定能够重复使用一个口令前必须经过的时间，单位为天

6、PASSWORD_REUSE_MAX： 用于指定在重复使用口令之前必须对口令进行修改的次数。 PASSWORD_REUSE_TIME 和PASSWORD_REUSE_MAX两个参数只能设置一个，另一个必须为UNLIMITED。

------

**案例1、**创建一个概要文件pro_teacher，限制用户登陆时，错误口令允许错误的次数为5次，并且限制会话的最长空闲时间为20分钟，代码如下：

```
create profile pro_teacher
limit
FAILED_LOGIN_ATTEMPTS 5
IDLE_TIME 20;
```

然后，可以通过数据字典DBA_PROFILES对创建好的概要文件进行查询，

数据字典DBA_PROFILES的resource_name字段代表的就是要限制的资源名称，在我们创建概要文件时，只设定了IDLE_TIME和FAILED_LOGIN_ATTEMPTS两个选项，其它限制选项不做限定时，就会默认使用系统的DEFAULT概要文件的配置信息（DEFAULT概要文件是系统创建时自动创建的，对大部分信息是不做限制的，读者可以自行查询数据字典查看配置信息）。

下一步，可以为teacher用户指定创建好的概要文件，代码如下：

```
alter user teacher 
profile  pro_teacher;
```

------

## Oracle修改概要文件

Oracle概要文件经常会随着系统业务需求的变化，有时候需要对其进行修改，主要是通过关键字ALTER PROFILE进行修改。和创建概要文件的语法一致，只需要对需要修改的限制资源进行重新设定即可。

案例2、把pro_teacher概要文件，增加一个口令有效期为365天的限制信息，修改代码如下：

```
alter profile pro_teacher
limit
PASSWORD_LIFE_TIME 365;
```

然后，再查看一下数据字典DBA_PROFILES中pro_teacher的信息，发现信息已经变化。

**总结：**

Oracle概要文件主要是做数据库资源和口令资源限制的配置，数据库管理员经常利用概要文件来批量管理不同用户的访问权限，从而达到合理分配数据库资源的目的。

# oracle体系结构概述

通过前面的学习，大概的知道Oracle常用的管理对象有哪些，比如[日志文件](https://www.oraclejsq.com/oraclegl/010300678.html)、[日志文件组](https://www.oraclejsq.com/oraclegl/010300718.html)、[控制文件](https://www.oraclejsq.com/oraclegl/010300654.html)、[表空间](https://www.oraclejsq.com/oraclegl/010300758.html)、[数据文件](https://www.oraclejsq.com/oraclegl/010300760.html)、[临时表空间](https://www.oraclejsq.com/oraclegl/010300764.html)、[用户](https://www.oraclejsq.com/oraclegl/010300768.html)、[权限](https://www.oraclejsq.com/oraclegl/010300774.html)、[角色](https://www.oraclejsq.com/oraclegl/010300778.html)、[概要文件](https://www.oraclejsq.com/oraclegl/010300798.html)等的管理。但是，对Oracle体系结构还不一定有宏观的了解，这里做一下概要的总结。

------

## 一、总体结构

Oracle体系结构总体可以分为：物理结构、系统全局区（System Global Area）、进程。

其中物理结构包含数据文件、日志文件、控制文件、参数文件等。系统全局区包含共享池、数据缓冲区、日子缓冲区。进程可以分为用户进程、服务进程和后台进程等。其总体结构如下图所示：

![image-20211211163453440](oacle.assets/image-20211211163453440.png)

### 物理结构 

数据文件（data file ）指的是oracle数据库直接的物理存储文件，每一个数据文件只能和一个数据库相关联且一旦建立则不能改变其大小。

日志文件（log file）指的是对数据库数据的修改，以备恢复数据之用。其中特点之处，在于日志文件一半包含两个或两个以上日志文件组，其中日志文件组以循环的方式来进行写的操作。

控制文件（control file）是一个描述数据库的二进制文件，它描述关于数据库的基本信息，如：建立时间，数据库名，数据库中的数据文件、日志文件的文件名以及所在路径信息，还有数据库恢复时候所需的同步信息。

参数文件（parameter file）是一个文本文件，该文件只有在建立数据库或者启动实例的时候才被访问。如：初始参数文件（init.ora）、生成参数文件（initSID.ora,config.ora）。它们的作用是用来确定存储结构的大小、设置数据库的全部缺省值、设置数据库的各种物理属性以及优化数据库性能。

------

### 系统全局区（SGA）

说系统全局区，先说明下实例：它是存取和控制数据库的软件机制，它由SGA和后台进程组成的，区别于数据库，数据库是物理存储文件。这个很多时候会混淆掉，其实他们是不同的概念。

SGA是Oracle系统为实例分配一组共享缓冲存储区、用于存放数据库的数据和控制信息，以实现对数据库的操作和管理。

共享池：由共享SQL区如我们使用SQL时候的语句文本、语法分析形势、执行方案等和数据字典区组成的。数据缓存区：用于存储从数据文件读出来的数据备份。日志缓冲存储区：以记录项的形式备份数据库缓冲区中被修改的缓冲块、而且这些记录都会被写入日志文件当中。

------

### 进程

Oracle实例一般可以分为单进程实例和多进程实例两种。单进程实例指的是一个进程执行全部的Oracle代码。而多进程的实例则是使用多个进程执行Oracle的不同代码，对于每一个连接的用户都有一个进程。其中又可以分为专用服务器的方式，为每个用户单独开设一个服务器进程，一般适用于实时系统。另一种就是多线索服务器方式，它是通过调度器为每个用户分配服务器进程。

用户进程指的是当用户运行一个应用程序时候所建立的进程，而服务器进程则是处理用户进程的请求，它一般分析[SQL命令](https://www.oraclejsq.com/article/010100136.html)并生成执行方案，从数据缓冲区中读取数据，将执行结果返回给用户。

后台进程指的是为所有数据库用户异步的完成各种任务，其中主要的后台进程有我们常见的数据库写进程，日志写进程，归档进程，恢复进程等等。

------

## 二、示例

如我们在sql编译器中输入select 查询的时候Oracle是如何执行的呢。如下我们输入

Select sal from emp where job =’clerk’。如下图：

![image-20211211163514068](oacle.assets/image-20211211163514068.png)

我们在看看UPDATE操作的区别：

![image-20211211163534872](oacle.assets/image-20211211163534872.png)

# Oracle系统基本轮廓和体系结构

一说到Oracle系统的体系结构和底层模块之间的关系，可能部分人会有概念，但是对于他们之间的联系和总体的轮廓可能不清晰，这里我们通过几个关系图来总结和分析Oracle系统的基本轮廓和体系结构。

## Oracle基本轮廓：

![image-20211211163600750](oacle.assets/image-20211211163600750.png)

## Oracle数据库管理系统体系结构：

![image-20211211163621166](oacle.assets/image-20211211163621166.png)

## Oracle数据库体系结构：

![image-20211211163638433](oacle.assets/image-20211211163638433.png)

## Oracle物理和逻辑实体关系图：

![image-20211211163652275](oacle.assets/image-20211211163652275.png)

## 名词解析：

**数据库：**存放数据的仓库，有组织的、可共享的、存储的数据集合，按数据模型组织、描述和存储、较小数据冗余度、较高数据独立性和易扩展性、可为用户共享。

**数据库实例：**存取和控制数据库的软件机制，SGA（System Global Area）和ORACLE进程的组合，内存和进程用于管理数据库的数据，并为数据库用户服务。

**物理结构：**由操作系统文件组成，包括[数据文件](https://www.oraclejsq.com/oraclegl/010300758.html)（一个或多个）、[重做日志文件](https://www.oraclejsq.com/oraclegl/010300678.html)（两个或多个）、[控制文件](https://www.oraclejsq.com/oraclegl/010300654.html)（一个或多个）。

**逻辑结构：**由[表空间](https://www.oraclejsq.com/oraclegl/010300760.html)、段、范围、数据块、模式对象组成，通过逻辑结构控制磁盘空间的使用。

**模式对象：**直接引用数据库数据的逻辑结构，包括[表](https://www.oraclejsq.com/article/010100139.html)、[视图](https://www.oraclejsq.com/article/010100440.html)、[序列](https://www.oraclejsq.com/article/010100438.html)、[存储过程](https://www.oraclejsq.com/plsql/010200560.html)、[触发器](https://www.oraclejsq.com/plsql/010200564.html)、同义词、[索引](https://www.oraclejsq.com/article/010100478.html)、集聚、数据库链、快照等。

# oracle实例是什么

Oracle实例指的是由Oracle内存结构（SGA）和Oracle进程组合在一起的统称。我们常说的Oracle数据库指的是Oracle数据库管理系统，它是由Oracle 数据库（数据存储）和管理数据库的实例组成的。

Oracle实例就是Oracle利用后台内存结构（SGA）和进程来管理oracle数据库，同时提供服务。我们可以做一个通俗的比喻，可以把Oracle数据库管理系统比作仓库，数据库（物理存储、数据文件、日志文件等）就是仓库当中的货物和记账本。而Oracle实例就是负责看管仓库，负责提货开门锁门的部门。

数据库实例，包括有数据库后台进程（PMON：进程监控、SMON：系统监控、DBWR：数据库读写、LGWR：日志读写、CKPT：检查点等）和内存区域SGA（包括shared pool：共享池、db buffer cache：数据库高速缓存、redo log buffer：重做日志缓存区、DC:数据字典高速缓存等）。

Oracle实例是一系列复杂的内存结构和操作系统进程 。在任何时刻一个Oracle实例只能和一个数据库关联，就比如一个门只能进一个数据库仓库一样。但是一个数据库是可以有多个实例的，就如一个仓库可能会有多个门一样。

------

## Oracle实例工作过程

Oracle实例是如何进行管理数据库的呢？大致步骤如下：

1、启动实例（分配SGA和启动Oracle进程）。

2、实例启动后，装配数据库，装配时实例会通过查询和加载控制文件，让实例和数据库进行关联。

3、数据库准备打开。

4、数据库打开（打开后，授权的相关用户就可以读取数据库了）

5、关闭数据库（先卸载数据库，后关闭实例。）

这里，我所说的数据库和实例的开启和关闭只能通过DBA权限的用户才能操作。

# oracle进程

Oracle进程是Oracle实例中主要的组成部分，是Oracle系统管理数据库必不可少的一部分，Oracle进程可以分为用户进程、Oracle进程两大部分。它们主要是维护数据库的稳定，相当于一个企业中的管理者，负责全局统筹的作用。

------

## 用户进程

用户进程指的是用户在运行程序或者Oracle工具时，需要通过建立用户进程和Oracle实例进行通信。

我们常说的connection(连接)就是用户进程和Oracle实例间建立的一个通信通道。Oracle的connection连接是允许一个用户可以同时多次连接到同一个数据库实例的。

还有，常说的session（会话）是用户在和Oracle服务器连接成功后，并通过了Oracle的身份验证后，用户会和Oracle服务器之间建立一个会话。同时同一个用户可以并发的和数据库建立多个会话。

## Oracle进程

Oracle进程又分为服务器进程（server process）和后台进程（background process）

**服务器进程：**用于处理连接到该实例的用户进程的请求。当应用程序和ORACLE运行在同一主机时，用户进程和相应的服务器进程可组合到单个进程，以减少系统开销；当应用程序和ORACLE运行在不同的主机时，用户进程将通过一个单独的服务器进程与ORACLE联系。

服务器进程主要用来分析和执行SQL语句、所需的数据不在SGA内存中时，从磁盘数据文件复制到SGA的共享数据缓冲区等工作。

**后台进程：**是在Oracle实例启动时建立的，用于优化性能和协调多用户连接通信的工作。常用的Oracle后台进程，可以通过数据字典V$bgprocess查询

这里，介绍一下一些常见的后台进程：

1、数据库写入（DBRn）进程：按照最近最少使用（LRU）算法，以批量（多块）方式，将“脏的”缓冲区的内容写入数据文件，保持缓冲区的“清洁”和数量。

2、日志写（LGWR）进程：将重做日志缓冲区中自上次写以来已经拷贝到缓冲区中的所有重做条目写入重做日志文件。重做日志缓冲区是一个循环缓冲区，LGWR正常写的速度很快。

3、检查点（CKPT）进程：发生检查点时，修改所有数据文件的标题和记录该检查点的细节。通常由LGWR完成，但有多个数据文件，而使用LGWR又明显降低系统性能时才使用CKPT。

4、系统监控（SMON）进程：在实例启动时执行实例恢复，整理不再使用的临时段，合并邻近的空闲空间获得更大的空闲可用块。

5、进程监控（PMON）进程：恢复出故障的用户进程，整理缓冲区的高速缓存和释放用户进程使用的资源。定期检查调度进程和服务器进程状态，重新启动非正常终止的进程。

6、存档（ARCH）进程：联机重做日志填满时，将日志内容拷贝到指定的存储设备中。

7、恢复（RECO）进程：在分布式数据库环境中自动解决分布式事务中的故障。

8、锁（LCKn）进程：在并行服务器系统中提供实例间的封锁。

9、作业队列（SNPn）进程：在分布式数据库环境中自动刷新表快照，还执行DBMS_JOB包创建的作业请求。

10、队列监控（QMn）进程：监控消息队列的ORACLE高级（AQ）队列。

11、调度（Dnnn）进程：通过允许用户进程共享限定数量的服务器进程来支持多线程配置。

12、共享服务器（Snnn）进程：在多线程配置模式下，每个服务器进程服务于多个客户请求。

------

# Oracle内存结构

Oracle内存结构是Oracle数据库重要组成部分，是 oracle 数据库重要的信息缓存和共享区域，和[Oracle后台进程](https://www.oraclejsq.com/oraclegl/010301318.html)一起组成[Oracle实例](https://www.oraclejsq.com/oraclegl/010301278.html)。Oracle用户的所有操作都会在内存当中进行一系列操作，然后再交给数据库后台进程，最后把数据持久化到相应的物理文件中进行保存。

Oracle内存是在数据库启动，或者是实例启动时进行分配的。Oracle内存的管理是由Oracle数据库本身进行管理的。我们常见的一些数据库性能有一大部分是由于随着数据库业务的发展，之前配置的Oracle内存不够合理或者太少，从而导致的数据库的性能下降。

Oracle内存结构主要有软件代码区（SCA）、系统全局区（SGA）、程序全局区（PGA）组成。

------

## 软件代码区（SCA）

Oracle软件代码区（SCA）主要是用于保存正在执行的和需要执行的代码，Oracle代码区是所有Oracle用户和实例都可以共享使用的。它的大小一般在数据库安装部署后SCA的内存大小就不会变化，并且SCA代码区只能进行读。

------

## 系统全局区（SGA）

系统全局区是共享内存结构，它是由Oracle系统进行分配的，主要包含Oracle实例和控制相关信息，用于提供查询性能，允许大量并发的数据库活动。当Oracle启动时，会分配制定的SGA。SGA中的数据供所有的服务器进程和后台进程共享，因此SGA 又常称为共享全局区。

SGA主要分为数据库缓冲区高速缓存、重做日志缓冲区、共享池、数据字典高速缓存和其它信息区。

**数据库高速缓存：**指的是从最近从数据文件中检索出来的数据，供所有用户共享使用。类似保存数据文件的一个副本，减少用户直接读取数据和操作数据文件的次数。提高读取和操作的速度。当用户要操作数据库中的数据时，Oracle使用高效的LRU算法，先由服务器进程将数据从磁盘的数据文件中读取到数据高速缓冲区中，然后进行处理。最后用户在缓冲区中处理的结果，最终会由数据库写入进程DBWR写到对应的数据文件中进行持久化。

[**重做日志**](https://www.oraclejsq.com/oraclegl/010300678.html)**缓冲区**：指的是用于记录用户对数据库的修改操作生成的操作记录。通过重做日志的概念我们可以了解到数据修改的记录并不是直接写入到重做日志文件中，而是先被写入到重做日志缓冲区，在特定条件下，才通过LGWR日志写入进程写到重做日志当中。

**共享池**：共享池主要是保存最近执行过的[SQL语句](https://www.oraclejsq.com/article/010100136.html)、[PL/SQL程序](https://www.oraclejsq.com/plsql/010200446.html)或数据字典缓存信息。主要用于提高SQL语句的解析效率，不需要每次执行过的语句或者程序块需要重新解析、编译再执行。

**数据字典高速缓存**：指的是保存数据库的对象信息、数据库结构信息等，当用户访问数据库时，可以快速的从数据字典缓存区中获取到常用的数据字典的信息，比如对象是否存在，用户对应的相关权限信息等等。

SGA是占用Oracle内存最大的一个区域，同时也是最影响Oracle性能的主要因素。我们可以通过数据字典v$sga查询数据库中SGA内存的相关信息，如下：

![image-20211211163833155](oacle.assets/image-20211211163833155.png)

## 程序全局区（PGA）

PGA顾名思义程序全局区，就是Oracle进程生成的同时会为它分配一个相应的内存区，该内存区就叫做PGA。因此PGA是一个私有的内存区，不能进行共享。所有的后台进程都会分配到自己的PGA区。它是随着后台进程的启动而分配，随着进程的终止而释放。

这里就会常用到的几个程序全局区，比如：

1、在进行复杂数据关联排序时，会用到**排序区**来存放排序操作所产生的临时数据。

2、PL/SQL程序块中使用到游标时，会用到[**游标**](https://www.oraclejsq.com/plsql/010200562.html)**信息区**来存放游标中的数据集。

3、用户登录后，会利用**会话信息区**来保存用户对应的会话所具有的权限或相关性能统计信息。

4、**堆栈区**：在PL/SQL程序块中，我们会经常使用到变量或者会话变量，这些会话中的变量信息会保存在堆栈区中。

![image-20211211163849748](oacle.assets/image-20211211163849748.png)

本文来源于：https://www.oraclejsq.com/