Trace系统添加新平台
===================

> 如果添加新的saas trace平台，需要注意下面的事情


1.接收js和api地址会改变，那么对应的需要修改

1.1、添加一个新的追踪js地址

js里面的内容也需要需要修改，将trace.js尾部的

```
img.src = '//120.24.37.249:3000/fec/trace?' + args;
```

改成当前的接收地址


1.2、handler/common/website.go文件里面修改：

```
var FecTraceJsUrl string = "trace.fecshop.com/fec_trace.js"
var FecTraceApiUrl string = "120.24.37.249:3000/fec/trace/api"
```

`FecTraceJsUrl`: 上面新建的js地址

`FecTraceApiUrl`: 相应的api地址，用于接收服务端传送的数据。

1.3、搭建go环境，安装mongodb，mysql

1.4、让新用户在这里注册，对接。




















