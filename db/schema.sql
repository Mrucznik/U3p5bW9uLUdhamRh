CREATE TABLE `urls`
(
    `id` int(11) AUTO_INCREMENT NOT NULL,
    `address` varchar(2048) NOT NULL,
    `interval` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `url_history`
(
    `id` int(11) AUTO_INCREMENT NOT NULL,
    `url_id` int(11) NOT NULL,
    `response` TEXT,
    `duration` DOUBLE NOT NULL,
    `created_at` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY(`url_id`) REFERENCES urls(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;