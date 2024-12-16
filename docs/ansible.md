# ansible

## 代码目录结构

1. ansible的内容是通过shell去执行ansible，这样可以在执行ansible之前进行基础环境的准备
2. 当前使用的是配置参数的方式，传参的方式如果有需要，可以提issue，我进行补充

```bash
.
├── config                        # 配置文件目录
│   ├── check_config.sh           # 检查配置文件是否配置正确
│   ├── config.ini                # 给用户的配置文件
│   ├── config.sh                 # 默认配置文件，start.sh的执行顺序在ini之前
│   └── set_config.sh             # 设置一些根据配置文件获得的变量
├── core                          # 主要函数内容
│   ├── common.sh                 # 工具内容，例如日志
│   └── gen_inventory_file.sh     # 生成inventory内容
├── inventory                     # 被生成的文件，第一次执行是不存在的
├── playbooks                     # 剧本内容
│   ├── all.yml                   # site 文件
│   └── roles                     # 固定目录roles
│       └── base                  # demo
│           └── tasks
│               └── main.yml
├── README.md                     # README
└── start.sh                      # 入口函数
```