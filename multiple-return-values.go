package main

import "fmt"

func main() {
    // Case 1
    min, max := getMinMax(4, 9, 8, 12)
    fmt.Println(min)
    fmt.Println(max)

    // Case 2
    _, max = getMinMax(4, 9, 8, 12)
    fmt.Println(max)
}

func getMinMax(numbers ...int) (int, int) {
    if len(numbers) == 0 {
      return 0, 0;
    }

    var min = numbers[0]
    var max = numbers[0]

    for _, num := range numbers {
        if num > max {
            max = num;
        }
        if num < min {
            min = num;
        }
    }

    return min, max
}
