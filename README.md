# 🌟 User Management System with Technical Challenges

Welcome to the **User Management System** repository! This project contains solutions to a variety of technical challenges, showcasing Go, TypeScript, React, and Next.js skills. Dive into the details of each task and explore how these technologies come together. 🚀

---

## 🧩 Challenges and Solutions

### 1️⃣ **Word Sorter by 'a' Count**

📚 **Description**: Sorts a list of words by the number of occurrences of the letter `a` (case-insensitive), then by word length, and finally lexicographically.

🔑 **Key Features**:
- Prioritizes words with more 'a' characters.
- Handles ties using word length and alphabetical order.

🎉 **Run the Solution**:
```bash
go run q1/main.go
```

🧪 **Test with**:
```bash
go test ./q1
```

---

### 2️⃣ **Recursive Sequence Generator**

🔢 **Description**: Generates a sequence starting from `2`, recursively calculating squares until the result exceeds a given number.

✨ **Example**:
- Input: `9`
- Output:
  ```
  2
  4
  9
  ```

🎉 **Run the Solution**:
```bash
go run q2/main.go
```

🧪 **Test with**:
```bash
go test ./q2
```

---

### 3️⃣ **Most Repeated Element Finder**

🔄 **Description**: Identifies the most frequently occurring element in a slice of strings.

✨ **Example**:
- Input: `["apple", "banana", "apple", "orange", "banana", "apple"]`
- Output: `"apple"`

🎉 **Run the Solution**:
```bash
go run q3/main.go
```

🧪 **Test with**:
```bash
go test ./q3
```

---

### 4️⃣ **User Management Application**

🧑‍💻 **Description**: A full-stack application to manage users with CRUD operations.

💻 **Tech Stack**:
- **Backend**: Go, SQLite
- **Frontend**: TypeScript, React, Next.js

🎨 **Features**:
- **Master View**: List all users, create new users, edit or delete existing ones.
- **Detail View**: Forms for creating or editing user details.
- **REST API**: Endpoints for user operations with proper HTTP status codes.

🎉 **Run the Application**:

#### With Docker:
1. Start the application:
   ```bash
   docker-compose up --build
   ```
2. Access the frontend at `http://localhost:3000` and the backend at `http://localhost:8080`.

#### Without Docker:
- **Backend**:
  ```bash
  cd backend
  go run main.go
  ```
- **Frontend**:
  ```bash
  cd frontend
  npm install
  npm run dev
  ```

🧪 **Test Backend**:
```bash
go test ./backend
```

---

## 🚀 Project Highlights

- **Go**: Clean, efficient backend with RESTful APIs.
- **SQLite**: Lightweight database for persistence.
- **React + Next.js**: Modern UI with responsive design.
- **TypeScript**: Ensures code quality with static typing.
- **Docker**: Simplified deployment with `docker-compose`.

---

## 📂 Repository Structure

```plaintext
.
├── q1/                  # Word Sorter by 'a' Count
│   ├── main.go
│   ├── main_test.go
├── q2/                  # Recursive Sequence Generator
│   ├── main.go
│   ├── main_test.go
├── q3/                  # Most Repeated Element Finder
│   ├── main.go
│   ├── main_test.go
├── q4/                  # User Management
│├── backend/             # User Management Backend
│   ├── main.go
│   ├── handlers/
│   ├── db/
│   ├── models/
│   ├── routes/
│├── frontend/            # User Management Frontend
│   ├── app/
│   ├── components/
│   ├── public/
│   ├── styles/
├── docker-compose.yml   # Docker Compose Configuration
└── README.md            # Repository Documentation
```
