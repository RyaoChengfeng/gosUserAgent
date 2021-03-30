# gosUserAgent

A simple way to get User-Agent.

The source code is easy to understand and modify.

（只因为我没有找到golang好用的fake-useragent包）

user-agent.txt来源于：[这里](https://gist.githubusercontent.com/pzb/b4b6f57144aea7827ae4/raw/cf847b76a142955b1410c8bcef3aabe221a63db1/user-agents.txt)

## Install
```bash
go get -u "github.com/RyaoChengfeng/gosUserAgent" 
```

## Usage

```go
UserAgent,err := gosUserAgent.GetRandomUserAgent()
```
