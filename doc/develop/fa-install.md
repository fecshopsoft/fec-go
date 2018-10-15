FA安装
======

### golang部分安装

1.库包安装

1.1安装xorm

```
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
```

2.安装gin

```
go get github.com/gin-gonic/gin
```

参考资料： [centos6 安装go框架gin的步骤，以及中间遇到的坑](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85go%e6%a1%86%e6%9e%b6gin%e7%9a%84%e6%ad%a5%e9%aa%a4%ef%bc%8c%e4%bb%a5%e5%8f%8a%e4%b8%ad%e9%97%b4%e9%81%87%e5%88%b0%e7%9a%84%e5%9d%91/)

上面如果都安装成功了, 下面安装`fec-go`

```
go get github.com/fecshopsoft/fec-go
```

下载完成后，即可下载玩所有的文件

该部分代码在 ``./github.com/fecshopsoft/fec-go`

### 安装 `Mysql`, `ElasticSearch`, `Mongodb`

1.安装`mysql`

1.1安装装文档：[mysql5.6安装](http://www.fancyecommerce.com/2016/04/29/linux-%E5%AE%89%E8%A3%85mysql5-6/)

1.2创建数据库`fa-go`

1.3导入数据库,mysql数据库文件如下：

```
example/my.sql
```


2.安装`elasticSearch6`

注意，这里是安装es6，可以参看文档[安装ElasticSearch](http://www.fancyecommerce.com/2016/11/09/%E5%AE%89%E8%A3%85elasticsearch-%EF%BC%8C%E4%BB%A5%E5%8F%8A%E5%9C%A8yii2%E4%B8%AD%E7%9A%84%E4%BD%BF%E7%94%A8/)

3.安装`mongodb`

> 下面的安装文档安装完第3步就可以了，后面是安装php mongodb扩展的部分，不需要安装

安装参考文档[安装mongodb](http://www.fancyecommerce.com/2016/05/03/yii2-mongodb%E7%9A%84%E5%AE%89%E8%A3%85%E5%92%8C%E9%85%8D%E7%BD%AE-mongo/)


### 配置

1.创建文件以及文件夹

```
mkdir -p /www/fec-go/etc
touch /www/fec-go/etc/config.ini

mkdir -p /www/fec-go/log
touch /www/fec-go/log/router_info.log
touch /www/fec-go/log/router_error.log
touch /www/fec-go/log/global.log
chmod 777 /www/fec-go/log/router_info.log /www/fec-go/log/router_error.log /www/fec-go/log/global.log

mkdir -p /www/fec-go/shell_log
touch /www/fec-go/shell_log/router_info.log
touch /www/fec-go/shell_log/router_error.log
touch /www/fec-go/shell_log/global.log
chmod 777 -R /www/fec-go/shell_log/router_info.log /www/fec-go/shell_log/router_error.log /www/fec-go/shell_log/global.log

mkdir -p /www/fec-go/xlsx
chmod 777 /www/fec-go/xlsx
```

2.将 config/config.ini 的内容，复制到 `/www/fec-go/etc/config.ini`，

按照里面的配置说明，配置 `mysql` `mongodb` `elasticsearch`, 以及监听的ip端口

**此部分非常重要，设置数据库连接参数和设置监听的ip端口，清务必填写正确**

3.创建main.go文件

3.1在src/main 下面创建   fec-go.go  和 fec-go-verify.go 文件

3.2将目录（github.com/fecshopsoft/fec-go）下的 `fec-go.go`  和 `fec-go-shell.go`,
复制到 `main/fec-go.go` 和 `main/fec-go-shell.go`。

4.运行

4.1运行web监听脚本（该脚本需要一直运行）

> 该脚本用来接收数据，以及为vue部分提供api，也就是通过web url访问的部分，由这里提供。

进入到main下面，运行

```
go run fec-go.go
```

如果启动成功, 最后会显示：

```
2018/10/15 16:43:07 /root/go/src/github.com/gin-gonic/gin/debug.go:45: [GIN-debug] Listening and serving HTTP on 0.0.0.0:3000
```

`ctrl + c` 会停止当前运行的脚本。

4.2运行数据处理脚本

> 该脚本是数据处理脚本，用来处理数据，也就是将fecshop接收的初始数据进行各个维度的数据聚合，通过
mongodb的mapreduce进行数据分析，得到统计后的数据，以及将数据同步到 ElasticSearch 中。

```
go run fec-go-shell.go
```


```
go run fec-go-shell.go 1 removeEsAllIndex

```

第一个参数： 选填，代表处理N天内的数据统计，默认为`1`

第二个参数： `removeEsAllIndex`,如果删除elasticSearch里面的数据，加上这个参数值，
一般情况下不加这个参数。


4.3cron

> 对于数据统计脚本，一般一天统计一次，因此通过cron可以更好的运行。

```
1 1 * * * /usr/bin/wget http://120.24.37.249:3000/fec/trace/cronssss > /dev/null
01 * * * *  /root/go/src/main/fec-go-shell   >> /www/web_logs/fec-go-shell.log  2>&1
```

第一个是更新数据，将ip和端口换成你自己的即可，一分钟运行一次，这个脚本相当于刷新缓存的性质，
定时的刷新golang中的全局变量（因为有一部分的变量是从数据库里面初始化的，当数据库的内容修改后，
通过这个脚本进行刷新）。

第二个是周期跑数据的脚本，一天跑一次，将`/root/go/src/main/fec-go-shell` 替换成您自己的
路径，后面是log输出文件，自行创建并设置成可写。

