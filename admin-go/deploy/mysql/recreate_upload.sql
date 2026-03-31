SET NAMES utf8mb4;

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
