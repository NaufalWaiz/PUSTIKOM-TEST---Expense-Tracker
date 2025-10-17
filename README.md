
# Expense Tracker

Project ini adalah test dari PUSTIKOM UNJ, dimana dibutuhkannya para peserta untuk membuat Fullstack WebApp pelacakan pengeluaran pribadi yang memungkinkan pengguna menambahkan, melihat, memfilter, dan menghapus pengeluaran.




## API Reference

#### Get all items

```http
GET    /api/expenses
POST   /api/expenses
PUT    /api/expenses/:id
DELETE /api/expenses/:id
```

## Deployment

Untuk deploy project ini, bisa mengikuti langkah-langkah berikut :

- Frontend
```bash
  cd frontend
  npm install
  npm run dev
```
- Backend
```bash
  cd backend
  go mod tidy/go get
  go run main.go
```


## Tech Stack

**Frontend (Client):** Svelte, Typescript

**Backend (Server):** Golang

**Database:** Supabase

