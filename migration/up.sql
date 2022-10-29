CREATE TABLE `banking`.`users` (
                                   `id` INT NOT NULL AUTO_INCREMENT,
                                   `name` VARCHAR(145) NULL,
                                   `email` VARCHAR(145) NULL,
                                   `password` VARCHAR(145) NULL,
                                   `username` VARCHAR(145) NULL,
                                   PRIMARY KEY (`id`),
                                   UNIQUE INDEX `email_UNIQUE` (`email` ASC),
                                   UNIQUE INDEX `username_UNIQUE` (`username` ASC));
