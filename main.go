package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type Data struct {
	message string
}
templ (d Data) Method() {
	<div>{d.message}</div>
}
templ Message(){
	<div>
	@Data{
		message:"you can implement method"
	}.Method()
	</div>
}

func main() {
//	component := hello("John")
//	http.Handle("/", templ.Handler(component))
//	fmt.Println("Listening on :3000")
//	http.ListenAndServe(":3000", nil)
Message().Render(context.Background(), os.Stdout)
}
}
