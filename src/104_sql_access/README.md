## create database: go_toturial_recordings and create an album table
At the command line, log into your DBMS, as in the following example for MySQL.
```
$ mysql -u root -p
Enter password:

mysql>
At the mysql command prompt, create a database.

mysql> create database recordings;
Change to the database you just created so you can add tables.

mysql> use go_tutorial_recordings;
Database changed
```

Use the source command to create tables
```
mysql> source /path/to/create-tables.sql
```


## allow the usage of setup.sh
```
chmod +x setup.sh
```

## run setup.sh
```
source setup.sh
```
