package todo

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(c *gin.Context) {
	var todoData CreateTodoRequest
	if err := c.Bind(&todoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}
	// userData, ok := c.Get("user")

	// if !ok {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }
	// user, ok := userData.(models.User)
	// if !ok {
	// 	fmt.Println("Failed to convert user data to User struct")
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(user)

	err := c.Request.ParseMultipartForm(10 << 20) // Max size: 10 MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}
	form := c.Request.MultipartForm
	values := form.Value

	// Print key-value data
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Retrieve all files from the form data
	filesMap := c.Request.MultipartForm.File
	var files []*multipart.FileHeader
	for _, fileHeaders := range filesMap {
		files = append(files, fileHeaders...)
	}
	// fmt.Println("FILES", filesMap)
	// Process each file
	// for _, file := range files {
	// 	// Open the file
	// 	f, err := file.Open()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
	// 		return
	// 	}
	// 	defer f.Close()

	// 	// Print file details
	// 	fmt.Println("File Name:", file.Filename)
	// 	fmt.Println("File Size:", file.Size)
	// 	fmt.Println("File MIME Type:", file.Header.Get("Content-Type"))

	// 	// Process the file as needed (e.g., save it to disk)
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Data and files uploaded successfully"})
}
