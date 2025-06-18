USE `servers`;
INSERT INTO `server` (`id`,`active`, `domain`, `hostname`, `os`) VALUES ('1', TRUE, 'amazon.com', 'Auth Server',  '0');
INSERT INTO `server` (`id`,`active`, `domain`, `hostname`, `os`) VALUES ('2', TRUE, 'amazon.com', 'Database Server',  '0');
INSERT INTO `server` (`id`,`active`, `domain`, `hostname`, `os`) VALUES ('3', TRUE, 'amazon.com', 'Payment Server',   '1');
INSERT INTO `server` (`id`,`active`, `domain`, `hostname`, `os`) VALUES ('4', TRUE, 'redhat.com', 'Network Server', '0');
INSERT INTO `server` (`id`,`active`, `domain`, `hostname`, `os`) VALUES ('5', TRUE, 'redhat.com', 'Backup Server', '2');

SET SQL_SAFE_UPDATES = 0;
UPDATE `hibernate_sequence` SET `next_val` = '6' WHERE `next_val` = '1';