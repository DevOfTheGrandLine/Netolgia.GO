package main

import "fmt"

type Employee struct {
	LastName  string
	FirstName string
	Age       int
	Position  string
	Salary    int
}

type Displayable interface {
	Display()
}

func (e Employee) Display() {
	fmt.Printf("Сотрудник: %s %s | Возраст: %d | Должность: %s | Зарплата: %d руб.\n",
		e.FirstName, e.LastName, e.Age, e.Position, e.Salary)
}

func FilterEmployees(employees []Employee, minAge int, minSalary int) []Displayable {
	var result []Displayable

	for _, emp := range employees {
		if emp.Age >= minAge && emp.Salary >= minSalary {
			result = append(result, emp)
		}
	}

	return result
}

func main() {
	// Инициализация списка сотрудников
	employees := []Employee{
		{"Иванов", "Иван", 30, "Разработчик", 120000},
		{"Петров", "Пётр", 45, "Менеджер", 90000},
		{"Сидорова", "Анна", 28, "Тестировщик", 95000},
		{"Кузнецова", "Ольга", 50, "Директор", 150000},
		{"Николаев", "Николай", 35, "Аналитик", 110000},
	}

	// Параметры фильтрации
	minAge := 30
	minSalary := 100000

	// Фильтрация и вывод
	filtered := FilterEmployees(employees, minAge, minSalary)

	// Вывод результатов
	fmt.Println("=== Отфильтрованные сотрудники ===")
	for _, d := range filtered {
		d.Display()
	}
}
