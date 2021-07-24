## Getting started
### Import
To import the library add the following import to your code.
```go
import "github.com/nishant8887/go-qrcode"
```

### Create a new QR code
Creating a new QR code is very simple.
```go
code, err := qrcode.New("HELLO WORLD", qrcode.Q)
if err != nil {
    return err
}
```

### Get code matrix
Get the 2D boolean array representing QR code.
```go
m := code.Matrix()
```

### Get code image
Get the image form of QR code.
```go
imageFile, err := os.Create("image.png")
if err != nil {
    return err
}

img := code.Image()
err := png.Encode(imageFile, code.Image())
if err != nil {
    return err
}
```

## Documentation
For API references checkout the [documentation](https://pkg.go.dev/github.com/nishant8887/go-qrcode).

## Tools
Go-qrcode provides handy command line tool.
- **qrgenerate** - Generates QR code image from text
```
go install github.com/nishant8887/go-qrcode/cmd/qrgenerate
qrgenerate --help
```

## Links
- Nice [tutorial](https://www.thonky.com/qr-code-tutorial/) for QR code reference.
- [Guide](https://www.nayuki.io/page/creating-a-qr-code-step-by-step) for creating a QR code.
