# igorm-mock-template

Wrapper template for [igorm](https://github.com/kacperjurak/igorm)

Creating your wrapper can be tiring because it has to implement 74 functions from the Gormw interface. That's why I did it for you. Fill out the methods with your code. See, for example, the First () function. 

### Usage

Do you remember your function which as an argument accepts any structure implementing the igorm.Gormw interface?

```go
func getUser(db igorm.Gormw) *User {
}
```
You've used a wrapper that supports native [gorm](https://github.com/jinzhu/gorm) package functions.

```go
db, err := igorm.Openw(dialect, path)
if err != nil {
    log.Fatal(err)
}
```

and you used it with your example getUser() function.

```go
user := getUser(db)
```

Now you are going to make your own wrapper. This package is named `template` but copy/paste the code to your own package, implement functions that you need and voila.

### Example with existing code

```go
db, err := template.New("I'm happy value")
if err != nil {
    log.Fatal(err)
}

user := getUser(db)
```

Look for example function First(). Every call of this function on your mocked database gives:

> Hey! First() method works, and the someField value is: I'm happy value

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
