CREATE DATABASE /*!32312 IF NOT EXISTS*/ `prompting` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE 'prompting';

--  用户表
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`(
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `nickname` varchar(30) NOT NULL,
    `email` varchar(256) NOT NULL,
    `phone` varchar(16),
    `introduction` varchar(255),
    `avatar` varchar(255),
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
)ENGINE = Innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4;

-- 实例表
DROP TABLE IF EXISTS `t_prompt`;
CREATE TABLE `t_prompt`(
    `id` int NOT NULL AUTO_INCREMENT COMMENT '实例id',
    `u_id` bigint(20) NOT NULL COMMENT '作者id',
--     `promptID` varchar(256) NOT NULL COMMENT ,
    `title` varchar(256) NOT NULL COMMENT '标题',
    `ask` text NOT NULL COMMENT '提问',
    `answer` text NOT NULL COMMENT '回答',
    `img` varchar(255) NOT NULL COMMENT '图片路径',
    `child` bigint(20) DEFAULT(-1) COMMENT '下一组对话id',
    `ifDelete` tinyint(1) DEFAULT(0) COMMENT '是否删除',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
--     UNIQUE KEY `promptID` (`promptID`),
)ENGINE = Innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4;

-- 评论表
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`(
    `id` int NOT NULL AUTO_INCREMENT COMMENT '评论id',
    `u_id` bigint(20) NOT NULL COMMENT '评论者id',
    `p_id` bigint(20) NOT NULL COMMENT '被评论的prompt的id',
    `content` text NOT NULL COMMENT '评论内容',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
)ENGINE = Innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4;

-- 收藏表
DROP TABLE IF EXISTS `t_collect`;
CREATE TABLE `t_collect`(
    `id` int NOT NULL AUTO_INCREMENT COMMENT '收藏id',
    `u_id` bigint(20) NOT NULL COMMENT '收藏者id',
    `p_id` bigint(20) NOT NULL COMMENT '被收藏的prompt的id',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
)ENGINE = Innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4;

-- 类型表
DROP TABLE IF EXISTS `t_type`;
CREATE TABLE `t_type`(
    `id` int NOT NULL AUTO_INCREMENT COMMENT '类型id',
    `name` varchar(255) NOT NULL COMMENT '类型名',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
)ENGINE = Innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4;

-- 类型实例关系表
DROP TABLE IF EXISTS `t_type_prompt`;
CREATE TABLE `t_type_prompt`(
    `t_id` int NOT NULL COMMENT '类型id',
    `p_id` int NOT NULL COMMENT '实例id',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
)ENGINE = Innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4;
