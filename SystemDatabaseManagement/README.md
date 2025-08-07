# Домашнее задание к занятию "База данных" - Конкин Дмитрий
![[DZ.png]]

*
```sql
-- Удаление таблиц (если уже существуют, для чистоты)
DROP TABLE IF EXISTS Employee_Project_Assignment;
DROP TABLE IF EXISTS Projects;
DROP TABLE IF EXISTS Employees;
DROP TABLE IF EXISTS Departments;
DROP TABLE IF EXISTS Branches;

-- Таблица: Branches (Филиалы)
CREATE TABLE Branches (
    branch_id SERIAL PRIMARY KEY,
    branch_address TEXT NOT NULL UNIQUE
);

-- Таблица: Departments (Структурные подразделения)
CREATE TABLE Departments (
    department_id SERIAL PRIMARY KEY,
    department_type VARCHAR(50) NOT NULL,
    department_name VARCHAR(255) NOT NULL,
    UNIQUE (department_type, department_name)
);

-- Таблица: Employees (Сотрудники)
CREATE TABLE Employees (
    employee_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    salary NUMERIC(10, 2) NOT NULL CHECK (salary >= 0),
    hire_date DATE NOT NULL,
    department_id INTEGER NOT NULL,
    branch_id INTEGER NOT NULL,
    FOREIGN KEY (department_id) REFERENCES Departments(department_id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES Branches(branch_id) ON DELETE CASCADE
);

-- Таблица: Projects (Проекты)
CREATE TABLE Projects (
    project_id SERIAL PRIMARY KEY,
    project_name TEXT NOT NULL UNIQUE
);

-- Таблица: Employee_Project_Assignment (Назначения на проекты)
CREATE TABLE Employee_Project_Assignment (
    assignment_id SERIAL PRIMARY KEY,
    employee_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES Employees(employee_id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES Projects(project_id) ON DELETE CASCADE,
    UNIQUE (employee_id, project_id)  -- Один сотрудник — один раз на проект
);
```