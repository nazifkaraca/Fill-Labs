# 📚 Word Sorter by 'a' Count

## ✨ Overview

This project implements a Go function to sort words based on:

1. **🅰️ Count of 'a'** in descending order.
2. **📏 Length of the word** (for ties).
3. **🔤 Lexicographical order** (for further ties).

## 🛠️ Features

- Counts occurrences of 'a' (case-insensitive) in words.
- Handles ties gracefully by using word length and alphabetical order.
- Includes comprehensive test cases to validate functionality.

## 🧩 Input Example

```go
["aaaasd", "a", "aab", "aaabcd", "ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"]
```

````

## 🎉 Output Example

```go
["aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "csssssssd", "fdz", "ef", "kf", "zc", "l"]
```

## 🚀 How to Run

1. Clone this repository:

   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. Run the program:

   ```bash
   go run main.go
   ```

3. Enter words separated by spaces when prompted.

## 🧪 Testing

1. Run tests to ensure the program works as expected:

   ```bash
   go test
   ```

2. Example test cases include:
   - Sorting by 'a' count.
   - Sorting by length for ties.
   - Sorting lexicographically for further ties.

## 📜 Sample Test Cases

### ✅ Test Case 1

**Input**:

```go
["aaaasd", "a", "aab", "aaabcd", "ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"]
```

**Output**:

```go
["aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "csssssssd", "fdz", "ef", "kf", "zc", "l"]
```

### ✅ Test Case 2

**Input**:

```go
["banana", "apple", "grape", "aaa", "aa", "a"]
```

**Output**:

```go
["banana", "aaa", "aa", "apple", "grape", "a"]
```

### ✅ Test Case 3

**Input**:

```go
["AAA", "AaA", "aaa", "A", "a"]
```

**Output**:

```go
["aaa", "AaA", "AAA", "a", "A"]
```

## 📂 Project Structure

```
/q1
│-- main.go        // Main sorting functionality
│-- main_test.go   // Unit tests
```
````
