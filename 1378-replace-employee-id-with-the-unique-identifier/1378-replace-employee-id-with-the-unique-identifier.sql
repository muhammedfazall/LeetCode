-- Write your PostgreSQL query statement below
SELECT EmployeeUni.unique_id , Employees.name FROM EmployeeUni RIGHT JOIN Employees ON EmployeeUni.id = Employees.id ;