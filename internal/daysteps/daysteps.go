package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// проверяем длину слайса
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат данных: ожидается 2 элемента")
	}
	// преобразуем кол-во шагов и проверяем на ошибку
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, errors.New("не удалось преобразовать кол-во шагов в число")
	}
	// проверка шагов
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов 0")
	}
	// преобразуем время и проверяем на ошибку
	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, errors.New("не удалось преобразовать время")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, _, err := parsePackage(data)
	//проверки
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	// применяем формулы
	distanceInMeters := float64(steps) * stepLength
	distanceInKm := distanceInMeters / mInKm
	calories := WalkingSpentCalories(steps, weight)
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceInKm, calories)

	return result
}
