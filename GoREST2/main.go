package main

func main() {
	app := &App{}
	app.Initialize(DB_Username, DB_Pass, DB_Name)
	app.Run("localhost:8000")
}
