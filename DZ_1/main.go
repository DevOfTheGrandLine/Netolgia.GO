package main

import (
	"fmt"
	"time"
)

// Глобальные константы для категорий товаров
const (
	CategoryElectronics = "Электроника"
	CategoryFood        = "Продукты"
	CategoryClothes     = "Одежда"
	MaxItems            = 100 // Максимальное количество товаров на складе
)

// Глобальная переменная для подсчета товаров
var totalItems int

// Счётчик ID для новых товаров
var currentID int64 = 1

// Структура товара
type Item struct {
	ID        int64
	Name      string
	Price     float64
	Quantity  int
	Available bool
	Category  string
	Color     string
	Weight    float32
	AddedDate time.Time
}

func main() {
	// Объявление и инициализация переменных разными способами

	// Способ 1: Объявление переменной с явным указанием типа и последующим присваиванием
	var itemName string
	itemName = "Смартфон"

	// Способ 2: Объявление переменной с явным указанием типа и инициализацией
	var itemPrice float64 = 999.99

	// Способ 3: Объявление нескольких переменных одного типа
	var minQuantity, maxQuantity int = 5, 20

	// Способ 4: Объявление переменной без явного указания типа (тип определяется компилятором)
	var isAvailable = true

	// Способ 5: Короткое объявление переменной (только внутри функций)
	quantity := 15

	// Способ 6: Объявление нескольких переменных разных типов одновременно
	var (
		itemID     int64  = 12345
		itemColor  string = "Черный"
		itemWeight float32
		dateAdded  time.Time = time.Now()
	)

	// Присваивание значения ранее объявленной переменной
	itemWeight = 0.3

	// Приведение типов
	percentInStock := float64(quantity) / float64(MaxItems) * 100

	// Использование констант
	category := CategoryElectronics

	// Вызов функции для вывода информации о товаре
	displayItemInfo(itemID, itemName, quantity, itemPrice, isAvailable, category)

	// Демонстрация работы с базовыми типами
	fmt.Println("\nДополнительная информация:")
	fmt.Printf("Цвет товара: %s\n", itemColor)
	fmt.Printf("Вес товара: %.2f кг\n", itemWeight)
	fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", minQuantity, maxQuantity)
	fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", percentInStock)
	fmt.Printf("Дата добавления: %s\n", dateAdded.Format("02-01-2006"))

	// Работа с целочисленными константами и их типами
	const (
		Small  = 1
		Medium = 5
		Large  = 10
	)

	// Демонстрация работы с различными числовыми типами
	var (
		shortValue   int8      = 127 // -128 до 127
		intValue     int       = 1000000
		uintValue    uint      = 10000 // только положительные
		floatValue   float32   = 123.456
		complexValue complex64 = 1 + 2i
	)

	fmt.Println("\nРазные числовые типы:")
	fmt.Printf("int8: %d\n", shortValue)
	fmt.Printf("int: %d\n", intValue)
	fmt.Printf("uint: %d\n", uintValue)
	fmt.Printf("float32: %f\n", floatValue)
	fmt.Printf("complex64: %v\n", complexValue)

	// Работа со строками и байтами
	productCode := "ЯЩИК-12345"
	fmt.Printf("\nКод товара: %s, длина: %d символов, длина: %d байт\n", productCode, len([]rune(productCode)), len(productCode))

	// Обновление общего количества товаров
	updateTotalItems(quantity)
	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)

	// Добавление нового товара
	newItem := addNewItem("Наушники", 499.99, 30, CategoryElectronics, "Белый", 0.1)
	fmt.Printf("\nНовый товар добавлен: %s (ID: %d)\n", newItem.Name, newItem.ID)

	// Расчёт цены со скидкой
	discountedPrice := calculateDiscount(newItem, 15)
	fmt.Printf("Цена со скидкой: %.2f руб.\n", discountedPrice)
}

// Функция для отображения информации о товаре
func displayItemInfo(id int64, name string, qty int, price float64, isAvailable bool, category string) {
	fmt.Println("=== Информация о товаре ===")
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Название: %s\n", name)
	fmt.Printf("Категория: %s\n", category)
	fmt.Printf("Количество: %d\n", qty)
	fmt.Printf("Цена: %.2f руб.\n", price)

	// Использование условного оператора с логическим типом
	if isAvailable {
		fmt.Println("Статус: В наличии")
	} else {
		fmt.Println("Статус: Нет в наличии")
	}
}

// Функция для обновления общего количества товаров
func updateTotalItems(qty int) {
	totalItems += qty

	// Проверка на превышение максимального количества
	if totalItems > MaxItems {
		fmt.Println("Предупреждение: Превышено максимальное количество товаров на складе!")
		totalItems = MaxItems
	}
}

// Функция для добавления нового товара
func addNewItem(name string, price float64, quantity int, category string, color string, weight float32) Item {
	// Генерация нового ID
	currentID++

	// Создание нового товара
	newItem := Item{
		ID:        currentID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		Available: quantity > 0,
		Category:  category,
		Color:     color,
		Weight:    weight,
		AddedDate: time.Now(),
	}

	// Обновление общего количества товаров
	updateTotalItems(quantity)

	return newItem
}

// Функция для расчёта скидки на товар
func calculateDiscount(item Item, discountPercent float64) float64 {
	if discountPercent < 0 {
		discountPercent = 0
	} else if discountPercent > 100 {
		discountPercent = 100
	}

	return item.Price - (item.Price * discountPercent / 100)
}
