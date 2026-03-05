-- Write your PostgreSQL query statement below
with cte as (
select customer_number,count(customer_number) from orders group by customer_number)
select customer_number  from cte where count = (select max(count) from cte) 