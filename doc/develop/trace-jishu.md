FA技术选型
===========


> FA系统使用的技术

### 技术结构

采用前后端彻底分离的方式，前端使用的是vue，后端使用的是golang语言

### 后端部分

`golang语言`：通过强劲的go做`数据接收`和`数据脚本计算`部分

`mongodb`：做`数据存储`和`数据计算`部分

`mysql`: `基础数据存储`

`elasticSearch`：计算的结果数据，最终放到`elasticsearch`中，
支持用户进行数据查询

### 前端部分

`vue element`：使用element ui作为前端部分