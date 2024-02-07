
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `avatar` varchar(255) NOT NULL COMMENT '头像',
    PRIMARY KEY (`id`),
    KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `uid` bigint NOT NULL COMMENT '用户ID',
    `play_url` varchar(255) NOT NULL COMMENT '视频url',
    `cover_url` varchar(255) NOT NULL COMMENT '封面url',
    `title` varchar(255) DEFAULT NULL COMMENT '标题',
    PRIMARY KEY (`id`),
--     KEY `time` (`publish_time`) USING BTREE,
    KEY `user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8 COMMENT='视频表';