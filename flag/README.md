# Usage

```
// go run main.go [OPTIONS]
// Options:
//  -e string   This is used by CAs, such as Let's Encrypt, to notify about problems with issued certificates.
//  -h string   This is used by autocert, controls which domains the Manager will attempt to retrieve new certificates for.
$ go run main.go -e youremail@gmail.com -h yourdomain.com
```