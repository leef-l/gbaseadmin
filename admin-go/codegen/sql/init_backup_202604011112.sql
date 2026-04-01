-- MySQL dump 10.13  Distrib 8.0.45, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: gbaseadmin
-- ------------------------------------------------------
-- Server version	8.0.45

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `play_activity`
--

DROP TABLE IF EXISTS `play_activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_activity` (
  `id` bigint unsigned NOT NULL COMMENT '活动ID（Snowflake）',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '活动名称',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '活动封面图',
  `desc_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '活动详情描述（富文本，支持图文混排）',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '活动类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动',
  `condition_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '参与条件:0=无条件,1=需报名,2=充值满额,3=下单满额,4=完成步骤',
  `condition_value` bigint NOT NULL DEFAULT '0' COMMENT '条件值（分/次，如充值满5000分、下单满3次）',
  `is_auto_reward` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否自动发奖:0=否（需审核）,1=是（用户完成即发）',
  `start_at` datetime NOT NULL COMMENT '活动开始时间',
  `end_at` datetime NOT NULL COMMENT '活动结束时间',
  `max_num` int NOT NULL DEFAULT '0' COMMENT '参与人数上限（0表示不限）',
  `join_num` int NOT NULL DEFAULT '0' COMMENT '已参与人数',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_type` (`type`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_activity`
--

LOCK TABLES `play_activity` WRITE;
/*!40000 ALTER TABLE `play_activity` DISABLE KEYS */;
INSERT INTO `play_activity` VALUES (314312344010231808,'测试步骤活动','/upload/2026-03-31/17749370568554486713.png','',4,0,0,1,'2026-03-31 06:18:11','2026-04-04 06:18:14',0,0,0,1,NULL,NULL,'2026-03-31 06:18:22','2026-03-31 06:18:22',NULL);
/*!40000 ALTER TABLE `play_activity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_activity_join`
--

DROP TABLE IF EXISTS `play_activity_join`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_activity_join` (
  `id` bigint unsigned NOT NULL COMMENT '记录ID（Snowflake）',
  `activity_id` bigint unsigned NOT NULL COMMENT '活动ID',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `join_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '参与状态:0=已报名,1=进行中,2=已完成,3=已领奖',
  `current_step` int NOT NULL DEFAULT '0' COMMENT '当前完成到第几步（步骤活动用）',
  `finish_at` datetime DEFAULT NULL COMMENT '完成时间',
  `reward_at` datetime DEFAULT NULL COMMENT '领奖时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_activity_member` (`activity_id`,`member_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_join_status` (`join_status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动参与记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_activity_join`
--

LOCK TABLES `play_activity_join` WRITE;
/*!40000 ALTER TABLE `play_activity_join` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_activity_join` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_activity_reward`
--

DROP TABLE IF EXISTS `play_activity_reward`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_activity_reward` (
  `id` bigint unsigned NOT NULL COMMENT '奖励ID（Snowflake）',
  `activity_id` bigint unsigned NOT NULL COMMENT '活动ID',
  `reward_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '奖励类型:1=余额,2=优惠券,3=经验值,4=会员等级天数',
  `reward_value` bigint NOT NULL DEFAULT '0' COMMENT '奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）',
  `reward_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '奖励名称（展示用，如"送50元余额"）',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_activity_id` (`activity_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动奖励表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_activity_reward`
--

LOCK TABLES `play_activity_reward` WRITE;
/*!40000 ALTER TABLE `play_activity_reward` DISABLE KEYS */;
INSERT INTO `play_activity_reward` VALUES (314312543306780672,314312344010231808,1,5,'5元现金，完成活动立刻到手',0,NULL,NULL,'2026-03-31 06:19:10','2026-03-31 06:19:10',NULL);
/*!40000 ALTER TABLE `play_activity_reward` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_activity_step`
--

DROP TABLE IF EXISTS `play_activity_step`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_activity_step` (
  `id` bigint unsigned NOT NULL COMMENT '步骤ID（Snowflake）',
  `activity_id` bigint unsigned NOT NULL COMMENT '活动ID',
  `step_num` int NOT NULL DEFAULT '1' COMMENT '步骤序号',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '步骤标题',
  `step_type` tinyint NOT NULL DEFAULT '1' COMMENT '步骤类型：1=文字 2=链接 3=图片',
  `example_text` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '示例文字或链接URL',
  `desc_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '步骤说明（富文本，支持图文）',
  `step_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '步骤示例图片',
  `is_required` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否需要填写:0=不需要,1=需要',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_activity_id` (`activity_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动步骤表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_activity_step`
--

LOCK TABLES `play_activity_step` WRITE;
/*!40000 ALTER TABLE `play_activity_step` DISABLE KEYS */;
INSERT INTO `play_activity_step` VALUES (314312696793141248,314312344010231808,1,'复制链接到拼多多app搜索',0,'','','',0,0,NULL,NULL,'2026-03-31 06:19:47','2026-03-31 06:20:10',NULL),(314313291071492096,314312344010231808,2,'aaaa',0,'','afdf','https://pic4.zhimg.com/v2-267217db1a16cf81a50184ea7d520289_1440w.jpg',0,0,NULL,NULL,'2026-03-31 06:22:08','2026-03-31 06:22:08',NULL),(314313340186791936,314312344010231808,3,'啊啊啊',0,'','啊啊啊','',0,0,NULL,NULL,'2026-03-31 06:22:20','2026-03-31 06:22:20',NULL);
/*!40000 ALTER TABLE `play_activity_step` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_activity_step_log`
--

DROP TABLE IF EXISTS `play_activity_step_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_activity_step_log` (
  `id` bigint unsigned NOT NULL COMMENT '记录ID（Snowflake）',
  `activity_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '活动ID',
  `step_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '步骤ID',
  `join_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '参与记录ID',
  `member_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '会员ID',
  `step_type` tinyint NOT NULL DEFAULT '1' COMMENT '步骤类型：1=文字 2=链接 3=图片',
  `submit_text` text COMMENT '用户提交的文字或链接',
  `submit_image` varchar(500) DEFAULT NULL COMMENT '用户提交的图片URL',
  `audit_status` tinyint NOT NULL DEFAULT '0' COMMENT '审核状态：0=待审核 1=通过 2=驳回',
  `audit_remark` varchar(500) DEFAULT NULL COMMENT '审核备注',
  `audit_by` bigint unsigned DEFAULT NULL COMMENT '审核人ID',
  `audit_at` datetime DEFAULT NULL COMMENT '审核时间',
  `created_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `dept_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_join_id` (`join_id`),
  KEY `idx_activity_step` (`activity_id`,`step_id`),
  KEY `idx_member` (`member_id`),
  KEY `idx_audit_status` (`audit_status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='活动步骤提交记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_activity_step_log`
--

LOCK TABLES `play_activity_step_log` WRITE;
/*!40000 ALTER TABLE `play_activity_step_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_activity_step_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_balance_log`
--

DROP TABLE IF EXISTS `play_balance_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_balance_log` (
  `id` bigint unsigned NOT NULL COMMENT '流水ID（Snowflake）',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `biz_type` tinyint(1) NOT NULL COMMENT '业务类型:1=充值,2=消费,3=退款,4=活动赠送,5=提现',
  `biz_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '关联业务ID（订单ID/充值订单ID/活动ID）',
  `change_amount` bigint NOT NULL COMMENT '变动金额（分，正数增加负数减少）',
  `before_balance` bigint NOT NULL COMMENT '变动前余额（分）',
  `after_balance` bigint NOT NULL COMMENT '变动后余额（分）',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注说明',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_biz_type` (`biz_type`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='余额流水表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_balance_log`
--

LOCK TABLES `play_balance_log` WRITE;
/*!40000 ALTER TABLE `play_balance_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_balance_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_category`
--

DROP TABLE IF EXISTS `play_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_category` (
  `id` bigint unsigned NOT NULL COMMENT '分类ID（Snowflake）',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上级分类ID，0 表示顶级分类',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分类图标',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分类封面图',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_category`
--

LOCK TABLES `play_category` WRITE;
/*!40000 ALTER TABLE `play_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_coach`
--

DROP TABLE IF EXISTS `play_coach`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_coach` (
  `id` bigint unsigned NOT NULL COMMENT '陪玩师ID（Snowflake）',
  `member_id` bigint unsigned NOT NULL COMMENT '关联会员ID',
  `coach_level_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '陪玩师等级ID',
  `shop_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '所属店铺ID（0表示无店铺）',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '真实姓名',
  `intro` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '个人简介',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '封面图',
  `total_orders` int NOT NULL DEFAULT '0' COMMENT '总接单数',
  `total_score` int NOT NULL DEFAULT '500' COMMENT '总评分（乘100，如 500=5.00）',
  `score_num` int NOT NULL DEFAULT '0' COMMENT '评分人数',
  `income_total` bigint NOT NULL DEFAULT '0' COMMENT '累计收入（分）',
  `income_balance` bigint NOT NULL DEFAULT '0' COMMENT '可提现余额（分）',
  `is_online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否在线:0=离线,1=在线',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=禁用,1=正常',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_member_id` (`member_id`),
  KEY `idx_coach_level_id` (`coach_level_id`),
  KEY `idx_shop_id` (`shop_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='陪玩师表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_coach`
--

LOCK TABLES `play_coach` WRITE;
/*!40000 ALTER TABLE `play_coach` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_coach` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_coach_apply`
--

DROP TABLE IF EXISTS `play_coach_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_coach_apply` (
  `id` bigint unsigned NOT NULL COMMENT '申请ID（Snowflake）',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '真实姓名',
  `id_card` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '身份证号',
  `id_card_front_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '身份证正面照',
  `id_card_back_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '身份证反面照',
  `skill_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '技能描述',
  `audit_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核状态:0=待审核,1=通过,2=拒绝',
  `audit_remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '审核备注',
  `audit_at` datetime DEFAULT NULL COMMENT '审核时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_audit_status` (`audit_status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='陪玩师申请表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_coach_apply`
--

LOCK TABLES `play_coach_apply` WRITE;
/*!40000 ALTER TABLE `play_coach_apply` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_coach_apply` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_coach_level`
--

DROP TABLE IF EXISTS `play_coach_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_coach_level` (
  `id` bigint unsigned NOT NULL COMMENT '等级ID（Snowflake）',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '等级名称',
  `level` tinyint NOT NULL DEFAULT '1' COMMENT '等级:1=青铜,2=白银,3=黄金,4=铂金,5=钻石',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '等级图标',
  `min_orders` int NOT NULL DEFAULT '0' COMMENT '所需最低接单数',
  `min_score` int NOT NULL DEFAULT '0' COMMENT '所需最低评分（乘100存储，如 450=4.50分）',
  `commission_rate` int NOT NULL DEFAULT '20' COMMENT '平台抽成比例（百分比，如 20 表示 20%）',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='陪玩师等级表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_coach_level`
--

LOCK TABLES `play_coach_level` WRITE;
/*!40000 ALTER TABLE `play_coach_level` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_coach_level` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_coupon`
--

DROP TABLE IF EXISTS `play_coupon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_coupon` (
  `id` bigint unsigned NOT NULL COMMENT '优惠券ID（Snowflake）',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '优惠券名称',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '优惠券类型:1=满减券,2=折扣券,3=无门槛券',
  `is_new_member` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否新人专享:0=否,1=是',
  `face_value` bigint NOT NULL DEFAULT '0' COMMENT '面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）',
  `min_amount` bigint NOT NULL DEFAULT '0' COMMENT '最低消费金额（分，0表示无门槛）',
  `total_num` int NOT NULL DEFAULT '0' COMMENT '发放总量（0表示不限）',
  `used_num` int NOT NULL DEFAULT '0' COMMENT '已使用数量',
  `claim_num` int NOT NULL DEFAULT '0' COMMENT '已领取数量',
  `per_limit` int NOT NULL DEFAULT '1' COMMENT '每人限领张数',
  `valid_start_at` datetime NOT NULL COMMENT '有效期开始时间',
  `valid_end_at` datetime NOT NULL COMMENT '有效期结束时间',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_is_new_member` (`is_new_member`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='优惠券模板表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_coupon`
--

LOCK TABLES `play_coupon` WRITE;
/*!40000 ALTER TABLE `play_coupon` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_coupon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_coupon_member`
--

DROP TABLE IF EXISTS `play_coupon_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_coupon_member` (
  `id` bigint unsigned NOT NULL COMMENT '记录ID（Snowflake）',
  `coupon_id` bigint unsigned NOT NULL COMMENT '优惠券模板ID',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `order_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '使用的订单ID（0表示未使用）',
  `use_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '使用状态:0=未使用,1=已使用,2=已过期',
  `claim_at` datetime DEFAULT NULL COMMENT '领取时间',
  `use_at` datetime DEFAULT NULL COMMENT '使用时间',
  `expire_at` datetime DEFAULT NULL COMMENT '过期时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_coupon_id` (`coupon_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_use_status` (`use_status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员优惠券表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_coupon_member`
--

LOCK TABLES `play_coupon_member` WRITE;
/*!40000 ALTER TABLE `play_coupon_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_coupon_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_goods`
--

DROP TABLE IF EXISTS `play_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_goods` (
  `id` bigint unsigned NOT NULL COMMENT '商品ID（Snowflake）',
  `category_id` bigint unsigned NOT NULL COMMENT '分类ID',
  `coach_id` bigint unsigned NOT NULL COMMENT '陪玩师ID',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品名称',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品封面图',
  `desc_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '商品详情描述',
  `price` bigint NOT NULL DEFAULT '0' COMMENT '单价（分）',
  `unit` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '局' COMMENT '计量单位（如：局、小时、把）',
  `sales_num` int NOT NULL DEFAULT '0' COMMENT '销量',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=下架,1=上架',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_coach_id` (`coach_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_goods`
--

LOCK TABLES `play_goods` WRITE;
/*!40000 ALTER TABLE `play_goods` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_member`
--

DROP TABLE IF EXISTS `play_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_member` (
  `id` bigint unsigned NOT NULL COMMENT '会员ID（Snowflake）',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码（bcrypt 加密）',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别:0=未知,1=男,2=女',
  `member_level_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '会员等级ID',
  `exp` int NOT NULL DEFAULT '0' COMMENT '经验值',
  `balance` bigint NOT NULL DEFAULT '0' COMMENT '账户余额（分）',
  `is_coach` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否陪玩师:0=否,1=是',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=禁用,1=正常',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone` (`phone`),
  KEY `idx_member_level_id` (`member_level_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_member`
--

LOCK TABLES `play_member` WRITE;
/*!40000 ALTER TABLE `play_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_member_level`
--

DROP TABLE IF EXISTS `play_member_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_member_level` (
  `id` bigint unsigned NOT NULL COMMENT '等级ID（Snowflake）',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '等级名称',
  `level` tinyint NOT NULL DEFAULT '1' COMMENT '等级:1=普通会员,2=白银会员,3=黄金会员,4=铂金会员,5=钻石会员',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '等级图标',
  `min_exp` int NOT NULL DEFAULT '0' COMMENT '所需最低经验值',
  `discount` int NOT NULL DEFAULT '100' COMMENT '折扣（百分比，如 90 表示九折）',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员等级表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_member_level`
--

LOCK TABLES `play_member_level` WRITE;
/*!40000 ALTER TABLE `play_member_level` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_member_level` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_oauth`
--

DROP TABLE IF EXISTS `play_oauth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_oauth` (
  `id` bigint unsigned NOT NULL COMMENT '记录ID（Snowflake）',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `provider` tinyint(1) NOT NULL COMMENT '第三方平台:1=微信,2=支付宝',
  `open_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '第三方OpenID',
  `union_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方UnionID',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方昵称',
  `avatar` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方头像',
  `access_token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '访问令牌',
  `refresh_token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '刷新令牌',
  `expire_at` datetime DEFAULT NULL COMMENT '令牌过期时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_provider_open_id` (`provider`,`open_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='第三方登录绑定表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_oauth`
--

LOCK TABLES `play_oauth` WRITE;
/*!40000 ALTER TABLE `play_oauth` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_oauth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_order`
--

DROP TABLE IF EXISTS `play_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_order` (
  `id` bigint unsigned NOT NULL COMMENT '订单ID（Snowflake）',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '订单编号',
  `member_id` bigint unsigned NOT NULL COMMENT '下单会员ID',
  `coach_id` bigint unsigned NOT NULL COMMENT '陪玩师ID',
  `shop_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '店铺ID（0表示无店铺）',
  `goods_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `goods_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品名称（冗余）',
  `goods_price` bigint NOT NULL COMMENT '商品单价（分，下单时快照）',
  `quantity` int NOT NULL DEFAULT '1' COMMENT '数量',
  `total_amount` bigint NOT NULL DEFAULT '0' COMMENT '订单总额（分）',
  `discount_amount` bigint NOT NULL DEFAULT '0' COMMENT '会员折扣金额（分）',
  `coupon_amount` bigint NOT NULL DEFAULT '0' COMMENT '优惠券抵扣金额（分）',
  `pay_amount` bigint NOT NULL DEFAULT '0' COMMENT '实付金额（分）',
  `coupon_member_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '使用的优惠券领取记录ID',
  `pay_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '支付方式:0=未支付,1=微信支付,2=支付宝支付,3=余额支付',
  `order_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '订单状态:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款',
  `pay_at` datetime DEFAULT NULL COMMENT '支付时间',
  `start_at` datetime DEFAULT NULL COMMENT '服务开始时间',
  `finish_at` datetime DEFAULT NULL COMMENT '服务完成时间',
  `cancel_at` datetime DEFAULT NULL COMMENT '取消时间',
  `cancel_reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '取消原因',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '订单备注',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_coach_id` (`coach_id`),
  KEY `idx_shop_id` (`shop_id`),
  KEY `idx_order_status` (`order_status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_order`
--

LOCK TABLES `play_order` WRITE;
/*!40000 ALTER TABLE `play_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_payment`
--

DROP TABLE IF EXISTS `play_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_payment` (
  `id` bigint unsigned NOT NULL COMMENT '支付记录ID（Snowflake）',
  `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `payment_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付流水号（平台内部）',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方交易号',
  `pay_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '支付方式:1=微信支付,2=支付宝支付,3=余额支付',
  `pay_amount` bigint NOT NULL DEFAULT '0' COMMENT '支付金额（分）',
  `pay_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '支付状态:0=待支付,1=支付成功,2=支付失败,3=已退款',
  `pay_at` datetime DEFAULT NULL COMMENT '支付成功时间',
  `refund_at` datetime DEFAULT NULL COMMENT '退款时间',
  `refund_amount` bigint NOT NULL DEFAULT '0' COMMENT '退款金额（分）',
  `callback_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '回调报文',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_payment_no` (`payment_no`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_payment`
--

LOCK TABLES `play_payment` WRITE;
/*!40000 ALTER TABLE `play_payment` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_profit_log`
--

DROP TABLE IF EXISTS `play_profit_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_profit_log` (
  `id` bigint unsigned NOT NULL COMMENT '流水ID（Snowflake）',
  `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '订单编号',
  `pay_amount` bigint NOT NULL DEFAULT '0' COMMENT '实付金额（分）',
  `coach_id` bigint unsigned NOT NULL COMMENT '陪玩师ID',
  `shop_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '店铺ID（0表示无店铺）',
  `platform_rate` int NOT NULL DEFAULT '0' COMMENT '平台抽成比例（百分比）',
  `platform_amount` bigint NOT NULL DEFAULT '0' COMMENT '平台抽成金额（分）',
  `shop_rate` int NOT NULL DEFAULT '0' COMMENT '店铺抽成比例（百分比）',
  `shop_amount` bigint NOT NULL DEFAULT '0' COMMENT '店铺抽成金额（分）',
  `coach_amount` bigint NOT NULL DEFAULT '0' COMMENT '陪玩师收入（分）',
  `settle_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '结算状态:0=待结算,1=已结算',
  `settle_at` datetime DEFAULT NULL COMMENT '结算时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_coach_id` (`coach_id`),
  KEY `idx_shop_id` (`shop_id`),
  KEY `idx_settle_status` (`settle_status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='利润分成流水表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_profit_log`
--

LOCK TABLES `play_profit_log` WRITE;
/*!40000 ALTER TABLE `play_profit_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_profit_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_recharge_order`
--

DROP TABLE IF EXISTS `play_recharge_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_recharge_order` (
  `id` bigint unsigned NOT NULL COMMENT '充值订单ID（Snowflake）',
  `order_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '充值订单号',
  `member_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `recharge_plan_id` bigint unsigned NOT NULL COMMENT '充值方案ID',
  `amount` bigint NOT NULL COMMENT '充值金额（分）',
  `gift_amount` bigint NOT NULL DEFAULT '0' COMMENT '赠送金额（分）',
  `pay_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '支付方式:1=微信支付,2=支付宝支付',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方交易号',
  `pay_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '支付状态:0=待支付,1=支付成功,2=支付失败',
  `pay_at` datetime DEFAULT NULL COMMENT '支付时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='充值订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_recharge_order`
--

LOCK TABLES `play_recharge_order` WRITE;
/*!40000 ALTER TABLE `play_recharge_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_recharge_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_recharge_plan`
--

DROP TABLE IF EXISTS `play_recharge_plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_recharge_plan` (
  `id` bigint unsigned NOT NULL COMMENT '方案ID（Snowflake）',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '方案名称',
  `amount` bigint NOT NULL COMMENT '充值金额（分）',
  `gift_amount` bigint NOT NULL DEFAULT '0' COMMENT '赠送金额（分）',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '方案封面图',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='充值方案表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_recharge_plan`
--

LOCK TABLES `play_recharge_plan` WRITE;
/*!40000 ALTER TABLE `play_recharge_plan` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_recharge_plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_review`
--

DROP TABLE IF EXISTS `play_review`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_review` (
  `id` bigint unsigned NOT NULL COMMENT '评价ID（Snowflake）',
  `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
  `member_id` bigint unsigned NOT NULL COMMENT '评价会员ID',
  `coach_id` bigint unsigned NOT NULL COMMENT '被评陪玩师ID',
  `score` int NOT NULL DEFAULT '500' COMMENT '评分（乘100，如 500=5.00分）',
  `review_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '评价内容',
  `review_image` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评价图片（多张逗号分隔）',
  `reply_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '陪玩师回复内容',
  `reply_at` datetime DEFAULT NULL COMMENT '回复时间',
  `is_anonymous` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否匿名:0=否,1=是',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=隐藏,1=显示',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_id` (`order_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_coach_id` (`coach_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评价表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_review`
--

LOCK TABLES `play_review` WRITE;
/*!40000 ALTER TABLE `play_review` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_review` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `play_shop`
--

DROP TABLE IF EXISTS `play_shop`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `play_shop` (
  `id` bigint unsigned NOT NULL COMMENT '店铺ID（Snowflake）',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺名称',
  `logo_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '店铺LOGO',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '封面图',
  `contact_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系人姓名',
  `contact_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系电话',
  `intro` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '店铺简介',
  `commission_rate` int NOT NULL DEFAULT '10' COMMENT '店铺抽成比例（百分比，如 10 表示 10%）',
  `coach_num` int NOT NULL DEFAULT '0' COMMENT '陪玩师数量',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='店铺表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `play_shop`
--

LOCK TABLES `play_shop` WRITE;
/*!40000 ALTER TABLE `play_shop` DISABLE KEYS */;
/*!40000 ALTER TABLE `play_shop` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_dept`
--

DROP TABLE IF EXISTS `system_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_dept` (
  `id` bigint unsigned NOT NULL COMMENT '部门ID（Snowflake）',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上级部门ID，0 表示顶级部门',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门名称',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门负责人姓名',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人邮箱',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间，非 NULL 表示已删除',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_dept`
--

LOCK TABLES `system_dept` WRITE;
/*!40000 ALTER TABLE `system_dept` DISABLE KEYS */;
INSERT INTO `system_dept` VALUES (1000000000000000001,0,'总公司','admin','admin@example.com',0,1,0,0,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL);
/*!40000 ALTER TABLE `system_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_menu`
--

DROP TABLE IF EXISTS `system_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_menu` (
  `id` bigint unsigned NOT NULL COMMENT '菜单ID（Snowflake）',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单ID，0 表示顶级菜单',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型:1=目录,2=菜单,3=按钮,4=外链,5=内链',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '前端路由路径',
  `component` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '前端组件路径',
  `permission` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '权限标识（如 system:dept:list）',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单图标（图标名称）',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示:0=隐藏,1=显示',
  `is_cache` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否缓存:0=不缓存,1=缓存',
  `link_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '外链/内链地址（type=4或5时有效）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间，非 NULL 表示已删除',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_menu`
--

LOCK TABLES `system_menu` WRITE;
/*!40000 ALTER TABLE `system_menu` DISABLE KEYS */;
INSERT INTO `system_menu` VALUES (314253730209861632,0,'上传管理',1,'/upload',NULL,'','CloudUploadOutlined',50,1,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730235027456,314253730209861632,'上传配置',2,'/upload/config','upload/config/index','upload:config:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730251804672,314253730235027456,'上传配置新增',3,NULL,NULL,'upload:config:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730268581888,314253730235027456,'上传配置修改',3,NULL,NULL,'upload:config:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730285359104,314253730235027456,'上传配置删除',3,NULL,NULL,'upload:config:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730344079360,314253730209861632,'文件目录',2,'/upload/dir','upload/dir/index','upload:dir:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730365050880,314253730344079360,'文件目录新增',3,NULL,NULL,'upload:dir:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730386022400,314253730344079360,'文件目录修改',3,NULL,NULL,'upload:dir:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730415382528,314253730344079360,'文件目录删除',3,NULL,NULL,'upload:dir:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730461519872,314253730209861632,'文件目录规则',2,'/upload/dir-rule','upload/dir_rule/index','upload:dir_rule:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730478297088,314253730461519872,'文件目录规则新增',3,NULL,NULL,'upload:dir_rule:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730490880000,314253730461519872,'文件目录规则修改',3,NULL,NULL,'upload:dir_rule:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730503462912,314253730461519872,'文件目录规则删除',3,NULL,NULL,'upload:dir_rule:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730566377472,314253730209861632,'文件记录',2,'/upload/file','upload/file/index','upload:file:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730583154688,314253730566377472,'文件记录新增',3,NULL,NULL,'upload:file:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730595737600,314253730566377472,'文件记录修改',3,NULL,NULL,'upload:file:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253730620903424,314253730566377472,'文件记录删除',3,NULL,NULL,'upload:file:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:27','2026-03-31 10:25:27',NULL),(314253751944744960,0,'陪玩管理',1,'/play',NULL,'','game-icons:joystick',50,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253751965716480,314253751944744960,'活动',2,'/play/activity','play/activity/index','play:activity:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253751986688000,314253751965716480,'活动新增',3,NULL,NULL,'play:activity:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752007659520,314253751965716480,'活动修改',3,NULL,NULL,'play:activity:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752024436736,314253751965716480,'活动删除',3,NULL,NULL,'play:activity:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752108322816,314253751944744960,'活动参与记录',2,'/play/activity-join','play/activity_join/index','play:activity_join:list','',0,0,0,'',1,0,0,'2026-03-31 10:25:33','2026-03-31 06:55:37',NULL),(314253752120905728,314253752108322816,'活动参与记录新增',3,NULL,NULL,'play:activity_join:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752141877248,314253752108322816,'活动参与记录修改',3,NULL,NULL,'play:activity_join:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752154460160,314253752108322816,'活动参与记录删除',3,NULL,NULL,'play:activity_join:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752196403200,314253751944744960,'活动奖励',2,'/play/activity-reward','play/activity_reward/index','play:activity_reward:list','',0,0,0,'',1,0,0,'2026-03-31 10:25:33','2026-03-31 06:55:52',NULL),(314253752208986112,314253752196403200,'活动奖励新增',3,NULL,NULL,'play:activity_reward:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752225763328,314253752196403200,'活动奖励修改',3,NULL,NULL,'play:activity_reward:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752238346240,314253752196403200,'活动奖励删除',3,NULL,NULL,'play:activity_reward:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752284483584,314253751944744960,'活动步骤',2,'/play/activity-step','play/activity_step/index','play:activity_step:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752297066496,314253752284483584,'活动步骤新增',3,NULL,NULL,'play:activity_step:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752318038016,314253752284483584,'活动步骤修改',3,NULL,NULL,'play:activity_step:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752339009536,314253752284483584,'活动步骤删除',3,NULL,NULL,'play:activity_step:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752435478528,314253751944744960,'余额流水',2,'/play/balance-log','play/balance_log/index','play:balance_log:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752460644352,314253752435478528,'余额流水新增',3,NULL,NULL,'play:balance_log:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752477421568,314253752435478528,'余额流水修改',3,NULL,NULL,'play:balance_log:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752490004480,314253752435478528,'余额流水删除',3,NULL,NULL,'play:balance_log:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752536141824,314253751944744960,'商品分类',2,'/play/category','play/category/index','play:category:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752557113344,314253752536141824,'商品分类新增',3,NULL,NULL,'play:category:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752569696256,314253752536141824,'商品分类修改',3,NULL,NULL,'play:category:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752586473472,314253752536141824,'商品分类删除',3,NULL,NULL,'play:category:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752661970944,314253751944744960,'陪玩师',2,'/play/coach','play/coach/index','play:coach:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752682942464,314253752661970944,'陪玩师新增',3,NULL,NULL,'play:coach:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752699719680,314253752661970944,'陪玩师修改',3,NULL,NULL,'play:coach:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752716496896,314253752661970944,'陪玩师删除',3,NULL,NULL,'play:coach:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752783605760,314253751944744960,'陪玩师申请',2,'/play/coach-apply','play/coach_apply/index','play:coach_apply:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752800382976,314253752783605760,'陪玩师申请新增',3,NULL,NULL,'play:coach_apply:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752821354496,314253752783605760,'陪玩师申请修改',3,NULL,NULL,'play:coach_apply:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752842326016,314253752783605760,'陪玩师申请删除',3,NULL,NULL,'play:coach_apply:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752884269056,314253751944744960,'陪玩师等级',2,'/play/coach-level','play/coach_level/index','play:coach_level:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752901046272,314253752884269056,'陪玩师等级新增',3,NULL,NULL,'play:coach_level:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752922017792,314253752884269056,'陪玩师等级修改',3,NULL,NULL,'play:coach_level:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752951377920,314253752884269056,'陪玩师等级删除',3,NULL,NULL,'play:coach_level:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253752989126656,314253751944744960,'优惠券模板',2,'/play/coupon','play/coupon/index','play:coupon:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753001709568,314253752989126656,'优惠券模板新增',3,NULL,NULL,'play:coupon:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753018486784,314253752989126656,'优惠券模板修改',3,NULL,NULL,'play:coupon:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753035264000,314253752989126656,'优惠券模板删除',3,NULL,NULL,'play:coupon:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753123344384,314253751944744960,'会员优惠券',2,'/play/coupon-member','play/coupon_member/index','play:coupon_member:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753140121600,314253753123344384,'会员优惠券新增',3,NULL,NULL,'play:coupon_member:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753152704512,314253753123344384,'会员优惠券修改',3,NULL,NULL,'play:coupon_member:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753173676032,314253753123344384,'会员优惠券删除',3,NULL,NULL,'play:coupon_member:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753244979200,314253751944744960,'商品',2,'/play/goods','play/goods/index','play:goods:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753261756416,314253753244979200,'商品新增',3,NULL,NULL,'play:goods:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753282727936,314253753244979200,'商品修改',3,NULL,NULL,'play:goods:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753299505152,314253753244979200,'商品删除',3,NULL,NULL,'play:goods:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753349836800,314253751944744960,'会员',2,'/play/member','play/member/index','play:member:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753362419712,314253753349836800,'会员新增',3,NULL,NULL,'play:member:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753383391232,314253753349836800,'会员修改',3,NULL,NULL,'play:member:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753400168448,314253753349836800,'会员删除',3,NULL,NULL,'play:member:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753437917184,314253751944744960,'会员等级',2,'/play/member-level','play/member_level/index','play:member_level:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753463083008,314253753437917184,'会员等级新增',3,NULL,NULL,'play:member_level:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753479860224,314253753437917184,'会员等级修改',3,NULL,NULL,'play:member_level:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753496637440,314253753437917184,'会员等级删除',3,NULL,NULL,'play:member_level:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753609883648,314253751944744960,'第三方登录绑定',2,'/play/oauth','play/oauth/index','play:oauth:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753626660864,314253753609883648,'第三方登录绑定新增',3,NULL,NULL,'play:oauth:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753651826688,314253753609883648,'第三方登录绑定修改',3,NULL,NULL,'play:oauth:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753668603904,314253753609883648,'第三方登录绑定删除',3,NULL,NULL,'play:oauth:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753798627328,314253751944744960,'订单',2,'/play/order','play/order/index','play:order:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753815404544,314253753798627328,'订单新增',3,NULL,NULL,'play:order:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753832181760,314253753798627328,'订单修改',3,NULL,NULL,'play:order:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753848958976,314253753798627328,'订单删除',3,NULL,NULL,'play:order:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753945427968,314253751944744960,'支付记录',2,'/play/payment','play/payment/index','play:payment:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753966399488,314253753945427968,'支付记录新增',3,NULL,NULL,'play:payment:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753983176704,314253753945427968,'支付记录修改',3,NULL,NULL,'play:payment:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253753999953920,314253753945427968,'支付记录删除',3,NULL,NULL,'play:payment:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754096422912,314253751944744960,'利润分成流水',2,'/play/profit-log','play/profit_log/index','play:profit_log:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754109005824,314253754096422912,'利润分成流水新增',3,NULL,NULL,'play:profit_log:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754125783040,314253754096422912,'利润分成流水修改',3,NULL,NULL,'play:profit_log:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754146754560,314253754096422912,'利润分成流水删除',3,NULL,NULL,'play:profit_log:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754222252032,314253751944744960,'充值订单',2,'/play/recharge-order','play/recharge_order/index','play:recharge_order:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754243223552,314253754222252032,'充值订单新增',3,NULL,NULL,'play:recharge_order:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754255806464,314253754222252032,'充值订单修改',3,NULL,NULL,'play:recharge_order:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754272583680,314253754222252032,'充值订单删除',3,NULL,NULL,'play:recharge_order:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754318721024,314253751944744960,'充值方案',2,'/play/recharge-plan','play/recharge_plan/index','play:recharge_plan:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754331303936,314253754318721024,'充值方案新增',3,NULL,NULL,'play:recharge_plan:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754343886848,314253754318721024,'充值方案修改',3,NULL,NULL,'play:recharge_plan:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754356469760,314253754318721024,'充值方案删除',3,NULL,NULL,'play:recharge_plan:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754465521664,314253751944744960,'评价',2,'/play/review','play/review/index','play:review:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754478104576,314253754465521664,'评价新增',3,NULL,NULL,'play:review:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754494881792,314253754465521664,'评价修改',3,NULL,NULL,'play:review:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754507464704,314253754465521664,'评价删除',3,NULL,NULL,'play:review:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754557796352,314253751944744960,'店铺',2,'/play/shop','play/shop/index','play:shop:list','',0,1,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754570379264,314253754557796352,'店铺新增',3,NULL,NULL,'play:shop:create','',1,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754591350784,314253754557796352,'店铺修改',3,NULL,NULL,'play:shop:update','',2,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314253754603933696,314253754557796352,'店铺删除',3,NULL,NULL,'play:shop:delete','',3,0,0,NULL,1,0,0,'2026-03-31 10:25:33','2026-03-31 10:25:33',NULL),(314616957221474304,314253751944744960,'活动步骤提交记录',2,'/play/activity-step-log','play/activity_step_log/index','play:activity_step_log:list','',0,1,0,NULL,1,0,0,'2026-04-01 10:28:47','2026-04-01 10:28:47',NULL),(314616957238251520,314616957221474304,'活动步骤提交记录新增',3,NULL,NULL,'play:activity_step_log:create','',1,0,0,NULL,1,0,0,'2026-04-01 10:28:47','2026-04-01 10:28:47',NULL),(314616957271805952,314616957221474304,'活动步骤提交记录修改',3,NULL,NULL,'play:activity_step_log:update','',2,0,0,NULL,1,0,0,'2026-04-01 10:28:47','2026-04-01 10:28:47',NULL),(314616957284388864,314616957221474304,'活动步骤提交记录删除',3,NULL,NULL,'play:activity_step_log:delete','',3,0,0,NULL,1,0,0,'2026-04-01 10:28:47','2026-04-01 10:28:47',NULL),(1000000000000000010,0,'系统管理',1,'/system',NULL,'','SettingOutlined',100,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000011,1000000000000000010,'部门管理',2,'/system/dept','system/dept/index','system:dept:list','ApartmentOutlined',1,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000012,1000000000000000010,'角色管理',2,'/system/role','system/role/index','system:role:list','TeamOutlined',2,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000013,1000000000000000010,'菜单管理',2,'/system/menu','system/menu/index','system:menu:list','MenuOutlined',3,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000014,1000000000000000010,'用户管理',2,'/system/users','system/users/index','system:user:list','UserOutlined',4,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000021,1000000000000000011,'部门新增',3,NULL,NULL,'system:dept:create','',1,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000022,1000000000000000011,'部门修改',3,NULL,NULL,'system:dept:update','',2,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000023,1000000000000000011,'部门删除',3,NULL,NULL,'system:dept:delete','',3,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000031,1000000000000000012,'角色新增',3,NULL,NULL,'system:role:create','',1,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000032,1000000000000000012,'角色修改',3,NULL,NULL,'system:role:update','',2,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000033,1000000000000000012,'角色删除',3,NULL,NULL,'system:role:delete','',3,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000034,1000000000000000012,'资源授权',3,NULL,NULL,'system:role:grant:menu','',4,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000035,1000000000000000012,'数据授权',3,NULL,NULL,'system:role:grant:dept','',5,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000041,1000000000000000013,'菜单新增',3,NULL,NULL,'system:menu:create','',1,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000042,1000000000000000013,'菜单修改',3,NULL,NULL,'system:menu:update','',2,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000043,1000000000000000013,'菜单删除',3,NULL,NULL,'system:menu:delete','',3,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000051,1000000000000000014,'用户新增',3,NULL,NULL,'system:user:create','',1,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000052,1000000000000000014,'用户修改',3,NULL,NULL,'system:user:update','',2,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000053,1000000000000000014,'用户删除',3,NULL,NULL,'system:user:delete','',3,0,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000060,0,'仪表盘',1,'/dashboard',NULL,'','DashboardOutlined',0,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000061,1000000000000000060,'分析页',2,'/analytics','dashboard/analytics/index','','AreaChartOutlined',1,1,1,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL),(1000000000000000062,1000000000000000060,'工作台',2,'/workspace','dashboard/workspace/index','','DesktopOutlined',2,1,0,NULL,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL);
/*!40000 ALTER TABLE `system_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role`
--

DROP TABLE IF EXISTS `system_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_role` (
  `id` bigint unsigned NOT NULL COMMENT '角色ID（Snowflake）',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上级角色ID，0 表示顶级角色',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `data_scope` tinyint NOT NULL DEFAULT '1' COMMENT '数据范围:1=全部,2=本部门及以下,3=本部门,4=仅本人,5=自定义',
  `is_admin` tinyint(1) NOT NULL DEFAULT '0',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序（升序）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间，非 NULL 表示已删除',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role`
--

LOCK TABLES `system_role` WRITE;
/*!40000 ALTER TABLE `system_role` DISABLE KEYS */;
INSERT INTO `system_role` VALUES (1000000000000000002,0,'超级管理员',1,1,0,1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL);
/*!40000 ALTER TABLE `system_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role_dept`
--

DROP TABLE IF EXISTS `system_role_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_role_dept` (
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `dept_id` bigint unsigned NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`),
  KEY `idx_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色自定义数据权限部门关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role_dept`
--

LOCK TABLES `system_role_dept` WRITE;
/*!40000 ALTER TABLE `system_role_dept` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_role_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role_menu`
--

DROP TABLE IF EXISTS `system_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_role_menu` (
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单权限关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role_menu`
--

LOCK TABLES `system_role_menu` WRITE;
/*!40000 ALTER TABLE `system_role_menu` DISABLE KEYS */;
INSERT INTO `system_role_menu` VALUES (1000000000000000002,314253730209861632),(1000000000000000002,314253730235027456),(1000000000000000002,314253730251804672),(1000000000000000002,314253730268581888),(1000000000000000002,314253730285359104),(1000000000000000002,314253730344079360),(1000000000000000002,314253730365050880),(1000000000000000002,314253730386022400),(1000000000000000002,314253730415382528),(1000000000000000002,314253730461519872),(1000000000000000002,314253730478297088),(1000000000000000002,314253730490880000),(1000000000000000002,314253730503462912),(1000000000000000002,314253730566377472),(1000000000000000002,314253730583154688),(1000000000000000002,314253730595737600),(1000000000000000002,314253730620903424),(1000000000000000002,314253751944744960),(1000000000000000002,314253751965716480),(1000000000000000002,314253751986688000),(1000000000000000002,314253752007659520),(1000000000000000002,314253752024436736),(1000000000000000002,314253752108322816),(1000000000000000002,314253752120905728),(1000000000000000002,314253752141877248),(1000000000000000002,314253752154460160),(1000000000000000002,314253752196403200),(1000000000000000002,314253752208986112),(1000000000000000002,314253752225763328),(1000000000000000002,314253752238346240),(1000000000000000002,314253752284483584),(1000000000000000002,314253752297066496),(1000000000000000002,314253752318038016),(1000000000000000002,314253752339009536),(1000000000000000002,314253752435478528),(1000000000000000002,314253752460644352),(1000000000000000002,314253752477421568),(1000000000000000002,314253752490004480),(1000000000000000002,314253752536141824),(1000000000000000002,314253752557113344),(1000000000000000002,314253752569696256),(1000000000000000002,314253752586473472),(1000000000000000002,314253752661970944),(1000000000000000002,314253752682942464),(1000000000000000002,314253752699719680),(1000000000000000002,314253752716496896),(1000000000000000002,314253752783605760),(1000000000000000002,314253752800382976),(1000000000000000002,314253752821354496),(1000000000000000002,314253752842326016),(1000000000000000002,314253752884269056),(1000000000000000002,314253752901046272),(1000000000000000002,314253752922017792),(1000000000000000002,314253752951377920),(1000000000000000002,314253752989126656),(1000000000000000002,314253753001709568),(1000000000000000002,314253753018486784),(1000000000000000002,314253753035264000),(1000000000000000002,314253753123344384),(1000000000000000002,314253753140121600),(1000000000000000002,314253753152704512),(1000000000000000002,314253753173676032),(1000000000000000002,314253753244979200),(1000000000000000002,314253753261756416),(1000000000000000002,314253753282727936),(1000000000000000002,314253753299505152),(1000000000000000002,314253753349836800),(1000000000000000002,314253753362419712),(1000000000000000002,314253753383391232),(1000000000000000002,314253753400168448),(1000000000000000002,314253753437917184),(1000000000000000002,314253753463083008),(1000000000000000002,314253753479860224),(1000000000000000002,314253753496637440),(1000000000000000002,314253753609883648),(1000000000000000002,314253753626660864),(1000000000000000002,314253753651826688),(1000000000000000002,314253753668603904),(1000000000000000002,314253753798627328),(1000000000000000002,314253753815404544),(1000000000000000002,314253753832181760),(1000000000000000002,314253753848958976),(1000000000000000002,314253753945427968),(1000000000000000002,314253753966399488),(1000000000000000002,314253753983176704),(1000000000000000002,314253753999953920),(1000000000000000002,314253754096422912),(1000000000000000002,314253754109005824),(1000000000000000002,314253754125783040),(1000000000000000002,314253754146754560),(1000000000000000002,314253754222252032),(1000000000000000002,314253754243223552),(1000000000000000002,314253754255806464),(1000000000000000002,314253754272583680),(1000000000000000002,314253754318721024),(1000000000000000002,314253754331303936),(1000000000000000002,314253754343886848),(1000000000000000002,314253754356469760),(1000000000000000002,314253754465521664),(1000000000000000002,314253754478104576),(1000000000000000002,314253754494881792),(1000000000000000002,314253754507464704),(1000000000000000002,314253754557796352),(1000000000000000002,314253754570379264),(1000000000000000002,314253754591350784),(1000000000000000002,314253754603933696),(1000000000000000002,1000000000000000010),(1000000000000000002,1000000000000000011),(1000000000000000002,1000000000000000012),(1000000000000000002,1000000000000000013),(1000000000000000002,1000000000000000014),(1000000000000000002,1000000000000000021),(1000000000000000002,1000000000000000022),(1000000000000000002,1000000000000000023),(1000000000000000002,1000000000000000031),(1000000000000000002,1000000000000000032),(1000000000000000002,1000000000000000033),(1000000000000000002,1000000000000000034),(1000000000000000002,1000000000000000035),(1000000000000000002,1000000000000000041),(1000000000000000002,1000000000000000042),(1000000000000000002,1000000000000000043),(1000000000000000002,1000000000000000051),(1000000000000000002,1000000000000000052),(1000000000000000002,1000000000000000053),(1000000000000000002,1000000000000000060),(1000000000000000002,1000000000000000061),(1000000000000000002,1000000000000000062);
/*!40000 ALTER TABLE `system_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_user_dept`
--

DROP TABLE IF EXISTS `system_user_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_user_dept` (
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `dept_id` bigint unsigned NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`user_id`,`dept_id`),
  KEY `idx_dept_id` (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户部门关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_user_dept`
--

LOCK TABLES `system_user_dept` WRITE;
/*!40000 ALTER TABLE `system_user_dept` DISABLE KEYS */;
INSERT INTO `system_user_dept` VALUES (1000000000000000003,1000000000000000001);
/*!40000 ALTER TABLE `system_user_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_user_role`
--

DROP TABLE IF EXISTS `system_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_user_role` (
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_user_role`
--

LOCK TABLES `system_user_role` WRITE;
/*!40000 ALTER TABLE `system_user_role` DISABLE KEYS */;
INSERT INTO `system_user_role` VALUES (1000000000000000003,1000000000000000002);
/*!40000 ALTER TABLE `system_user_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_users`
--

DROP TABLE IF EXISTS `system_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_users` (
  `id` bigint unsigned NOT NULL COMMENT '用户ID（Snowflake）',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码（SHA-256 加密）',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称/显示名',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱地址',
  `avatar` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像图片 URL',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=关闭,1=开启',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '所属部门ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间，非 NULL 表示已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_dept_id` (`dept_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_users`
--

LOCK TABLES `system_users` WRITE;
/*!40000 ALTER TABLE `system_users` DISABLE KEYS */;
INSERT INTO `system_users` VALUES (1000000000000000003,'admin','240be518fabd2724ddb6f04eeb1da5967448d7e831c08c8fa822809f74c720a9','超级管理员','admin@example.com','',1,0,1000000000000000001,'2026-03-30 21:20:22','2026-03-30 21:20:22',NULL);
/*!40000 ALTER TABLE `system_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `upload_config`
--

DROP TABLE IF EXISTS `upload_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `upload_config` (
  `id` bigint unsigned NOT NULL COMMENT 'ID',
  `name` varchar(100) NOT NULL COMMENT '配置名称',
  `storage` tinyint(1) NOT NULL DEFAULT '1' COMMENT '存储类型:1=本地,2=阿里云OSS,3=腾讯云COS',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认:0=否,1=是',
  `local_path` varchar(500) DEFAULT '' COMMENT '本地存储路径',
  `oss_endpoint` varchar(255) DEFAULT '' COMMENT 'OSS Endpoint',
  `oss_bucket` varchar(255) DEFAULT '' COMMENT 'OSS Bucket',
  `oss_access_key` varchar(255) DEFAULT '' COMMENT 'OSS AccessKey',
  `oss_secret_key` varchar(255) DEFAULT '' COMMENT 'OSS SecretKey',
  `cos_region` varchar(100) DEFAULT '' COMMENT 'COS Region',
  `cos_bucket` varchar(255) DEFAULT '' COMMENT 'COS Bucket',
  `cos_secret_id` varchar(255) DEFAULT '' COMMENT 'COS SecretId',
  `cos_secret_key` varchar(255) DEFAULT '' COMMENT 'COS SecretKey',
  `max_size` int DEFAULT '10' COMMENT '最大文件大小(MB)',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=禁用,1=启用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '部门ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='上传配置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `upload_config`
--

LOCK TABLES `upload_config` WRITE;
/*!40000 ALTER TABLE `upload_config` DISABLE KEYS */;
INSERT INTO `upload_config` VALUES (314309590294466560,'本地存储',1,1,'uploads','','','','','','','','',2,1,'2026-03-31 06:07:26','2026-03-31 06:07:26',NULL,NULL,NULL);
/*!40000 ALTER TABLE `upload_config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `upload_dir`
--

DROP TABLE IF EXISTS `upload_dir`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `upload_dir` (
  `id` bigint unsigned NOT NULL COMMENT 'ID',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '上级目录',
  `name` varchar(100) NOT NULL COMMENT '目录名称',
  `path` varchar(500) NOT NULL COMMENT '目录路径',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=禁用,1=启用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '部门ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件目录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `upload_dir`
--

LOCK TABLES `upload_dir` WRITE;
/*!40000 ALTER TABLE `upload_dir` DISABLE KEYS */;
/*!40000 ALTER TABLE `upload_dir` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `upload_dir_rule`
--

DROP TABLE IF EXISTS `upload_dir_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `upload_dir_rule` (
  `id` bigint unsigned NOT NULL COMMENT 'ID',
  `dir_id` bigint unsigned NOT NULL COMMENT '目录ID',
  `category` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类别:1=默认,2=类型,3=接口',
  `save_path` varchar(500) DEFAULT '' COMMENT '保存目录',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:0=禁用,1=启用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '部门ID',
  PRIMARY KEY (`id`),
  KEY `idx_dir_id` (`dir_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件目录规则';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `upload_dir_rule`
--

LOCK TABLES `upload_dir_rule` WRITE;
/*!40000 ALTER TABLE `upload_dir_rule` DISABLE KEYS */;
/*!40000 ALTER TABLE `upload_dir_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `upload_file`
--

DROP TABLE IF EXISTS `upload_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `upload_file` (
  `id` bigint unsigned NOT NULL COMMENT 'ID',
  `dir_id` bigint unsigned DEFAULT '0' COMMENT '所属目录',
  `name` varchar(255) NOT NULL COMMENT '文件名称',
  `url` varchar(500) NOT NULL COMMENT '文件地址',
  `ext` varchar(20) DEFAULT '' COMMENT '文件扩展名',
  `size` bigint unsigned DEFAULT '0' COMMENT '文件大小',
  `mime` varchar(100) DEFAULT '' COMMENT 'MIME类型',
  `storage` tinyint(1) NOT NULL DEFAULT '1' COMMENT '存储类型:1=本地,2=阿里云OSS,3=腾讯云COS',
  `is_image` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否图片:0=否,1=是',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '部门ID',
  PRIMARY KEY (`id`),
  KEY `idx_dir_id` (`dir_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `upload_file`
--

LOCK TABLES `upload_file` WRITE;
/*!40000 ALTER TABLE `upload_file` DISABLE KEYS */;
INSERT INTO `upload_file` VALUES (314308797315158016,0,'17749370568554486713.png','/upload/2026-03-31/17749370568554486713.png','png',93017,'image/png',1,1,'2026-03-31 06:04:17','2026-03-31 06:04:17',NULL,NULL,NULL),(314309094477402112,0,'17749371277032938579.png','/upload/2026-03-31/17749371277032938579.png','png',1793,'image/png',1,1,'2026-03-31 06:05:28','2026-03-31 06:05:28',NULL,NULL,NULL);
/*!40000 ALTER TABLE `upload_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'gbaseadmin'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-04-01 11:12:18
