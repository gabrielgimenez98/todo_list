package main

import (
	"bufio"
	"fmt"
	"os"
)

type task struct {
	title  string
	detail string
	status int
}

var task_list []task

func main() {
	opt := 0
	for opt != 9 {
		fmt.Println("escolha sua opção:")
		fmt.Println("1 - Listar tarefas")
		fmt.Println("2 - Adicionar tarefa")
		fmt.Println("3 - Concluir tarefa")
		fmt.Println("4 - Excluir tarefa")
		fmt.Println("9 - Sair")

		fmt.Scan(&opt)

		handle_option(opt)

	}

}

func handle_option(option int) {
	switch option {
	case 1:
		show_tasks()
	case 2:
		handle_add_task()
	case 3:
		handle_finish_task()
	case 4:
		handle_exclude_task()
	case 9:
		fmt.Println("Saindo.")
	default:
		fmt.Println("Tente outra opção.")
	}
}

func handle_add_task() {
	var title string
	var description string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Titulo da tarefa")
	title, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	fmt.Println("Descrição da tarefa")
	description, err = reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	add_task(task{title, description, 0})

}

func add_task(new_task task) {
	task_list = append(task_list, new_task)
}

func show_tasks() {
	fmt.Println("Suas tarefas:")
	for i, task := range task_list {
		fmt.Println(i, " - ", task.title)
	}
}

func handle_finish_task() {
	var title string
	fmt.Println("qual titulo da tarefa concluida?")
	reader := bufio.NewReader(os.Stdin)

	title, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	finish_task((title))

}

func finish_task(title string) {
	for i := 0; i < len(task_list); i++ {
		if task_list[i].title == title {
			task_list[i].status = 1
			fmt.Println("tarefa concluida")
			break

		}
	}

}

func handle_exclude_task() {
	var title string
	fmt.Println("qual titulo da tarefa a ser excluida?")
	reader := bufio.NewReader(os.Stdin)

	title, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	exclude_task((title))

}

func exclude_task(title string) {
	for i := 0; i < len(task_list); i++ {
		if task_list[i].title == title {
			task_list[i].status = 2
			fmt.Println("tarefa excluida")
			break

		}
	}

}
