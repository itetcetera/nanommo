package main

import (
	"fmt"
	"time"
)

const (
	MAX_CLIENTS    = 100 //Столько клиентов мы готовы обслуживать одновременно
	MAX_FPS        = 60
	FRAME_DURATION = time.Second / MAX_FPS
)

var characters map[string]*Character

func updateCharacters(k float64) {
	for _, c := range characters {
		c.update(k)
	}
}
func mainLoop() {
	// Мы хотим чтобы персонажи двигались в не зависимости от скорости железа и
	// загруженности системы.
	// Поэтому привязываем движение объектов ко времени при помощи этого коэфицента
	var k float64
	for {
		frameStart := time.Now()

		updateCharacters(k)

		duration := time.Now().Sub(frameStart)
		// Если кадр просчитался быстрее, чем необходимо подождем оставшееся время
		if duration > 0 && duration < FRAME_DURATION {
			time.Sleep(FRAME_DURATION - duration)
		}
		ellapsed := time.Now().Sub(frameStart)
		// Коэфицент это отношение времени, потраченного на обработку одного кадра к секунде
		// Время в гоу измеряется в наносекундах
		// time.Second это количество наносекунд в секунде
		k = float64(ellapsed) / float64(time.Second)
	}
}

func main() {
	characters = make(map[string]*Character, MAX_CLIENTS)
	fmt.Println("Server started at ", time.Now())

	// Запускаем обработчик вебсокетов
	go NanoHandler()
	mainLoop()
}
