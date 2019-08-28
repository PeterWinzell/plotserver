package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// MAIN SECTION HTML CODE
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	fmt.Fprintf(w, "<title>Go</title>")
	fmt.Fprintf(w, "<img src='assets/plotter.png' alt='gopher' style='width:1024px;height:800px;'>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	// ABOUT SECTION HTML CODE
	fmt.Fprintf(w, "<title>Go/about/</title>")
	fmt.Fprintf(w, "Plotting data from signal broker")
}

func dynamicHandler(w http.ResponseWriter, r *http.Request){
	html := `<!DOCTYPE html> 
				<html> <title>Plotter...</title> 
				<head> 
					<p><h2>Plotter</h2></p>
				</head>

				
                <script> 
						function refresh() {
							var x = location.pathname	
							document.getElementById("demo").innerHTML = x;
							document.getElementById('plotter').src = 'plotter.png';
						}
                	</script>

				<body onLoad="setInterval(refresh, 1000)">
                    <p id="demo"></p>
					<img id="plotter" height="800" width=""/>
				</body>

			</html>`
	fmt.Fprintf(w,html);

}

func main() {
	http.HandleFunc("/mupp", index_handler)
	http.HandleFunc("/about/", about_handler)
	//http.HandleFunc("/", dynamicHandler)

	http.Handle("/",http.FileServer(http.Dir("./assets")));

	http.ListenAndServe(":9000", nil)
}
