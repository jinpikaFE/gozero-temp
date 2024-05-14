CREATE TABLE `sys_role_user`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`role_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和用户关联表' ROW_FORMAT = COMPACT;