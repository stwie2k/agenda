# CLI 命令行实用程序开发实战 - Agenda

此次作业使用Cobra实现了终端的Agenda系统

## 主要功能
- 注册账户
- 登录账户
- 登出账户
- 删除账户
- 查询账户
- 创建会议
- 查询会议
- 取消会议
- 退出会议
- 清空所有会议
- 增加会议参与者
- 删除会议参与者

##  构架设计

主要可以分为三层构架
#### 表示层 cmd
- 负责与用户的交互操作和指令提示。
-  接收用户输入的指令与参数，并将数据传输给业务逻辑层。
 - 负责输出语句执行情况。
#### 业务逻辑层 service
- 业务逻辑的执行，调取实体层提供的相关API进行操作。
- 判断表示层传输进来数据、命令的合法性。
#### 实体层 entity
 -  提供相关数据操作接口。
 -  文件读取与存储。

## 命令与参数设计
### register

> 用户注册

参数列表
username(-u --user)
password(-p --password)
email(-e --email)
phonenumber(-n --phonenumber)

### login

> 用户登录

参数列表

username(-u --user)
password(-p --password)

###  logout
> 用户登出

参数列表：空


### queryuser
> 查询用户

参数列表
username (-u --user)

### deleteuser
> 删除用户

参数列表
username (-u --user)


### createmeeting
> 创建会议

参数列表

会议标题(-t --title)
开始时间(-s --starttime)
结束时间(-e --endtime)
首个参与者(-p --participator)

###  querymeeting
> 查询某个时段的会议

参数列表

开始时间(-s --starttime)
结束时间(-e --endtime)

###  deletemeeting
> 取消某个会议

参数列表

会议标题(-t --title)

### quitmeeting
> 退出某个会议

参数列表

会议标题(-t --title)


### addperson
> 添加一个与会者

参数列表

会议标题(-t --title)
参与者(-p --participator)

### deleteperson
> 删除一个与会者

会议标题(-t --title)
参与者(-p --participator)
