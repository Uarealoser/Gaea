# gaea 分片规则示例说明

## 导航
- [gaea kingshard hash分片示例](#gaea_kingshard_hash)

<h3 id="gaea_kingshard_hash">gaea kingshard hash分片示例</h3>
#### 创建数据库表
我们预定义两个分片slice-0、slice-1，分别位于两个数据库实例端口为3307、3308
```shell script
#连接3307数据库实例
mysql -h127.0.0.1 -P3307 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_hash_0000、shard_hash_0001
for i in `seq 0 1`;do  mysql -h127.0.0.1 -P3307 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_hash_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done

#连接3306数据库实例
mysql -h127.0.0.1 -P3308 -uroot -p1234
#创建数据库
create database db_kingshard;
#在命令行执行以下命令，创建分表,shard_hash_0002、shard_hash_0003
for i in `seq 2 3`;do  mysql -h127.0.0.1 -P3308 -uroot -p1234  db_kingshard -e "CREATE TABLE IF NOT EXISTS shard_hash_000"${i}" ( id INT(64) NOT NULL, col1 VARCHAR(256),PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8;";done
```

#### namespace配置
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

#### 插入数据
```shell script
for i in `seq 1 10`;do mysql -h127.0.0.1 -P13306 -utest -p1234  db_kingshard -e "insert into shard_hash (id, col1) values(${i}, 'test$i')";done
```

#### 查看数据
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