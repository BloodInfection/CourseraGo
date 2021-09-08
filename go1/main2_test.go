package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing" /* пакет для тестирования*/
)

/* создаддим тестовые данные*/
var testOk = `1
2
3
4
5
`

func TestOk(t *testing.T) { /* принимает на вход единственный аргумент*/
	in := bufio.NewReader(strings.NewReader(testOk)) /* новый ридер ньюридер, принимающий стрингс.ньюридер, в него передадим тестовую строку*/
	out := new(bytes.Buffer)                         /* создадим буфер байт, куда будет писаться вывод*/
	err := uniq(in, out)                             /* готовый ввод  и вывод передаем в юник*/
	if err != nil {
		t.Errorf("test for OK Failed")

	}

}

/* go test -v main2_test.go
-v(мы хотим видеть результаты, какие тесты заработали, а какие закончились неудачей)*/
