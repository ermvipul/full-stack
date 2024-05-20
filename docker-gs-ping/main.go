/**** Uncomment below code if you wanna run on server that is with frontend & backend config not wordpress app ***///


// package main

// import (
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {
// 	e := echo.New()

// 	// Use the default CORS middleware
// 	e.Use(middleware.CORS())

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Docker! <3"})
// 	})	

// 	e.GET("/health", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
// 	})

// 	httpPort := os.Getenv("PORT")
// 	if httpPort == "" {
// 		httpPort = "1190"
// 	}

// 	e.Logger.Fatal(e.Start(":" + httpPort))
// }



/**** Uncomment below code if you wanna run docker-compose config with wordpress app ***///

package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

type Post struct {
    ID    int    `json:"id"`
    Title struct {
        Rendered string `json:"rendered"`
    } `json:"title"`
    Content struct {
        Rendered string `json:"rendered"`
    } `json:"content"`
}

func main() {
    http.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
        // Add CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // Fetch posts
        posts := fetchPosts()

        // Set Content-Type header
        w.Header().Set("Content-Type", "application/json")
        
        // Encode and write posts to response
        json.NewEncoder(w).Encode(posts)
    })

    log.Println("Server started on :1190")
    log.Fatal(http.ListenAndServe(":1190", nil))
}

func fetchPosts() []Post {
    resp, err := http.Get("http://wordpress/wp-json/wp/v2/posts")
    if err != nil {
        log.Fatalf("Error fetching posts: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }

    var posts []Post
    err = json.Unmarshal(body, &posts)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON: %v", err)
    }

    return posts
}
