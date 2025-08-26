-- Задание 2.1: Создание пользователя MyUser
CREATE USER "MyUser" WITH LOGIN;

-- Задание 2.2: Установка пароля с ограничением по сроку

ALTER USER "MyUser" 
WITH PASSWORD '8961' 
VALID UNTIL '2025-12-31';


-- Задание 2.3: Предоставление прав на чтение

GRANT SELECT ON TABLE employees TO "MyUser";
GRANT SELECT ON TABLE departments TO "MyUser";

-- Задание 2.4: Отзыв прав на чтение

REVOKE SELECT ON TABLE employees FROM "MyUser";
REVOKE SELECT ON TABLE departments FROM "MyUser";

-- Задание 2.5: Удаление пользователя

DROP USER "MyUser";

-- Задание 3: Работа с транзакциями

BEGIN;

-- 3.2: Добавление новой записи в таблицу projects

INSERT INTO projects (project_id, project_name, start_date, end_date)
VALUES (9999, 'Test', CURRENT_DATE, CURRENT_DATE + INTERVAL '30 days');

-- 3.3: Создание точки сохранения

SAVEPOINT insert_project;

-- 3.4: Удаление добавленной записи

DELETE FROM projects WHERE project_id = 9999;

-- 3.5: Откат к точке сохранения

ROLLBACK TO SAVEPOINT insert_project;

-- 3.6: Фиксация изменений

COMMIT;