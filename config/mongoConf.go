package config

func TestMongo() string {
	return ""
}

func DevMongo() string {
	return "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"
}

func ProMongo() string {
	return ""
}
