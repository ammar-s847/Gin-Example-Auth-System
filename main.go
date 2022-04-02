package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    //"errors"
)

type user struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

var userList = []user{
    {ID: "1", Name: "John", Email: "johndoe@example.com", Age: 25},
    {ID: "2", Name: "Jack", Email: "jack124@example.com", Age: 31},
}

func getUsers(c *gin.Context) {
    c.JSON(http.StatusOK, userList)
}

func getUserById(c *gin.Context) {
    index, _ := strconv.Atoi(c.Param("id"))
    c.JSON(http.StatusOK, userList[index + 1])
}

func main() {
    router := gin.Default()

    router.GET("/users", getUsers)
    router.GET("/users/:id", getUserById)

    router.Run("localhost:5000")
}

// ------ TEST CODE ------
// var x int = 12
// y := 2
// fmt.Println(x + y)
//
// asdf := true
//
// if asdf {
//     fmt.Println("as" + " df")
// }
