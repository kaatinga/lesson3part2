package main

import (
	"fmt"
	"log"
)

var queue []string     // очередь тут
var messages [3]string // сообщения пользователю тут
var action byte        // хранение пользовательского действия

// LIFO implementation

func push(newString string) {
	p := log.Println // the alias for log.Println in order to simplify the code

	queue = append(queue, newString)

	p("--------------------------------")
	p("Добавлено сообщение в очередь:", newString)
	p("--------------------------------")

}

func pop() (takenString string) {
	takenString = queue[0]
	queue = queue[1:]
	return
}

// UserInput implements user data input
func userInput() (input string) {
	p := log.Println // the alias for log.Println in order to simplify the code

	p("Please, enter a string:")
	for {
		fmt.Scan(&input)
		if input != "" {
			p("The data input is done")
			break
		} else {
			p("The data input has been wrong. Please, repeat:")
		}
	}
	return
}

func main() {

	p := log.Println // the alias for log.Print in order to simplify the code

	p("Welcome to a LIFO demo")

	messages = [3]string{
		"Добавить сообщение в очередь",
		"Забрать сообщение из очереди",
		"Просмотр сообщений в очереди",
	}

	for {
		p("--------------------------------")
		p("Список доступных действий:")
		for key, value := range messages {
			p(key+1, "-", value)
		}
		p("Определите ваше действие:")
		fmt.Scan(&action)
		switch action {
		case 1:
			push(userInput())
		case 2:
			p("Сообщение типа забрано из очереди:", pop())
		case 3:
			{
				p("--------------------------------")
				if len(queue) != 0 {
					p("Список сообщений в очереди:")
					for key, value := range queue {
						p(key, "-", value)
					}
				} else {
					p("--------------------------------")
					p("Очередь не содержит сообщений")
				}
			}
		default:
			p("Неверный ввод. Повторите пожалуйста")
		}
	}
}
