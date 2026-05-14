package main

import "notepad-minus-minus/dll"

type Character struct {
	Rune rune
}

type Line = dll.DoubleLinkedList[Character]

func NewLine() *Line {
	return &Line{}
}
