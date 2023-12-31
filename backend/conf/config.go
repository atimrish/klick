package conf

var Config = map[string]string{
	//POSTGRES
	"postgres_address":  "postgres",
	"postgres_user":     "root",
	"postgres_password": "root",
	"postgres_database": "social_network",
	"postgres_port":     "5432",

	//MONGO
	"mongo_address":  "mongo",
	"mongo_user":     "root",
	"mongo_password": "example",
	"mongo_port":     "27017",

	//JWT SECRET
	"jwt_secret": "ZTNiMGM0NDI5OGZjMWMxNDlhZmJmNGM4OTk2ZmI5MjQyN2FlNDFlNDY0OWI5MzRjYTQ5NTk5MWI3ODUyYjg1NQ==",
}

func GetConfig() map[string]string {
	return Config
}
