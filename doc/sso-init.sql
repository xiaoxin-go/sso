-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: sso
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `casbin_rule`
--

--
-- Dumping data for table `casbin_rule`
--

--
-- Table structure for table `t_api`
--

DROP TABLE IF EXISTS `t_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_api` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `name` varchar(50) NOT NULL COMMENT 'api名称',
  `uri` varchar(255) NOT NULL COMMENT 'api路径',
  `method` varchar(10) NOT NULL COMMENT 'api请求方式',
  `description` varchar(255) DEFAULT NULL COMMENT 'api描述',
  `enabled` tinyint(1) DEFAULT '1' COMMENT 'api是否启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uri` (`uri`,`method`)
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_api`
--

LOCK TABLES `t_api` WRITE;
/*!40000 ALTER TABLE `t_api` DISABLE KEYS */;
INSERT INTO `t_api` VALUES (1,'2024-01-13 21:00:21','2024-01-13 21:00:21','','','获取用户列表','/system/users','GET',NULL,1),(2,'2024-01-13 21:32:04','2024-01-13 21:32:04','','','添加用户','/system/user','POST',NULL,1),(5,'2024-01-13 21:33:09','2024-01-13 21:33:09','','','修改用户信息','/system/user','PUT',NULL,1),(6,'2024-01-13 21:33:24','2024-01-13 21:33:24','','','删除用户','/system/user','DELETE',NULL,1),(7,'2024-01-13 21:33:40','2024-01-13 21:33:40','','','获取角色列表','/system/roles','GET',NULL,1),(67,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'添加角色','/system/role','POST',NULL,1),(68,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新角色信息','/system/role','PUT',NULL,1),(69,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'删除角色','/system/role','DELETE',NULL,1),(70,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'获取角色权限','/system/role/permission','GET',NULL,1),(71,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新角色权限','/system/role/permission','PUT',NULL,1),(72,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'获取菜单列表','/system/menus','GET',NULL,1),(73,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'添加菜单','/system/menu','POST',NULL,1),(74,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新菜单信息','/system/menu','PUT',NULL,1),(75,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'删除菜单','/system/menu','DELETE',NULL,1),(76,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'获取接口列表','/system/apis','GET',NULL,1),(77,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'添加接口','/system/api','POST',NULL,1),(78,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新接口信息','/system/api','PUT',NULL,1),(79,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'删除接口','/system/api','DELETE',NULL,1),(80,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'获取租户列表','/tenements','GET',NULL,1),(81,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'添加租户','/tenement','POST',NULL,1),(82,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新租户信息','/tenement','PUT',NULL,1),(83,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'删除租户','/tenement','DELETE',NULL,1),(84,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'获取平台列表','/platforms','GET',NULL,1),(85,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'添加平台','/platform','POST',NULL,1),(86,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新平台信息','/platform','PUT',NULL,1),(87,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'删除平台','/platform','DELETE',NULL,1),(88,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'获取平台用户列表','/platform_users','GET',NULL,1),(89,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'添加平台用户','/platform_user','POST',NULL,1),(90,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'更新平台用户信息','/platform_user','PUT',NULL,1),(91,'2024-01-13 21:41:34','2024-01-13 21:41:34',NULL,NULL,'删除平台用户','/platform_user','DELETE',NULL,1),(92,'2024-01-14 11:12:40','2024-01-14 11:12:40',NULL,NULL,'获取日志列表','/system/logs','GET',NULL,1),(93,'2024-01-16 15:16:19','2024-01-16 15:16:19',NULL,NULL,'获取平台分类列表','/platform_kinds','GET',NULL,1),(94,'2024-01-16 15:16:19','2024-01-16 15:16:19',NULL,NULL,'添加平台分类','/platform_kind','POST',NULL,1),(95,'2024-01-16 15:16:19','2024-01-16 15:16:19',NULL,NULL,'更新平台分类信息','/platform_kind','PUT',NULL,1),(96,'2024-01-16 15:16:19','2024-01-16 15:16:19',NULL,NULL,'删除平台分类','/platform_kind','DELETE',NULL,1);
/*!40000 ALTER TABLE `t_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_log`
--

DROP TABLE IF EXISTS `t_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `operator` varchar(50) DEFAULT NULL COMMENT '操作人',
  `content` varchar(5000) DEFAULT NULL COMMENT '操作内容',
  `created_by` varchar(64) DEFAULT NULL,
  `updated_by` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_menu`
--

DROP TABLE IF EXISTS `t_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `name` varchar(50) NOT NULL COMMENT '菜单名称',
  `name_en` varchar(50) NOT NULL COMMENT '菜单英文名称',
  `path` varchar(255) NOT NULL COMMENT '菜单路径',
  `icon` varchar(255) DEFAULT NULL COMMENT '菜单图标',
  `parent_id` int DEFAULT NULL COMMENT '父菜单id',
  `sort` int DEFAULT NULL COMMENT '菜单排序',
  `enabled` tinyint(1) DEFAULT '1' COMMENT '菜单是否启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_menu`
