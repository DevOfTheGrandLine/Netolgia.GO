--Задание 1

-- Создание пользователя
CREATE USER 'sys_temp'@'%' IDENTIFIED BY 'password';

-- Просмотр всех пользователей
SELECT User, Host FROM mysql.user;

-- Выдача всех прав
GRANT ALL PRIVILEGES ON *.* TO 'sys_temp'@'%';
FLUSH PRIVILEGES;

-- Смена метода аутентификации
ALTER USER 'sys_temp'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
FLUSH PRIVILEGES;

-- Просмотр прав пользователя
SHOW GRANTS FOR 'sys_temp'@'%';

-- Отзыв прав на изменение данных
REVOKE INSERT, UPDATE, DELETE ON sakila.* FROM 'sys_temp'@'%';
FLUSH PRIVILEGES;


--Задание 2

-- Составьте таблицу, в которой должно быть два столбца: 
-- в первом должны быть названия таблиц восстановленной базы, 
-- во втором названия первичных ключей этих таблиц.
SELECT 
    TABLE_NAME AS 'Tabls',
    COLUMN_NAME AS 'PK'
FROM 
    INFORMATION_SCHEMA.KEY_COLUMN_USAGE
WHERE 
    TABLE_SCHEMA = 'sakila'
    AND CONSTRAINT_NAME = 'PRIMARY'
ORDER BY 
    TABLE_NAME;

--Задание 3

-- Отзываем все права на базу sakila
REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'sys_temp'@'%';
FLUSH PRIVILEGES;

-- проверяем
SHOW GRANTS FOR 'sys_temp'@'%';
