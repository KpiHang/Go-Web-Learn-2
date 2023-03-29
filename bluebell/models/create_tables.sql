DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(64) COLLATE utf8mb4_general_ci,
                        `gender` tinyint(4) NOT NULL DEFAULT '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

# 这是一段MySQL语句，用于创建一个名为user的表。
# 表包含了id、user_id、username、password、email、gender、create_time和update_time八个字段，
# 其中id为主键，并自动增长。该表使用的引擎是InnoDB，字符集为utf8mb4，排序规则为utf8mb4_general_ci。
# 同时还添加了对于username和user_id两个字段的唯一索引，确保这两个字段的唯一性。此外，该表的结构也充分考虑到了用户信息的存储需求。