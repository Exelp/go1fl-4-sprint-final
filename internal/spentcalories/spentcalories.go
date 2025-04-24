package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Недопустимый формат данных")
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("Недопустимый формат шагов")
	}
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка продолжительности")
	}
	activity := strings.TrimSpace(parts[1])
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distanceMeters := (height * stepLengthCoefficient) * float64(steps)
	distanceKM := distanceMeters / mInKm
	return distanceKM
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKM := distance(steps, height)
	speed := distanceKM / duration.Hours()
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	if steps <= 0 || weight <= 0 || height <= 0 {
		return "", fmt.Errorf("Количество шагов, веса или продолжительности меньше нуля")
	}
	var distan, speed, calories float64
	switch strings.ToLower(activity) {
	case "Ходьба":
		distan = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		distan = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("Неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", activity, duration.Hours(), distan, speed, calories), nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Количество шагов, веса, роста или продолжительности меньше нуля")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Количество шагов, веса, роста или продолжительности меньше нуля")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH * walkingCaloriesCoefficient
	return calories, nil
}
