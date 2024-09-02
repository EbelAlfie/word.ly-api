-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 02, 2024 at 02:52 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `wordlydb`
--

-- --------------------------------------------------------

--
-- Table structure for table `quiz_table`
--

CREATE TABLE `quiz_table` (
  `QuizId` int(11) NOT NULL,
  `TeacherId` int(11) NOT NULL,
  `ChoiceId` int(11) NOT NULL,
  `Soal` varchar(256) NOT NULL,
  `CorrectAnswer` varchar(100) NOT NULL,
  `Hint` varchar(256) NOT NULL,
  `Score` int(11) NOT NULL,
  `Type` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `quiz_table`
--
ALTER TABLE `quiz_table`
  ADD PRIMARY KEY (`QuizId`),
  ADD KEY `TeacherId` (`TeacherId`),
  ADD KEY `ChoiceId` (`ChoiceId`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `quiz_table`
--
ALTER TABLE `quiz_table`
  MODIFY `QuizId` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `quiz_table`
--
ALTER TABLE `quiz_table`
  ADD CONSTRAINT `quiz_table_ibfk_1` FOREIGN KEY (`TeacherId`) REFERENCES `user_data` (`Uid`),
  ADD CONSTRAINT `quiz_table_ibfk_2` FOREIGN KEY (`ChoiceId`) REFERENCES `choice_table` (`ChoiceId`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
