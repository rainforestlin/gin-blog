-- `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    -- `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    -- `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    -- `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人'
    -- `deleted_on` INT (10) UNSIGNED DEFAULT '0' COMMENT '删除时间'
    -- `is_delete` TINYINT(1) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已经删除'

CREATE TABLE if NOT exists `blog_tag` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';


create table IF not exists `blog_article` (
    `id` INT(10) UNSIGNED not null AUTO_INCREMENT,
    `title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
    `desc` VARCHAR(255) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面图片地址',
    `content` longtext COMMENT '文章内容',

    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT (10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_delete` TINYINT(1) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已经删除',
    `state` TINYINT(1) UNSIGNED DEFAULT '1' COMMENT '1为启用，0为禁用',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

CREATE TABLE if NOT exists `blog_article_tag` (
    `id` INT(10) UNSIGNED not NULL AUTO_INCREMENT,
    `article_id` INT(10) NOT NULL  COMMENT '文章ID',
    `tag_id` INT(10) NOT  NULL  COMMENT  '标签id',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
    `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT (10) UNSIGNED DEFAULT '0' COMMENT '删除时间',
    `is_delete` TINYINT(1) UNSIGNED DEFAULT '0' COMMENT '是否删除 0为未删除，1为已经删除',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联表';
