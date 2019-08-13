# Write your MySQL querxy statement below
SELECT
    (SELECT DISTINCT
            salary
        FROM Employee
        ORDER BY
            salary desc
        LIMIT 1
        OFFSET 1) AS SecondHighestSalary