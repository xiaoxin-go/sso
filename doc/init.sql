# 创建日志表
create table t_log
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    operator   varchar(50) comment '操作人',
    content    varchar(5000) comment '操作内容'
);

# 创建用户表
create table t_user
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    username   varchar(50)  not null unique comment '用户名',
    email      varchar(50)  not null unique comment '用户邮箱地址',
    name_cn    varchar(255) not null comment '用户中文名',
    otp_secret varchar(64)  not null comment 'otp_secret',
    enabled    boolean  default 1 comment '用户是否启用',
    password_updated_at datetime comment '密码更新时间'
) comment '用户表';

# 创建菜单表
create table t_menu
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    name       varchar(50)  not null comment '菜单名称',
    name_en    varchar(50)  not null comment '菜单英文名称',
    path       varchar(255) not null comment '菜单路径',
    icon       varchar(255) comment '菜单图标',
    parent_id  int comment '父菜单id',
    sort       int comment '菜单排序',
    enabled    boolean  default 1 comment '菜单是否启用'
);
# 创建api表

create table t_api
(
    id          int primary key auto_increment,
    created_at  datetime default current_timestamp,
    updated_at  datetime default current_timestamp,
    created_by   varchar(64) comment '创建人',
    updated_by   varchar(64) comment '更新人',
    name        varchar(50)  not null comment 'api名称',
    uri         varchar(255) unique not null comment 'api路径',
    method      varchar(10)  not null comment 'api请求方式',
    description varchar(255) comment 'api描述',
    enabled     boolean  default 1 comment 'api是否启用'
);

# 创建角色表
create table t_role
(
    id          int primary key auto_increment,
    created_at  datetime default current_timestamp,
    updated_at  datetime default current_timestamp,
    created_by   varchar(64) comment '创建人',
    updated_by   varchar(64) comment '更新人',
    name        varchar(50) unique not null comment '角色名称',
    description varchar(255) comment '角色描述'
) comment '角色表';

# 创建用户角色关系表
create table t_user_role
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    user_id    int not null comment '用户id',
    role_id    int not null comment '角色id'
) comment '用户与角色关联表';

# 创建角色菜单关系表
create table t_role_menu
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    role_id    int not null comment '角色id',
    menu_id    int not null comment '菜单id'
) comment '角色与菜单关联表';

# 创建菜单和api关系表
create table t_menu_api
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    menu_id    int not null comment '菜单id',
    api_id     int not null comment 'apiid'
) comment '菜单与api关联表';

# 创建角色api关系表
create table t_role_api
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    role_id    int not null comment '角色id',
    api_id     int not null comment 'apiid'
) comment '角色与api关联表';

# 创建平台信息
create table t_platform
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    name varchar(64) unique comment '平台名',
    name_cn varchar(64) comment '平台中文名',
    description varchar(255) comment '描述',
    url varchar(64) comment '平台地址',
    index_url varchar(64) comment '平台首页',
    type int comment '平台类型',
    login_func varchar(20) comment '登录函数',
    enabled boolean comment '是否启用'
) comment '平台表';

# 租户表
create table t_tenement
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    name varchar(64) unique comment '租户名',
    description varchar(255) comment '描述'
) comment '租户表';

# 租户关联平台
create table t_tenement_platform
(
    id         int primary key auto_increment,
    tenement_id int comment '租户ID',
    platform_id int comment '平台ID',
    created_at datetime default current_timestamp,
    created_by varchar(64) comment '创建人'
) comment '租户关联平台';

# 租户关联用户
create table t_tenement_user
(
    id         int primary key auto_increment,
    tenement_id int comment '租户ID',
    user_id int comment '用户ID',
    created_at datetime default current_timestamp,
    created_by varchar(64) comment '创建人'
) comment '租户关联用户';

# 平台账号管理
create table t_platform_user
(
    id         int primary key auto_increment,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    created_by  varchar(64) comment '创建人',
    updated_by  varchar(64) comment '更新人',
    platform_id int comment '平台ID',
    username varchar(64) comment '用户名',
    password varchar(64) comment '密码',
    is_default boolean comment '是否是默认账号'
) comment '平台用户表';

insert into t_api (name, uri, method, enabled)
values
('获取角色列表', '/system/roles', 'GET', 1),
('添加角色', '/system/role', 'POST', 1),
('更新角色信息', '/system/role', 'PUT', 1),
('删除角色', '/system/role', 'DELETE', 1),
('获取角色权限', '/system/role/permission', 'GET', 1),
('更新角色权限', '/system/role/permission', 'PUT', 1),

('获取菜单列表', '/system/menus', 'GET', 1),
('添加菜单', '/system/menu', 'POST', 1),
('更新菜单信息', '/system/menu', 'PUT', 1),
('删除菜单', '/system/menu', 'DELETE', 1),

('获取接口列表', '/system/apis', 'GET', 1),
('添加接口', '/system/api', 'POST', 1),
('更新接口信息', '/system/api', 'PUT', 1),
('删除接口', '/system/api', 'DELETE', 1),

('获取租户列表', '/tenements', 'GET', 1),
('添加租户', '/tenement', 'POST', 1),
('更新租户信息', '/tenement', 'PUT', 1),
('删除租户', '/tenement', 'DELETE', 1),

('获取平台列表', '/platforms', 'GET', 1),
('添加平台', '/platform', 'POST', 1),
('更新平台信息', '/platform', 'PUT', 1),
('删除平台', '/platform', 'DELETE', 1),

('获取平台用户列表', '/platform_users', 'GET', 1),
('添加平台用户', '/platform_user', 'POST', 1),
('更新平台用户信息', '/platform_user', 'PUT', 1),
('删除平台用户', '/platform_user', 'DELETE', 1),

('获取日志列表', '/system/logs', 'GET', 1);

