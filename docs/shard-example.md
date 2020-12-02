# gaea 分片规则示例说明

## 导航
- [gaea kingshard hash分片示例](#gaea_kingshard_hash)
- [gaea kingshard mod分片示例](#gaea_kingshard_mod)
- [gaea kingshard range分片示例](#gaea_kingshard_range)
- [gaea kingshard date year分片示例](#gaea_kingshard_date_year)

<h2 id="gaea_kingshard_hash">gaea kingshard hash分片示例</h2>

### 创建数据库表
我们预定义两个分片slice-0、slice-1，分别位于两个数据库实例端口为3307、3308，每个slice预定义2张表

```shell script
#连接3307数据库实例
mysql -h127.0.0.1 -P3307 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表shard_hash_0000、shard_hash_0001
for i in `seq 0 1`;do  mysql -h127.0.0.1 -P3307 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_hash_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done

#连接3306数据库实例
mysql -h127.0.0.1 -P3308 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表shard_hash_0002、shard_hash_0003
for i in `seq 2 3`;do  mysql -h127.0.0.1 -P3308 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_hash_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done
#登录3307示例，查询slice-0分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_hash_0000        |
| shard_hash_0001        |
+------------------------+
2 rows in set (0.01 sec)
#登录3308示例，查询slice-1分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_hash_0002        |
| shard_hash_0003        |
+------------------------+
2 rows in set (0.00 sec)
```

### namespace配置
```json
{
    "name": "test_kingshard_hash",
    "online": true,
    "read_only": false,
    "allowed_dbs": {
        "db_kingshard": true
    },
    "default_phy_dbs": {
        "db_kingshard": "db_kingshard"
    },
    "slow_sql_time": "1000",
    "black_sql": [
        ""
    ],
    "allowed_ip": null,
    "slices": [
        {
            "name": "slice-0",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3307",
            "slaves": [],
            "statistic_slaves": null,
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        },
        {
            "name": "slice-1",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3308",
            "slaves": [],
            "statistic_slaves": [],
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        }
    ],
    "shard_rules": [
        {
            "db": "db_kingshard",
            "table": "shard_hash",
            "type": "hash",
            "key": "id",
            "locations": [
                2,
                2
            ],
            "slices": [
                "slice-0",
                "slice-1"
            ]
        }
    ],
    "users": [
        {
            "user_name": "test",
            "password": "1234",
            "namespace": "test_kingshard_hash",
            "rw_flag": 2,
            "rw_split": 1,
            "other_property": 0
        }
    ],
    "default_slice": "slice-1",
    "global_sequences": null
}
```

### 插入数据
```shell script
#命令行执行：
for i in `seq 1 10`;do mysql -h127.0.0.1 -P13306 -utest -p1234  db_kingshard -e "insert into shard_hash (id, col1) values(${i}, 'test$i')";done
```

### 查看数据
```shell script
#连接gaea，进行数据查询：
mysql> select * from shard_hash;
+----+--------+
| id | col1   |
+----+--------+
|  4 | test4  |
|  8 | test8  |
|  1 | test1  |
|  5 | test5  |
|  9 | test9  |
|  2 | test2  |
|  6 | test6  |
| 10 | test10 |
|  3 | test3  |
|  7 | test7  |
+----+--------+
10 rows in set (0.03 sec)
#连接3307数据库实例，对slice-0分表数据进行查询：
mysql> select * from shard_hash_0000;
+----+-------+
| id | col1  |
+----+-------+
|  4 | test4 |
|  8 | test8 |
+----+-------+
2 rows in set (0.00 sec)
mysql> select * from shard_hash_0001;
+----+-------+
| id | col1  |
+----+-------+
|  1 | test1 |
|  5 | test5 |
|  9 | test9 |
+----+-------+
3 rows in set (0.01 sec)
#连接3308数据库实例，对slice-1分表数据进行查询：
mysql>  select * from shard_hash_0002;
+----+--------+
| id | col1   |
+----+--------+
|  2 | test2  |
|  6 | test6  |
| 10 | test10 |
+----+--------+
3 rows in set (0.01 sec)
mysql>  select * from shard_hash_0003;
+----+-------+
| id | col1  |
+----+-------+
|  3 | test3 |
|  7 | test7 |
+----+-------+
2 rows in set (0.01 sec)
```

<h2 id="gaea_kingshard_mod">gaea kingshard mod分片示例</h2>

### 创建数据库表
我们预定义两个分片slice-0、slice-1，分别位于两个数据库实例端口为3307、3308，每个slice预定义2张表

```shell script
#连接3307数据库实例
mysql -h127.0.0.1 -P3307 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_mod_0000、shard_mod_0001
for i in `seq 0 1`;do  mysql -h127.0.0.1 -P3307 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_mod_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done

#连接3306数据库实例
mysql -h127.0.0.1 -P3308 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_mod_0002、shard_mod_0003
for i in `seq 2 3`;do  mysql -h127.0.0.1 -P3308 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_mod_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done
#登录3307实例，查询slice-0分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_mod_0000         |
| shard_mod_0001         |
+------------------------+
2 rows in set (0.01 sec)
#登录3308示例，查询slice-1分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_mod_0002         |
| shard_mod_0003         |
+------------------------+
2 rows in set (0.00 sec)
```

### namespace配置
```json
{
    "name": "test_kingshard_mod",
    "online": true,
    "read_only": false,
    "allowed_dbs": {
        "db_kingshard": true
    },
    "default_phy_dbs": {
        "db_kingshard": "db_kingshard"
    },
    "slow_sql_time": "1000",
    "black_sql": [
        ""
    ],
    "allowed_ip": null,
    "slices": [
        {
            "name": "slice-0",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3307",
            "slaves": [],
            "statistic_slaves": null,
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        },
        {
            "name": "slice-1",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3308",
            "slaves": [],
            "statistic_slaves": [],
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        }
    ],
    "shard_rules": [
        {
            "db": "db_kingshard",
            "table": "shard_mod",
            "type": "mod",
            "key": "id",
            "locations": [
                2,
                2
            ],
            "slices": [
                "slice-0",
                "slice-1"
            ]
        }
    ],
    "users": [
        {
            "user_name": "test",
            "password": "1234",
            "namespace": "test_kingshard_mod",
            "rw_flag": 2,
            "rw_split": 1,
            "other_property": 0
        }
    ],
    "default_slice": "slice-1",
    "global_sequences": null
}

```

### 插入数据
```shell script
#命令行执行：
for i in `seq 1 10`;do mysql -h127.0.0.1 -P13306 -utest -p1234  db_kingshard -e "insert into shard_mod (id, col1) values(${i}, 'test$i')";done
```

### 查看数据
```shell script
#连接gaea，进行数据查询：
mysql> select * from shard_mod;
+----+--------+
| id | col1   |
+----+--------+
|  4 | test4  |
|  8 | test8  |
|  1 | test1  |
|  5 | test5  |
|  9 | test9  |
|  2 | test2  |
|  6 | test6  |
| 10 | test10 |
|  3 | test3  |
|  7 | test7  |
+----+--------+
10 rows in set (0.03 sec)

#连接3307数据库实例，对slice-0分表数据进行查询：
mysql> select * from shard_mod_0000;
+----+-------+
| id | col1  |
+----+-------+
|  4 | test4 |
|  8 | test8 |
+----+-------+
2 rows in set (0.00 sec)
mysql> select * from shard_mod_0001;
+----+-------+
| id | col1  |
+----+-------+
|  1 | test1 |
|  5 | test5 |
|  9 | test9 |
+----+-------+
3 rows in set (0.01 sec)

#连接3308数据库实例，对slice-1分表数据进行查询：
mysql> select * from shard_mod_0002;
+----+--------+
| id | col1   |
+----+--------+
|  2 | test2  |
|  6 | test6  |
| 10 | test10 |
+----+--------+
3 rows in set (0.00 sec)
mysql> select * from shard_mod_0003;
+----+-------+
| id | col1  |
+----+-------+
|  3 | test3 |
|  7 | test7 |
+----+-------+
2 rows in set (0.01 sec)

# 对分片表数据进行跨节点更新，连接Gaea，执行：
mysql> update shard_mod set col1="test" where id in(4,2);
Query OK, 2 rows affected (0.02 sec)

mysql> select * from shard_mod;
+----+--------+
| id | col1   |
+----+--------+
|  2 | test   |
|  6 | test6  |
| 10 | test10 |
|  3 | test3  |
|  7 | test7  |
|  4 | test   |
|  8 | test8  |
|  1 | test1  |
|  5 | test5  |
|  9 | test9  |
+----+--------+
10 rows in set (0.04 sec)
```

<h2 id="gaea_kingshard_range">gaea kingshard range分片示例</h2>

### 创建数据库表
我们预定义两个分片slice-0、slice-1，分别位于两个数据库实例端口为3307、3308，每个slice预定义2张表

```shell script
#连接3307数据库实例
mysql -h127.0.0.1 -P3307 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_range_0000、shard_range_0001
for i in `seq 0 1`;do  mysql -h127.0.0.1 -P3307 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_range_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done

#连接3306数据库实例
mysql -h127.0.0.1 -P3308 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_range_0002、shard_range_0003
for i in `seq 2 3`;do  mysql -h127.0.0.1 -P3308 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_range_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done
#登录3307实例，查询slice-0分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_range_0000       |
| shard_range_0001       |
+------------------------+
2 rows in set (0.00 sec)

#登录3308示例，查询slice-1分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_range_0002       |
| shard_range_0003       |
+------------------------+
2 rows in set (0.01 sec)
```

### namespace配置
```json
{
    "name": "test_kingshard_range",
    "online": true,
    "read_only": false,
    "allowed_dbs": {
        "db_kingshard": true
    },
    "default_phy_dbs": {
        "db_kingshard": "db_kingshard"
    },
    "slow_sql_time": "1000",
    "black_sql": [
        ""
    ],
    "allowed_ip": null,
    "slices": [
        {
            "name": "slice-0",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3307",
            "slaves": [],
            "statistic_slaves": null,
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        },
        {
            "name": "slice-1",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3308",
            "slaves": [],
            "statistic_slaves": [],
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        }
    ],
    "shard_rules": [
        {
            "db": "db_kingshard",
            "table": "shard_range",
            "type": "range",
            "key": "id",
            "locations": [
                2,
                2
            ],
            "slices": [
                "slice-0",
                "slice-1"
            ],
            "table_row_limit": 3
        }
    ],
    "users": [
        {
            "user_name": "test",
            "password": "1234",
            "namespace": "test_kingshard_range",
            "rw_flag": 2,
            "rw_split": 1,
            "other_property": 0
        }
    ],
    "default_slice": "slice-1",
    "global_sequences": null
}

```
其中，"table_row_limit:3"配置含义为:每张子表的记录数，分表字段位于区间[0,3)在shard_range_0000上，分表字段位于区间[3,6)在子表shard_range_0001上，依此类推...

### 插入数据
```shell script
#命令行执行：
for i in `seq 1 10`;do mysql -h127.0.0.1 -P13306 -utest -p1234  db_kingshard -e "insert into shard_range (id, col1) values(${i}, 'test$i')";done
```

### 查看数据
```shell script
#连接gaea，进行数据查询：
mysql> select * from shard_range;
+----+--------+
| id | col1   |
+----+--------+
|  1 | test1  |
|  2 | test2  |
|  3 | test3  |
|  4 | test4  |
|  5 | test5  |
|  6 | test6  |
|  7 | test7  |
|  8 | test8  |
|  9 | test9  |
| 10 | test10 |
+----+--------+
10 rows in set (0.03 sec)
#连接3307数据库实例，对slice-0分表数据进行查询：
mysql> select * from shard_range_0000;
+----+-------+
| id | col1  |
+----+-------+
|  1 | test1 |
|  2 | test2 |
+----+-------+
2 rows in set (0.01 sec)
mysql> select * from shard_range_0001;
+----+-------+
| id | col1  |
+----+-------+
|  3 | test3 |
|  4 | test4 |
|  5 | test5 |
+----+-------+
3 rows in set (0.01 sec)
#连接3308数据库实例，对slice-1分表数据进行查询：
mysql> select * from shard_range_0002;
+----+-------+
| id | col1  |
+----+-------+
|  6 | test6 |
|  7 | test7 |
|  8 | test8 |
+----+-------+
3 rows in set (0.01 sec)
mysql> select * from shard_range_0003;
+----+--------+
| id | col1   |
+----+--------+
|  9 | test9  |
| 10 | test10 |
+----+--------+
2 rows in set (0.00 sec)
```

<h2 id="gaea_kingshard_date_year">gaea kingshard date year分片示例</h2>

### 创建数据库表
我们预定义两个分片slice-0、slice-1，分别位于两个数据库实例端口为3307、3308，每个slice预定义2张表

```shell script
#连接3307数据库实例
mysql -h127.0.0.1 -P3307 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_range_0000、shard_range_0001
for i in `seq 6 7`;do  mysql -h127.0.0.1 -P3307 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_year_201"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),create_time datetime DEFAULT NULL,PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done

#连接3306数据库实例
mysql -h127.0.0.1 -P3308 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_range_0002、shard_range_0003
for i in `seq 8 9`;do  mysql -h127.0.0.1 -P3308 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_year_201"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),create_time datetime DEFAULT NULL,PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done
#登录3307实例，查询slice-0分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_year_2016        |
| shard_year_2017        |
+------------------------+
2 rows in set (0.00 sec)

#登录3308示例，查询slice-1分片表展示：
mysql> show tables;
+------------------------+
| Tables_in_db_kingshard |
+------------------------+
| shard_year_2018        |
| shard_year_2019        |
+------------------------+
2 rows in set (0.01 sec)
```

### namespace配置
```json
{
    "name": "test_kingshard_date_year",
    "online": true,
    "read_only": false,
    "allowed_dbs": {
        "db_kingshard": true
    },
    "default_phy_dbs": {
        "db_kingshard": "db_kingshard"
    },
    "slow_sql_time": "1000",
    "black_sql": [
        ""
    ],
    "allowed_ip": null,
    "slices": [
        {
            "name": "slice-0",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3307",
            "slaves": [],
            "statistic_slaves": null,
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        },
        {
            "name": "slice-1",
            "user_name": "root",
            "password": "1234",
            "master": "127.0.0.1:3308",
            "slaves": [],
            "statistic_slaves": [],
            "capacity": 12,
            "max_capacity": 24,
            "idle_timeout": 60
        }
    ],
    "shard_rules": [
        {
            "db": "db_kingshard",
            "table": "shard_year",
            "type": "date_year",
            "key": "create_time",
            "slices": [
                "slice-0",
                "slice-1"
            ],
            "date_range": [
                "2016-2017",
                "2018-2019"
            ]
        }
    ],
    "users": [
        {
            "user_name": "test",
            "password": "1234",
            "namespace": "test_kingshard_date_year",
            "rw_flag": 2,
            "rw_split": 1,
            "other_property": 0
        }
    ],
    "default_slice": "slice-1",
    "global_sequences": null
}


```


### 插入数据
```shell script
#命令行执行：
for i in `seq 6 9`;do mysql -h127.0.0.1 -P13306 -utest -p1234  db_kingshard -e "insert into shard_year (id, col1,create_time) values(${i}, 'test$i','201$i-07-01')";done
```

### 查看数据
```shell script
#连接gaea，进行数据查询：
mysql> select * from shard_year;
+----+-------+---------------------+
| id | col1  | create_time         |
+----+-------+---------------------+
|  6 | test6 | 2016-07-01 00:00:00 |
|  7 | test7 | 2017-07-01 00:00:00 |
|  8 | test8 | 2018-07-01 00:00:00 |
|  9 | test9 | 2019-07-01 00:00:00 |
+----+-------+---------------------+
4 rows in set (0.03 sec)

#连接3307数据库实例，对slice-0分表数据进行查询：
mysql> select * from shard_year_2016;
+----+-------+---------------------+
| id | col1  | create_time         |
+----+-------+---------------------+
|  6 | test6 | 2016-07-01 00:00:00 |
+----+-------+---------------------+
1 row in set (0.01 sec)
mysql> select * from shard_year_2017;
+----+-------+---------------------+
| id | col1  | create_time         |
+----+-------+---------------------+
|  7 | test7 | 2017-07-01 00:00:00 |
+----+-------+---------------------+
1 row in set (0.01 sec)
#连接3308数据库实例，对slice-1分表数据进行查询：
mysql> select * from shard_year_2018;
+----+-------+---------------------+
| id | col1  | create_time         |
+----+-------+---------------------+
|  8 | test8 | 2018-07-01 00:00:00 |
+----+-------+---------------------+
1 row in set (0.01 sec)
mysql> select * from shard_year_2019;
+----+-------+---------------------+
| id | col1  | create_time         |
+----+-------+---------------------+
|  9 | test9 | 2019-07-01 00:00:00 |
+----+-------+---------------------+
1 row in set (0.00 sec)
```
