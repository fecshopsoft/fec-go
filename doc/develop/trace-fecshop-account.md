trace 后台账户介绍
================




trace安装好后，您可以通过  `admin  admin123`
进行登录


对于trace系统，有3中账户级别

> 因为原来想做成saas的方式，因此是三层用户，后来想改成这种方式，三层用户模式也保留下来了

`super admin`: 超级admin账户，用来管理资源，创建网站，将网站的的操作权限指定到某个common admin

`common admin`：普通admin账户，通过`super admin`创建，这个用户用来管理商城的数据，相当于
admin用户

	
`Child Account`：普通员工账户，通过 common admin 创建，给予各个具体 员工。


### 添加 common admin账户

因为网站是和common admin 绑定的，因此，您需要先创建一个 common admin账户

1.您登陆您的admin超级账户

2.点击菜单 控制面板  --> 账户列表 ， 

![xx](images/a0.png)

点击添加，新建一个 common admin账户，
账户类型选择 Common Admin，如图 

![xx](images/a1.png)

保存即可，这样就添加了一个common admin账户，


### 添加 Child Account 账户

登录上面创建的common admin账户


![xx](images/b3.png)


总之：


`super admin`你只用他来创建管理资源，创建common admin账户即可，
不要用他进行其他的操作


common admin：相当于网站的admin账户，通过该账户，可以
添加`Child Account`，设置权限，创建一些编辑信息等等。













