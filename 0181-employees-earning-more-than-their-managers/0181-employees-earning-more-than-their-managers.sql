-- Write your PostgreSQL query statement below
SELECT e.name AS Employee from employee AS e
join employee AS m
on e.managerId = m.id
WHERE e.salary > m.salary