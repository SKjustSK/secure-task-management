# Secure Task Manager

![Live Demo](https://img.shields.io/badge/Live_Demo-Available-success?style=for-the-badge)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

A full-stack, secure task management application featuring a high-contrast modern UI, JWT-based authentication, and a robust RESTful API. 

🚀 **Live Application:** [https://secure-task-management-iota.vercel.app/](https://secure-task-management-iota.vercel.app/)

---

## ✨ Features

* **Secure Authentication:** User registration and login protected by bcrypt password hashing and JSON Web Tokens (JWT).
* **Full CRUD Operations:** Create, read, update, and delete daily tasks seamlessly.
* **Modern Interface:** A clean, responsive, "neo-brutalist" design built with Tailwind CSS.
* **Smart Form Context:** Dynamic UI that transitions seamlessly between task creation and editing modes.
* **Real-time Synchronization:** Frontend fetches the absolute latest data on edit, preventing state conflicts.

## 🛠️ Technology Stack

### Frontend
* **Framework:** React 18 + TypeScript + Vite
* **Styling:** Tailwind CSS
* **Icons:** Lucide React
* **Network:** Axios (with automated JWT interceptors)
* **Hosting:** Vercel

### Backend & Database
* **Language:** Go (Golang)
* **Framework:** Gin Web Framework
* **ORM:** GORM
* **Security:** `golang-jwt/jwt`, `golang.org/x/crypto/bcrypt`
* **Database:** PostgreSQL (Hosted on Render)
* **Hosting:** Render
