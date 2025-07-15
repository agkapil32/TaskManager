# 📫 Task Manager API – Quick Reference

This document includes all available endpoints, their parameters, and how to use them with Postman or curl.

---

## 🔧 How to Run the Server

```bash
# Step 1: Make sure Go is installed
go version

# Step 2: Run the server
go run main.go
```

Server will be up at: http://localhost:8080

---

##  API Endpoints

### ▶ POST /tasks  
Create a new task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "userId": 1,
    "name": "Demo Task",
    "description": "This is a test",
    "priority": "medium"
  }'
```

---

### ▶ GET /tasks  
Get a list of tasks (supports filters & pagination)

Query Parameters:

- userId: filter by user
- status: filter by status (`Backlog`, `WIP`, `Completed`)
- page: default = 1
- pageSize: default = 10
- sort: optional (`id`, `startTime`, `createdAt`)

```bash
curl "http://localhost:8080/tasks?userId=1&status=Backlog&page=1&pageSize=5"
```

---

### ▶ GET /tasks/{id}  
Get a task by ID

```bash
curl http://localhost:8080/tasks/1
```

---

### ▶ PUT /tasks/{id}  
Update an existing task

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "userId": 1,
    "name": "Updated Task Name",
    "description": "Updated description",
    "status": "Completed",
    "priority": "high"
  }'
```

---

### ▶ DELETE /tasks/{id}  
Delete a task

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## 🧪 Testing with Postman

1. Open Postman
2. Import the file: `TaskManager.postman_collection.json`
3. Use the "Task Manager" collection to try out all endpoints
4. Make sure the base URL is: http://localhost:8080

---