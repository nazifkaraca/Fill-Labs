# 🔢 Recursive Sequence Generator

## ✨ Overview

This Go program generates a sequence of numbers based on a recursive function. The sequence starts from the square of 2 and continues with the squares of subsequent integers until the square exceeds the input number.

## 🛠️ Features

- Generates a sequence of squares recursively starting from 2.
- Includes interactive input or accepts a command-line argument.
- Validates input to ensure it is a positive integer.

## 🧩 Input Example

The program accepts a single integer (e.g., 9) via:

1. **Command-line argument**:
   ```bash
   go run main.go 9
   ```
2. **Interactive input**:
   ```bash
   go run main.go
   Enter a positive integer: 9
   ```

## 🎉 Output Example

For an input of `9`, the program outputs:

```
2
4
9
```

## 🚀 How to Run

1. Build or run the program:

   - **With Command-Line Argument**:
     ```bash
     go run main.go 25
     ```
   - **Interactive Input**:
     ```bash
     go run main.go
     Enter a positive integer: 25
     ```

2. The output will display the sequence.

## 🧪 How It Works

1. The sequence starts with the number `2`.
2. For each integer, it prints:
   - The integer itself (only for the initial value of `2`).
   - The square of the integer.
3. The recursion stops when the square of the current number exceeds the input value.

## ⚠️ Error Handling

- If a non-positive integer is provided as input, the program will terminate with:
  ```
  Please provide a valid positive integer.
  ```

## 📂 Project Structure

```
/recursive-sequence
│-- main.go        // Main program for sequence generation
```

## 🧪 Testing

Testing can be performed by running the program with various input values to confirm:

- Proper sequence generation.
- Input validation (handles invalid inputs gracefully).
