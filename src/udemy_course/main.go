package main

import "fmt"

func main() {
	name := "Mahidul Moon"

	html := `
		<!DOCTYPE html>
		<html lang="en">
		<title>Hello World</title>
		<body>
			<h1>` + name + `</h1>
		</body>
		</html>
	`
	fmt.Println(html)
}
