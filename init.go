package owmjp

import "github.com/joho/godotenv"

func Init() {
	godotenv.Load(".env")
}
