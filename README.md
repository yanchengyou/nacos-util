# nacos-util
nacos 工具类，主要功能包括查看服务列表，服务上线线等

# 使用方法

## 配置补齐

```shell script
echo  "source  <(nacos-util completion bash)"  >>  ~/.bashrc
```

## nacos配置

### 配置nacos地址

```shell script
nacos-util config set-host 192.168.0.100:8848
```

### 配置nacos用户

```shell script
nacos-util config set-username nacos
```

### 配置nacos密码

```shell script
nacos-util config set-password nacos
```

### 测试nacos连接

```shell script
nacos-util config connect-test
```

### 查看当前配置

```shell script
nacos-util config view
```