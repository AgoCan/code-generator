# ansible

## 代码目录结构

```bash
.
├── config                        # 配置文件目录
│   ├── config.ini                # 给用户的配置文件
│   ├── config.sh                 # 保存一些常用的配置
│   └── set_config.sh             # 设置一些根据配置文件获得的变量
                                  # 这里应该得再来一个 check_config.sh 检测配置文件是否符合要求
├── playbooks                     # ansible的剧本
│   ├── all.yml                   # site 文件
│   └── roles                     # 角色，简单的一个打印hello的任务
│       └── base
│           └── tasks
│               └── main.yml
├── README.md                     # 看心情写吧
├── start.sh                      # 入口函数
└── utils                         # 一些工具
    ├── _common.sh                # 保存了几个日志函数
    └── _gen_inventory_file.sh    # 生成inventory文件
```