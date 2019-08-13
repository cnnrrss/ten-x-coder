# Write your MySQL query statement below

# Query Table Twice Solution
SELECT a.Name As Employee
FROM 
    Employee AS a,
    Employee AS b
WHERE
    a.ManagerId = b.Id
        AND  a.Salary > b.Salary
;

# JOIN Solution
SELECT 
    a.Name As Employee
FROM Employee AS a JOIN Employee AS b
    ON a.ManagerId = b.Id
    AND  a.Salary > b.Salary
;