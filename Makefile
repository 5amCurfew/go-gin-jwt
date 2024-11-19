rebuild:
	rm -f *.db
	sqlite3 auth.db "select true;"
	go run main.go