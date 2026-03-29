-- ─────────────────────────────────────────────────────────────
-- GBaseAdmin 数据库初始化脚本
-- 数据库：gbaseadmin
-- 说明：此文件在 MySQL 容器首次启动时自动执行
-- ─────────────────────────────────────────────────────────────

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ─────────────────────────────────────────────
-- 2.1 部门表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `dept` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT '部门ID（Snowflake）',
    `parent_id`  BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '上级部门ID，0 表示顶级部门',
    `title`      VARCHAR(50)     NOT NULL               COMMENT '部门名称',
    `username`   VARCHAR(50)                            COMMENT '部门负责人姓名',
    `email`      VARCHAR(100)                           COMMENT '负责人邮箱',
    `sort`       INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at` DATETIME                               COMMENT '创建时间',
    `updated_at` DATETIME                               COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '软删除时间，非 NULL 表示已删除',
    PRIMARY KEY (`id`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门表';

-- ─────────────────────────────────────────────
-- 2.2 角色表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `role` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT '角色ID（Snowflake）',
    `parent_id`  BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '上级角色ID，0 表示顶级角色',
    `title`      VARCHAR(50)     NOT NULL               COMMENT '角色名称',
    `data_scope` TINYINT         NOT NULL DEFAULT 1     COMMENT '数据范围:1=全部,2=本部门及以下,3=本部门,4=仅本人,5=自定义',
    `sort`       INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at` DATETIME                               COMMENT '创建时间',
    `updated_at` DATETIME                               COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '软删除时间，非 NULL 表示已删除',
    PRIMARY KEY (`id`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- ─────────────────────────────────────────────
-- 2.3 角色-部门关联表（自定义数据权限）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `role_dept` (
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    `dept_id` BIGINT UNSIGNED NOT NULL COMMENT '部门ID',
    PRIMARY KEY (`role_id`, `dept_id`),
    KEY `idx_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色自定义数据权限部门关联表';

-- ─────────────────────────────────────────────
-- 2.4 角色-菜单关联表（资源权限）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `role_menu` (
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    `menu_id` BIGINT UNSIGNED NOT NULL COMMENT '菜单ID',
    PRIMARY KEY (`role_id`, `menu_id`),
    KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单权限关联表';

-- ─────────────────────────────────────────────
-- 2.5 菜单表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `menu` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT '菜单ID（Snowflake）',
    `parent_id`  BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '上级菜单ID，0 表示顶级菜单',
    `title`      VARCHAR(50)     NOT NULL               COMMENT '菜单名称',
    `type`       TINYINT         NOT NULL DEFAULT 1     COMMENT '类型:1=目录,2=菜单,3=按钮,4=外链,5=内链',
    `path`       VARCHAR(200)                           COMMENT '前端路由路径',
    `component`  VARCHAR(200)                           COMMENT '前端组件路径',
    `permission` VARCHAR(100)                           COMMENT '权限标识（如 system:dept:list）',
    `icon`       VARCHAR(100)                           COMMENT '菜单图标（图标名称）',
    `sort`       INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `is_show`    TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '是否显示:0=隐藏,1=显示',
    `is_cache`   TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否缓存:0=不缓存,1=缓存',
    `link_url`   VARCHAR(500)                           COMMENT '外链/内链地址（type=4或5时有效）',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at` DATETIME                               COMMENT '创建时间',
    `updated_at` DATETIME                               COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '软删除时间，非 NULL 表示已删除',
    PRIMARY KEY (`id`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';

-- ─────────────────────────────────────────────
-- 2.6 用户表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `users` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT '用户ID（Snowflake）',
    `username`   VARCHAR(50)     NOT NULL               COMMENT '登录用户名',
    `password`   VARCHAR(255)    NOT NULL               COMMENT '密码（bcrypt 加密）',
    `nickname`   VARCHAR(50)                            COMMENT '昵称/显示名',
    `email`      VARCHAR(100)                           COMMENT '邮箱地址',
    `avatar`     VARCHAR(500)                           COMMENT '头像图片 URL',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at` DATETIME                               COMMENT '创建时间',
    `updated_at` DATETIME                               COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '软删除时间，非 NULL 表示已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_username` (`username`),
    KEY `idx_dept_id` (`dept_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ─────────────────────────────────────────────
-- 2.7 用户-部门关联表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `user_dept` (
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `dept_id` BIGINT UNSIGNED NOT NULL COMMENT '部门ID',
    PRIMARY KEY (`user_id`, `dept_id`),
    KEY `idx_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户部门关联表';

-- ─────────────────────────────────────────────
-- 2.8 用户-角色关联表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `user_role` (
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    PRIMARY KEY (`user_id`, `role_id`),
    KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';


-- ═════════════════════════════════════════════
-- 初始化数据
-- ═════════════════════════════════════════════

-- ─────────────────────────────────────────────
-- 根部门（总公司）
-- ID 使用固定 Snowflake 值，确保跨环境一致
-- ─────────────────────────────────────────────
INSERT INTO `dept` (`id`, `parent_id`, `title`, `username`, `email`, `sort`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (1000000000000000001, 0, '总公司', 'admin', 'admin@example.com', 0, 1, 0, 0, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 超级管理员角色（data_scope=1 全部数据权限）
-- ─────────────────────────────────────────────
INSERT INTO `role` (`id`, `parent_id`, `title`, `data_scope`, `sort`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (1000000000000000002, 0, '超级管理员', 1, 0, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 超级管理员用户
-- 用户名：admin
-- 密码：admin123  →  bcrypt 加密后的哈希值
-- bcrypt(admin123, cost=10) = $2a$10$xVpo7FECYLBDq9NlExvX4.dkNa7JJ/3k/p14v/t4ZTQ4XvNWXrNuS
-- ─────────────────────────────────────────────
INSERT INTO `users` (`id`, `username`, `password`, `nickname`, `email`, `avatar`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (
    1000000000000000003,
    'admin',
    '$2a$10$xVpo7FECYLBDq9NlExvX4.dkNa7JJ/3k/p14v/t4ZTQ4XvNWXrNuS',
    '超级管理员',
    'admin@example.com',
    '',
    1,
    0,
    1000000000000000001,
    NOW(),
    NOW(),
    NULL
);

-- ─────────────────────────────────────────────
-- 用户-部门关联（admin 属于总公司）
-- ─────────────────────────────────────────────
INSERT INTO `user_dept` (`user_id`, `dept_id`)
VALUES (1000000000000000003, 1000000000000000001);

-- ─────────────────────────────────────────────
-- 用户-角色关联（admin 拥有超级管理员角色）
-- ─────────────────────────────────────────────
INSERT INTO `user_role` (`user_id`, `role_id`)
VALUES (1000000000000000003, 1000000000000000002);

-- ─────────────────────────────────────────────
-- 基础菜单树
-- 结构：
--   系统管理（目录, id=1000000000000000010）
--     ├── 部门管理（菜单, id=1000000000000000011）
--     ├── 角色管理（菜单, id=1000000000000000012）
--     ├── 菜单管理（菜单, id=1000000000000000013）
--     └── 用户管理（菜单, id=1000000000000000014）
-- ─────────────────────────────────────────────
INSERT INTO `menu` (`id`, `parent_id`, `title`, `type`, `path`, `component`, `permission`, `icon`, `sort`, `is_show`, `is_cache`, `link_url`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
-- 系统管理（目录）
(1000000000000000010, 0,                    '系统管理', 1, '/system',       NULL,                          '',                     'SettingOutlined', 100, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 部门管理（菜单）
(1000000000000000011, 1000000000000000010,  '部门管理', 2, '/system/dept',  'system/dept/index',           'system:dept:list',     'ApartmentOutlined', 1, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 角色管理（菜单）
(1000000000000000012, 1000000000000000010,  '角色管理', 2, '/system/role',  'system/role/index',           'system:role:list',     'TeamOutlined',    2, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 菜单管理（菜单）
(1000000000000000013, 1000000000000000010,  '菜单管理', 2, '/system/menu',  'system/menu/index',           'system:menu:list',     'MenuOutlined',    3, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 用户管理（菜单）
(1000000000000000014, 1000000000000000010,  '用户管理', 2, '/system/user',  'system/user/index',           'system:user:list',     'UserOutlined',    4, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 按钮级权限菜单（归属于各功能菜单下）
-- ─────────────────────────────────────────────
INSERT INTO `menu` (`id`, `parent_id`, `title`, `type`, `path`, `component`, `permission`, `icon`, `sort`, `is_show`, `is_cache`, `link_url`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
-- 部门管理按钮
(1000000000000000021, 1000000000000000011, '部门新增', 3, NULL, NULL, 'system:dept:create', '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000022, 1000000000000000011, '部门修改', 3, NULL, NULL, 'system:dept:update', '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000023, 1000000000000000011, '部门删除', 3, NULL, NULL, 'system:dept:delete', '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 角色管理按钮
(1000000000000000031, 1000000000000000012, '角色新增', 3, NULL, NULL, 'system:role:create',     '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000032, 1000000000000000012, '角色修改', 3, NULL, NULL, 'system:role:update',     '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000033, 1000000000000000012, '角色删除', 3, NULL, NULL, 'system:role:delete',     '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000034, 1000000000000000012, '资源授权', 3, NULL, NULL, 'system:role:grant:menu', '', 4, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000035, 1000000000000000012, '数据授权', 3, NULL, NULL, 'system:role:grant:dept', '', 5, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 菜单管理按钮
(1000000000000000041, 1000000000000000013, '菜单新增', 3, NULL, NULL, 'system:menu:create', '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000042, 1000000000000000013, '菜单修改', 3, NULL, NULL, 'system:menu:update', '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000043, 1000000000000000013, '菜单删除', 3, NULL, NULL, 'system:menu:delete', '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 用户管理按钮
(1000000000000000051, 1000000000000000014, '用户新增', 3, NULL, NULL, 'system:user:create', '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000052, 1000000000000000014, '用户修改', 3, NULL, NULL, 'system:user:update', '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000053, 1000000000000000014, '用户删除', 3, NULL, NULL, 'system:user:delete', '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 角色-菜单关联（超级管理员拥有所有菜单权限）
-- ─────────────────────────────────────────────
INSERT INTO `role_menu` (`role_id`, `menu_id`)
VALUES
-- 功能目录和菜单
(1000000000000000002, 1000000000000000010),
(1000000000000000002, 1000000000000000011),
(1000000000000000002, 1000000000000000012),
(1000000000000000002, 1000000000000000013),
(1000000000000000002, 1000000000000000014),
-- 部门管理按钮
(1000000000000000002, 1000000000000000021),
(1000000000000000002, 1000000000000000022),
(1000000000000000002, 1000000000000000023),
-- 角色管理按钮
(1000000000000000002, 1000000000000000031),
(1000000000000000002, 1000000000000000032),
(1000000000000000002, 1000000000000000033),
(1000000000000000002, 1000000000000000034),
(1000000000000000002, 1000000000000000035),
-- 菜单管理按钮
(1000000000000000002, 1000000000000000041),
(1000000000000000002, 1000000000000000042),
(1000000000000000002, 1000000000000000043),
-- 用户管理按钮
(1000000000000000002, 1000000000000000051),
(1000000000000000002, 1000000000000000052),
(1000000000000000002, 1000000000000000053);

SET FOREIGN_KEY_CHECKS = 1;
