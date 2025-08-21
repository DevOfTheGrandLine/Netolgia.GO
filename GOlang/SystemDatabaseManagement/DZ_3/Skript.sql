-- Задание 1
SELECT DISTINCT district
FROM address
WHERE district LIKE 'K%a'
  AND district NOT LIKE '% %';

-- Задание 2
SELECT *
FROM payment
WHERE DATE(payment_date) BETWEEN '2005-06-15' AND '2005-06-18'
  AND amount > 10.00;

-- Задание 3
SELECT *
FROM rental
ORDER BY rental_date DESC
LIMIT 5;

-- Задание 4
SELECT
    LOWER(REPLACE(first_name, 'LL', 'pp')) AS first_name,
    LOWER(last_name) AS last_name
FROM customer
WHERE active = 1
  AND UPPER(first_name) IN ('KELLY', 'WILLIE');

-- Задание 5*
SELECT
  SUBSTRING_INDEX(email, '@', 1) AS username,
  SUBSTRING_INDEX(email, '@', -1) AS domain
FROM customer;

-- Задание 6*
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