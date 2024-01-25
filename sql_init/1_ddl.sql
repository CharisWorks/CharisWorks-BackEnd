-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: CharisWorks
-- ------------------------------------------------------
-- Server version	11.2.2-MariaDB-1:11.2.2+maria~ubu2204

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

CREATE DATABASE IF NOT EXISTS `CharisWorks`;

USE `CharisWorks`;

--
-- Table structure for table `carts`
--

DROP TABLE IF EXISTS `carts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `carts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `purchaser_user_id` varchar(100) NOT NULL,
  `item_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `carts_items_FK` (`item_id`),
  KEY `carts_users_FK` (`purchaser_user_id`),
  CONSTRAINT `carts_items_FK` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `carts_users_FK` FOREIGN KEY (`purchaser_user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `history_items`
--

DROP TABLE IF EXISTS `history_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `history_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `price` int(11) NOT NULL,
  `status` varchar(100) NOT NULL,
  `stock` int(11) NOT NULL,
  `size` int(11) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `tags` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `history_items_items_FK` (`item_id`),
  CONSTRAINT `history_items_items_FK` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `history_users`
--

DROP TABLE IF EXISTS `history_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `history_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(100) NOT NULL,
  `real_name` varchar(100) NOT NULL,
  `display_name` varchar(100) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `history_users_users_FK` (`user_id`),
  CONSTRAINT `history_users_users_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `manufacturer_user_id` varchar(100) NOT NULL,
  `history_item_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `history_item_id` (`history_item_id`),
  KEY `items_users_FK` (`manufacturer_user_id`),
  CONSTRAINT `items_history_items_FK` FOREIGN KEY (`history_item_id`) REFERENCES `history_items` (`id`),
  CONSTRAINT `items_users_FK` FOREIGN KEY (`manufacturer_user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `shippings`
--

DROP TABLE IF EXISTS `shippings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shippings` (
  `id` varchar(100) NOT NULL,
  `zip_code` varchar(100) NOT NULL,
  `address_1` varchar(100) NOT NULL,
  `address_2` varchar(100) NOT NULL,
  `address_3` varchar(100) DEFAULT NULL,
  `phone_number` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `shippings_users_FK` FOREIGN KEY (`id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `manufacturer_user_id` varchar(100) NOT NULL,
  `purchaser_user_id` varchar(100) NOT NULL,
  `item_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `tracking_id` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL,
  `zip_code` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `phone_number` varchar(100) NOT NULL,
  `history_manufacturer_user_id` int(11) NOT NULL,
  `history_purchaser_user_id` int(11) NOT NULL,
  `history_item_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `transactions_items_FK` (`item_id`),
  KEY `transactions_history_users_FK` (`history_manufacturer_user_id`),
  KEY `transactions_history_users_FK_1` (`history_purchaser_user_id`),
  KEY `transactions_history_items_FK` (`history_item_id`),
  KEY `transactions_users_FK` (`manufacturer_user_id`),
  KEY `transactions_users_FK_1` (`purchaser_user_id`),
  CONSTRAINT `transactions_history_items_FK` FOREIGN KEY (`history_item_id`) REFERENCES `history_items` (`id`),
  CONSTRAINT `transactions_history_users_FK` FOREIGN KEY (`history_manufacturer_user_id`) REFERENCES `history_users` (`id`),
  CONSTRAINT `transactions_history_users_FK_1` FOREIGN KEY (`history_purchaser_user_id`) REFERENCES `history_users` (`id`),
  CONSTRAINT `transactions_items_FK` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `transactions_users_FK` FOREIGN KEY (`manufacturer_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `transactions_users_FK_1` FOREIGN KEY (`purchaser_user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(100) NOT NULL,
  `stripe_account_id` varchar(100) DEFAULT NULL,
  `history_user_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `history_user_id` (`history_user_id`),
  UNIQUE KEY `stripe_account_id` (`stripe_account_id`),
  CONSTRAINT `users_history_users_FK` FOREIGN KEY (`history_user_id`) REFERENCES `history_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'CharisWorks'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-25 14:18:37
