关于平台账户权限
==========

> 品台账户描述


### 用户类型

`super admin`：平台总账户，超级账户admin，可以看所有的信息，也就是账号 `admin`

`common admin`：普通admin账户，只可以看到相应的own_id = customer_id 对应的数据

`common user`：普通admin子账户，也就是 普通admin账户 添加的子账户

### 数据库表customer

1.表加上类型字段

1.1表字段`type`: 1代表`super admin`，2代表普通admin账户`common admin`，3代表普通admin子账户`common user`

1.2表字段`parent_id`，此字段只针对`common user`有效，也就是该用户只能
看到`own_id = parent_id`的数据（最多看到的数据）

### 账户登录

1.1根据登录的账户，获取到账户的type，将该信息保存到jwt-token里面，登录后，通过token解码，
获取当前用户的type

1.2获取相应的菜单权限，也就是能看到的菜单，渲染前端vue

### 请求数据

1.vue点击菜单访问页面，进行后端数据请求

2.go部分，先判断是否有url请求权限，如果没有权限，直接返回权限拒绝信息

3.如果用户有url请求权限，则进行数据请求，然后查看用户的数据获取权限

3.1如果用户是`super admin`，则没有限制

3.2如果用户是`common admin`，则需要在数据获取部分的where部分加上 own_id = customer_id 

3.3如果用户是`common user`，如果该功能不需要数据做区分，也就是各个子账户之间看所有的数据，则
在数据获取部分的where部分加上 own_id = customer_id 

3.4如果用户是`common user`，如果该功能需要数据做区分，也就是子账户之间不能看相互的数据，则
在where部分加上 `own_id = customer_id and created_customer_id = common_user_id` 即可

4返回数据。

### 实现

1.数据库表添加字段，增删改查修改

2.后端gin对每一个url做好权限控制，那些url允许上面的那些角色访问，需要角色访问受限的，就加上。
譬如，只允许超级账户访问就加上`handle.PermissionSuperAdmin`
, 通用的权限限制就用`handle.PermissionAdmin`

3.除了Super Admin独有访问的部分，就是`common admin`访问的菜单了，
`common admin`对于开的每一个子账户，都可以设置相应的访问权限。

3.1`common admin`可以创建用户组，为每一个用户组勾选权限，保存

3.2为开启的子用户添加用户组，可以勾选多个，最终的权限为各个权限组的合集

3.3在`handle.PermissionAdmin`的执行过程中，如果是`common admin`，则直接通过，
如果是`common user`，则查询该用户的是否有权限,先通过查询得到用户的role_array

```
role_id in (role_array)  and url_key = current_url_key and request_method = 'POST'
```

如果没有权限，则返回，如果有权限，则可以访问

4.新建表

4.1权限组表:role

id, role_name, own_id,

4.2用户和权限对应表：customer_role

id, own_id, customer_child_id, role_id

4.3权限和资源对应表  role_resource

id, own_id, role_id, resource_id

4.4资源表

resource_id, name, url_key, request_method，group(资源分组)

5.编辑

5.1【`common admin`】编辑权限组，对权限组表，进行增删改查

5.2【`common admin`】在子账户页面，编辑用户和权限对应表，在列表中查看相应的子账户，点击编辑权限，为用户勾选权限组，
可以多选

5.3【`super admin`】资源表编辑，编辑所有的`common admin`可访问资源。

5.4权限和资源对应，在权限组列表中，点击添加资源，为权限组添加资源。

6.子账户验证权限

6.1子账户获取own_id和customer_child_id

6.2customer_role中通过own_id, customer_child_id，得到自己的role_id数组

6.3通过role_id数组和own_id，获取到所有的resource_id数组，然后将数组进行处理，得到唯一resource_id数组

6.4资源表，通过resource_id数组，得到当前用户所有的允许资源

6.5将这些资源保存到缓存中存储。

6.6从数组中取出，进行循环匹配，匹配成功则有权限，匹配失败则无权限

7.功能内数据过滤权限

个性化配置

7.1某些可以进行数据，按照customer_child_id进行过滤的设置。

id, own_id, 功能名称 func_name, 是否需要过滤 filter_status (1代表需要按照customer_child_id过滤，2代表不需要按照customer_child_id过滤)

7.2在where中进行封装，通过传值决定，是否需要按照customer_child_id过滤数据

8.对于菜单，先让所有用户可以看到所有的菜单，如果权限不允许，则点击菜单无效，返回
权限不足即可。也就是先做后端的权限验证，前端的先不管。






























