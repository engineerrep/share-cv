SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
BEGIN;

--
-- 数据库： `app_share_cv`
--

CREATE DATABASE IF NOT EXISTS app_share_cv DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
use app_share_cv;

-- app用户表
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user`
(
    `id`                    bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `uuid`                  varchar(128) NOT NULL DEFAULT '' COMMENT 'uuid',
    `apple_uid`             varchar(500) NOT NULL DEFAULT '' COMMENT '苹果授权ID',
    `google_uid`            varchar(500) NOT NULL DEFAULT '' COMMENT '谷歌授权ID',
    `email`                 varchar(100) NOT NULL DEFAULT '' COMMENT '邮件',
    `mobile`                char(20)     NOT NULL DEFAULT '' COMMENT '手机号',
    `mobile_country_id`     bigint       NOT NULL DEFAULT 0  COMMENT '手机国家ID',
    `mobile_country_num`    char(20)     NOT NULL DEFAULT '' COMMENT '手机国家区号',
    `locale`                varchar(200) NOT NULL DEFAULT '' COMMENT '本地语言',
    `user_type`             tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '账户类型 1=邮箱 2=手机 3=谷歌 4=苹果',
    `password`              varchar(32)  NOT NULL COMMENT '密码',
    `invitation_code`       varchar(15)  NOT NULL DEFAULT '' COMMENT '邀请码',
    `create_at`             timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`             timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`            timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`),
    KEY               `idx_google_uid` (`google_uid`),
    KEY               `idx_email` (`email`),
    KEY               `idx_mobile` (`mobile`),
    KEY               `idx_mobile_country_id` (`mobile_country_id`),
    KEY               `idx_user_type` (`user_type`),
    KEY               `idx_invitation_code` (`invitation_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app用户表';

-- 用户更多信息表
DROP TABLE IF EXISTS `app_user_info`;
CREATE TABLE `app_user_info`
(
    `user_id`               bigint       NOT NULL COMMENT '用户id',
    `first_name`            varchar(200) NOT NULL DEFAULT '' COMMENT '姓',
    `last_name`             varchar(200) NOT NULL DEFAULT '' COMMENT '名',
    `nickname`              varchar(32)  NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`                varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
    `background_img`        varchar(200) NOT NULL DEFAULT '' COMMENT '背景图片',
    `gender`                tinyint(1) UNSIGNED NOT NULL DEFAULT 3 COMMENT '1=女 2=男 3=保密',
    `mode_type`             tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否公开 1=公开 2=隐藏',
    `height`                int          NOT NULL DEFAULT 0 COMMENT '身高',
    `weight`                int          NOT NULL DEFAULT 0 COMMENT '体重',
    `birthday`              varchar(20)  NOT NULL DEFAULT '' COMMENT '生日',
    `country_id`            bigint       NOT NULL DEFAULT 0  COMMENT '国家表id',
    `country_name`          varchar(200) NOT NULL DEFAULT '' COMMENT '国家名称',
    `identity`              tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户身份 1=求职者 2=学生 3=招聘方',
    `province`              varchar(200) NOT NULL DEFAULT '' COMMENT '省份',
    `city`                  varchar(200) NOT NULL DEFAULT '' COMMENT '城市',
    `address`               varchar(500) NOT NULL DEFAULT '' COMMENT '详细地点',
    `personal_strength`     varchar(2000)   NOT NULL  COMMENT '个人实力',
    `purpose_city`          varchar(2000)   NOT NULL  COMMENT '意向城市 多个用英语逗号分割',
    `interests`             varchar(2000)   NOT NULL  COMMENT '兴趣爱好 多个用英语逗号分割',
    `languages`             varchar(2000)   NOT NULL  COMMENT '会的语言 多个用英语逗号分割',
    `dating_targets`        varchar(2000)   NOT NULL  COMMENT '交友的目 多个用英语逗号分割',
    `personal_skills`       varchar(2000)   NOT NULL  COMMENT '个人技能 多个用英语逗号分割',
    `create_at`             timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`             timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`            timestamp                DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_id`),
    KEY               `idx_first_name` (`first_name`),
    KEY               `idx_last_name` (`last_name`),
    KEY               `idx_nickname` (`nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户更多信息表';

-- 用户浏览记录表
DROP TABLE IF EXISTS `app_user_views`;
CREATE TABLE `app_user_views`
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL COMMENT '登录用户id',
    `views_user_id`     bigint       NOT NULL COMMENT '被浏览用户id',
    `views_type`        tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '浏览类型 1=查看',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户浏览记录表';

-- 用户定位信息表
DROP TABLE IF EXISTS `app_user_location`;
CREATE TABLE `app_user_location`
(
    `id`                bigint          NOT NULL AUTO_INCREMENT,
    `user_id`           bigint          NOT NULL COMMENT 'app_user表id',
    `time_zone`         varchar(100)    NOT NULL DEFAULT '' COMMENT '时区',
    `country`           varchar(200)    NOT NULL DEFAULT '' COMMENT '国家',
    `country_code`      varchar(50)     NOT NULL DEFAULT '' COMMENT '国家代码',
    `province`          varchar(200)    NOT NULL DEFAULT '' COMMENT '省份',
    `city`              varchar(200)    NOT NULL DEFAULT '' COMMENT '城市',
    `state`             varchar(200)    NOT NULL DEFAULT '' COMMENT '县',
    `area`              varchar(200)    NOT NULL DEFAULT '' COMMENT '区域',
    `street`            varchar(500)    NOT NULL DEFAULT '' COMMENT '街道',
    `latitude`          decimal(20, 8)  NOT NULL DEFAULT 0 COMMENT '纬度',
    `longitude`         decimal(20, 8)  NOT NULL DEFAULT 0 COMMENT '经度',
    `postal_code`       varchar(50)     NOT NULL DEFAULT '' COMMENT '邮政编号',
    `address`           varchar(1000)   NOT NULL DEFAULT '' COMMENT '详细地点',
    `create_at`         timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp                DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户定位信息表';

-- 邀请管理表
DROP TABLE IF EXISTS `app_user_invitation`;
CREATE TABLE `app_user_invitation`
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL COMMENT '登录用户id(app_user表id)',
    `super_user_id`     bigint       NOT NULL COMMENT '上级用户id(app_user表id)',
    `invitation_code`   varchar(100) NOT NULL DEFAULT '' COMMENT '邀请码',
    `channel_name`      varchar(100) NOT NULL DEFAULT '' COMMENT '渠道名称',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`),
    KEY               `idx_super_user_id` (`super_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='邀请管理表';

-- 用户登录日志
DROP TABLE IF EXISTS `app_user_login_log`;
CREATE TABLE `app_user_login_log`
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL COMMENT '用户id(app_user表id)',
    `ip`                varchar(100) NOT NULL DEFAULT '' COMMENT 'ip',
    `device_id`         varchar(128) NOT NULL DEFAULT '' COMMENT '设备唯一标识',
    `app_version`       varchar(100) NOT NULL DEFAULT '' COMMENT 'app版本',
    `system_version`    varchar(100) NOT NULL DEFAULT '' COMMENT 'system版本',
    `platform`          int          NOT NULL DEFAULT '1' COMMENT '登录平台 1 pc 2 win 3 iOS 4 android 5 iPad 6 mac',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户登录日志';

-- 用户相册表
DROP TABLE IF EXISTS `app_user_photo`;
CREATE TABLE `app_user_photo`
(
    `id`                bigint       NOT NULL AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `title`             varchar(50)  NOT NULL DEFAULT '' COMMENT '图片标题',
    `file_url`          varchar(500) NOT NULL DEFAULT '' COMMENT '图片URL',
    `order_by`          int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户相册表';

-- 用户视频表
DROP TABLE IF EXISTS `app_user_video`;
CREATE TABLE `app_user_video`
(
    `id`                bigint          NOT NULL AUTO_INCREMENT,
    `user_id`           bigint          NOT NULL DEFAULT 0 COMMENT '用户id',
    `title`             varchar(50)     NOT NULL DEFAULT '' COMMENT '视频标题',
    `cover_url`         varchar(500)    NOT NULL DEFAULT '' COMMENT '视频封面URL',
    `file_url`          varchar(500)    NOT NULL DEFAULT '' COMMENT '视频URL',
    `play_time`         decimal(20, 10) NOT NULL DEFAULT 0 COMMENT '播放时长',
    `video_size`        decimal(20, 10) NOT NULL DEFAULT 0 COMMENT '视频大小kb',
    `order_by`          int             NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp                DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户视频表';

-- app自定义属性
DROP TABLE IF EXISTS `app_user_custom_attributes`;
CREATE TABLE `app_user_custom_attributes`
(
    `id`                bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `title`             varchar(100) NOT NULL DEFAULT '' COMMENT '自定义属性标题',
    `content`           varchar(1000) NOT NULL DEFAULT '' COMMENT '自定义属性内容',
    `summary`           varchar(200) NOT NULL DEFAULT '' COMMENT '自定义属性摘要',
    `icon`              varchar(500) NOT NULL DEFAULT '' COMMENT '自定义属性图标',
    `start_time`        varchar(20)  NOT NULL DEFAULT '' COMMENT '开始时间',
    `end_time`          varchar(20)  NOT NULL DEFAULT '' COMMENT '结束时间',
    `order_by`          int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app自定义属性';

-- app志愿者经验
DROP TABLE IF EXISTS `app_user_volunteer`;
CREATE TABLE `app_user_volunteer`
(
    `id`                bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `title`             varchar(100) NOT NULL DEFAULT '' COMMENT '志愿者标题',
    `content`           varchar(1000) NOT NULL DEFAULT '' COMMENT '志愿者内容',
    `start_time`        varchar(20)  NOT NULL DEFAULT '' COMMENT '开始时间',
    `end_time`          varchar(20)  NOT NULL DEFAULT '' COMMENT '结束时间',
    `order_by`          int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app志愿者经验';

-- app用户经验表
DROP TABLE IF EXISTS `app_user_experience`;
CREATE TABLE `app_user_experience`
(
    `id`                        bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`                   bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `experience_organization`   varchar(100) NOT NULL DEFAULT '' COMMENT '所在组织',
    `experience_position`       varchar(100) NOT NULL DEFAULT '' COMMENT '担任职位',
    `experience_title`          varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
    `experience_summary`        varchar(200) NOT NULL DEFAULT '' COMMENT '摘要',
    `experience_content`        varchar(1000) NOT NULL DEFAULT '' COMMENT '内容',
    `start_time`        varchar(20)  NOT NULL DEFAULT '' COMMENT '开始时间',
    `end_time`          varchar(20)  NOT NULL DEFAULT '' COMMENT '结束时间',
    `order_by`          int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app用户经验表';

-- app用户教育经历表
DROP TABLE IF EXISTS `app_user_education`;
CREATE TABLE `app_user_education`
(
    `id`                    bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`               bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `organization_name`     varchar(200) NOT NULL DEFAULT '' COMMENT '组织名称',
    `organization_country`  varchar(200) NOT NULL DEFAULT '' COMMENT '组织国家',
    `organization_city`     varchar(200) NOT NULL DEFAULT '' COMMENT '组织城市',
    `organization_level`    varchar(100) NOT NULL DEFAULT '' COMMENT '组织级别',
    `start_time`            varchar(20)  NOT NULL DEFAULT '' COMMENT '入学时间',
    `end_time`              varchar(20)  NOT NULL DEFAULT '' COMMENT '毕业时间',
    `study_specialty`       varchar(100) NOT NULL DEFAULT '' COMMENT '学习专业',
    `study_achievement`     varchar(2000)NOT NULL DEFAULT '' COMMENT '学习成就',
    `degree_education`      varchar(100) NOT NULL DEFAULT '' COMMENT '教育学位',
    `order_by`              int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`             timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`             timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`            timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app用户教育经历表';

-- app用户公司表
DROP TABLE IF EXISTS `app_user_company`;
CREATE TABLE `app_user_company`
(
    `id`                bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `company_name`      varchar(200) NOT NULL DEFAULT '' COMMENT '公司名称',
    `company_country`   varchar(200) NOT NULL DEFAULT '' COMMENT '公司国家',
    `company_city`      varchar(200) NOT NULL DEFAULT '' COMMENT '公司城市',
    `company_industry`  varchar(200) NOT NULL DEFAULT '' COMMENT '公司行业',
    `start_time`        varchar(20)  NOT NULL DEFAULT '' COMMENT '开始时间',
    `end_time`          varchar(20)  NOT NULL DEFAULT '' COMMENT '结束时间',
    `job_position`      varchar(100) NOT NULL DEFAULT '' COMMENT '工作职位',
    `job_content`       varchar(1000)NOT NULL DEFAULT '' COMMENT '工作内容',
    `job_summary`       varchar(1000)NOT NULL DEFAULT '' COMMENT '工作总结',
    `job_skills`        varchar(1000)NOT NULL DEFAULT '' COMMENT '使用技能 多个用逗号分隔',
    `order_by`          int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app用户公司表';

-- app用户项目表
DROP TABLE IF EXISTS `app_user_project`;
CREATE TABLE `app_user_project`
(
    `id`                bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`           bigint       NOT NULL DEFAULT 0 COMMENT '用户id',
    `project_name`      varchar(100) NOT NULL DEFAULT '' COMMENT '项目名称',
    `project_content`   varchar(1000)NOT NULL DEFAULT '' COMMENT '项目内容',
    `job_position`      varchar(100) NOT NULL DEFAULT '' COMMENT '工作职位',
    `job_content`       varchar(1000)NOT NULL DEFAULT '' COMMENT '工作内容',
    `start_time`        varchar(20)  NOT NULL DEFAULT '' COMMENT '启动时间',
    `end_time`          varchar(20)  NOT NULL DEFAULT '' COMMENT '结束时间',
    `project_link_platform`      varchar(50) NOT NULL DEFAULT '' COMMENT '项目链接平台',
    `project_link`      varchar(500) NOT NULL DEFAULT '' COMMENT '项目链接',
    `project_skills`    varchar(1000)NOT NULL DEFAULT '' COMMENT '使用技能 多个用逗号分隔',
    `order_by`          int          NOT NULL DEFAULT 0 COMMENT '排序从小到大',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app用户项目表';

-- 用户意见反馈表
DROP TABLE IF EXISTS `app_user_feedback`;
CREATE TABLE `app_user_feedback`
(
    `id`                bigint          NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`           bigint          NOT NULL COMMENT '用户id',
    `type`              varchar(200)    NOT NULL DEFAULT '' COMMENT '意见反馈类型',
    `content`           varchar(1000)   NOT NULL DEFAULT '' COMMENT '意见反馈内容',
    `file_urls`         varchar(2000)   NOT NULL DEFAULT '' COMMENT '图片地址 多张用逗号分割',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户意见反馈表';

-- 用户浏览统计
DROP TABLE IF EXISTS `app_user_views`;
CREATE TABLE `app_user_views`
(
    `id`                bigint          NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `user_id`           bigint          NOT NULL COMMENT '用户id',
    `view_user_id`      bigint          NOT NULL COMMENT '被用户id',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY               `idx_user_id` (`user_id`),
    KEY               `idx_view_user_id` (`view_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户浏览统计';

COMMIT;