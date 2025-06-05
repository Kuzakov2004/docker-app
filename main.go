package main

import (
    "fmt"



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





    // Собираем интерфейс
    content := container.NewVBox(
        widget.NewLabel("=== Калькулятор ==="),
        calcEntry,
        calcButton,
        calcResult,


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


