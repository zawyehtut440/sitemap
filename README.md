# Sitemap builder

## 說明
給定網址的root, 輸出該網址的sitemap

## 執行
flag domain: 輸入網址, 要是完整網址, 例如: https://www.example.com, http://www.example.com
```
$ go run main.go -domain=website_domain
```

flag depth: 參訪網址的深度, 最少等於1
```
$ go run main.go -domain=https://www.example.com -depth=1
```
上述範例會參訪https://www.example.com/網頁上的所有domain是www.example.com的網址
