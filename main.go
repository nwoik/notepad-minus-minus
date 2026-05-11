package main

import (
	"bufio"
	"notepad-minus-minus/dll"
	"os"
)

func main() {
	OpenFile("text.txt")

	// a := app.New()
	// w := a.NewWindow("Hello World")

	// w.SetContent(widget.NewLabel("Hello World!"))
	// w.ShowAndRun()

	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	key, _ := reader.ReadString('\n')
	// 	println(key)
	// }
}

func OpenFile(path string) (*dll.DoubleLinkedList[*dll.DoubleLinkedList[Character]], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	text := dll.NewDLL[*dll.DoubleLinkedList[Character]]()

	for scanner.Scan() {
		scannerText := scanner.Text()
		line := dll.NewDLL[Character]()
		text.Add(line)
		for _, char := range scannerText {
			character := Character{Rune: char}
			line.Add(character)
		}

		line.PrintElements()
	}

	return text, nil
}
