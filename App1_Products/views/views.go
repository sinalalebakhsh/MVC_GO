package App1_Products

import (
    "html/template"
    "net/http"
)

func RenderHome(w http.ResponseWriter, data interface{}) {
    // // render the home page using a template
    // tmpl := `
    //     <html>
    //         <head>
    //             <title>Home Page</title>
    //         </head>
    //         <body>
    //             <h1>Welcome!</h1>
    //             <!-- Add your HTML content here -->
    //         </body>
    //     </html>
    // `
    // t := template.Must(template.New("home").Parse(tmpl))
    // t.Execute(w, data)
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../views/templates/index.html"))
		tmpl.Execute(w, nil)		
	}
	http.HandleFunc("/", h1)
}