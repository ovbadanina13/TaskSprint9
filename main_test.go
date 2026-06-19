package main

import (
	"math"
	"testing"
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
			if len(actualSlice) != tc.expectedLen {
				t.Errorf("generateRandomElements(%d) length = %d, want %d", tc.size, len(actualSlice), tc.expectedLen)
			}

			// 2. Проверяем, что все элементы находятся в ожидаемом диапазоне [0, 100)
			for i, v := range actualSlice {
				if v < 0 || v >= 100 {
					t.Errorf("generateRandomElements(%d)[%d] = %d, want in range [0, 100)", tc.size, i, v)
				}
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
		{"Пустой слайс", []int{}, math.MinInt},
		{"Nil слайс", nil, math.MinInt},

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
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, result, tt.expected)
			}
		})
	}
}