--

LOCK TABLES `t_menu` WRITE;
/*!40000 ALTER TABLE `t_menu` DISABLE KEYS */;
INSERT INTO `t_menu` VALUES (1,'2024-01-13 15:26:42','2024-01-14 19:54:11',NULL,'admin','我的平台','MyPlatform','/platform/self','menu-model',NULL,1,1),(2,'2024-01-13 17:40:38','2024-01-16 15:19:01','','admin','平台管理','platform','/platform','cloud-admin-box',0,2,1),(3,'2024-01-13 17:41:17','2024-01-16 09:34:14','','admin','租户管理','tenement','/tenement','cloud-admin-user',0,3,1),(4,'2024-01-13 17:41:50','2024-01-16 09:35:03','','admin','平台用户管理','platformUser','/platform_user','cloud-admin-user',0,4,1),(5,'2024-01-13 18:04:46','2024-01-13 18:04:46','','','系统管理','system','/system','cloud-admin-setting',0,4,1),(6,'2024-01-13 18:05:27','2024-01-14 11:14:39','','admin','用户管理','user','/system/user','',5,1,1),(7,'2024-01-13 18:05:46','2024-01-14 11:13:43','','admin','角色管理','role','/system/role','',5,2,1),(8,'2024-01-13 18:06:08','2024-01-14 11:13:29','','admin','菜单管理','menu','/system/menu','',5,3,1),(9,'2024-01-13 18:06:25','2024-01-14 10:46:47','','admin','接口管理','api','/system/api','',5,4,1),(10,'2024-01-13 18:06:43','2024-01-14 11:12:58','','admin','日志管理','log','/system/log','',5,5,1),(11,'2024-01-16 15:18:35','2024-01-16 15:20:21','','admin','平台类别','PlatformKind','/platform_kind','cloud-admin-relation',0,4,1);
/*!40000 ALTER TABLE `t_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_menu_api`
--

DROP TABLE IF EXISTS `t_menu_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_menu_api` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `menu_id` int NOT NULL COMMENT '菜单id',
  `api_id` int NOT NULL COMMENT 'apiid',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_menu_api`
--

LOCK TABLES `t_menu_api` WRITE;
/*!40000 ALTER TABLE `t_menu_api` DISABLE KEYS */;
INSERT INTO `t_menu_api` VALUES (10,'2024-01-14 10:46:47','2024-01-14 10:46:47','','',9,76),(11,'2024-01-14 10:46:47','2024-01-14 10:46:47','','',9,77),(12,'2024-01-14 10:46:47','2024-01-14 10:46:47','','',9,78),(13,'2024-01-14 10:46:47','2024-01-14 10:46:47','','',9,79),(14,'2024-01-14 11:12:58','2024-01-14 11:12:58','','',10,92),(15,'2024-01-14 11:13:29','2024-01-14 11:13:29','','',8,72),(16,'2024-01-14 11:13:29','2024-01-14 11:13:29','','',8,73),(17,'2024-01-14 11:13:29','2024-01-14 11:13:29','','',8,74),(18,'2024-01-14 11:13:29','2024-01-14 11:13:29','','',8,75),(19,'2024-01-14 11:13:29','2024-01-14 11:13:29','','',8,76),(20,'2024-01-14 11:13:43','2024-01-14 11:13:43','','',7,7),(21,'2024-01-14 11:13:43','2024-01-14 11:13:43','','',7,67),(22,'2024-01-14 11:13:43','2024-01-14 11:13:43','','',7,68),(23,'2024-01-14 11:13:43','2024-01-14 11:13:43','','',7,69),(24,'2024-01-14 11:13:43','2024-01-14 11:13:43','','',7,70),(25,'2024-01-14 11:13:43','2024-01-14 11:13:43','','',7,71),(40,'2024-01-14 11:14:39','2024-01-14 11:14:39','','',6,1),(41,'2024-01-14 11:14:39','2024-01-14 11:14:39','','',6,2),(42,'2024-01-14 11:14:39','2024-01-14 11:14:39','','',6,5),(43,'2024-01-14 11:14:39','2024-01-14 11:14:39','','',6,6),(44,'2024-01-14 11:14:39','2024-01-14 11:14:39','','',6,7),(45,'2024-01-14 11:14:39','2024-01-14 11:14:39','','',6,80),(63,'2024-01-16 09:34:14','2024-01-16 09:34:14','','',3,80),(64,'2024-01-16 09:34:14','2024-01-16 09:34:14','','',3,81),(65,'2024-01-16 09:34:14','2024-01-16 09:34:14','','',3,82),(66,'2024-01-16 09:34:14','2024-01-16 09:34:14','','',3,83),(67,'2024-01-16 09:35:03','2024-01-16 09:35:03','','',4,88),(68,'2024-01-16 09:35:03','2024-01-16 09:35:03','','',4,89),(69,'2024-01-16 09:35:03','2024-01-16 09:35:03','','',4,90),(70,'2024-01-16 09:35:03','2024-01-16 09:35:03','','',4,91),(71,'2024-01-16 09:35:03','2024-01-16 09:35:03','','',4,84),(72,'2024-01-16 15:19:01','2024-01-16 15:19:01','','',2,84),(73,'2024-01-16 15:19:01','2024-01-16 15:19:01','','',2,85),(74,'2024-01-16 15:19:01','2024-01-16 15:19:01','','',2,86),(75,'2024-01-16 15:19:01','2024-01-16 15:19:01','','',2,87),(76,'2024-01-16 15:19:01','2024-01-16 15:19:01','','',2,93);
/*!40000 ALTER TABLE `t_menu_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_platform`
--

DROP TABLE IF EXISTS `t_platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_platform` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `name` varchar(64) DEFAULT NULL COMMENT '平台名',
  `name_cn` varchar(64) DEFAULT NULL COMMENT '平台中文名',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `url` varchar(64) DEFAULT NULL COMMENT '平台地址',
  `index_url` varchar(64) DEFAULT NULL COMMENT '平台首页',
  `kind_id` int DEFAULT NULL,
  `type` int DEFAULT NULL COMMENT '平台类型',
  `login_func` varchar(20) DEFAULT NULL COMMENT '登录函数',
  `enabled` tinyint(1) DEFAULT NULL COMMENT '是否启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COMMENT='平台表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_platform`
--

LOCK TABLES `t_platform` WRITE;
/*!40000 ALTER TABLE `t_platform` DISABLE KEYS */;
INSERT INTO `t_platform` VALUES (1,'2024-01-14 09:41:13','2024-01-16 17:18:39',NULL,'admin','netops','网络自动化','网络自动化111','sso.com','sso.com/index.html',1,1,NULL,1),(2,'2024-01-16 17:12:44','2024-01-16 17:14:11','','admin','gitlab','代码仓库','代码托管平台','https://gitlab.com','https://gitlab.com/index.html',3,1,'gitlab',1),(3,'2024-01-16 17:13:45','2024-01-16 17:13:45','','','zabbix','监控平台','zabbix监控平台','https://zabbix.com','https://zabbix.com/index.html',4,2,'zabbix',1),(4,'2024-01-16 22:30:08','2024-01-18 14:19:42','','admin','zabbix-sh','上海zabbix','上海zabbix监控平台','https://www.zabbix-sh.tool.com.cn','https://zabbix-sh.com/index.html',4,2,'zabbix',1),(5,'2024-01-16 22:30:53','2024-01-16 22:30:53','','','zabbix-nj','南京zabbix','南京zabbix平台','https://zabbix-nj.com','https://zabbix-nj.com/index.html',4,2,'zabbix',1),(6,'2024-01-16 22:31:30','2024-01-16 22:31:30','','','zabbix-bj','北京zabbix','北京zabbix监控平台','https://zabbix-bj.com','https://zabbix-bj.com/index.html',4,2,'',1);
/*!40000 ALTER TABLE `t_platform` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_platform_kind`
--

DROP TABLE IF EXISTS `t_platform_kind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_platform_kind` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL COMMENT '类型名',
  `description` varchar(255) DEFAULT NULL COMMENT '类型描述',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_platform_kind`
--

LOCK TABLES `t_platform_kind` WRITE;
/*!40000 ALTER TABLE `t_platform_kind` DISABLE KEYS */;
INSERT INTO `t_platform_kind` VALUES (1,'网络管理','网络相关平台','2024-01-16 15:23:24','2024-01-16 15:23:24','',''),(2,'中间件','中间件相关平台','2024-01-16 15:23:34','2024-01-16 15:23:34','',''),(3,'数据库管理','数据库相关管理平台','2024-01-16 15:24:04','2024-01-16 15:24:04','',''),(4,'监控','监控类平台','2024-01-16 15:24:14','2024-01-16 15:24:14','','');
/*!40000 ALTER TABLE `t_platform_kind` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_platform_record`
--

DROP TABLE IF EXISTS `t_platform_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_platform_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL COMMENT '用户ID',
  `platform_id` int DEFAULT NULL COMMENT '平台ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_platform_record`
--

LOCK TABLES `t_platform_record` WRITE;
/*!40000 ALTER TABLE `t_platform_record` DISABLE KEYS */;
INSERT INTO `t_platform_record` VALUES (1,1008,2,'2024-01-18 14:08:10','2024-01-18 14:08:10','',''),(2,1008,2,'2024-01-18 14:08:35','2024-01-18 14:08:35','',''),(3,1008,2,'2024-01-18 14:09:21','2024-01-18 14:09:21','',''),(4,1008,4,'2024-01-18 14:13:52','2024-01-18 14:13:52','',''),(5,1008,2,'2024-01-18 15:34:29','2024-01-18 15:34:29','',''),(6,1009,2,'2024-01-21 17:22:41','2024-01-21 17:22:41','',''),(7,1009,3,'2024-01-21 17:22:48','2024-01-21 17:22:48','',''),(8,1009,1,'2024-01-21 17:22:50','2024-01-21 17:22:50','','');
/*!40000 ALTER TABLE `t_platform_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_platform_user`
--

DROP TABLE IF EXISTS `t_platform_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_platform_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `platform_id` int DEFAULT NULL COMMENT '平台ID',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `is_default` tinyint(1) DEFAULT NULL COMMENT '是否是默认账号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='平台用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_platform_user`
--

LOCK TABLES `t_platform_user` WRITE;
/*!40000 ALTER TABLE `t_platform_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_platform_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role`
--

DROP TABLE IF EXISTS `t_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `description` varchar(255) DEFAULT NULL COMMENT '角色描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_role_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role`
--

LOCK TABLES `t_role` WRITE;
/*!40000 ALTER TABLE `t_role` DISABLE KEYS */;
INSERT INTO `t_role` VALUES (1,'2024-01-07 09:09:38','2024-01-07 09:12:29','',NULL,'admin','超级管理员1');
/*!40000 ALTER TABLE `t_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role_api`
--

DROP TABLE IF EXISTS `t_role_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_role_api` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `role_id` int NOT NULL COMMENT '角色id',
  `api_id` int NOT NULL COMMENT 'apiid',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role_api`
--

LOCK TABLES `t_role_api` WRITE;
/*!40000 ALTER TABLE `t_role_api` DISABLE KEYS */;
INSERT INTO `t_role_api` VALUES (1,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,6),(2,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,75),(3,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,92),(4,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,86),(5,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,83),(6,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,71),(7,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,80),(8,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,89),(9,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,7),(10,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,69),(11,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,70),(12,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,85),(13,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,91),(14,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,78),(15,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,90),(16,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,68),(17,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,79),(18,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,84),(19,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,87),(20,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,81),(21,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,77),(22,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,88),(23,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,5),(24,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,72),(25,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,67),(26,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,73),(27,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,74),(28,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,76),(29,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,82),(30,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,1),(31,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,2);
/*!40000 ALTER TABLE `t_role_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role_menu`
--

DROP TABLE IF EXISTS `t_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `role_id` int NOT NULL COMMENT '角色id',
  `menu_id` int NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role_menu`
--

LOCK TABLES `t_role_menu` WRITE;
/*!40000 ALTER TABLE `t_role_menu` DISABLE KEYS */;
INSERT INTO `t_role_menu` VALUES (1,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,1),(2,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,2),(3,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,3),(4,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,4),(5,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,5),(6,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,6),(7,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,7),(8,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,8),(9,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,9),(10,'2024-01-14 11:17:37','2024-01-14 11:17:37',NULL,NULL,1,10);
/*!40000 ALTER TABLE `t_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_tenement`
--

DROP TABLE IF EXISTS `t_tenement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_tenement` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `name` varchar(64) DEFAULT NULL COMMENT '租户名',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 COMMENT='租户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_tenement`
--

LOCK TABLES `t_tenement` WRITE;
/*!40000 ALTER TABLE `t_tenement` DISABLE KEYS */;
INSERT INTO `t_tenement` VALUES (1,'2024-01-14 19:52:00','2024-01-14 19:52:00','','','sre','应用运维'),(2,'2024-01-14 19:52:11','2024-01-14 19:52:11','','','网络组','网络管理员'),(3,'2024-01-14 19:52:26','2024-01-14 19:52:26','','','admin','管理员');
/*!40000 ALTER TABLE `t_tenement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_tenement_platform`
--

DROP TABLE IF EXISTS `t_tenement_platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_tenement_platform` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tenement_id` int DEFAULT NULL COMMENT '租户ID',
  `platform_id` int DEFAULT NULL COMMENT '平台ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3 COMMENT='租户关联平台';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_tenement_platform`
--

LOCK TABLES `t_tenement_platform` WRITE;
/*!40000 ALTER TABLE `t_tenement_platform` DISABLE KEYS */;
INSERT INTO `t_tenement_platform` VALUES (5,3,6,'2024-01-16 22:31:38','','2024-01-16 22:31:38',''),(6,3,5,'2024-01-16 22:31:38','','2024-01-16 22:31:38',''),(7,3,4,'2024-01-16 22:31:38','','2024-01-16 22:31:38',''),(8,3,1,'2024-01-16 22:31:38','','2024-01-16 22:31:38',''),(9,3,2,'2024-01-16 22:31:38','','2024-01-16 22:31:38',''),(10,3,3,'2024-01-16 22:31:38','','2024-01-16 22:31:38','');
/*!40000 ALTER TABLE `t_tenement_platform` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_tenement_user`
--

DROP TABLE IF EXISTS `t_tenement_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_tenement_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tenement_id` int DEFAULT NULL COMMENT '租户ID',
  `user_id` int DEFAULT NULL COMMENT '用户ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 COMMENT='租户关联用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_tenement_user`
--

LOCK TABLES `t_tenement_user` WRITE;
/*!40000 ALTER TABLE `t_tenement_user` DISABLE KEYS */;
INSERT INTO `t_tenement_user` VALUES (3,3,1,'2024-01-21 11:55:12','yangmaoqi','2024-01-21 11:55:12','');
/*!40000 ALTER TABLE `t_tenement_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `email` varchar(50) NOT NULL COMMENT '用户邮箱地址',
  `name_cn` varchar(255) NOT NULL COMMENT '用户中文名',
  `otp_secret` varchar(64) NOT NULL COMMENT 'otp_secret',
  `enabled` tinyint(1) DEFAULT '1' COMMENT '用户是否启用',
  `password_updated_at` datetime DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1010 DEFAULT CHARSET=utf8mb3 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_user`
--

LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
INSERT INTO `t_user` VALUES (1,'2024-01-14 11:31:14','2024-01-14 11:32:55',NULL,'admin','admin','admin@qq.com','管理员','123456',1,NULL,'123456');
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_user_role`
--

DROP TABLE IF EXISTS `t_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `user_id` int NOT NULL COMMENT '用户id',
  `role_id` int NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_user_role`
--

LOCK TABLES `t_user_role` WRITE;
/*!40000 ALTER TABLE `t_user_role` DISABLE KEYS */;
INSERT INTO `t_user_role` VALUES (1,'2024-01-14 11:32:55','2024-01-14 11:32:55',NULL,NULL,1,1);
/*!40000 ALTER TABLE `t_user_role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-21 20:31:32
