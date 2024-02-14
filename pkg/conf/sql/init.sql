
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- BTREE的用途
--     匹配模糊查询;全值匹配的查询(where);

-- DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `avatar` varchar(255) NOT NULL COMMENT '头像',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- DROP TABLE IF EXISTS `video`;
CREATE TABLE video (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `uid` bigint NOT NULL COMMENT '用户ID',
    `play_url` varchar(255) NOT NULL COMMENT '视频url',
    `cover_url` varchar(255) NOT NULL COMMENT '封面url',
    `title` varchar(255) DEFAULT NULL COMMENT '标题',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '发布时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `time` (`created_at`) USING BTREE,
    KEY `author` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

-- DROP TABLE IF EXISTS `like`;
CREATE TABLE `like` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `uid` bigint NOT NULL COMMENT '点赞用户id',
    `videoId` bigint NOT NULL COMMENT '被点赞的视频id',
    `status` bigint NOT NULL DEFAULT 1 COMMENT '点赞：1 取消：0',#因为取消点赞不是真的删除行所以就多个status
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '点赞时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `favorite_user`
        foreign key (`uid`)
            references `user` (`id`)
            on delete cascade,
    constraint `favorite_video`
        foreign key (`videoId`)
            references video (`id`)
            on delete cascade,
    KEY `userIdtoVideoIdIdx` (`uid`,`videoId`) USING BTREE,
    KEY `uid_idx` (`uid`) USING BTREE,
    KEY `video_idx` (`videoId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';