package labs_1

import (
	"fmt"
	"time"
)

func timeModule() {
	fmt.Println("\nTime module:")

	currentTime := time.Now()

	fmt.Println("Текущая дата и время:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Дата:", currentTime.Format("2006-01-02"))
	fmt.Println("Время:", currentTime.Format("15:04:05"))
	fmt.Println("Дата (дд.мм.гггг):", currentTime.Format("02.01.2006"))
	fmt.Println("Дата с названием месяца:", currentTime.Format("02 January 2006"))
	fmt.Println("Дата с часовым поясом:", currentTime.Format("2006-01-02 15:04:05 MST"))
}
