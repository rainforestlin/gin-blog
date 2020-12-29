
    -- `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    -- `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    -- `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    -- `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人'
    -- `deleted_on` INT (10) UNSIGNED DEFAULT '0' COMMENT '删除时间'
    -- `is_delete` TINYINT(1) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已经删除'

CREATE TABLE `blog_tag` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT '' COMMENT '标签名字',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT (10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_delete` TINYINT(1) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已经删除',
    `state` TINYINT(1) UNSIGNED DEFAULT '1' COMMENT '1为启用，0为禁用',
    PRIMARY KEY (`id`)