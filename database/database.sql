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
  `name` varchar(50) NOT NULL,
  `time` float DEFAULT '0',
  `number1` int DEFAULT '0',
  `ingredient1` varchar(50) DEFAULT '',
  `number2` int DEFAULT '0',
  `ingredient2` varchar(50) DEFAULT '',
  `number3` int DEFAULT '0',
  `ingredient3` varchar(50) DEFAULT '',
  `result` int NOT NULL,
  `machineType` varchar(30) NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES ('Advanced circuit',6,4,'Copper cable',2,'Electronic circuit',2,'Plastic bar',1,'Assembling'),('Coal',0,0,'',0,'',0,'',1,'Mining'),('Copper cable',0.5,1,'Copper plate',0,'',0,'',2,'Assembling'),('Copper ore',0,0,'',0,'',0,'',1,'Mining'),('Copper plate',3.2,1,'Copper ore',0,'',0,'',1,'Furnace'),('Crude oil',0,0,'',0,'',0,'',1,'Pumping'),('Electronic circuit',0.5,3,'Copper cable',1,'Iron plate',0,'',1,'Assembling'),('Iron gear wheel',0.5,2,'Iron ore',0,'',0,'',1,'Assembling'),('Iron ore',0,0,'',0,'',0,'',1,'Mining'),('Iron plate',3.2,1,'Iron ore',0,'',0,'',1,'Furnace'),('Petroleum gas',5,100,'Crude oil',0,'',0,'',45,'Refinery'),('Plastic bar',1,1,'Coal',20,'Petroleum gas',0,'',2,'Chemical'),('Steel plate',16,5,'Iron plate',0,'',0,'',1,'Furnace'),('Stone',0,0,'',0,'',0,'',1,'Mining'),('Stone brick',3.2,2,'Stone',0,'',0,'',1,'Furnace'),('Transport belt',0.5,1,'Iron gear wheel',1,'Iron plate',0,'',1,'Assembling');
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `machines`
--

DROP TABLE IF EXISTS `machines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `machines` (
  `name` varchar(50) NOT NULL,
  `type` varchar(50) NOT NULL,
  `number1` int DEFAULT '0',
  `ingredient1` varchar(50) DEFAULT '',
  `number2` int DEFAULT '0',
  `ingredient2` varchar(50) DEFAULT '',
  `number3` int DEFAULT '0',
  `ingredient3` varchar(50) DEFAULT '',
  `time` float DEFAULT '0',
  `speed` float NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `machines`
--

LOCK TABLES `machines` WRITE;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */;
INSERT INTO `machines` VALUES ('Assembling machine 1','Assembling',3,'Electronic circuit',5,'Iron gear wheel',9,'Iron plate',0.5,0.5),('Assembling machine 2','Assembling',3,'Electronic circuit',5,'Iron gear wheel',2,'Steel plate',0.5,0.75),('Electric furnace','Furnace',5,'Advanced circuit',10,'Steel plate',10,'Stone brick',5,2),('Electric mining drill','Mining',3,'Electronic circuit',5,'Iron gear wheel',10,'Iron plate',2,0.5);
/*!40000 ALTER TABLE `machines` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-16  9:36:25