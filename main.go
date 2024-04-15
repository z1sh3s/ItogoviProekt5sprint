package main

import (
	"fmt"
	"math"
	"time"
)

// Общие константы для вычислений.
const (
	MInKm      = 1000 // количество метров в одном километре
	MinInHours = 60   // количество минут в одном часе
	LenStep    = 0.65 // длина одного шага
	CmInM      = 100  // количество сантиметров в одном метре
)

// Training общая структура для всех тренировок
type Training struct {
	TrainingType string        // тип тренировки
	Action       int           // количество повторов(шаги, гребки при плавании)
	LenStep      float64       // длина одного шага или гребка в м
	Duration     time.Duration // продолжительность тренировки
	Weight       float64       // вес пользователя в кг
}

// distance возвращает дистанцию, которую преодолел пользователь.
// Формула расчета:
// количество_повторов * длина_шага / м_в_км
func (t Training) distance() float64 {
	// вставьте ваш код ниже
	calcDistance := float64(t.Action) * (t.LenStep) / float64(MInKm)
	return calcDistance
}

// meanSpeed возвращает среднюю скорость бега или ходьбы.
func (t Training) meanSpeed() float64 {
	// вставьте ваш код ниже
	calcMeanSpeed := (float64(t.Action) * (t.LenStep / MInKm)) / t.Duration.Hours()
	return calcMeanSpeed
}

// Calories возвращает количество потраченных килокалорий на тренировке.
// Пока возвращаем 0, так как этот метод будет переопределяться для каждого типа тренировки.
func (t Training) Calories() float64 {
	// вставьте ваш код ниже
	return 0
}

// InfoMessage содержит информацию о проведенной тренировке.
type InfoMessage struct {
	// добавьте необходимые поля в структуру
	TrainingType string        // тип тренировки
	Duration     time.Duration // длительность тренировки
	Distance     float64       // расстояние, которое преодолел пользователь
	Speed        float64       // средняя скорость, с которой двигался пользователь
	Calories     float64       // количество потраченных килокалорий на тренировке
}

// TrainingInfo возвращает труктуру InfoMessage, в которой хранится вся информация о проведенной тренировке.
func (t Training) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return InfoMessage{TrainingType: t.TrainingType, Duration: t.Duration, Distance: t.distance(), Speed: t.meanSpeed()}
}

// String возвращает строку с информацией о проведенной тренировке.
func (i InfoMessage) String() string {
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n",
		i.TrainingType,
		i.Duration.Minutes(),
		i.Distance,
		i.Speed,
		i.Calories,
	)
}

// CaloriesCalculator интерфейс для структур: Running, Walking и Swimming.
type CaloriesCalculator interface {
	// добавьте необходимые методы в интерфейс
	Calories() float64
	TrainingInfo() InfoMessage
}

// Константы для расчета потраченных килокалорий при беге.
const (
	CaloriesMeanSpeedMultiplier = 18   // множитель средней скорости бега
	CaloriesMeanSpeedShift      = 1.79 // коэффициент изменения средней скорости
)

// Running структура, описывающая тренировку Бег.
type Running struct {
	// добавьте необходимые поля в структуру
	Training
}

// Calories возввращает количество потраченных килокалория при беге.
// Формула расчета:
// ((18 * средняя_скорость_в_км/ч + 1.79) * вес_спортсмена_в_кг / м_в_км * время_тренировки_в_часах * мин_в_часе)
// Это переопределенный метод Calories() из Training.
func (r Running) Calories() float64 {
	// вставьте ваш код ниже
	calcCaloriesRun := ((CaloriesMeanSpeedMultiplier*r.meanSpeed() + CaloriesMeanSpeedShift) * r.Weight / MInKm * r.Duration.Hours() * MinInHours)
	return calcCaloriesRun
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (r Running) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return r.Training.TrainingInfo()
}

// Константы для расчета потраченных килокалорий при ходьбе.
const (
	CaloriesWeightMultiplier      = 0.035 // коэффициент для веса
	CaloriesSpeedHeightMultiplier = 0.029 // коэффициент для роста
	KmHInMsec                     = 0.278 // коэффициент для перевода км/ч в м/с
)

