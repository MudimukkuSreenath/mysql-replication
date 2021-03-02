***********STEP 1: INSTALL MYSQL SERVER ON SERVER 1 (MASTER)*********

user@user-sreenath-456:~$ sudo apt-get update
[sudo] password for user: 
Hit:1 http://in.archive.ubuntu.com/ubuntu groovy InRelease
Hit:2 http://in.archive.ubuntu.com/ubuntu groovy-updates InRelease                              
Get:3 http://security.ubuntu.com/ubuntu groovy-security InRelease [110 kB]                                                                   
Hit:4 http://in.archive.ubuntu.com/ubuntu groovy-backports InRelease                                                                         
Hit:5 https://linux.teamviewer.com/deb stable InRelease                                                                                      
Ign:6 http://ppa.launchpad.net/nae-team/ppa/ubuntu groovy InRelease                                                                          
Hit:7 https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/groovy pgadmin4 InRelease
Err:8 http://ppa.launchpad.net/nae-team/ppa/ubuntu groovy Release                  
  404  Not Found [IP: 2001:67c:1560:8008::19 80]
Reading package lists... Done
E: The repository 'http://ppa.launchpad.net/nae-team/ppa/ubuntu groovy Release' does not have a Release file.
N: Updating from such a repository can't be done securely, and is therefore disabled by default.
N: See apt-secure(8) manpage for repository creation and user configuration details.
user@user-sreenath-456:~$ sudo apt-get install mysql-server
Reading package lists... Done
Building dependency tree       
Reading state information... Done
mysql-server is already the newest version (8.0.23-0ubuntu0.20.10.1).
0 upgraded, 0 newly installed, 0 to remove and 46 not upgraded.



***********STEP 2: EDIT MYSQL CONFIGURATION FILE ON SERVER 1*********



user@user-sreenath-456:~$ sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf
user@user-sreenath-456:~$ sudo systemctl restart mysql
user@user-sreenath-456:~$ sudo mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.23-0ubuntu0.20.10.1 (Ubuntu)

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

user@user-sreenath-456:~$ sudo systemctl status mysql
● mysql.service - MySQL Community Server
     Loaded: loaded (/lib/systemd/system/mysql.service; enabled; vendor preset: enabled)
     Active: active (running) since Mon 2021-03-01 11:51:51 IST; 36s ago
    Process: 10234 ExecStartPre=/usr/share/mysql/mysql-systemd-start pre (code=exited, status=0/SUCCESS)
   Main PID: 10243 (mysqld)
     Status: "Server is operational"
      Tasks: 38 (limit: 1871)
     Memory: 274.8M
     CGroup: /system.slice/mysql.service
             └─10243 /usr/sbin/mysqld

Mar 01 11:51:48 user-sreenath-456 systemd[1]: Starting MySQL Community Server...
Mar 01 11:51:51 user-sreenath-456 systemd[1]: Started MySQL Community Server.
user@user-sreenath-456:~$  sudo mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.23-0ubuntu0.20.10.1 (Ubuntu)

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> CREATE USER 'replication_user'@'192.168.43.74' IDENTIFIED BY 'sreenath15';
Query OK, 0 rows affected (0.16 sec)

mysql>  GRANT REPLICATION SLAVE ON *.* TO 'replication_user '@'192.168.43.74';
Query OK, 0 rows affected (0.08 sec)

mysql> SHOW MASTER STATUS\G
*************************** 1. row ***************************
             File: mysql-bin.000001
         Position: 737
     Binlog_Do_DB: 
 Binlog_Ignore_DB: 
Executed_Gtid_Set: 
1 row in set (0.14 sec)

mysql> sudo vim  /etc/mysql/mysql.conf.d/mysqld.cnf



**********Step 3: Create a New User for Replication Services on Server 1**********


