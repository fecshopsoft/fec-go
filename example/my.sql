-- phpMyAdmin SQL Dump
-- version 3.3.10
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2018 年 03 月 19 日 10:19
-- 服务器版本: 5.6.14
-- PHP 版本: 5.4.34

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";

--
-- 数据库: `fec-go`
--

-- --------------------------------------------------------

--
-- 表的结构 `customer`
--

CREATE TABLE IF NOT EXISTS `customer` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `created_at` int(30) DEFAULT NULL,
  `updated_at` int(30) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `sex` int(5) DEFAULT NULL,
  `telephone` varchar(20) DEFAULT NULL,
  `access_token` varchar(200) DEFAULT NULL,
  `status` int(5) DEFAULT NULL,
  `age` int(5) DEFAULT NULL,
  `remark` text,
  `birth_date` int(20) DEFAULT NULL COMMENT '出生年月',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=88 ;

--
-- 转存表中的数据 `customer`
--

INSERT INTO `customer` (`id`, `username`, `password`, `created_at`, `updated_at`, `email`, `name`, `sex`, `telephone`, `access_token`, `status`, `age`, `remark`, `birth_date`) VALUES
(15, 'admin', 'xxxx', 1519978859, 1521334376, 'zqy234@126.com', 'Terry', 2, '1855343432', NULL, 1, 32, '<!DOCTYPE html><br /><html><br /><head><br /></head><br /><body>\nFecshop 创始人\n</body><br /></html>', 534096000),
(87, 'editor', 'xxxx', 1521329333, 1521329333, '2358269014@qq.com', '编辑人', 1, '1545454545', NULL, 1, 18, '<!DOCTYPE html><br /><html><br /><head><br /></head><br /><body>\n小表年纪\n</body><br /></html>', 946915200);
