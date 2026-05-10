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

func OpenFile(path string) (*dll.DoubleLinkedList[Line], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	text := dll.NewDLL[Line]()

	for scanner.Scan() {
		scannerText := scanner.Text()
		line := NewLine()
		text.Add(line)
		for _, char := range scannerText {
			character := Character{Rune: char}
			line.Characters.Add(character)
		}
		line.Characters.Insert(line.Characters.Start(), Character{Rune: 32345})

		line.Characters.PrintElements()
	}

	return text, nil
}
