# 🚀 Мои решения и эксперименты на Go

[![GitHub License](https://img.shields.io/github/license/Andrey-Zobnin/EasyProgrammOnGo?style=flat-square)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg?style=flat-square)](https://golang.org/dl/)
[![Stepik Course](https://img.shields.io/badge/Stepik-Course-009688.svg?style=flat-square)](https://stepik.org/course/ПРО-GO-Основы-программирования-1158)

Мои решения задач и эксперименты во время изучения Go. Здесь я сохраняю свой прогресс и наработки.

## 🔥 Примеры моего кода

### 1. Основы вывода (Section 1.3.2)
```go
package main

import "fmt"

func main() {
    // Изучаю комментарии и вывод
    fmt.Println("Go") // Мой выбор языка
    fmt.Println("Python") // Для сравнения
    fmt.Println("other")
}
```
В этом примере я экспериментировал с комментариями и выводом разных языков.

### 2. Разные способы вывода (Section 2.1)
```go
package main

import "fmt"

func main() {
    // Сравниваю Println и Print
    fmt.Println("row1") // Вывод с новой строки
    fmt.Print("row2")   // Без перевода строки
    fmt.Println("row3") // Снова с новой строки
}
```
Здесь я изучал разницу между `Println` и `Print`.

### 3. Работа с переменными (Section 2.2)
```go
package main

import "fmt"

func main() {
    var a string = "25" // Объявление переменной
    var b string = "4"
    fmt.Println(a + b) // Конкатенация строк
    // Хотел получить 29, но вышло "254"
}
```
Этот пример показал мне разницу между строковыми и числовыми типами.

### 4. Ввод данных (Section 2.3)
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Учусь работать с вводом
    scan := bufio.NewScanner(os.Stdin)
    var name string
    scan.Scan()
    name = scan.Text()
    fmt.Println("Привет,", name)
}
```
Здесь я освоил базовый ввод данных от пользователя.

### 5. Числовые операции (Section 2.4)
```go
package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    // Изучаю математические операции
    cube := num * num * num
    result := cube * cube
    fmt.Println(result) // num^6
}
```
В этих задачах я практиковался в математических вычислениях.

## 📊 Мой прогресс

Недавно я начал изучать Go и уже освоил:

- Базовый вывод (`Println`, `Print`)
- Работу с переменными
- Основы ввода данных
- Простые математические операции

Сейчас перехожу к более сложным темам: условные операторы и циклы.

## 🛠️ Как использовать мой код

1. Клонируйте репозиторий:
```bash
git clone https://github.com/Andrey-Zobnin/EasyProgrammOnGo.git
```

2. Запустите любой пример:
```bash
go run stepik_course/section2/task2.1/main.go
```

## 📝 Мои заметки

В процессе изучения я обнаружил:

1. В Go нельзя иметь несколько `main()` в одном пакете
2. `Println` автоматически добавляет пробелы между аргументами
3. Конкатенация строк и сложение чисел - разные операции
4. Важно проверять ошибки при вводе данных

## 📜 Лицензия

Этот проект лицензирован под MIT License. Все учебные материалы принадлежат Stepik и авторам курса. Здесь только мои личные решения и эксперименты.
