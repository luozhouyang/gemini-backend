# Gemini-backend
This project is used as a backend service for **gemini**.


## Create databases
To ensure that this program runs correctly, you need to create databases that this program needs.
Here are the commands to create mysql databases:
　　
```bash
mysql -u root -p

mysql> create user 'gemini'@'localhost' identified by 'usergemini';
mysql> create database gemini_db;
mysql> grant all privileges on gemini_db.* to gemini@ocalhost identified by 'usergemini';
mysql>
mysql> create database gemini_db_dev;
mysql> grant all privileges on gemini_db_dev.* to gemini@ocalhost identified by 'usergemini';
mysql>
mysql> create database gemini_db_test;
mysql> grant all privileges on gemini_db_test.* to gemini@ocalhost identified by 'usergemini';
mysql>
```
