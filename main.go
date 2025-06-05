package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Многофункциональное приложение")
    myWindow.Resize(fyne.NewSize(400, 300))

    // 1. Калькулятор (как в статье)
    calcEntry := widget.NewEntry()
    calcEntry.SetPlaceHolder("Введите выражение, например: 2+2")
    calcResult := widget.NewLabel("Результат:")
    calcButton := widget.NewButton("Вычислить", func() {
        expr := calcEntry.Text
        result, err := eval(expr)
        if err != nil {
            calcResult.SetText("Ошибка: " + err.Error())
        } else {
            calcResult.SetText(fmt.Sprintf("Результат: %v", result))
        }
    })

    // 2. Генератор случайного числа (1-10)
    randomResult := widget.NewLabel("Число: -")
    randomButton := widget.NewButton("Случайное число (1-10)", func() {
        rand.Seed(time.Now().UnixNano())
        num := rand.Intn(10) + 1
        randomResult.SetText(fmt.Sprintf("Число: %d", num))
    })

    // 3. Подбрасывание монетки
    coinResult := widget.NewLabel("Монетка: -")
    coinButton := widget.NewButton("Подбросить монетку", func() {
        rand.Seed(time.Now().UnixNano())
        side := "Орёл"
        if rand.Intn(2) == 1 {
            side = "Решка"
        }
        coinResult.SetText("Монетка: " + side)
    })

    // 4. Проверка чётности числа
    evenEntry := widget.NewEntry()
    evenEntry.SetPlaceHolder("Введите число")
    evenResult := widget.NewLabel("Чётность: -")
    evenButton := widget.NewButton("Проверить чётность", func() {
        num, err := strconv.Atoi(evenEntry.Text)
        if err != nil {
            evenResult.SetText("Ошибка: введите число!")
        } else {
            if num%2 == 0 {
                evenResult.SetText("Чётность: Чётное")
            } else {
                evenResult.SetText("Чётность: Нечётное")
            }
        }
    })

    // Собираем интерфейс
    content := container.NewVBox(
        widget.NewLabel("=== Калькулятор ==="),
        calcEntry,
        calcButton,
        calcResult,

        widget.NewLabel("=== Генератор случайного числа ==="),
        randomButton,
        randomResult,

        widget.NewLabel("=== Подбрасывание монетки ==="),
        coinButton,
        coinResult,

        widget.NewLabel("=== Проверка чётности ==="),
        evenEntry,
        evenButton,
        evenResult,
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}

// Функция для вычисления выражения (как в статье)
func eval(expr string) (float64, error) {
    // Простейший калькулятор (без обработки ошибок ввода)
    // В реальном приложении лучше использовать парсер выражений
    var result float64
    _, err := fmt.Sscanf(expr, "%f+%f", &result, &result)
    if err == nil {
        return result + result, nil
    }
    _, err = fmt.Sscanf(expr, "%f-%f", &result, &result)
    if err == nil {
        return result - result, nil
    }
    _, err = fmt.Sscanf(expr, "%f*%f", &result, &result)
    if err == nil {
        return result * result, nil
    }
    _, err = fmt.Sscanf(expr, "%f/%f", &result, &result)
    if err == nil {
        return result / result, nil
    }
    return 0, fmt.Errorf("неверное выражение")
}

