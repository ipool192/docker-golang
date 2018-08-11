CREATE SCHEMA `db_exchange_rate` ;

CREATE TABLE `db_exchange_rate`.`daily_exchange_rate` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `exchange_rate_id` INT(5) NOT NULL,
  `rate` DOUBLE(9,6) NOT NULL,
  `date` DATE NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `db_exchange_rate`.`exchange_rates` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ex_from` VARCHAR(6) NOT NULL,
  `ex_to` VARCHAR(6) NOT NULL,
  PRIMARY KEY (`id`));
