Fecshop和Trace系统对接
=====================


> fecshop和Trace系统的对接配置的介绍


1.申请

1.1申请trace账户

1.2联系管理员，提交下面的资料，后面的值示例数据

`网站名称`：fancyecommerce

`网站域名`：fecshop.appfront.fancyecommerce.com

`网站图片地址`：fecshop.appfront.fancyecommerce.com/catalog/product/image

1.3提交后，管理员会给予用户名和密码，您登陆后，
点击菜单：基础信息-->网站管理，就可以看到新增的追踪网站。


2.打开fecshop的@common/config/fecshop_local_services/Page

```
'trace' => [
    'class' => 'fecshop\services\page\Trace',
    // 关闭和打开Trace功能，默认关闭，打开前，请先联系申请下面的信息，QQ：2358269014
    'traceJsEnable' => true,
    // trace系统的 站点唯一标示  website id
    'website_id'    => '9b17f5b4-b96f-46fd-abe6-a579837ccdd9',
    // trace系统的Token，当fecshop给trace通过curl发送数据的时候，需要使用该token进行安全认证。
    'access_token'  => 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ3ZWJzaXRlX3VpZCI6IjliMTdmNWI0LWI5NmYtNDZmZC1hYmU2LWE1Nzk4MzdjY2RkOSJ9.-HsUq-qKcn2dhvGoxSYHVqMxNTH0cBcLsUl-R_utaCo',
    // 当fecshop给trace通过curl发送数据，最大的超时时间，该时间是为了防止
    'api_time_out' => 1.5, // 秒
    // 追踪js url，这个是在统计系统，由管理员提供
    'trace_url'     => 'trace.fecshop.com/fec_trace.js',
    // 管理员提供，用于发送登录注册邮件，下单信息等。
    'trace_api_url' => 'http://120.24.37.249:3000/fec/trace/api',
],
```

登录trace系统，切换成中文，点击菜单：基础信息-->网站管理，然后点击编辑，弹出编辑框，
将编辑框的内容一一填写。

`traceJsEnable`: 设置为`true`

`website_id`：`站点唯一标示`

`access_token`：`验证 Token`

`api_time_out`：服务端通过api给trace系统
发送数据，使用的是curl，该值是设置curl的最大超时时间，默认为`1.5`秒

`trace_url`：`追踪Js Url`

`trace_api_url`：`追踪Api Url`

填写完成后，保存即可完成

> fecshop已经默认将埋点写入系统中，只需要配置即可使用，如果
您想关闭trace功能，只需要将配置中的
`traceJsEnable`: 设置为`false`即可。