mysql> CHANGE MASTER TO MASTER_HOST ='192.168.43.74', MASTER_USER ='replication_user', MASTER_PASSWORD ='sreenath15', MASTER_LOG_FILE = 'mysql-bin.000001',MASTER_LOG_POS = 737;
Query OK, 0 rows affected, 8 warnings (0.21 sec)

mysql> START SLAVE;
Query OK, 0 rows affected, 1 warning (0.23 sec)

mysql> CREATE DATABASE replication_db;
Query OK, 1 row affected (0.11 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| authentication     |
| information_schema |
| myisam             |
| mysql              |
| performance_schema |
| sreenath           |
| sys                |
+--------------------+
7 rows in set (0.84 sec)




mysql> use myisam;
Database changed
mysql> show tables;
Empty set (0.04 sec)

mysql> create database replica_demo;
Query OK, 1 row affected (0.10 sec)

mysql> use replica_demo;
Database changed
mysql>  CREATE TABLE students ( id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, student_name VARCHAR(30) NOT NULL ,student_age VARCHAR(40),student_gender VARCHAR(20));
Query OK, 0 rows affected, 1 warning (0.30 sec)

mysql> insert into student(id, student_name,student_age,student_gender) values ('420', 'John Doe', '20', 'male');
Query OK, 1 row affected (0.04 sec)

mysql> insert into student(id, student_name,student_age,student_gender) values ('421', 'John', '22', 'male');
Query OK, 1 row affected (0.15 sec)

mysql> insert into student(id, student_name,student_age,student_gender) values ('422', 'jothi', '24', 'female');
Query OK, 1 row affected (0.03 sec)

mysql> insert into student(id, student_name,student_age,student_gender) values ('424', 'lavi', '24', 'female');
Query OK, 1 row affected (0.02 sec)

mysql> insert into student(id, student_name,student_age,student_gender) values ('426', 'raj', '25', 'male');
Query OK, 1 row affected (0.12 sec)

mysql> show tables;
+------------------------+
| Tables_in_replica_demo |
+------------------------+
| student                |
| students               |
+------------------------+
2 rows in set (0.11 sec)

mysql> select * from student;
+-----+--------------+-------------+----------------+
| id  | student_name | student_age | student_gender |
+-----+--------------+-------------+----------------+
| 420 | John Doe     |          20 | male           |
| 421 | John         |          22 | male           |
| 422 | jothi        |          24 | female         |
| 424 | lavi         |          24 | female         |
| 426 | raj          |          25 | male           |
+-----+--------------+-------------+----------------+
5 rows in set (0.15 sec)

mysql> show slave status;
+----------------------+---------------+------------------+-------------+---------------+------------------+---------------------+------------------------+---------------+-----------------------+------------------+-------------------+-----------------+---------------------+--------------------+------------------------+-------------------------+-----------------------------+------------+------------+--------------+---------------------+-----------------+-----------------+----------------+---------------+--------------------+--------------------+--------------------+-----------------+-------------------+----------------+-----------------------+-------------------------------+---------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------+----------------+-----------------------------+------------------+-------------+-------------------------+-----------+---------------------+--------------------------------------------------------+--------------------+-------------+-------------------------+--------------------------+----------------+--------------------+--------------------+-------------------+---------------+----------------------+--------------+--------------------+------------------------+-----------------------+-------------------+
| Slave_IO_State       | Master_Host   | Master_User      | Master_Port | Connect_Retry | Master_Log_File  | Read_Master_Log_Pos | Relay_Log_File         | Relay_Log_Pos | Relay_Master_Log_File | Slave_IO_Running | Slave_SQL_Running | Replicate_Do_DB | Replicate_Ignore_DB | Replicate_Do_Table | Replicate_Ignore_Table | Replicate_Wild_Do_Table | Replicate_Wild_Ignore_Table | Last_Errno | Last_Error | Skip_Counter | Exec_Master_Log_Pos | Relay_Log_Space | Until_Condition | Until_Log_File | Until_Log_Pos | Master_SSL_Allowed | Master_SSL_CA_File | Master_SSL_CA_Path | Master_SSL_Cert | Master_SSL_Cipher | Master_SSL_Key | Seconds_Behind_Master | Master_SSL_Verify_Server_Cert | Last_IO_Errno | Last_IO_Error                                                                                                                                                                                                   | Last_SQL_Errno | Last_SQL_Error | Replicate_Ignore_Server_Ids | Master_Server_Id | Master_UUID | Master_Info_File        | SQL_Delay | SQL_Remaining_Delay | Slave_SQL_Running_State                                | Master_Retry_Count | Master_Bind | Last_IO_Error_Timestamp | Last_SQL_Error_Timestamp | Master_SSL_Crl | Master_SSL_Crlpath | Retrieved_Gtid_Set | Executed_Gtid_Set | Auto_Position | Replicate_Rewrite_DB | Channel_Name | Master_TLS_Version | Master_public_key_path | Get_master_public_key | Network_Namespace |
+----------------------+---------------+------------------+-------------+---------------+------------------+---------------------+------------------------+---------------+-----------------------+------------------+-------------------+-----------------+---------------------+--------------------+------------------------+-------------------------+-----------------------------+------------+------------+--------------+---------------------+-----------------+-----------------+----------------+---------------+--------------------+--------------------+--------------------+-----------------+-------------------+----------------+-----------------------+-------------------------------+---------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------+----------------+-----------------------------+------------------+-------------+-------------------------+-----------+---------------------+--------------------------------------------------------+--------------------+-------------+-------------------------+--------------------------+----------------+--------------------+--------------------+-------------------+---------------+----------------------+--------------+--------------------+------------------------+-----------------------+-------------------+
| Connecting to master | 192.168.43.74 | replication_user |        3306 |            60 | mysql-bin.000001 |                 737 | mysql-relay-bin.000001 |             4 | mysql-bin.000001      | Connecting       | Yes               |                 |                     |                    |                        |                         |                             |          0 |            |            0 |                 737 |             156 | None            |                |             0 | No                 |                    |                    |                 |                   |                |                  NULL | No                            |          2061 | error connecting to master 'replication_user@192.168.43.74:3306' - retry-time: 60 retries: 16 message: Authentication plugin 'caching_sha2_password' reported error: Authentication requires secure connection. |              0 |                |                             |                0 |             | mysql.slave_master_info |         0 |                NULL | Slave has read all relay log; waiting for more updates |              86400 |             | 210301 12:35:34         |                          |                |                    |                    |                   |             0 |                      |              |                    |                        |                     0 |                   |
+----------------------+---------------+------------------+-------------+---------------+------------------+---------------------+------------------------+---------------+-----------------------+------------------+-------------------+-----------------+---------------------+--------------------+------------------------+-------------------------+-----------------------------+------------+------------+--------------+---------------------+-----------------+-----------------+----------------+---------------+--------------------+--------------------+--------------------+-----------------+-------------------+----------------+-----------------------+-------------------------------+---------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------+----------------+-----------------------------+------------------+-------------+-------------------------+-----------+---------------------+--------------------------------------------------------+--------------------+-------------+-------------------------+--------------------------+----------------+--------------------+--------------------+-------------------+---------------+----------------------+--------------+--------------------+------------------------+-----------------------+-------------------+
1 row in set, 1 warning (0.05 sec)

mysql> create database raju;
Query OK, 1 row affected (0.14 sec)

mysql> use raju;
Database changed
mysql> create table EmployeeRecords ( EmpId int,EmpName varchar(100),EmpAge int, EmpSalary float)ENGINE=INNODB; 
Query OK, 0 rows affected (0.12 sec)

mysql> DESC EmployeeRecords;
+-----------+--------------+------+-----+---------+-------+
| Field     | Type         | Null | Key | Default | Extra |
+-----------+--------------+------+-----+---------+-------+
| EmpId     | int          | YES  |     | NULL    |       |
| EmpName   | varchar(100) | YES  |     | NULL    |       |
| EmpAge    | int          | YES  |     | NULL    |       |
| EmpSalary | float        | YES  |     | NULL    |       |
+-----------+--------------+------+-----+---------+-------+
4 rows in set (0.18 sec)

mysql> insert into EmployeeRecords (Empid, EmpName, Empage, EmpSalary) values ('101', 'rajesh','21', '21000');
Query OK, 1 row affected (0.15 sec)

mysql> insert into EmployeeRecords (Empid, EmpName, Empage, EmpSalary) values ('102', 'raj','22', '20000');
Query OK, 1 row affected (0.10 sec)

mysql> insert into EmployeeRecords (Empid, EmpName, Empage, EmpSalary) values ('103', 'raja','23', '25000');
Query OK, 1 row affected (0.04 sec)

mysql> insert into EmployeeRecords (Empid, EmpName, Empage, EmpSalary) values ('104', 'kumar','23', '25000');
Query OK, 1 row affected (0.01 sec)

mysql> insert into EmployeeRecords (Empid, EmpName, Empage, EmpSalary) values ('105', 'raj kumar','25', '30000');
Query OK, 1 row affected (0.08 sec)

mysql> select * from EmployeeRecords;
+-------+-----------+--------+-----------+
| EmpId | EmpName   | EmpAge | EmpSalary |
+-------+-----------+--------+-----------+
|   101 | rajesh    |     21 |     21000 |
|   102 | raj       |     22 |     20000 |
|   103 | raja      |     23 |     25000 |
|   104 | kumar     |     23 |     25000 |
|   105 | raj kumar |     25 |     30000 |
+-------+-----------+--------+-----------+
5 rows in set (0.08 sec)

mysql> create database Human;
Query OK, 1 row affected (0.14 sec)

mysql> use Human;
Database changed
mysql> CREATE TABLE students ( id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, student_name VARCHAR(30) NOT NULL, student_age int(30), student_gender VARCHAR(20))ENGINE=INNODB;
Query OK, 0 rows affected, 2 warnings (0.51 sec)

mysql> insert into students(id, student_name,student_age,student_gender) values ('401', ' Doe', '20', 'female');
Query OK, 1 row affected (0.17 sec)

mysql> insert into students(id, student_name,student_age,student_gender) values ('402', 'raj', '22', 'male');
Query OK, 1 row affected (0.01 sec)

mysql> insert into students(id, student_name,student_age,student_gender) values ('403', 'raju', '24', 'male');
Query OK, 1 row affected (0.02 sec)

mysql> insert into students(id, student_name,student_age,student_gender) values ('405', 'rajsh', '25', 'male');
Query OK, 1 row affected (0.06 sec)

mysql> insert into students(id, student_name,student_age,student_gender) values ('406', 'sudha', '26', 'female');
Query OK, 1 row affected (0.03 sec)

mysql> Desc students;
+----------------+--------------+------+-----+---------+----------------+
| Field          | Type         | Null | Key | Default | Extra          |
+----------------+--------------+------+-----+---------+----------------+
| id             | int unsigned | NO   | PRI | NULL    | auto_increment |
| student_name   | varchar(30)  | NO   |     | NULL    |                |
| student_age    | int          | YES  |     | NULL    |                |
| student_gender | varchar(20)  | YES  |     | NULL    |                |
+----------------+--------------+------+-----+---------+----------------+
4 rows in set (0.29 sec)

mysql> select * from students;
+-----+--------------+-------------+----------------+
| id  | student_name | student_age | student_gender |
+-----+--------------+-------------+----------------+
| 401 |  Doe         |          20 | female         |
| 402 | raj          |          22 | male           |
| 403 | raju         |          24 | male           |
| 405 | rajsh        |          25 | male           |
| 406 | sudha        |          26 | female         |
+-----+--------------+-------------+----------------+
5 rows in set (0.08 sec)

mysql> create database samera;
Query OK, 1 row affected (0.12 sec)

mysql> use samera;
Database changed
mysql> CREATE TABLE birds ( id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, student_name VARCHAR(30) NOT NULL )ENGINE=INNODB;
Query OK, 0 rows affected, 1 warning (0.19 sec)

mysql> insert into birds(id, student_name) values ('6', 'parrot');
Query OK, 1 row affected (0.11 sec)

mysql> insert into birds(id, student_name) values ('7', 'sparrow');
Query OK, 1 row affected (0.02 sec)

mysql> insert into birds(id, student_name) values ('8', 'pigeon');
Query OK, 1 row affected (0.01 sec)

mysql> insert into birds(id, student_name) values ('9', 'bug');
Query OK, 1 row affected (0.03 sec)

mysql> insert into birds(id, student_name) values ('9', 'crow');
ERROR 1062 (23000): Duplicate entry '9' for key 'birds.PRIMARY'
mysql> insert into birds(id, student_name) values ('10', 'crow');
Query OK, 1 row affected (0.01 sec)

mysql> select * from birds;
+----+--------------+
| id | student_name |
+----+--------------+
|  6 | parrot       |
|  7 | sparrow      |
|  8 | pigeon       |
|  9 | bug          |
| 10 | crow         |
+----+--------------+
5 rows in set (0.06 sec)

mysql> RESET MASTER;
Query OK, 0 rows affected (0.40 sec)

mysql> FLUSH TABLES WITH READ LOCK;
Query OK, 0 rows affected (0.02 sec)

mysql> SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000002 |      156 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
1 row in set (0.01 sec)


mysql> show global variables like 'GTID_EXECUTED';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| gtid_executed |       |
+---------------+-------+
1 row in set (0.28 sec)

mysql> START SLAVE;
Query OK, 0 rows affected, 2 warnings (0.00 sec)

mysql> SHOW SLAVE STATUS;
+----------------------+---------------+------------------+-------------+---------------+------------------+---------------------+------------------------+---------------+-----------------------+------------------+-------------------+-----------------+---------------------+--------------------+------------------------+-------------------------+-----------------------------+------------+------------+--------------+---------------------+-----------------+-----------------+----------------+---------------+--------------------+--------------------+--------------------+-----------------+-------------------+----------------+-----------------------+-------------------------------+---------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------+----------------+-----------------------------+------------------+-------------+-------------------------+-----------+---------------------+--------------------------------------------------------+--------------------+-------------+-------------------------+--------------------------+----------------+--------------------+--------------------+-------------------+---------------+----------------------+--------------+--------------------+------------------------+-----------------------+-------------------+
| Slave_IO_State       | Master_Host   | Master_User      | Master_Port | Connect_Retry | Master_Log_File  | Read_Master_Log_Pos | Relay_Log_File         | Relay_Log_Pos | Relay_Master_Log_File | Slave_IO_Running | Slave_SQL_Running | Replicate_Do_DB | Replicate_Ignore_DB | Replicate_Do_Table | Replicate_Ignore_Table | Replicate_Wild_Do_Table | Replicate_Wild_Ignore_Table | Last_Errno | Last_Error | Skip_Counter | Exec_Master_Log_Pos | Relay_Log_Space | Until_Condition | Until_Log_File | Until_Log_Pos | Master_SSL_Allowed | Master_SSL_CA_File | Master_SSL_CA_Path | Master_SSL_Cert | Master_SSL_Cipher | Master_SSL_Key | Seconds_Behind_Master | Master_SSL_Verify_Server_Cert | Last_IO_Errno | Last_IO_Error                                                                                                                                                                                                   | Last_SQL_Errno | Last_SQL_Error | Replicate_Ignore_Server_Ids | Master_Server_Id | Master_UUID | Master_Info_File        | SQL_Delay | SQL_Remaining_Delay | Slave_SQL_Running_State                                | Master_Retry_Count | Master_Bind | Last_IO_Error_Timestamp | Last_SQL_Error_Timestamp | Master_SSL_Crl | Master_SSL_Crlpath | Retrieved_Gtid_Set | Executed_Gtid_Set | Auto_Position | Replicate_Rewrite_DB | Channel_Name | Master_TLS_Version | Master_public_key_path | Get_master_public_key | Network_Namespace |
+----------------------+---------------+------------------+-------------+---------------+------------------+---------------------+------------------------+---------------+-----------------------+------------------+-------------------+-----------------+---------------------+--------------------+------------------------+-------------------------+-----------------------------+------------+------------+--------------+---------------------+-----------------+-----------------+----------------+---------------+--------------------+--------------------+--------------------+-----------------+-------------------+----------------+-----------------------+-------------------------------+---------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------+----------------+-----------------------------+------------------+-------------+-------------------------+-----------+---------------------+--------------------------------------------------------+--------------------+-------------+-------------------------+--------------------------+----------------+--------------------+--------------------+-------------------+---------------+----------------------+--------------+--------------------+------------------------+-----------------------+-------------------+
| Connecting to master | 192.168.43.74 | replication_user |        3306 |            60 | mysql-bin.000001 |                 737 | mysql-relay-bin.000001 |             4 | mysql-bin.000001      | Connecting       | Yes               |                 |                     |                    |                        |                         |                             |          0 |            |            0 |                 737 |             156 | None            |                |             0 | No                 |                    |                    |                 |                   |                |                  NULL | No                            |          2061 | error connecting to master 'replication_user@192.168.43.74:3306' - retry-time: 60 retries: 34 message: Authentication plugin 'caching_sha2_password' reported error: Authentication requires secure connection. |              0 |                |                             |                0 |             | mysql.slave_master_info |         0 |                NULL | Slave has read all relay log; waiting for more updates |              86400 |             | 210301 12:53:35         |                          |                |                    |                    |                   |             0 |                      |              |                    |                        |                     0 |                   |
+----------------------+---------------+------------------+-------------+---------------+------------------+---------------------+------------------------+---------------+-----------------------+------------------+-------------------+-----------------+---------------------+--------------------+------------------------+-------------------------+-----------------------------+------------+------------+--------------+---------------------+-----------------+-----------------+----------------+---------------+--------------------+--------------------+--------------------+-----------------+-------------------+----------------+-----------------------+-------------------------------+---------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------+----------------+-----------------------------+------------------+-------------+-------------------------+-----------+---------------------+--------------------------------------------------------+--------------------+-------------+-------------------------+--------------------------+----------------+--------------------+--------------------+-------------------+---------------+----------------------+--------------+--------------------+------------------------+-----------------------+-------------------+
1 row in set, 1 warning (0.03 sec)

mysql> SHOW SLAVE STATUS\G
*************************** 1. row ***************************
               Slave_IO_State: Connecting to master
                  Master_Host: 192.168.43.74
                  Master_User: replication_user
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: mysql-bin.000001
          Read_Master_Log_Pos: 737
               Relay_Log_File: mysql-relay-bin.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File: mysql-bin.000001
             Slave_IO_Running: Connecting
            Slave_SQL_Running: Yes
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 737
              Relay_Log_Space: 156
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 2061
                Last_IO_Error: error connecting to master 'replication_user@192.168.43.74:3306' - retry-time: 60 retries: 34 message: Authentication plugin 'caching_sha2_password' reported error: Authentication requires secure connection.
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 0
                  Master_UUID: 
             Master_Info_File: mysql.slave_master_info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: Slave has read all relay log; waiting for more updates
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 210301 12:53:35
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
         Replicate_Rewrite_DB: 
                 Channel_Name: 
           Master_TLS_Version: 
       Master_public_key_path: 
        Get_master_public_key: 0
            Network_Namespace: 
1 row in set, 1 warning (0.01 sec)

mysql> 

