package main

import "fmt"

func main() {
	node1, node2 := NewNode("n1"), NewNode("n2")
	node1.AddLabel("Person")
	node2.AddLabel("Person")

	node1.SetProperty("Age", 21)
	node2.SetProperty("Age", 25)

	fmt.Println(node1.GetProperty("Age"))

	brothers := NewRelationship("rel1", "BROTHERS", node1, node2)
	brothers.SetProperty("talking", true)

	fmt.Println(brothers.GetProperty("talking"))
}