// Walking структура описывающая тренировку Ходьба
type Walking struct {
	// добавьте необходимые поля в структуру
	Training
	Height float64 // рост пользователя
}

// Calories возвращает количество потраченных килокалорий при ходьбе.
// Формула расчета:
// ((0.035 * вес_спортсмена_в_кг + (средняя_скорость_в_метрах_в_секунду**2 / рост_в_метрах)
// * 0.029 * вес_спортсмена_в_кг) * время_тренировки_в_часах * мин_в_ч)
// Это переопределенный метод Calories() из Training.
func (w Walking) Calories() float64 {
	// вставьте ваш код ниже
	Speed := w.meanSpeed() * KmHInMsec
	calcCaloriesWalk := ((CaloriesWeightMultiplier*w.Weight + (math.Pow(Speed, 2)/w.Height)*CaloriesSpeedHeightMultiplier*w.Weight) * w.Training.Duration.Hours() * MinInHours)
	return calcCaloriesWalk
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (w Walking) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return w.Training.TrainingInfo()
}

// Константы для расчета потраченных килокалорий при плавании.
const (
	SwimmingLenStep                  = 1.38 // длина одного гребка
	SwimmingCaloriesMeanSpeedShift   = 1.1  // коэффициент изменения средней скорости
	SwimmingCaloriesWeightMultiplier = 2    // множитель веса пользователя
)

// Swimming структура, описывающая тренировку Плавание
type Swimming struct {
	// добавьте необходимые поля в структуру
	Training
	LengthPool int // длина бассейна
	CountPool  int // количество пересечений бассейна
}

// meanSpeed возвращает среднюю скорость при плавании.
// Формула расчета:
// длина_бассейна * количество_пересечений / м_в_км / продолжительность_тренировки
// Это переопределенный метод Calories() из Training.
func (s Swimming) meanSpeed() float64 {
	// вставьте ваш код ниже
	calcMeanSpeedSwimm := s.LengthPool * s.CountPool / MInKm / int(s.Duration)
	return float64(calcMeanSpeedSwimm)
}

// Calories возвращает количество калорий, потраченных при плавании.
// Формула расчета:
// (средняя_скорость_в_км/ч + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * вес_спортсмена_в_кг * время_тренировки_в_часах
// Это переопределенный метод Calories() из Training.
func (s Swimming) Calories() float64 {
	// вставьте ваш код ниже
	calcCaloriesSwimm := (s.meanSpeed() + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * s.Weight * s.Training.Duration.Hours()
	return calcCaloriesSwimm
}

// TrainingInfo returns info about swimming training.
// Это переопределенный метод TrainingInfo() из Training.
func (s Swimming) TrainingInfo() InfoMessage {
	// вставьте ваш код ниже
	return s.Training.TrainingInfo()
}

// ReadData возвращает информацию о проведенной тренировке.
func ReadData(training CaloriesCalculator) string {
	// получите количество затраченных калорий

	// получите информацию о тренировке
	info := training.TrainingInfo()
	// добавьте полученные калории в структуру с информацией о тренировке
	info.Calories = training.Calories()
	return fmt.Sprint(info)
}

func main() {

	swimming := Swimming{
		Training: Training{
			TrainingType: "Плавание",
			Action:       2000,
			LenStep:      SwimmingLenStep,
			Duration:     90 * time.Minute,
			Weight:       85,
		},
		LengthPool: 50,
		CountPool:  5,
	}

	fmt.Println(ReadData(swimming))

	walking := Walking{
		Training: Training{
			TrainingType: "Ходьба",
			Action:       20000,
			LenStep:      LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85,
		},
		Height: 185,
	}

	fmt.Println(ReadData(walking))

	running := Running{
		Training: Training{
			TrainingType: "Бег",
			Action:       5000,
			LenStep:      LenStep,
			Duration:     30 * time.Minute,
			Weight:       85,
		},
	}

	fmt.Println(ReadData(running))

}
