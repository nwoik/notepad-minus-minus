package main

import "notepad-minus-minus/dll"

type Line struct {
	Characters *dll.DoubleLinkedList[Character]
}

func NewLine() Line {
	return Line{Characters: dll.NewDLL[Character]()}
}
