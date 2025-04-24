package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Недопустимый формат даннных")
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("Недопустимый формат шагов")
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов меньше нуля")
	}
	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка продолжительности")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceMeters := float64(steps) * stepLength
	distanceKM := distanceMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка подсчета калорий ", err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d. \nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKM, calories)
}
