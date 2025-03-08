package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Zheltyj/go1fl-4-sprint-final/internal/spentcalories"
)

const mInKm = 1000

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	splitedStr := strings.Split(data, ",")
	if len(splitedStr) != 2 {
		return 0, 0, errors.New("wrong daysteps data string")
	}

	steps, err := strconv.Atoi(splitedStr[0])
	if err != nil {
		return 0, 0, err
	}

	walkDur, err := time.ParseDuration(splitedStr[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, walkDur, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	daySteps, walkDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	distM := float64(daySteps) * StepLength
	distKm := distM / mInKm

	calories := spentcalories.WalkingSpentCalories(daySteps, weight, height, walkDuration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.\n", daySteps, distKm, calories)
}
