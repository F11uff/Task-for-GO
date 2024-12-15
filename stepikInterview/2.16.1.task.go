package main

/*
Write the functions produce and main.

Produce function:
- receives the pipe channel as input
- writes integers to the pipe endlessly, starting from 0
- on a signal from main should terminate work

When finished, it should fall asleep for 3 seconds, then print “produce finished”

Main function:

- should create a pipe channel
- run produceCount functions produce and start reading from the pipe, printing each number
- when receiving a number, produceStop from pipe should stop reading new numbers from the pipe and should send
- a signal in the produce function that completes their work
should wait for all "produce finished" messages and print "main finished"

To implement the requirements, you can add additional arguments to the produce function
*/

const (
	produceCount = 3
	produceStop  = 10
)

func produce(pipe chan<- int) { // допускается добавить доп. аргументы
	// напишите свой код здесь
}
func main() {
	// напишите свой код здесь
}
