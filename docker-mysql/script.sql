DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255),
    `description` VARCHAR(255)
);


INSERT INTO tasks (title, description) VALUES ("title1", "desc1");
INSERT INTO tasks (title, description) VALUES ("title2", "desc2");
INSERT INTO tasks (title, description) VALUES ("title3", "desc3");