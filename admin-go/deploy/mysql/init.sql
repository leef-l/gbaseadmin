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
CREATE TABLE IF NOT EXISTS `system_dept` (
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
CREATE TABLE IF NOT EXISTS `system_role` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT '角色ID（Snowflake）',
    `parent_id`  BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '上级角色ID，0 表示顶级角色',
    `title`      VARCHAR(50)     NOT NULL               COMMENT '角色名称',
    `data_scope` TINYINT         NOT NULL DEFAULT 1     COMMENT '数据范围:1=全部,2=本部门及以下,3=本部门,4=仅本人,5=自定义',
    `sort`       INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `is_admin`   TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否超级管理员:0=否,1=是',
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
CREATE TABLE IF NOT EXISTS `system_role_dept` (
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    `dept_id` BIGINT UNSIGNED NOT NULL COMMENT '部门ID',
    PRIMARY KEY (`role_id`, `dept_id`),
    KEY `idx_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色自定义数据权限部门关联表';

-- ─────────────────────────────────────────────
-- 2.4 角色-菜单关联表（资源权限）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `system_role_menu` (
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    `menu_id` BIGINT UNSIGNED NOT NULL COMMENT '菜单ID',
    PRIMARY KEY (`role_id`, `menu_id`),
    KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单权限关联表';

-- ─────────────────────────────────────────────
-- 2.5 菜单表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `system_menu` (
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
CREATE TABLE IF NOT EXISTS `system_users` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT '用户ID（Snowflake）',
    `username`   VARCHAR(50)     NOT NULL               COMMENT '登录用户名',
    `password`   VARCHAR(255)    NOT NULL               COMMENT '密码（SHA-256 加密）',
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
CREATE TABLE IF NOT EXISTS `system_user_dept` (
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `dept_id` BIGINT UNSIGNED NOT NULL COMMENT '部门ID',
    PRIMARY KEY (`user_id`, `dept_id`),
    KEY `idx_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户部门关联表';

-- ─────────────────────────────────────────────
-- 2.8 用户-角色关联表
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `system_user_role` (
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
INSERT INTO `system_dept` (`id`, `parent_id`, `title`, `username`, `email`, `sort`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (1000000000000000001, 0, '总公司', 'admin', 'admin@example.com', 0, 1, 0, 0, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 超级管理员角色（data_scope=1 全部数据权限）
-- ─────────────────────────────────────────────
INSERT INTO `system_role` (`id`, `parent_id`, `title`, `data_scope`, `sort`, `status`, `is_admin`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (1000000000000000002, 0, '超级管理员', 1, 0, 1, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 超级管理员用户
-- 用户名：admin
-- 密码：admin123  →  SHA-256 加密后的哈希值
-- SHA-256(admin123) = 240be518fabd2724ddb6f04eeb1da5967448d7e831c08c8fa822809f74c720a9
-- ─────────────────────────────────────────────
INSERT INTO `system_users` (`id`, `username`, `password`, `nickname`, `email`, `avatar`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (
    1000000000000000003,
    'admin',
    '240be518fabd2724ddb6f04eeb1da5967448d7e831c08c8fa822809f74c720a9',
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
INSERT INTO `system_user_dept` (`user_id`, `dept_id`)
VALUES (1000000000000000003, 1000000000000000001);

-- ─────────────────────────────────────────────
-- 用户-角色关联（admin 拥有超级管理员角色）
-- ─────────────────────────────────────────────
INSERT INTO `system_user_role` (`user_id`, `role_id`)
VALUES (1000000000000000003, 1000000000000000002);

-- ─────────────────────────────────────────────
-- 基础菜单树
-- 结构：
--   仪表盘（目录, id=1000000000000000060）
--     ├── 分析页（菜单, id=1000000000000000061）
--     └── 工作台（菜单, id=1000000000000000062）
--   系统管理（目录, id=1000000000000000010）
--     ├── 部门管理（菜单, id=1000000000000000011）
--     ├── 角色管理（菜单, id=1000000000000000012）
--     ├── 菜单管理（菜单, id=1000000000000000013）
--     └── 用户管理（菜单, id=1000000000000000014）
-- ─────────────────────────────────────────────
INSERT INTO `system_menu` (`id`, `parent_id`, `title`, `type`, `path`, `component`, `permission`, `icon`, `sort`, `is_show`, `is_cache`, `link_url`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
-- 仪表盘（目录）
(1000000000000000060, 0,                    '仪表盘',   1, '/dashboard',    NULL,                          '',                     'DashboardOutlined', 0, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 分析页（菜单）
(1000000000000000061, 1000000000000000060,  '分析页',   2, '/analytics',    'dashboard/analytics/index',   '',                     'AreaChartOutlined',  1, 1, 1, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 工作台（菜单）
(1000000000000000062, 1000000000000000060,  '工作台',   2, '/workspace',    'dashboard/workspace/index',   '',                     'DesktopOutlined',    2, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 系统管理（目录）
(1000000000000000010, 0,                    '系统管理', 1, '/system',       NULL,                          '',                     'SettingOutlined', 100, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 部门管理（菜单）
(1000000000000000011, 1000000000000000010,  '部门管理', 2, '/system/dept',  'system/dept/index',           'system:dept:list',     'ApartmentOutlined', 1, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 角色管理（菜单）
(1000000000000000012, 1000000000000000010,  '角色管理', 2, '/system/role',  'system/role/index',           'system:role:list',     'TeamOutlined',    2, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 菜单管理（菜单）
(1000000000000000013, 1000000000000000010,  '菜单管理', 2, '/system/menu',  'system/menu/index',           'system:menu:list',     'MenuOutlined',    3, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),

-- 用户管理（菜单）
(1000000000000000014, 1000000000000000010,  '用户管理', 2, '/system/users', 'system/users/index',          'system:user:list',     'UserOutlined',    4, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- ─────────────────────────────────────────────
-- 按钮级权限菜单（归属于各功能菜单下）
-- ─────────────────────────────────────────────
INSERT INTO `system_menu` (`id`, `parent_id`, `title`, `type`, `path`, `component`, `permission`, `icon`, `sort`, `is_show`, `is_cache`, `link_url`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
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
INSERT INTO `system_role_menu` (`role_id`, `menu_id`)
VALUES
-- 仪表盘
(1000000000000000002, 1000000000000000060),
(1000000000000000002, 1000000000000000061),
(1000000000000000002, 1000000000000000062),
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

-- ═════════════════════════════════════════════════════════════════
-- 陪玩平台 数据库表
-- 应用前缀：play
-- 所有表名格式：play_{module}
-- ═════════════════════════════════════════════════════════════════

-- ─────────────────────────────────────────────
-- 1. 会员等级表（play_member_level）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_member_level` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '等级ID（Snowflake）',
    `title`           VARCHAR(50)     NOT NULL               COMMENT '等级名称',
    `level`           TINYINT         NOT NULL DEFAULT 1     COMMENT '等级:1=普通会员,2=白银会员,3=黄金会员,4=铂金会员,5=钻石会员',
    `icon`            VARCHAR(100)                           COMMENT '等级图标',
    `min_exp`         INT             NOT NULL DEFAULT 0     COMMENT '所需最低经验值',
    `discount`        INT             NOT NULL DEFAULT 100   COMMENT '折扣（百分比，如 90 表示九折）',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员等级表';

-- ─────────────────────────────────────────────
-- 2. 会员表（play_member）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_member` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '会员ID（Snowflake）',
    `phone`           VARCHAR(20)     NOT NULL               COMMENT '手机号',
    `password`        VARCHAR(255)    NOT NULL               COMMENT '密码（bcrypt 加密）',
    `nickname`        VARCHAR(50)                            COMMENT '昵称',
    `avatar`          VARCHAR(500)                           COMMENT '头像',
    `gender`          TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '性别:0=未知,1=男,2=女',
    `member_level_id` BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '会员等级ID',
    `exp`             INT             NOT NULL DEFAULT 0     COMMENT '经验值',
    `balance`         BIGINT          NOT NULL DEFAULT 0     COMMENT '账户余额（分）',
    `is_coach`        TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否陪玩师:0=否,1=是',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=禁用,1=正常',
    `last_login_at`   DATETIME                               COMMENT '最后登录时间',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_phone` (`phone`),
    KEY `idx_member_level_id` (`member_level_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员表';

-- ─────────────────────────────────────────────
-- 3. 陪玩师等级表（play_coach_level）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_coach_level` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '等级ID（Snowflake）',
    `title`           VARCHAR(50)     NOT NULL               COMMENT '等级名称',
    `level`           TINYINT         NOT NULL DEFAULT 1     COMMENT '等级:1=青铜,2=白银,3=黄金,4=铂金,5=钻石',
    `icon`            VARCHAR(100)                           COMMENT '等级图标',
    `min_orders`      INT             NOT NULL DEFAULT 0     COMMENT '所需最低接单数',
    `min_score`       INT             NOT NULL DEFAULT 0     COMMENT '所需最低评分（乘100存储，如 450=4.50分）',
    `commission_rate` INT             NOT NULL DEFAULT 20    COMMENT '平台抽成比例（百分比，如 20 表示 20%）',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='陪玩师等级表';

-- ─────────────────────────────────────────────
-- 4. 陪玩师申请表（play_coach_apply）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_coach_apply` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '申请ID（Snowflake）',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `real_name`       VARCHAR(50)     NOT NULL               COMMENT '真实姓名',
    `id_card`         VARCHAR(30)     NOT NULL               COMMENT '身份证号',
    `id_card_front_image` VARCHAR(500) NOT NULL              COMMENT '身份证正面照',
    `id_card_back_image`  VARCHAR(500) NOT NULL              COMMENT '身份证反面照',
    `skill_desc`      TEXT                                   COMMENT '技能描述',
    `audit_status`    TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '审核状态:0=待审核,1=通过,2=拒绝',
    `audit_remark`    VARCHAR(500)                           COMMENT '审核备注',
    `audit_at`        DATETIME                               COMMENT '审核时间',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_audit_status` (`audit_status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='陪玩师申请表';

-- ─────────────────────────────────────────────
-- 5. 陪玩师表（play_coach）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_coach` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '陪玩师ID（Snowflake）',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '关联会员ID',
    `coach_level_id`  BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '陪玩师等级ID',
    `shop_id`         BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '所属店铺ID（0表示无店铺）',
    `real_name`       VARCHAR(50)     NOT NULL               COMMENT '真实姓名',
    `intro`           VARCHAR(500)                           COMMENT '个人简介',
    `cover_image`     VARCHAR(500)                           COMMENT '封面图',
    `total_orders`    INT             NOT NULL DEFAULT 0     COMMENT '总接单数',
    `total_score`     INT             NOT NULL DEFAULT 500   COMMENT '总评分（乘100，如 500=5.00）',
    `score_num`       INT             NOT NULL DEFAULT 0     COMMENT '评分人数',
    `income_total`    BIGINT          NOT NULL DEFAULT 0     COMMENT '累计收入（分）',
    `income_balance`  BIGINT          NOT NULL DEFAULT 0     COMMENT '可提现余额（分）',
    `is_online`       TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否在线:0=离线,1=在线',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=禁用,1=正常',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_member_id` (`member_id`),
    KEY `idx_coach_level_id` (`coach_level_id`),
    KEY `idx_shop_id` (`shop_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='陪玩师表';

-- ─────────────────────────────────────────────
-- 6. 店铺表（play_shop）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_shop` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '店铺ID（Snowflake）',
    `title`           VARCHAR(100)    NOT NULL               COMMENT '店铺名称',
    `logo_image`      VARCHAR(500)                           COMMENT '店铺LOGO',
    `cover_image`     VARCHAR(500)                           COMMENT '封面图',
    `contact_name`    VARCHAR(50)                            COMMENT '联系人姓名',
    `contact_phone`   VARCHAR(20)                            COMMENT '联系电话',
    `intro`           VARCHAR(500)                           COMMENT '店铺简介',
    `commission_rate` INT             NOT NULL DEFAULT 10    COMMENT '店铺抽成比例（百分比，如 10 表示 10%）',
    `coach_num`       INT             NOT NULL DEFAULT 0     COMMENT '陪玩师数量',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='店铺表';

-- ─────────────────────────────────────────────
-- 7. 商品分类表（play_category）— 树形结构
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_category` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '分类ID（Snowflake）',
    `parent_id`       BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '上级分类ID，0 表示顶级分类',
    `title`           VARCHAR(50)     NOT NULL               COMMENT '分类名称',
    `icon`            VARCHAR(100)                           COMMENT '分类图标',
    `cover_image`     VARCHAR(500)                           COMMENT '分类封面图',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品分类表';

-- ─────────────────────────────────────────────
-- 8. 商品表（play_goods）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_goods` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '商品ID（Snowflake）',
    `category_id`     BIGINT UNSIGNED NOT NULL               COMMENT '分类ID',
    `coach_id`        BIGINT UNSIGNED NOT NULL               COMMENT '陪玩师ID',
    `title`           VARCHAR(100)    NOT NULL               COMMENT '商品名称',
    `cover_image`     VARCHAR(500)                           COMMENT '商品封面图',
    `desc_content`    TEXT                                   COMMENT '商品详情描述',
    `price`           BIGINT          NOT NULL DEFAULT 0     COMMENT '单价（分）',
    `unit`            VARCHAR(20)     NOT NULL DEFAULT '局'  COMMENT '计量单位（如：局、小时、把）',
    `sales_num`       INT             NOT NULL DEFAULT 0     COMMENT '销量',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=下架,1=上架',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`),
    KEY `idx_coach_id` (`coach_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品表';

-- ─────────────────────────────────────────────
-- 9. 订单表（play_order）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_order` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '订单ID（Snowflake）',
    `order_no`        VARCHAR(32)     NOT NULL               COMMENT '订单编号',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '下单会员ID',
    `coach_id`        BIGINT UNSIGNED NOT NULL               COMMENT '陪玩师ID',
    `shop_id`         BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '店铺ID（0表示无店铺）',
    `goods_id`        BIGINT UNSIGNED NOT NULL               COMMENT '商品ID',
    `goods_title`     VARCHAR(100)    NOT NULL               COMMENT '商品名称（冗余）',
    `goods_price`     BIGINT          NOT NULL               COMMENT '商品单价（分，下单时快照）',
    `quantity`        INT             NOT NULL DEFAULT 1     COMMENT '数量',
    `total_amount`    BIGINT          NOT NULL DEFAULT 0     COMMENT '订单总额（分）',
    `discount_amount` BIGINT          NOT NULL DEFAULT 0     COMMENT '会员折扣金额（分）',
    `coupon_amount`   BIGINT          NOT NULL DEFAULT 0     COMMENT '优惠券抵扣金额（分）',
    `pay_amount`      BIGINT          NOT NULL DEFAULT 0     COMMENT '实付金额（分）',
    `coupon_member_id` BIGINT UNSIGNED NOT NULL DEFAULT 0    COMMENT '使用的优惠券领取记录ID',
    `pay_type`        TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '支付方式:0=未支付,1=微信支付,2=支付宝支付,3=余额支付',
    `order_status`    TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '订单状态:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款',
    `pay_at`          DATETIME                               COMMENT '支付时间',
    `start_at`        DATETIME                               COMMENT '服务开始时间',
    `finish_at`       DATETIME                               COMMENT '服务完成时间',
    `cancel_at`       DATETIME                               COMMENT '取消时间',
    `cancel_reason`   VARCHAR(500)                           COMMENT '取消原因',
    `remark`          VARCHAR(500)                           COMMENT '订单备注',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_no` (`order_no`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_coach_id` (`coach_id`),
    KEY `idx_shop_id` (`shop_id`),
    KEY `idx_order_status` (`order_status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单表';

-- ─────────────────────────────────────────────
-- 10. 支付记录表（play_payment）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_payment` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '支付记录ID（Snowflake）',
    `order_id`        BIGINT UNSIGNED NOT NULL               COMMENT '订单ID',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `payment_no`      VARCHAR(64)     NOT NULL               COMMENT '支付流水号（平台内部）',
    `trade_no`        VARCHAR(64)                            COMMENT '第三方交易号',
    `pay_type`        TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '支付方式:1=微信支付,2=支付宝支付,3=余额支付',
    `pay_amount`      BIGINT          NOT NULL DEFAULT 0     COMMENT '支付金额（分）',
    `pay_status`      TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '支付状态:0=待支付,1=支付成功,2=支付失败,3=已退款',
    `pay_at`          DATETIME                               COMMENT '支付成功时间',
    `refund_at`       DATETIME                               COMMENT '退款时间',
    `refund_amount`   BIGINT          NOT NULL DEFAULT 0     COMMENT '退款金额（分）',
    `callback_content` TEXT                                  COMMENT '回调报文',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_payment_no` (`payment_no`),
    KEY `idx_order_id` (`order_id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付记录表';

-- ─────────────────────────────────────────────
-- 11. 充值方案表（play_recharge_plan）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_recharge_plan` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '方案ID（Snowflake）',
    `title`           VARCHAR(50)     NOT NULL               COMMENT '方案名称',
    `amount`          BIGINT          NOT NULL               COMMENT '充值金额（分）',
    `gift_amount`     BIGINT          NOT NULL DEFAULT 0     COMMENT '赠送金额（分）',
    `cover_image`     VARCHAR(500)                           COMMENT '方案封面图',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='充值方案表';

-- ─────────────────────────────────────────────
-- 12. 充值订单表（play_recharge_order）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_recharge_order` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '充值订单ID（Snowflake）',
    `order_no`        VARCHAR(32)     NOT NULL               COMMENT '充值订单号',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `recharge_plan_id` BIGINT UNSIGNED NOT NULL              COMMENT '充值方案ID',
    `amount`          BIGINT          NOT NULL               COMMENT '充值金额（分）',
    `gift_amount`     BIGINT          NOT NULL DEFAULT 0     COMMENT '赠送金额（分）',
    `pay_type`        TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '支付方式:1=微信支付,2=支付宝支付',
    `trade_no`        VARCHAR(64)                            COMMENT '第三方交易号',
    `pay_status`      TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '支付状态:0=待支付,1=支付成功,2=支付失败',
    `pay_at`          DATETIME                               COMMENT '支付时间',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_no` (`order_no`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='充值订单表';

-- ─────────────────────────────────────────────
-- 13. 余额流水表（play_balance_log）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_balance_log` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '流水ID（Snowflake）',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `biz_type`        TINYINT(1)      NOT NULL               COMMENT '业务类型:1=充值,2=消费,3=退款,4=活动赠送,5=提现',
    `biz_id`          BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '关联业务ID（订单ID/充值订单ID/活动ID）',
    `change_amount`   BIGINT          NOT NULL               COMMENT '变动金额（分，正数增加负数减少）',
    `before_balance`  BIGINT          NOT NULL               COMMENT '变动前余额（分）',
    `after_balance`   BIGINT          NOT NULL               COMMENT '变动后余额（分）',
    `remark`          VARCHAR(200)                           COMMENT '备注说明',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_biz_type` (`biz_type`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='余额流水表';

-- ─────────────────────────────────────────────
-- 14. 活动表（play_activity）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_activity` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '活动ID（Snowflake）',
    `title`           VARCHAR(100)    NOT NULL               COMMENT '活动名称',
    `cover_image`     VARCHAR(500)                           COMMENT '活动封面图',
    `desc_content`    TEXT                                   COMMENT '活动详情描述（富文本，支持图文混排）',
    `type`            TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '活动类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动',
    `condition_type`  TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '参与条件:0=无条件,1=需报名,2=充值满额,3=下单满额,4=完成步骤',
    `condition_value` BIGINT          NOT NULL DEFAULT 0     COMMENT '条件值（分/次，如充值满5000分、下单满3次）',
    `is_auto_reward`  TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '是否自动发奖:0=否（需审核）,1=是（用户完成即发）',
    `start_at`        DATETIME        NOT NULL               COMMENT '活动开始时间',
    `end_at`          DATETIME        NOT NULL               COMMENT '活动结束时间',
    `max_num`         INT             NOT NULL DEFAULT 0     COMMENT '参与人数上限（0表示不限）',
    `join_num`        INT             NOT NULL DEFAULT 0     COMMENT '已参与人数',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_type` (`type`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动表';

-- ─────────────────────────────────────────────
-- 14b. 活动奖励表（play_activity_reward）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_activity_reward` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '奖励ID（Snowflake）',
    `activity_id`     BIGINT UNSIGNED NOT NULL               COMMENT '活动ID',
    `reward_type`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '奖励类型:1=余额,2=优惠券,3=经验值,4=会员等级天数',
    `reward_value`    BIGINT          NOT NULL DEFAULT 0     COMMENT '奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）',
    `reward_name`     VARCHAR(100)    NOT NULL               COMMENT '奖励名称（展示用，如"送50元余额"）',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_activity_id` (`activity_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动奖励表';

-- ─────────────────────────────────────────────
-- 14c. 活动步骤表（play_activity_step）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_activity_step` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '步骤ID（Snowflake）',
    `activity_id`     BIGINT UNSIGNED NOT NULL               COMMENT '活动ID',
    `step_num`        INT             NOT NULL DEFAULT 1     COMMENT '步骤序号',
    `title`           VARCHAR(100)    NOT NULL               COMMENT '步骤标题',
    `step_type`       TINYINT         NOT NULL DEFAULT 1     COMMENT '步骤类型：1=文字 2=链接 3=图片',
    `example_text`    VARCHAR(500)    NOT NULL DEFAULT ''    COMMENT '示例文字或链接URL',
    `desc_content`    TEXT                                   COMMENT '步骤说明（富文本，支持图文）',
    `step_image`      VARCHAR(500)                           COMMENT '步骤示例图片',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_activity_id` (`activity_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动步骤表';

-- ─────────────────────────────────────────────
-- 14d. 活动参与记录表（play_activity_join）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_activity_join` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '记录ID（Snowflake）',
    `activity_id`     BIGINT UNSIGNED NOT NULL               COMMENT '活动ID',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `join_status`     TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '参与状态:0=已报名,1=进行中,2=已完成,3=已领奖',
    `current_step`    INT             NOT NULL DEFAULT 0     COMMENT '当前完成到第几步（步骤活动用）',
    `finish_at`       DATETIME                               COMMENT '完成时间',
    `reward_at`       DATETIME                               COMMENT '领奖时间',
    `remark`          VARCHAR(500)                           COMMENT '备注',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_activity_member` (`activity_id`, `member_id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_join_status` (`join_status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动参与记录表';

-- ─────────────────────────────────────────────
-- 15. 优惠券模板表（play_coupon）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_coupon` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '优惠券ID（Snowflake）',
    `title`           VARCHAR(100)    NOT NULL               COMMENT '优惠券名称',
    `type`            TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '优惠券类型:1=满减券,2=折扣券,3=无门槛券',
    `is_new_member`   TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否新人专享:0=否,1=是',
    `face_value`      BIGINT          NOT NULL DEFAULT 0     COMMENT '面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）',
    `min_amount`      BIGINT          NOT NULL DEFAULT 0     COMMENT '最低消费金额（分，0表示无门槛）',
    `total_num`       INT             NOT NULL DEFAULT 0     COMMENT '发放总量（0表示不限）',
    `used_num`        INT             NOT NULL DEFAULT 0     COMMENT '已使用数量',
    `claim_num`       INT             NOT NULL DEFAULT 0     COMMENT '已领取数量',
    `per_limit`       INT             NOT NULL DEFAULT 1     COMMENT '每人限领张数',
    `valid_start_at`  DATETIME        NOT NULL               COMMENT '有效期开始时间',
    `valid_end_at`    DATETIME        NOT NULL               COMMENT '有效期结束时间',
    `sort`            INT             NOT NULL DEFAULT 0     COMMENT '排序（升序）',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=关闭,1=开启',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_is_new_member` (`is_new_member`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='优惠券模板表';

-- ─────────────────────────────────────────────
-- 16. 会员优惠券表（play_coupon_member）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_coupon_member` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '记录ID（Snowflake）',
    `coupon_id`       BIGINT UNSIGNED NOT NULL               COMMENT '优惠券模板ID',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `order_id`        BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '使用的订单ID（0表示未使用）',
    `use_status`      TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '使用状态:0=未使用,1=已使用,2=已过期',
    `claim_at`        DATETIME                               COMMENT '领取时间',
    `use_at`          DATETIME                               COMMENT '使用时间',
    `expire_at`       DATETIME                               COMMENT '过期时间',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_coupon_id` (`coupon_id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_use_status` (`use_status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员优惠券表';

-- ─────────────────────────────────────────────
-- 17. 第三方登录绑定表（play_oauth）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_oauth` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '记录ID（Snowflake）',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '会员ID',
    `provider`        TINYINT(1)      NOT NULL               COMMENT '第三方平台:1=微信,2=支付宝',
    `open_id`         VARCHAR(128)    NOT NULL               COMMENT '第三方OpenID',
    `union_id`        VARCHAR(128)                           COMMENT '第三方UnionID',
    `nickname`        VARCHAR(50)                            COMMENT '第三方昵称',
    `avatar`          VARCHAR(500)                           COMMENT '第三方头像',
    `access_token`    VARCHAR(500)                           COMMENT '访问令牌',
    `refresh_token`   VARCHAR(500)                           COMMENT '刷新令牌',
    `expire_at`       DATETIME                               COMMENT '令牌过期时间',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_provider_open_id` (`provider`, `open_id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='第三方登录绑定表';

-- ─────────────────────────────────────────────
-- 18. 评价表（play_review）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_review` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '评价ID（Snowflake）',
    `order_id`        BIGINT UNSIGNED NOT NULL               COMMENT '订单ID',
    `member_id`       BIGINT UNSIGNED NOT NULL               COMMENT '评价会员ID',
    `coach_id`        BIGINT UNSIGNED NOT NULL               COMMENT '被评陪玩师ID',
    `score`           INT             NOT NULL DEFAULT 500   COMMENT '评分（乘100，如 500=5.00分）',
    `review_content`  TEXT                                   COMMENT '评价内容',
    `review_image`    VARCHAR(2000)                          COMMENT '评价图片（多张逗号分隔）',
    `reply_content`   TEXT                                   COMMENT '陪玩师回复内容',
    `reply_at`        DATETIME                               COMMENT '回复时间',
    `is_anonymous`    TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否匿名:0=否,1=是',
    `status`          TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=隐藏,1=显示',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_id` (`order_id`),
    KEY `idx_member_id` (`member_id`),
    KEY `idx_coach_id` (`coach_id`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评价表';

-- ─────────────────────────────────────────────
-- 19. 利润分成流水表（play_profit_log）
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS `play_profit_log` (
    `id`              BIGINT UNSIGNED NOT NULL               COMMENT '流水ID（Snowflake）',
    `order_id`        BIGINT UNSIGNED NOT NULL               COMMENT '订单ID',
    `order_no`        VARCHAR(32)     NOT NULL               COMMENT '订单编号',
    `pay_amount`      BIGINT          NOT NULL DEFAULT 0     COMMENT '实付金额（分）',
    `coach_id`        BIGINT UNSIGNED NOT NULL               COMMENT '陪玩师ID',
    `shop_id`         BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT '店铺ID（0表示无店铺）',
    `platform_rate`   INT             NOT NULL DEFAULT 0     COMMENT '平台抽成比例（百分比）',
    `platform_amount` BIGINT          NOT NULL DEFAULT 0     COMMENT '平台抽成金额（分）',
    `shop_rate`       INT             NOT NULL DEFAULT 0     COMMENT '店铺抽成比例（百分比）',
    `shop_amount`     BIGINT          NOT NULL DEFAULT 0     COMMENT '店铺抽成金额（分）',
    `coach_amount`    BIGINT          NOT NULL DEFAULT 0     COMMENT '陪玩师收入（分）',
    `settle_status`   TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '结算状态:0=待结算,1=已结算',
    `settle_at`       DATETIME                               COMMENT '结算时间',
    `created_by`      BIGINT UNSIGNED                        COMMENT '创建人ID',
    `dept_id`         BIGINT UNSIGNED                        COMMENT '所属部门ID',
    `created_at`      DATETIME                               COMMENT '创建时间',
    `updated_at`      DATETIME                               COMMENT '更新时间',
    `deleted_at`      DATETIME                               COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_order_id` (`order_id`),
    KEY `idx_coach_id` (`coach_id`),
    KEY `idx_shop_id` (`shop_id`),
    KEY `idx_settle_status` (`settle_status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='利润分成流水表';

-- =============================================
-- Upload 应用 — 文件管理系统
-- =============================================

-- 16a. 文件目录（upload_dir）
CREATE TABLE IF NOT EXISTS `upload_dir` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT 'ID',
    `parent_id`  BIGINT UNSIGNED DEFAULT 0              COMMENT '上级目录',
    `name`       VARCHAR(100)    NOT NULL               COMMENT '目录名称',
    `path`       VARCHAR(500)    NOT NULL               COMMENT '目录路径',
    `sort`       INT             DEFAULT 0              COMMENT '排序',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=禁用,1=启用',
    `created_at` DATETIME        DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '删除时间',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '部门ID',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件目录';

-- 16b. 文件记录（upload_file）
CREATE TABLE IF NOT EXISTS `upload_file` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT 'ID',
    `dir_id`     BIGINT UNSIGNED DEFAULT 0              COMMENT '所属目录',
    `name`       VARCHAR(255)    NOT NULL               COMMENT '文件名称',
    `url`        VARCHAR(500)    NOT NULL               COMMENT '文件地址',
    `ext`        VARCHAR(20)     DEFAULT ''             COMMENT '文件扩展名',
    `size`       BIGINT UNSIGNED DEFAULT 0              COMMENT '文件大小(字节)',
    `mime`       VARCHAR(100)    DEFAULT ''             COMMENT 'MIME类型',
    `storage`    TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '存储类型:1=本地,2=阿里云OSS,3=腾讯云COS',
    `is_image`   TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否图片:0=否,1=是',
    `created_at` DATETIME        DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '删除时间',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '部门ID',
    PRIMARY KEY (`id`),
    KEY `idx_dir_id` (`dir_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件记录';

-- 16c. 上传配置（upload_config）
CREATE TABLE IF NOT EXISTS `upload_config` (
    `id`             BIGINT UNSIGNED NOT NULL               COMMENT 'ID',
    `name`           VARCHAR(100)    NOT NULL               COMMENT '配置名称',
    `storage`        TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '存储类型:1=本地,2=阿里云OSS,3=腾讯云COS',
    `is_default`     TINYINT(1)      NOT NULL DEFAULT 0     COMMENT '是否默认:0=否,1=是',
    `local_path`     VARCHAR(500)    DEFAULT ''             COMMENT '本地存储路径',
    `oss_endpoint`   VARCHAR(255)    DEFAULT ''             COMMENT 'OSS Endpoint',
    `oss_bucket`     VARCHAR(255)    DEFAULT ''             COMMENT 'OSS Bucket',
    `oss_access_key` VARCHAR(255)    DEFAULT ''             COMMENT 'OSS AccessKey',
    `oss_secret_key` VARCHAR(255)    DEFAULT ''             COMMENT 'OSS SecretKey',
    `cos_region`     VARCHAR(100)    DEFAULT ''             COMMENT 'COS Region',
    `cos_bucket`     VARCHAR(255)    DEFAULT ''             COMMENT 'COS Bucket',
    `cos_secret_id`  VARCHAR(255)    DEFAULT ''             COMMENT 'COS SecretId',
    `cos_secret_key` VARCHAR(255)    DEFAULT ''             COMMENT 'COS SecretKey',
    `max_size`       INT             DEFAULT 10             COMMENT '最大文件大小(MB)',
    `status`         TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=禁用,1=启用',
    `created_at`     DATETIME        DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     DATETIME        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`     DATETIME                               COMMENT '删除时间',
    `created_by`     BIGINT UNSIGNED                        COMMENT '创建人',
    `dept_id`        BIGINT UNSIGNED                        COMMENT '部门ID',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='上传配置';

-- 16d. 文件目录规则（upload_dir_rule）
CREATE TABLE IF NOT EXISTS `upload_dir_rule` (
    `id`         BIGINT UNSIGNED NOT NULL               COMMENT 'ID',
    `dir_id`     BIGINT UNSIGNED NOT NULL               COMMENT '目录ID',
    `category`   TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '类别:1=默认,2=类型,3=接口',
    `save_path`  VARCHAR(500)    DEFAULT ''             COMMENT '保存目录',
    `status`     TINYINT(1)      NOT NULL DEFAULT 1     COMMENT '状态:0=禁用,1=启用',
    `created_at` DATETIME        DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME                               COMMENT '删除时间',
    `created_by` BIGINT UNSIGNED                        COMMENT '创建人',
    `dept_id`    BIGINT UNSIGNED                        COMMENT '部门ID',
    PRIMARY KEY (`id`),
    KEY `idx_dir_id` (`dir_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件目录规则';

-- ─────────────────────────────────────────────
-- Upload 应用菜单
-- ─────────────────────────────────────────────
INSERT INTO `system_menu` (`id`, `parent_id`, `title`, `type`, `path`, `component`, `permission`, `icon`, `sort`, `is_show`, `is_cache`, `link_url`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
-- 文件管理（目录）
(1000000000000000070, 0,                    '文件管理',     1, '/upload',           NULL,                        '',                       'CloudUploadOutlined', 90, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 文件目录（菜单）
(1000000000000000071, 1000000000000000070,  '文件目录',     2, '/upload/dir',       'upload/dir/index',          'upload:dir:list',        '', 1, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 文件列表（菜单）
(1000000000000000072, 1000000000000000070,  '文件列表',     2, '/upload/file',      'upload/file/index',         'upload:file:list',       '', 2, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 上传配置（菜单）
(1000000000000000073, 1000000000000000070,  '上传配置',     2, '/upload/config',    'upload/config/index',       'upload:config:list',     '', 3, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 目录规则（菜单）
(1000000000000000074, 1000000000000000070,  '目录规则',     2, '/upload/dir-rule',  'upload/dir_rule/index',     'upload:dir_rule:list',   '', 4, 1, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- Upload 按钮权限
INSERT INTO `system_menu` (`id`, `parent_id`, `title`, `type`, `path`, `component`, `permission`, `icon`, `sort`, `is_show`, `is_cache`, `link_url`, `status`, `created_by`, `dept_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
-- 文件目录按钮
(1000000000000000081, 1000000000000000071, '文件目录新增', 3, NULL, NULL, 'upload:dir:create',      '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000082, 1000000000000000071, '文件目录修改', 3, NULL, NULL, 'upload:dir:update',      '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000083, 1000000000000000071, '文件目录删除', 3, NULL, NULL, 'upload:dir:delete',      '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 文件列表按钮
(1000000000000000084, 1000000000000000072, '文件记录新增', 3, NULL, NULL, 'upload:file:create',     '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000085, 1000000000000000072, '文件记录修改', 3, NULL, NULL, 'upload:file:update',     '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000086, 1000000000000000072, '文件记录删除', 3, NULL, NULL, 'upload:file:delete',     '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 上传配置按钮
(1000000000000000087, 1000000000000000073, '上传配置新增', 3, NULL, NULL, 'upload:config:create',   '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000088, 1000000000000000073, '上传配置修改', 3, NULL, NULL, 'upload:config:update',   '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000089, 1000000000000000073, '上传配置删除', 3, NULL, NULL, 'upload:config:delete',   '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
-- 目录规则按钮
(1000000000000000090, 1000000000000000074, '目录规则新增', 3, NULL, NULL, 'upload:dir_rule:create', '', 1, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000091, 1000000000000000074, '目录规则修改', 3, NULL, NULL, 'upload:dir_rule:update', '', 2, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL),
(1000000000000000092, 1000000000000000074, '目录规则删除', 3, NULL, NULL, 'upload:dir_rule:delete', '', 3, 0, 0, NULL, 1, 0, 1000000000000000001, NOW(), NOW(), NULL);

-- Upload 角色-菜单关联（超级管理员）
INSERT INTO `system_role_menu` (`role_id`, `menu_id`)
VALUES
(1000000000000000002, 1000000000000000070),
(1000000000000000002, 1000000000000000071),
(1000000000000000002, 1000000000000000072),
(1000000000000000002, 1000000000000000073),
(1000000000000000002, 1000000000000000074),
(1000000000000000002, 1000000000000000081),
(1000000000000000002, 1000000000000000082),
(1000000000000000002, 1000000000000000083),
(1000000000000000002, 1000000000000000084),
(1000000000000000002, 1000000000000000085),
(1000000000000000002, 1000000000000000086),
(1000000000000000002, 1000000000000000087),
(1000000000000000002, 1000000000000000088),
(1000000000000000002, 1000000000000000089),
(1000000000000000002, 1000000000000000090),
(1000000000000000002, 1000000000000000091),
(1000000000000000002, 1000000000000000092);

SET FOREIGN_KEY_CHECKS = 1;
