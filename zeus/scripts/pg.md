# pg设置

## 新建数据库和用户

```sh
# 登录
sudo -u postgres psql postgres

# 新建用户
CREATE USER username WITH PASSWORD 'xxx' CREATEDB;

# 新建数据库
CREATE DATABASE databasename OWNER username TABLESPACE username;
```

## 导入表结构和预置数据

```sh
sudo -u postgres psql databasename < ./zeus/data/schema.sql
```
