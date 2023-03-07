-- MySQL dump 10.13  Distrib 8.0.23, for Linux (x86_64)
--
-- Host: localhost    Database: newdockerdb
-- ------------------------------------------------------
-- Server version       8.0.23

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
-- Current Database: `newdockerdb`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `newdockerdb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `newdockerdb`;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `time` float DEFAULT '0',
  `result` int NOT NULL,
  `machineType` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (1,'Advanced circuit',6,1,'Assembling'),(2,'Coal',0,1,'Mining'),(3,'Copper cable',0.5,2,'Assembling'),(4,'Copper ore',0,1,'Mining'),(5,'Copper plate',3.2,1,'Furnace'),(6,'Crude oil',0,1,'Pumping'),(7,'Electronic circuit',0.5,1,'Assembling'),(8,'Iron gear wheel',0.5,1,'Assembling'),(9,'Iron ore',0,1,'Mining'),(10,'Iron plate',3.2,1,'Furnace'),(11,'Petroleum gas',5,45,'Refinery'),(12,'Plastic bar',1,2,'Chemical'),(13,'Steel plate',16,1,'Furnace'),(14,'Stone',0,1,'Mining'),(15,'Stone brick',3.2,1,'Furnace'),(16,'Transport belt',0.5,1,'Assembling'),(17,'Assembling machine 1',0.5,1,'Assembling'),(18,'Inserter',0.5,1,'Assembling');
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `machines`
--

DROP TABLE IF EXISTS `machines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `machines` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `type` varchar(50) NOT NULL,
  `time` float DEFAULT '0',
  `speed` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `machines`
--

LOCK TABLES `machines` WRITE;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */;
INSERT INTO `machines` VALUES (2,'Assembling machine 2','Assembling',0.5,0.75),(3,'Electric furnace','Furnace',5,2),(4,'Electric mining drill','Mining',2,0.5),(5,'Assembling machine 1','Assembling',0.5,0.5);
/*!40000 ALTER TABLE `machines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recipes`
--

DROP TABLE IF EXISTS `recipes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `recipes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `item` varchar(50) NOT NULL,
  `number` int NOT NULL,
  `ingredient` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1035 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recipes`
--

LOCK TABLES `recipes` WRITE;
/*!40000 ALTER TABLE `recipes` DISABLE KEYS */;
INSERT INTO `recipes` VALUES (1001,'Advanced circuit',4,'Copper cable'),(1002,'Advanced circuit',2,'Electronic circuit'),(1004,'Copper cable',2,'Copper plate'),(1005,'Copper plate',1,'Copper 
ore'),(1006,'Iron gear wheel',2,'Iron plate'),(1007,'Iron plate',1,'Iron ore'),(1008,'Petroleum gas',100,'Crude oil'),(1009,'Plastic bar',1,'Coal'),(1010,'Plastic bar',20,'Petroleum gas'),(1011,'Steel plate',5,'Iron plate'),(1012,'Stone brick',2,'Stone'),(1013,'Transport belt',1,'Iron gear wheel'),(1014,'Transport belt',1,'Iron plate'),(1015,'Advanced circuit',2,'Plastic bar'),(1016,'Assembling machine 1',3,'Electronic circuit'),(1017,'Assembling machine 1',5,'Iron gear wheel'),(1018,'Assembling machine 1',9,'Iron plate'),(1019,'Electric furnace',5,''),(1020,'Electric furnace',10,'Steel plate'),(1021,'Electric furnace',10,'Stone brick'),(1022,'Electric mining drill',3,'Electronic circuit'),(1023,'Electric mining drill',5,'Iron gear wheel'),(1024,'Electric mining drill',10,'Iron plate'),(1028,'Assembling machine 2',1,'Assembling machine 1'),(1029,'Assembling machine 2',3,'Electronic circuit'),(1030,'Assembling machine 2',5,'Iron gear wheel'),(1031,'Assembling machine 2',2,'Steel plate'),(1032,'Inserter',1,'Electronic circuit'),(1033,'Inserter',1,'Iron gear wheel'),(1034,'Inserter',1,'Iron plate'),(1035,'Electronic circuit',2,'Copper cable'),(1036,'Electronic circuit',1,'Iron plate');
/*!40000 ALTER TABLE `recipes` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-03 13:06:37