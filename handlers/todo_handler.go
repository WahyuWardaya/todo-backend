package handlers

import (
	"log"
	"net/http"
	"strconv"

	"todo-backend/database"
	"todo-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	rows, err := database.DB.Query("SELECT id, text, completed FROM todos")
	if err != nil {
		log.Println("Error fetching todos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.Completed); err != nil {
			log.Println("Error scanning todo:", err)
			continue
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.DB.Exec("INSERT INTO todos (text, completed) VALUES (?, ?)", todo.Text, todo.Completed)
	if err != nil {
		log.Println("Error creating todo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get new todo ID"})
		return
	}
	todo.ID = int(id)

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = id // Pastikan ID di model sesuai dengan ID dari URL

	_, err = database.DB.Exec("UPDATE todos SET text=?, completed=? WHERE id=?", todo.Text, todo.Completed, id)
	if err != nil {
		log.Println("Error updating todo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		log.Println("Error deleting todo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.Status(http.StatusNoContent)
}