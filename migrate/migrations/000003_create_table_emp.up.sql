CREATE TABLE emp(
empno INT PRIMARY KEY,
ename VARCHAR(10),
job VARCHAR(9),
mgr INT,
hiredate DATE,
sal DOUBLE,
comm DOUBLE,
deptno INT REFERENCES dept);
