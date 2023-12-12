package main

import "fmt"

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
		fmt.Println("Vamos listar.")
	case 2:
		fmt.Println("Vamos adicionar.")
	case 3:
		fmt.Println("Vamos concluir.")
	case 4:
		fmt.Println("Vamos excluir.")
	case 9:
		fmt.Println("Saindo.")
	default:
		fmt.Println("Tente outra opção.")
	}
}

func add_task(new_task task) {
	task_list = append(task_list, new_task)
}

func show_tasks() {
	for i := 0; i < len(task_list); i++ {
		fmt.Println(task_list[i])
	}
}
