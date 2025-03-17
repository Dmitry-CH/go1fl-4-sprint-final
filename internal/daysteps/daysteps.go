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
	stepLength = 0.65 // длина шага в метрах
	mInKm      = 1000 // количество метров в километре.
)

var (
	errConvDuration = errors.New("invalid duration conversion")
	errConvSteps    = errors.New("invalid step conversion")
	errCountSteps   = errors.New("count of steps must be greater than 0")
	errSplitStr     = errors.New("string does not match the format: 678,0h50m")
)

func parsePackage(data string) (int, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 2 {
		return 0, 0, errSplitStr
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, errConvSteps
	}
	if steps <= 0 {
		return 0, 0, errCountSteps
	}

	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, errConvDuration
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
		fmt.Println(errCountSteps.Error())
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.\n", steps, distance, calories)
}
