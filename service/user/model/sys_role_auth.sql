CREATE TABLE `sys_role_auth`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `auth_id` bigint(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `auth_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和用户关联表' ROW_FORMAT = COMPACT;
