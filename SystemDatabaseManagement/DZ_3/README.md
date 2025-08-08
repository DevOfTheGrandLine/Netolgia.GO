# Домашнее задание к занятию "SQL. Часть 1" - `Конкин Дмитрий`

Задание 1
Получите уникальные названия районов из таблицы с адресами, которые начинаются на “K” и заканчиваются на “a” и не содержат пробелов.
```sql
SELECT DISTINCT district
FROM address
WHERE district LIKE 'K%a'
  AND district NOT LIKE '% %';
  ```
![txt](img/{2DAA9717-A95D-4972-9701-53B8EA987673}.png)

Задание 2
Получите из таблицы платежей за прокат фильмов информацию по платежам, которые выполнялись в промежуток с 15 июня 2005 года по 18 июня 2005 года включительно и стоимость которых превышает 10.00.
```sql
SELECT *
FROM payment
WHERE DATE(payment_date) BETWEEN '2005-06-15' AND '2005-06-18'
  AND amount > 10.00;
  ```
![txt](img/{1659DE7F-CB28-47F6-80FB-757F54BF3E49}.png)

Задание 3
Получите последние пять аренд фильмов.
```sql
SELECT *
FROM rental
ORDER BY rental_date DESC
LIMIT 5;
```
![txt](img/{B8D29B0C-83C2-431C-AD1B-AA0535CE6A60}.png)

Задание 4
Одним запросом получите активных покупателей, имена которых Kelly или Willie.

Сформируйте вывод в результат таким образом:

все буквы в фамилии и имени из верхнего регистра переведите в нижний регистр,
замените буквы 'll' в именах на 'pp'.

```sql
SELECT
    LOWER(REPLACE(first_name, 'LL', 'pp')) AS first_name,
    LOWER(last_name) AS last_name
FROM customer
WHERE active = 1
  AND UPPER(first_name) IN ('KELLY', 'WILLIE');
```
![txt](img/{E7686D0B-8A91-4403-9FAC-1098F90BEB31}.png)

Задание 5*
Выведите Email каждого покупателя, разделив значение Email на две отдельных колонки: в первой колонке должно быть значение, указанное до @, во второй — значение, указанное после @.
```sql
  SUBSTRING_INDEX(email, '@', 1) AS username,
  SUBSTRING_INDEX(email, '@', -1) AS domain
FROM customer;
```
![txt](img/{0E405881-D5C1-47A5-BD1C-41F79E1D044D}.png)
Задание 6*
Доработайте запрос из предыдущего задания, скорректируйте значения в новых колонках: первая буква должна быть заглавной, остальные — строчными.
```sql
SELECT
  CONCAT(
    UPPER(LEFT(SUBSTRING_INDEX(email, '@', 1), 1)),
    LOWER(SUBSTRING(SUBSTRING_INDEX(email, '@', 1), 2))
  ) AS username_formatted,
  CONCAT(
    UPPER(LEFT(SUBSTRING_INDEX(email, '@', -1), 1)),
    LOWER(SUBSTRING(SUBSTRING_INDEX(email, '@', -1), 2))
  ) AS domain_formatted
FROM customer;
```
![txt](img/{6E402841-FE55-441F-AE3C-B04A7E9DA894}.png
)
