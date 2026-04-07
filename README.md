# api-estudantes

Routes:

- GET /students - List all students
- GET /students?active=<true/false> - List all active/non-active students
- GET /students/:id - Get informations about a specific student
- POST /students - Create student
- PUT /students/:id - Update informations about student
- DELET /students/:id - Delete student

Struct Student:
- Name (string)
- CPF (int)
- Email (string)
- Age (int)
- Active (bool)
