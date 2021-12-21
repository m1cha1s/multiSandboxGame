package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Object struct {
	id         int
	properties []string
	position   rl.Vector2
}

func dumpObject(obj *Object) {
	fmt.Printf("Object ID: %d", obj.id)
	fmt.Println("Object properties: ")

	for _, property := range obj.properties {
		fmt.Printf("         - %s", property)
	}

	fmt.Printf("Object position: %f, %f", obj.position.X, obj.position.Y)
}

func getPositionOnGrid(obj *Object) (int, int) {
	return int(obj.position.X), int(obj.position.Y)
}
