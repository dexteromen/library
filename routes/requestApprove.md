# **📌 How to Test Issue Request API on Postman**

#### **🔹 Prerequisites:**

1. **Start Your Server:** Ensure your Go Gin server is running.
2. **Database Setup:** Ensure your PostgreSQL database is running and has some books added.
3. **Authentication:** Use a valid JWT token (if authentication is required).

---

## **📌 1️⃣ Reader - Raise an Issue Request**

### **🔹 Endpoint:**

```
POST http://localhost:8080/reader/books/request
```

### **🔹 Headers:**

| Key           | Value                     |
| ------------- | ------------------------- |
| Authorization | Bearer **YOUR_JWT_TOKEN** |
| Content-Type  | application/json          |

### **🔹 Body (Raw, JSON):**

```json
{
  "book_id": 1
}
```

### **✅ Expected Responses:**

| **Scenario**       | **Response**                           | **Status Code**    |
| ------------------ | -------------------------------------- | ------------------ |
| Book is available  | `Issue request submitted successfully` | `200 OK`           |
| Book not available | `Book is not available`                | `409 Conflict`     |
| Invalid book ID    | `Book not found`                       | `404 Not Found`    |
| No token provided  | `Unauthorized access`                  | `401 Unauthorized` |

---

## **📌 2️⃣ Admin - List Issue Requests**

### **🔹 Endpoint:**

```
GET http://localhost:8080/admin/books/requests
```

### **🔹 Headers:**

| Key           | Value                           |
| ------------- | ------------------------------- |
| Authorization | Bearer **YOUR_ADMIN_JWT_TOKEN** |

### **✅ Expected Responses:**

| **Scenario**      | **Response**           | **Status Code**    |
| ----------------- | ---------------------- | ------------------ |
| Requests exist    | List of issue requests | `200 OK`           |
| No requests found | `[]` (empty list)      | `200 OK`           |
| No token provided | `Unauthorized access`  | `401 Unauthorized` |

---

## **📌 3️⃣ Admin - Approve/Reject Issue Request**

### **🔹 Endpoint:**

```
PUT http://localhost:8080/admin/books/requests/1
```

### **🔹 Headers:**

| Key           | Value                           |
| ------------- | ------------------------------- |
| Authorization | Bearer **YOUR_ADMIN_JWT_TOKEN** |
| Content-Type  | application/json                |

### **🔹 Body (Raw, JSON):**

```json
{
  "status": "approved"
}
```

**OR**

```json
{
  "status": "rejected"
}
```

### **✅ Expected Responses:**

| **Scenario**       | **Response**                           | **Status Code**    |
| ------------------ | -------------------------------------- | ------------------ |
| Request approved   | `Issue request processed successfully` | `200 OK`           |
| Request rejected   | `Issue request processed successfully` | `200 OK`           |
| Invalid request ID | `Issue request not found`              | `404 Not Found`    |
| Invalid status     | `Invalid status value`                 | `400 Bad Request`  |
| No token provided  | `Unauthorized access`                  | `401 Unauthorized` |

---

## **📌 How to Authenticate in Postman?**

1. **Go to "Authorization" Tab in Postman.**
2. **Select "Bearer Token".**
3. **Paste Your JWT Token.**
   - If you don't have a token, first **login** and get a token.

---

## **✅ Summary**

| **User**   | **Action**             | **Method** | **Endpoint**                | **Status Codes**     |
| ---------- | ---------------------- | ---------- | --------------------------- | -------------------- |
| **Reader** | Raise an Issue Request | `POST`     | `/reader/books/request`     | `200, 404, 409, 401` |
| **Admin**  | List Issue Requests    | `GET`      | `/admin/books/requests`     | `200, 401`           |
| **Admin**  | Approve/Reject Request | `PUT`      | `/admin/books/requests/:id` | `200, 400, 404, 401` |

🚀 Now, you can **test your API properly using Postman**! Let me know if you need more help! 😊
