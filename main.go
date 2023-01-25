package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	before *Task
	after  *Task
	task   interface{}
}

type List struct {
	head *Task
	tail *Task
}

const (
	OPTION_VIEWS  = 1
	OPTION_CREATE = 2
	OPTION_UPDATE = 3
	OPTION_DELETE = 4
	OPTION_EXIT   = 5
)

func main() {
	menu()
}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	link := List{}
	for {
		fmt.Println("Drag & Drop Sortable Task List")
		fmt.Println("----------------------------------")
		fmt.Println("1. Views Task")
		fmt.Println("2. Create Task")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose menu do you want (number) ? ")
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		option, err := strconv.Atoi(text)
		fmt.Println()
		if err != nil {
			fmt.Println("Inputan harus berupa angka!")
			fmt.Println()
			menu()
		}

		// Option menu
		switch option {
		case OPTION_VIEWS:
			var head interface{}
			var tail interface{}
			if link.head != nil {
				head = link.head.task
			}
			if link.tail != nil {
				tail = link.tail.task
			}

			fmt.Println("==============================")
			fmt.Printf("Head: %v\n", head)
			fmt.Printf("Tail: %v\n", tail)
			fmt.Println("==============================\n")
			link.Display()
			link.OrderDisplay()
		case OPTION_CREATE:
			fmt.Print("Insert task	: ")
			task, _ := reader.ReadString('\n')
			// convert CRLF to LF
			task = strings.Replace(task, "\n", "", -1)
			fmt.Println()

			if err != nil {
				fmt.Println("Inputan harus berupa angka!")
				fmt.Println()
				menu()
			}

			link.Insert(task)
		case OPTION_UPDATE:
			fmt.Println("UPDATE")

		case OPTION_DELETE:
			fmt.Println("DELETE")
		case OPTION_EXIT:
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Option tidak tersedia!")
		}
	}
}

func (L *List) Insert(task interface{}) {
	list := &Task{
		after: L.head,
		task:  task,
	}
	if L.head != nil {
		L.head.before = list
	}
	L.head = list

	l := L.head
	for l.after != nil {
		l = l.after
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head

	for list != nil {
		if list.after == nil {
			fmt.Printf("%v \n", list.task)
			fmt.Println()
		} else {
			fmt.Printf("%v -> ", list.task)
		}
		list = list.after
	}
}

func (l *List) OrderDisplay() {
	list := l.head

	count := 1
	for list != nil {
		fmt.Printf("%d. %v \n", count, list.task)
		count++
		list = list.after
	}
	fmt.Println()
}

func (l *List) Update() {
	curr := l.head
	var before *Task
	l.tail = l.head

	for curr != nil {
		after := curr.after
		curr.after = before
		before = curr
		curr = after
	}
	l.head = before
}
