# interactive_inputs

A Go library for creating interactive terminal user interfaces, providing radio button and checkbox selectors for easy item selection from lists.

## Features

- **Radio Selector**: Single selection from a list of items
- **Checkbox Selector**: Multiple selection with configurable min/max limits
- **Text Transformation**: Support for uppercase, lowercase, and capitalize transformations
- **Scrolling**: Automatic scrolling for large lists
- **Cross-platform**: Works on Windows, macOS, and Linux
- **Keyboard Navigation**: Navigate with arrow keys, select with Enter/Space

## Installation

```bash
go get github.com/digvijay-tech/interactive_inputs
```

## Requirements

- Go 1.25.0 or later
- A terminal that supports raw mode (most modern terminals do)

## Usage

### Radio Selector

The radio selector allows users to pick a single item from a list.

```go
package main

import (
    "fmt"
    "log"

    "github.com/digvijay-tech/interactive_inputs"
)

func main() {
    options := []string{"Option 1", "Option 2", "Option 3"}

    radioOpts := &interactive_inputs.RadioOptions{
        Title:         "Choose an option:",
        Description:   "Select one item from the list.",
        TextTransform: interactive_inputs.CAPITALISE,
    }

    selected, err := interactive_inputs.Radio(options, radioOpts)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("You selected: %s\n", selected)
}
```

### Checkbox Selector

The checkbox selector allows users to pick multiple items from a list with optional min/max selection limits.

```go
package main

import (
    "fmt"
    "log"

    "github.com/digvijay-tech/interactive_inputs"
)

func main() {
    items := []string{"Apple", "Banana", "Cherry", "Date"}

    checkboxOpts := &interactive_inputs.CheckboxOptions{
        Title:         "Select fruits:",
        Description:   "Choose your favorite fruits.",
        TextTransform: interactive_inputs.CAPITALISE,
        MinSelection:  1,
        MaxSelection:  3,
    }

    selected, err := interactive_inputs.Checkbox(items, checkboxOpts)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("You selected:")
    for _, item := range selected {
        fmt.Printf("- %s\n", item)
    }
}
```

## Reference

### Types

#### `AcceptedListType`

The library supports the following types for list items:
- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`

#### `TextTransform`

Constants for text transformation:
- `NONE`: No transformation
- `CAPITALISE`: Capitalize first letter of each word
- `LOWERCASE`: Convert to lowercase
- `UPPERCASE`: Convert to uppercase

### RadioOptions

```go
type RadioOptions struct {
    Title         string        // Title displayed above the list
    Description   string        // Description displayed below the title
    TextTransform TextTransform // Text transformation to apply
}
```

### CheckboxOptions

```go
type CheckboxOptions struct {
    Title         string        // Title displayed above the list
    Description   string        // Description displayed below the title
    MinSelection  uint          // Minimum number of items that must be selected
    MaxSelection  uint          // Maximum number of items that can be selected
    TextTransform TextTransform // Text transformation to apply
}
```

## Controls

- **↑/↓**: Navigate through items
- **Enter**: Confirm selection (Radio) or finish selection (Checkbox when min requirements met)
- **Space**: Toggle selection (Checkbox only)
- **Ctrl+C**: Exit

## Examples

See the `examples/` directory for complete working examples:
- [Radio](./examples/radio/main.go)
- [Checkbox](./examples/checkbox//main.go)


## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.
