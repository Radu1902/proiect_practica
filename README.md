# Practice Management Application

This web application was built using the Go programming language [Golang](https://go.dev/), HTML, and [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/). For the visual styling of the pages, i also used [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS).

---

## Backend of the Application

The backend is written in Golang. I utilized the [`net/http`](https://pkg.go.dev/net/http) package to create the web application.

### Key Functionalities
- **Web Server:**
  - We used the [`http.ListenAndServe()`](https://pkg.go.dev/net/http#ListenAndServe) function to start the web server.
  - Handlers were added to the server using the `http.HandleFunc()` function. 

- **Handler Functions:**
  - These functions take two parameters:
    1. `http.ResponseWriter` - used to send responses to clients.
    2. `*http.Request` - a pointer to the request object, used to receive data from clients.

### Database Connection
- We used the [`database/sql`](https://pkg.go.dev/database/sql) package for database connectivity.
- A PostgreSQL driver from GitHub ([lib/pq](https://github.com/lib/pq)) was employed for database operations.

### Database Schema
The database used is [PostgreSQL 16](https://www.postgresql.org/docs/current/release-16.html). The schema for the database named `practica` is as follows:

#### Table: `studenti` (students)
```sql
CREATE TABLE studenti (
    email VARCHAR(255) PRIMARY KEY,
    nume VARCHAR(255),
    parola VARCHAR(255),
    grupa VARCHAR(255),
    specializare VARCHAR(255),
    firma VARCHAR(255),
    start_date DATE,
    end_date DATE
);
```

#### Table: `admini` (admins)
```sql
CREATE TABLE admini (
    email VARCHAR(255) PRIMARY KEY,
    nume VARCHAR(255),
    parola VARCHAR(255)
);
```

---

## Frontend of the Application

The layout of the pages is written in HTML, and the visual styling of some page elements is done using CSS (within the `<style>` tag in the HTML files). Data fetched from the backend is processed and displayed in the frontend using JavaScript (within the `<script>` tag in the HTML files).

---

## Usage

The application consists of three web pages:

### 1. Login Page
- Users authenticate using their email and password.
- A user can either be a **student** or an **admin**.

### 2. Student Page
- Displays the student’s details, including:
  - Email, specialization, group, company, start date, and end date.
- If the student does not have a practice placement, fields like company, start date, and end date will be marked as `unspecified`.
- Contains a button to open a dialog where the student can update or delete their practice details.

### 3. Admin Page
- Displays a list of all students with their details:
  - Email, name, group, specialization, company, start date, and end date.

---
