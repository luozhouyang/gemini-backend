# Init postgreSQL

Create user and database:

```bash
sudo -u postgres createuser --superuser gemini
sudo -u postgres psql
postgres=# \password gemini
Enter new password:$PASSWORD
Enter it again:$PASSWORD
postgres=# CREATE DATABASE geminidb OWNER gemini;
postgres=# GRANT ALL PRIVILEGES ON DATABASE geminidb TO gemini;
```

Setup tables:

```bash
cd $GOPATH/src/gemini/db
psql -U gemini -d geminidb -h 127.0.0.1 -p 5432 < setup.sql
```


