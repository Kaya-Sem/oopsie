# Oopsie - A Simple Error Indicator for Go

Oopsie is a lightweight package designed to enhance error message
in Go CLI applications by providing colorfull, customizable
indicators.

## Installation

```bash
go get github.com/Kaya-Sem/oopsie

```

## Usage

##### Creating an Error Indicator
```go
package main

import (
    "fmt"
    "github.com/Kaya-Sem/oopsie"
)

func main() {
    err := fmt.Errorf("file not found")
    oops := oopsie.CreateOopsie().Title("File Read
        Error").IndicatorMessage("CRITICAL").IndicatorColors(oopsie.BRIGHT_WHITE,
        oopsie.RED).Error(err)

    fmt.Println(oops.Render)
}

```

##### Customizing the Indicator

- **Title**: Set a custom error title using `Title(string)`
- **Indicator Message**: Modify the indicator message using
`IndicatorMessage(string)`
- **Indicator Colors**: Change foreground and background colors
using `IndicatorColors(fg, bg)`.
- **Disable Indicator**: Hide the indicator using
`DisableIndicator(true)`.

### Available Colors


| Color Name       | Foreground Code | Background Code |
|-----------------|----------------|----------------|
| BLACK          | `033[30m`     | `033[40m`    |
| RED            | `033[31m`     | `033[41m`    |
| GREEN          | `033[32m`     | `033[42m`    |
| YELLOW         | `033[33m`     | `033[43m`    |
| BLUE           | `033[34m`     | `033[44m`    |
| PURPLE         | `033[35m`     | `033[45m`    |
| CYAN           | `033[36m`     | `033[46m`    |
| WHITE          | `033[37m`     | `033[47m`    |
| BRIGHT_BLACK   | `033[90m`     | `033[100m`   |
| BRIGHT_RED     | `033[91m`     | `033[101m`   |
| BRIGHT_GREEN   | `033[92m`     | `033[102m`   |
| BRIGHT_YELLOW  | `033[93m`     | `033[103m`   |
| BRIGHT_BLUE    | `033[94m`     | `033[104m`   |
| BRIGHT_PURPLE  | `033[95m`     | `033[105m`   |
| BRIGHT_CYAN    | `033[96m`     | `033[106m`   |
| BRIGHT_WHITE   | `033[97m`     | `033[107m`   |


## Guidelines

##### Best Practices

- Use `CreateOopsie()` to generate a default error object.
- Chain methods to customize error messages.
- Use appropriate colors to improve readability.

