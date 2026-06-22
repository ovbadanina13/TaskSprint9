package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateRandomElements проверяет генерацию случайных чисел
func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		name        string
		size        int
		expectedLen int
	}{
		{"Нулевой размер", 0, 0},
		{"Отрицательный размер", -10, 0},
		{"Положительный размер", 100, 100},
		{"Большой размер", 10000, 10000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualSlice := generateRandomElements(tc.size)

			// 1. Проверяем длину возвращаемого слайса
			assert.Equal(t, tc.expectedLen, len(actualSlice), "Длина слайса не совпадает для размера %d", tc.size)

			// 2. Проверяем, что все элементы больше 0
			for i, v := range actualSlice {
				assert.GreaterOrEqual(t, v, 0, "Элемент %d для размера %d оказался меньше нуля", i, tc.size)
			}
		})
	}
}

// TestMaximum проверяет поиск максимума в слайсе
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		// Крайние случаи с пустыми данными
		{"Пустой слайс", []int{}, 0},
		{"Nil слайс", nil, 0},

		// Слайсы с одним элементом
		{"Один положительный элемент", []int{42}, 42},
		{"Один отрицательный элемент", []int{-42}, -42},
		{"Один нулевой элемент", []int{0}, 0},

		// Расположение максимума в разных частях слайса
		{"Максимум в начале", []int{10, 5, 2}, 10},
		{"Максимум в конце", []int{1, 2, 10}, 10},
		{"Максимум в середине", []int{1, 10, 2}, 10},

		// Специфические наборы данных
		{"Все элементы одинаковы", []int{5, 5, 5}, 5},
		{"Все элементы отрицательные", []int{-10, -5, -20}, -5},

		// Экстремальные значения (проверка на переполнения и корректность инициализации)
		{"Содержит MinInt и MaxInt", []int{math.MinInt, 0, math.MaxInt}, math.MaxInt},
		{"Только MinInt", []int{math.MinInt}, math.MinInt},
		{"Только MaxInt", []int{math.MaxInt}, math.MaxInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			assert.Equal(t, tt.expected, result, "maximum(%v) вернул неверный результат", tt.data)
		})
	}
}
