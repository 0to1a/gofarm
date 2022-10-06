CREATE TABLE access_list (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL DEFAULT '' COLLATE 'utf8mb4_general_ci',
    api_key VARCHAR(50) NOT NULL DEFAULT '' COLLATE 'utf8mb4_general_ci',

    PRIMARY KEY (id) USING BTREE
) COLLATE='utf8mb4_general_ci' ENGINE=InnoDB;