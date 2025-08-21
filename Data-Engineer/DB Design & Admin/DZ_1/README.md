#### **Задание 1. Работа с командной строкой**  

1.1. Создайте новую базу данных с названием **`hr_db`**:  
```bash
createdb hr_db
```

1.2. Восстановите бэкап учебной БД из файла `hr.sql` в созданную базу:  
```bash
psql -d hr_db -U postgres -h localhost -p 5432 -f hr.sql
```

1.3. Подключитесь к базе `hr_db` и выведите список всех таблиц:  
```bash
psql -d hr_db
```
```sql
\dt  -- Показать все таблицы в текущей схеме
```

1.4. Выполните запрос на выборку всех данных из таблицы `employees`:  
```sql
SELECT * FROM employees;  -- Пример для таблицы employees
```

---

#### **Задание 2. Работа с пользователями**  

2.1. Создайте пользователя `myuser` с разрешением на вход, без пароля и прав:  
```sql
CREATE USER myuser WITH LOGIN;
```

2.2. Установите пароль `secure_password` с сроком действия до конца текущего месяца (пример для мая 2024):  
```sql
ALTER USER myuser 
WITH PASSWORD 'secure_password' 
VALID UNTIL '2024-05-31';  -- Замените на актуальную дату
```

2.3. Предоставьте права на чтение для таблиц `employees` и `departments`:  
```sql
GRANT SELECT ON TABLE employees TO myuser;
GRANT SELECT ON TABLE departments TO myuser;
```

2.4. Отзовите выданные права:  
```sql
REVOKE SELECT ON TABLE employees FROM myuser;
REVOKE SELECT ON TABLE departments FROM myuser;
```

2.5. Удалите пользователя:  
```sql
DROP USER myuser;
```

---

#### **Задание 3. Работа с транзакциями**  

```sql
BEGIN;  -- Начало транзакции

-- Добавление записи в таблицу projects (пример для HR-схемы)
INSERT INTO projects (project_id, project_name, start_date, end_date)
VALUES (9999, 'Test Project', CURRENT_DATE, CURRENT_DATE + INTERVAL '30 days');

SAVEPOINT insert_project;  -- Создание точки сохранения

DELETE FROM projects WHERE project_id = 9999;  -- Удаление добавленной записи

ROLLBACK TO SAVEPOINT insert_project;  -- Откат к точке сохранения

COMMIT;  -- Фиксация изменений
```
