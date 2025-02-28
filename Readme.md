API Endpoints

1. Libraries
   Create Library
   POST /libraries
   Request Body:

json
Copy
Edit
{
"name": "Central Library"
}
Get Library
GET /libraries/:id
Example: GET /libraries/1

Update Library
PUT /libraries/:id
Request Body:

json
Copy
Edit
{
"name": "Updated Central Library"
}
Delete Library
DELETE /libraries/:id
Example: DELETE /libraries/1

2. Users
   Create User
   POST /users
   Request Body:

json
Copy
Edit
{
"name": "John Doe",
"email": "john.doe@example.com",
"contact_number": "+1234567890",
"role": "Librarian",
"lib_id": 1
}
Get User
GET /users/:id
Example: GET /users/1

Update User
PUT /users/:id
Request Body:

json
Copy
Edit
{
"name": "Johnathan Doe",
"email": "johnathan.doe@example.com",
"contact_number": "+1234567890",
"role": "Manager",
"lib_id": 1
}
Delete User
DELETE /users/:id
Example: DELETE /users/1

3. Books
   Create Book
   POST /books
   Request Body:

json
Copy
Edit
{
"isbn": "978-3-16-148410-0",
"lib_id": 1,
"title": "Introduction to Go",
"authors": "John Smith, Jane Doe",
"publisher": "TechBooks",
"version": "1.0",
"total_copies": 10,
"available_copies": 10
}
Get Book
GET /books/:isbn
Example: GET /books/978-3-16-148410-0

Update Book
PUT /books/:isbn
Request Body:

json
Copy
Edit
{
"isbn": "978-3-16-148410-0",
"lib_id": 1,
"title": "Advanced Go Programming",
"authors": "John Smith, Jane Doe",
"publisher": "TechBooks",
"version": "2.0",
"total_copies": 15,
"available_copies": 15
}
Delete Book
DELETE /books/:isbn
Example: DELETE /books/978-3-16-148410-0

4. Request Events
   Create Request Event
   POST /requestevents
   Request Body:
   json
   Copy
   Edit
   {
   "book_id": "978-3-16-148410-0",
   "reader_id": 1,
   "request_date": "2025-02-28T15:04:05Z",
   "approver_id": 2,
   "request_type": "Reserve"
   }
5. Issue Registry
   Create Issue Registry
   POST /issueregistry
   Request Body:
   json
   Copy
   Edit
   {
   "isbn": "978-3-16-148410-0",
   "reader_id": 1,
   "issue_approver_id": 2,
   "issue_status": "Issued",
   "issue_date": "2025-02-28T15:04:05Z",
   "expected_return_date": "2025-03-07",
   "return_date": null,
   "return_approver_id": null
   }
