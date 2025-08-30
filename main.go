package main

import (
	"fmt"
	taskone "web-three-task/go/task-one"
)

func main() {
	fmt.Println(taskone.RemoveDuplicates([]int{1, 2, 3, 3, 4}))
}
