CREATE TABLE `urls`
(
    `id`          int(11) AUTO_INCREMENT NOT NULL,
    `address`        varchar(2048)            NOT NULL,
    `interval` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `url_history`
(
    `id`          int(11) AUTO_INCREMENT NOT NULL,
    `url_id`          int(11) FOREIGN KEY,
    `response` TEXT NULL,
    `duration` DOUBLE,
    `created_at` TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY(`url_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;