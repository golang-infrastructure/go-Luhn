# Luhn 

# 一、这是什么？
一个银行卡校验算法，用于离线计算银行卡是否合法，关于Luhn的详情请移步[https://www.cnblogs.com/cc11001100/p/9357177.html](https://www.cnblogs.com/cc11001100/p/9357177.html) 

# 二、安装

```bash
go get -u github.com/golang-infrastructure/go-Luhn
```

# 三、示例
```go
package main

import (
	"fmt"
	luhn "github.com/golang-infrastructure/go-Luhn"
)

func main() {
	b := luhn.Check("8703008427912866")
	fmt.Println(b)
	// Output:
	// true

	b = luhn.Check("8703008427912861")
	fmt.Println(b)
	// Output:
	// false
}
```

