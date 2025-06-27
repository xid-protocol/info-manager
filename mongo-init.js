// 切换到 xidp 数据库
db = db.getSiblingDB('xidp');

// 创建应用用户
db.createUser({
  user: 'xidp_user',
  pwd: 'xidp_password',
  roles: [
    {
      role: 'readWrite',
      db: 'xidp'
    }
  ]
});

// 创建集合和索引
db.createCollection('xid_info');
db.createCollection('users');

// 为 xid_info 集合创建索引
db.xid_info.createIndex({ "xid": 1, "metadata.path": 1 }, { unique: true });
db.xid_info.createIndex({ "name": 1, "path": 1 });

print('MongoDB 初始化完成'); 