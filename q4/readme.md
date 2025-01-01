# 🧑‍💻 User Management Application

## Description

This project is a **User Management System** that implements full CRUD functionality using:

- **Backend**: Written in Go with a REST API using SQLite for data persistence.
- **Frontend**: Built with TypeScript, React, and Next.js (App Router) for a responsive and dynamic UI.

The application is designed for managing users, featuring:

- A **Master View** with a data grid to list users and perform `New`, `Edit`, and `Delete` operations.
- A **Detail View** with a form for creating and editing users.

---

## Features

### Backend Features:

- **Endpoints**:
  - `GET /users`: Fetch all users.
  - `GET /users/{id}`: Fetch a user by ID.
  - `POST /users`: Create a new user.
  - `PUT /users/{id}`: Update a user by ID.
  - `DELETE /users/{id}`: Delete a user by ID.
- SQLite database for persistent storage.
- Follows REST API standards with proper HTTP status codes.
- CORS support for seamless communication with the frontend.

### Frontend Features:

- **Master View**:
  - Displays users in a **data grid**.
  - Buttons for creating a new user (`New`), editing a user (`Edit`), and deleting a user (`Delete`).
  - User IDs are dynamically displayed starting from 1, irrespective of database IDs.
- **Detail View**:
  - **Create**: Action button text displays "Create" for new users.
  - **Edit**: Action button text displays "Save" for editing users.
- Responsive design for desktop and mobile devices.

---

## Project Structure

### Backend

```
backend/
├── db/             # Database setup and management
│   └── database.go # SQLite initialization
├── handlers/       # REST API handlers
│   └── user_handlers.go
├── models/         # Data models
├── routes/         # Route definitions
│   └── routes.go
├── main.go         # Entry point for the backend
├── go.mod          # Dependency management
└── Dockerfile      # Docker file
```

### Frontend

```
frontend/
├── app/                # Next.js App Router
│   ├── layout.tsx      # Global layout
│   ├── page.tsx        # Master view
│   ├── user/           # User-specific routes
│   │   ├── [id]/       # Edit user page
│   │   │   └── page.tsx
│   │   ├── new/        # Create new user page
│   │   │   └── page.tsx
├── components/         # Reusable UI components
│   ├── DataGrid.tsx    # Master view data grid
│   └── UserForm.tsx    # Detail view user form
├── styles/             # Global styles
│   └── globals.css
├── public/             # Static assets
├── package.json        # Frontend dependencies
├── tsconfig.json       # TypeScript configuration
└── Dockerfile          # Docker file
```

---

## Technologies Used

### Backend:

- **Go**: Backend language.
- **SQLite**: Lightweight relational database.
- **Gorilla Mux**: HTTP router for REST API.
- **Gorilla Handlers**: CORS middleware.

### Frontend:

- **React**: User interface.
- **Next.js**: Framework for React with App Router.
- **TypeScript**: For static typing.
- **Axios**: For making HTTP requests.

---

## Installation & Setup

### Docker Compose Setup

1. Ensure you have **Docker** and **Docker Compose** installed.
2. Use the provided `docker-compose.yml` file to spin up both the frontend and backend:
   ```bash
   docker-compose up --build
   ```
3. Access the application:
   - Frontend: `http://localhost:3000`
   - Backend: `http://localhost:8080`

---

### Manual Setup (Without Docker)

#### Backend

1. Navigate to the `backend` folder:
   ```bash
   cd backend
   ```
2. Install Go dependencies:
   ```bash
   go mod tidy
   ```
3. Run the backend server:
   ```bash
   go run main.go
   ```
4. The server will start on `http://localhost:8080`.

#### Frontend

1. Navigate to the `frontend` folder:
   ```bash
   cd frontend
   ```
2. Install Node.js dependencies:
   ```bash
   npm install
   ```
3. Run the frontend development server:
   ```bash
   npm run dev
   ```
4. Access the application on `http://localhost:3000`.

---

## Docker Compose Configuration

### `docker-compose.yml`

```yaml
services:
  backend:
    build:
      context: ./backend
    container_name: user-management-backend
    ports:
      - "8080:8080"
    volumes:
      - ./backend:/app
    working_dir: /app
    command: ["go", "run", "main.go"]

  frontend:
    build:
      context: ./frontend
    container_name: user-management-frontend
    ports:
      - "3000:3000"
    volumes:
      - ./frontend:/app
      - /app/node_modules
    working_dir: /app
    command: ["npm", "run", "dev"]

networks:
  default:
    driver: bridge
```

---

## How to Use

1. Open the application in your browser at `http://localhost:3000`.
2. Use the **Master View** to:
   - View all users.
   - Create a new user by clicking `New`.
   - Edit an existing user by clicking `Edit`.
   - Delete a user by clicking `Delete` (with confirmation).
3. Use the **Detail View** for:
   - Creating a user (fields: `Name`, `Email`, `Age`).
   - Editing a user’s details.

---

## API Endpoints

| HTTP Method | Endpoint      | Description                   |
| ----------- | ------------- | ----------------------------- |
| GET         | `/users`      | Fetch all users               |
| GET         | `/users/{id}` | Fetch a specific user by ID   |
| POST        | `/users`      | Create a new user             |
| PUT         | `/users/{id}` | Update an existing user by ID |
| DELETE      | `/users/{id}` | Delete a user by ID           |

---

## Screenshots

### Master View

![image](https://github.com/user-attachments/assets/599ec8e1-78fa-4cbf-b3cd-22a04847c9ca)

### Detail View

![image](https://github.com/user-attachments/assets/ca83d3f6-e8df-42e8-8c82-080a49d48fc0)

---

## Additional Notes

- Ensure the backend (`http://localhost:8080`) is running when using the frontend.
- CORS is enabled to allow communication between the frontend and backend during development.
