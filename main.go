package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

// Переменные для считывая флагов
var (
	wordsFlag = flag.Bool("words", false, "Показать количество слов")
	linesFlag = flag.Bool("lines", false, "Показать количество строк")
	charsFlag = flag.Bool("chars", false, "Показать самую частую букву")
)

func main() {
	log.SetFlags(0)
	// Запарсим флаги и прокинем их в слайс
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Не указан путь к файлу")
	}

	file := args[0]

	// Прочитаем файл в переменную
	content := readFile(file)

	if len(content) == 0 {
		log.Fatal("Файл пуст")
	}

	// Если флагов нет, считаем всё
	showAll := !(*wordsFlag || *linesFlag || *charsFlag)

	// проверяем флаги и выводим результат
	if showAll || *linesFlag {
		fmt.Printf("Количество строк: %d\n", countLines(content))
	}

	if showAll || *wordsFlag {
		fmt.Printf("Количество слов: %d\n", countWords(content))
	}

	if showAll || *charsFlag {
		letters, count := countFreqLet(content)
		if count == 0 {
			fmt.Println("Букв не найдено")
		} else {
			if len(letters) == 1 {
				fmt.Printf("Самая частая буква: %c (%d раз)\n", letters[0], count)
			} else {
				fmt.Printf("Самые частые буквы: ")
				for i, letter := range letters {
					if i > 0 {
						fmt.Print(", ")
					}
					fmt.Printf("%c", letter)
				}
				fmt.Printf(" (по %d раз)\n", count)
			}
		}
	}

}

// Функция принимает путь до файла и возвращает строку контента
func readFile(file string) string {
	data, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("Ошибка чтения файла. \n %v", err)
	}

	return string(data)

}

// Функция считает количество строк в файле и возвращает их число
func countLines(content string) int {
	split := strings.Split(content, "\n")

	return len(split)
}

// Функция считает количество слов в файле и возвращает их число
func countWords(content string) int {
	split := strings.Fields(content)

	return len(split)
}

// Функция считает количество букв и возвращает её Unicode и кол-во повторений
func countFreqLet(content string) ([]rune, int) {
	cnt := make(map[rune]int)

	for _, char := range strings.ToLower(content) {
		if unicode.IsLetter(char) { // Игнорируем не-буквы
			cnt[char]++
		}
	}

	maxCount := 0
	var letters []rune

	for char, count := range cnt {
		if count > maxCount {
			maxCount = count
			letters = []rune{char}
		} else if count == maxCount {
			letters = append(letters, char)
		}
	}

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	return letters, maxCount

}
