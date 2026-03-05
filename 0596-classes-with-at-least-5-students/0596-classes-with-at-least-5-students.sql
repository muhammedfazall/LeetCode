-- Write your PostgreSQL query statement below
select class from Courses Group by class having Count(*) > 4