# 📚 Bookstore App

A full-stack bookstore application that allows users to manage books, perform CRUD operations, and interact with a backend API.

---

## 🚀 Features

* 📖 View list of books
* ➕ Add new books
* ✏️ Update book information
* ❌ Delete books
* 🔎 REST API for managing data

---

## 🛠 Tech Stack

**Backend:**

* Go (Golang)
* REST API
* JSON

**Frontend:**

* (допиши: React / HTML / нет фронта)

**Other:**

* Git & GitHub

---

## 📂 Project Structure

```
backend/
 ├── config/         # Configuration files
 ├── controllers/    # Business logic
 ├── middleware/     # Middleware functions
 ├── models/         # Data models
 ├── routes/         # API routes
 ├── main.go         # Entry point
```

---

## ⚙️ Installation & Setup

### 1. Clone repository

```bash
git clone https://github.com/bisembaevaaigerim/bookstore.git
cd bookstore
```

### 2. Setup backend

```bash
cd backend
go mod tidy
go run main.go
```

---

## 🌐 API Endpoints (пример)

| Method | Endpoint   | Description     |
| ------ | ---------- | --------------- |
| GET    | /books     | Get all books   |
| GET    | /books/:id | Get book by ID  |
| POST   | /books     | Create new book |
| PUT    | /books/:id | Update book     |
| DELETE | /books/:id | Delete book     |

---

## 📌 Example Request

```json
{
  "title": "Book Name",
  "author": "Author Name",
  "year": 2024
}
```

---

## 🧠 Project Goal

This project demonstrates building a RESTful API using Go and structuring a backend application with clean architecture principles.

---

## 📷 Screenshots

*(добавь сюда скрины, если есть)*

---

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first.

---

## 📄 License

This project is open-source and available under the MIT License.
