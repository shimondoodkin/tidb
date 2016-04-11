# TiDB on HBase


0) simple installation of HBase, using cloudera cdh (with steps for ubuntu):

1. http://www.cloudera.com/documentation/enterprise/latest/topics/cm_ig_path_b_manual.html#concept_p4z_qxt_25_unique_3__ol_jhv_kdd_d5_unique_2 (do the 1 click install to install the repositories,  richt click on links  of your dist, and wget download it, dpkg -i it :  wget https://archive.cloudera.com/cdh5/one-click-install/trusty/amd64/cdh5-repository_1.0_all.deb ; sudo dpkg -i cdh5-repository_1.0_all.deb
)
2. http://www.cloudera.com/documentation/enterprise/latest/topics/cdh_ig_hbase_install.html (sudo apt-get install hbase )
3. http://www.cloudera.com/documentation/enterprise/latest/topics/cdh_ig_hbase_config.html (edit the files as written)
2. http://www.cloudera.com/documentation/enterprise/latest/topics/cdh_ig_hbase_standalone_start.html (sudo apt-get install hbase-master)

**TiDB only supports HBase >= 0.98.5**

1) Build && Install pingcap/themis coprocessor to HBase:

1. git clone https://github.com/pingcap/themis.git
2. cd themis && mvn clean package -DskipTests
4. cp themis-coprocessor/target/themis-coprocessor-1.0-SNAPSHOT-jar-with-dependencies.jar /usr/lib/hbase/lib/  ( copy jar to $HBASE_ROOT/lib )
5. Add configurations for themis coprocessor in hbase-site.xml:

3)
1. service hbase-master start
```
<property>
    <name>hbase.coprocessor.user.region.classes</name>
    <value>org.apache.hadoop.hbase.themis.cp.ThemisEndpoint,org.apache.hadoop.hbase.themis.cp.ThemisScanObserver,org.apache.hadoop.hbase.regionserver.ThemisRegionObserver</value>
</property>
<property>
    <name>hbase.coprocessor.master.classes</name>
    <value>org.apache.hadoop.hbase.master.ThemisMasterObserver</value>
</property>
```

and then restart HBase.


2) Build TiDB:

```
git clone https://github.com/pingcap/tidb.git $GOPATH/src/github.com/pingcap/tidb
cd $GOPATH/src/github.com/pingcap/tidb
make
```

Run command line interpreter:

```
make interpreter
cd interpreter && ./interpreter -store hbase -dbpath localhost/tidb -dbname test
```

Enjoy it.

```
Welcome to the TiDB.
Version:
Git Commit Hash: f37bd11d16c79a3db1cdb068ef7a6c872f682cda
UTC Build Time:  2015-12-07 08:10:25

tidb> create table t1(id int, name text, key id(id));
Query OK, 0 row affected (2.12 sec)
tidb> insert into t1 values(1, "hello");
Query OK, 1 row affected (0.01 sec)
tidb> insert into t1 values(2, "world");
Query OK, 1 row affected (0.00 sec)
tidb> select * from t1;
+----+-------+
| id | name  |
+----+-------+
| 1  | hello |
| 2  | world |
+----+-------+
2 rows in set (0.00 sec)
tidb>
```

Run TiDB server:

```
make server
cd tidb-server
./tidb-server -store=hbase -path="zkaddrs/hbaseTbl?tso=tsoType" -P=4000
DSN parameters:
zkaddrs is the address of zookeeper.
hbaseTbl is the table in hbase to store TiDB data.
tsoaddr is the type of tso sever. Its value could be zk or local.
Here is an example of dsn:
./tidb-server -store=hbase -path="zk1,zk2/test?tso=zk" -P=5000
```
