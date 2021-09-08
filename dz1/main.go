package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

/*func ShowMeFiles() {  функция для отображения содержимого папок
	fmt.Println("nothing here yet")
}*/

func main() {

	out := os.Stdout
	path := os.Args[1] /* путь = второму аргументу(например точка это путь к текущему файлу) */
	/* будем печатать файлы или нет */

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("Неверное кол-во аргументов. Укажите имя файла и путь. 'main.go . [-f]' ")
	}
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f" /*true or false */
	fmt.Println(path, printFiles, out)
	dirTree(out, path, printFiles)
}

func dirTree(out io.Writer, path string, printFiles bool) (err error) {
	/* ShowFiles := flag.Bool("f", false, "shows files")
	flag.Parse()  флаг для вывода не только папок, но и их содержимого  */

	var MassiveOfFiles []string              // заносим названия ФАЙЛОВ в директории
	var MassiveOfDirs []string               /* заносим названия ПАПОК в директории */
	DirsAndFiles, err := ioutil.ReadDir(".") /* пояснение внизу */
	if err != nil {                          /* какая то ошибка при чтении директории */
		log.Fatal(err)
	}

	for _, file := range DirsAndFiles {
		if file.IsDir() { /* если папка */
			MassiveOfDirs = append(MassiveOfDirs, file.Name()) /* добавляем в массив строки = названию сожержимого */
		} else {
			MassiveOfFiles = append(MassiveOfFiles, file.Name()+" ("+strconv.Itoa(int(file.Size()))+"b)")
		}
		/* fmt.Println(f.Name()) */

	}
	sort.Strings(MassiveOfFiles)
	fmt.Println(MassiveOfFiles)
	fmt.Println(MassiveOfDirs)
	return nil
}

/*Результаты ( список папок-файлов ) должны быть отсортированы по алфавиту.
Т.е. у вас должен быть код который отсортирует уровень.
Смотрите для этого пакет sort.
Это самая частая причина непрохождения тестов. Тесты запускаются в среде linux. В задании есть докер-файл для тестов ровно в тех же условиях, он сразу выявит все проблемы.

https://pkg.go.dev/io/ioutil#example-ReadDir

STRCONV
i, err := strconv.Atoi("-42") (string to int)
s := strconv.Itoa(-42) (int to string)

" ("+strconv.Itoa(int(file.Size()))+"b)") - если делать без strconv, выдает ошибку, несовместимые типы данных стринг и инт

*/
