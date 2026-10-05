-- Seckill 初始化 schema
-- MySQL 容器首次启动时自动执行(通过 /docker-entrypoint-initdb.d/ 挂载)
-- 在 MYSQL_DATABASE=seckill 环境变量创建的数据库内执行

USE `seckill`;

-- -----------------------------------------------------------
-- 秒杀活动
-- -----------------------------------------------------------
CREATE TABLE IF NOT EXISTS `activity` (
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(128)    NOT NULL                COMMENT '活动名称',
    `start_time`  BIGINT          NOT NULL DEFAULT 0      COMMENT '开始时间(秒)',
    `end_time`    BIGINT          NOT NULL DEFAULT 0      COMMENT '结束时间(秒)',
    `status`      TINYINT         NOT NULL DEFAULT 1      COMMENT '1=筹备 2=上线 3=已结束',
    `create_time` BIGINT          NOT NULL DEFAULT 0,
    `update_time` BIGINT          NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY `idx_status_time` (`status`, `start_time`, `end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀活动';

-- -----------------------------------------------------------
-- 活动商品 (Activity 子项,一个活动包含多个商品)
-- -----------------------------------------------------------
CREATE TABLE IF NOT EXISTS `activity_goods` (
    `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `activity_id`    BIGINT UNSIGNED NOT NULL                COMMENT '活动ID',
    `goods_id`       VARCHAR(64)     NOT NULL                COMMENT '商品ID',
    `name`           VARCHAR(128)    NOT NULL DEFAULT ''     COMMENT '商品名',
    `price`          BIGINT          NOT NULL DEFAULT 0      COMMENT '原价(分)',
    `activity_price` BIGINT          NOT NULL DEFAULT 0      COMMENT '活动价(分)',
    `stock`          BIGINT          NOT NULL DEFAULT 0      COMMENT '秒杀库存',
    `limit_per_user` INT             NOT NULL DEFAULT 1      COMMENT '单用户限购',
    `create_time`    BIGINT          NOT NULL DEFAULT 0,
    `update_time`    BIGINT          NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_activity_goods` (`activity_id`, `goods_id`),
    KEY `idx_activity` (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动商品';

-- -----------------------------------------------------------
-- 秒杀订单
-- -----------------------------------------------------------
CREATE TABLE IF NOT EXISTS `seckill_order` (
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `uid`         VARCHAR(64)     NOT NULL                COMMENT '用户ID',
    `activity_id` BIGINT UNSIGNED NOT NULL                COMMENT '活动ID',
    `goods_id`    VARCHAR(64)     NOT NULL                COMMENT '商品ID',
    `price`       BIGINT          NOT NULL DEFAULT 0      COMMENT '下单价(分)',
    `status`      TINYINT         NOT NULL DEFAULT 1      COMMENT '1=待支付 2=已支付 3=已取消',
    `create_time` BIGINT          NOT NULL DEFAULT 0,
    `update_time` BIGINT          NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_activity_goods` (`uid`, `activity_id`, `goods_id`),
    KEY `idx_user` (`uid`),
    KEY `idx_activity` (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀订单';
