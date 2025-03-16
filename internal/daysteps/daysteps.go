package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

const (
	StepLength = 0.65 // длина шага в метрах
)

var (
	ErrConvDuration = errors.New("ошибка преобразования времени тренировки")
	ErrConvSteps    = errors.New("ошибка преобразования шагов")
	ErrCountSteps   = errors.New("ошибка количество шагов меньше или равно 0")
	ErrSplitStr     = errors.New("ошибка строка не соответствует формату: 678,0h50m")
)

func parsePackage(data string) (int, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 2 {
		return 0, 0, ErrSplitStr
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, ErrConvSteps
	}
	if steps <= 0 {
		return 0, 0, ErrCountSteps
	}

	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, ErrConvDuration
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * StepLength / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	fSteps := fmt.Sprintf("Количество шагов: %d.\n", steps)
	fDistance := fmt.Sprintf("Дистанция составила: %.2f км.\n", distance)
	fCalories := fmt.Sprintf("Вы сожгли: %.2f ккал.\n", calories)

	return fmt.Sprint(fSteps, fDistance, fCalories)
}
