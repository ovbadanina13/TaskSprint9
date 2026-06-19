package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	// проверка размера на равенство 0, в этом случае вернем пустой слайс
	if size <= 0 {
		return []int{}
	}

	// Инициализируем глобальный генератор случайных чисел
	rand.Seed(time.Now().UnixNano())
	// Создаем слайс заданного размера
	result := make([]int, size)

	// Заполняем слайс
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(100)
	}
	return result

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	// Слайс пустой или равен nil
	if len(data) == 0 {
		// Возвращает минимально возможное значение типа int
		return math.MinInt
	}
	// Инициализация max значением первого элемента слайса ( проверка выше, что слайс не пустой)
	max := data[0]

	// Перебирвем слайс со второго элемента
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	// Возвращаем максимум
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	// Проверка массива, что не пустой и не равен nil
	if len(data) == 0 {
		return math.MinInt
	}
	// Вычисляем размер каждого среза и делим на количество частей
	chunkSize := len(data) / CHUNKS

	// Создаем слай для хранения максимумов каждой части
	maxes := make([]int, CHUNKS)
	for i := range maxes {
		maxes[i] = math.MinInt
	}

	// Создаем переменную типа sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Интерация цикла создает горутину для обработки своей части слайса
	for i := 0; i < CHUNKS; i++ {
		// Увеличиваем счетчик на 1 перед запуском, чтобы сообщить, что мы ожидаем завершение 1 горутины
		wg.Add(1)

		// Вычисляем начальный индекс текущего среза
		start := i * chunkSize

		// Вычисляем конечный индекс текущего среза
		end := start + chunkSize

		// Захватываем все элементы до конца слайса, если размер части = 0 или мы на последней итерации

		if i == CHUNKS-1 || chunkSize == 0 {
			end = len(data)
		}

		// Если start > end , значит часть пустая, горутина завершится ничего не делая
		if start >= end {
			wg.Done()
			continue
		}

		// Запускаем анонимную функцию асинхронно
		// Передаем в горутину необходимые параметры: срез данных, индекс для записи результата, начальный и конечный индексы среза

		go func(data []int, index int, start int, end int) {
			defer wg.Done()

			// Инициализируем локальный максимум первым элементом текущего среза
			localMax := data[start]

			// Поиск максимума по всем элементам текущего среза
			for j := start + 1; j < end; j++ {
				if data[j] > localMax {
					localMax = data[j]
				}
			}

			// Запись максимума в слайс
			maxes[index] = localMax

		}(data, i, start, end)
	}

	wg.Wait()

	// Слайс maxes заполнен 8 значениями, ищем максимум
	finalMax := math.MinInt
	for i := 0; i < len(maxes); i++ {
		if maxes[i] > finalMax {
			finalMax = maxes[i]
		}
	}
	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	// записываем момент времени перед началом выполнения функции maximum
	startSingle := time.Now()

	// Вызываем функцию maximum для поиска максимума в слайсе data в один поток
	max := maximum(data)

	// Вычисляем прошедшее время, вычитая момент начала из текущего момента
	elapsedSingle := time.Since(startSingle)

	elapsed := elapsedSingle.Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	startMulti := time.Now()
	// Вызываем функцию maxChunks для параллельного поиска максимума в слайсе data
	max = maxChunks(data)

	elapsedMulti := time.Since(startMulti)
	elapsed = elapsedMulti.Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

}
