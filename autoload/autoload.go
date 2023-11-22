package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/g41797/godotenv/autoload"
*/

import "github.com/g41797/godotenv"

func init() {
	godotenv.Load()
}
