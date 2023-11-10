Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
3 2 3

Нужно помнить, что при работе со слайсами и передаче их в функцию мы осуществляем
копирование. Таким образом 0 элемент меняется, тк ссылка в срезе совпадает с 
исходным при аппенде же ссылка на массив меняется, тк создается новый срез и 
все последующие действия не влияют на исходный срез
```