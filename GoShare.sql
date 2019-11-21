-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost
-- Tiempo de generación: 21-11-2019 a las 09:19:48
-- Versión del servidor: 10.4.8-MariaDB
-- Versión de PHP: 7.3.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `GoShare`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `Files`
--

CREATE TABLE `Files` (
  `id_file` int(11) NOT NULL,
  `nombre` varchar(100) NOT NULL,
  `hash` varchar(64) NOT NULL,
  `size` bigint(20) NOT NULL,
  `type` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `Peers`
--

CREATE TABLE `Peers` (
  `id_peer` int(11) NOT NULL,
  `ip` varchar(16) NOT NULL,
  `port` varchar(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `Peer_File`
--

CREATE TABLE `Peer_File` (
  `id_pf` int(11) NOT NULL,
  `id_file` int(11) NOT NULL,
  `id_peer` int(11) NOT NULL,
  `last_seen` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `Files`
--
ALTER TABLE `Files`
  ADD PRIMARY KEY (`id_file`);

--
-- Indices de la tabla `Peers`
--
ALTER TABLE `Peers`
  ADD PRIMARY KEY (`id_peer`);

--
-- Indices de la tabla `Peer_File`
--
ALTER TABLE `Peer_File`
  ADD PRIMARY KEY (`id_pf`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `Files`
--
ALTER TABLE `Files`
  MODIFY `id_file` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `Peers`
--
ALTER TABLE `Peers`
  MODIFY `id_peer` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `Peer_File`
--
ALTER TABLE `Peer_File`
  MODIFY `id_pf` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
