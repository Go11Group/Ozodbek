package main

import (

	postgress "example.com/n11/postgres"
)

func main() {
	_, err := postgress.ConnectDB()

	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Ozodbek", "Shamsiddinov", "ozodbek@gmail.com", "12345", 19, "Golang and python", "Male", true)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Saidazim", "Qilichov", "saidazim@gmail.com", "54321", 18, "Golang", "Male", false)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Akbar", "Akbarov", "akbar@gmail.com", "012", 20, ".net", "Male", true)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Asilbek", "Asilbekov", "asilbek@gmail.com", "1911", 16, "python", "Male", false)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Abbos", "Abbosbekov", "Abbos@gmail.com", "a123", 24, "Flutter", "Male", true)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Nigini", "Abdusalimov", "nigini@gmail.com", "n12354", 17, "Dezayner", "Female", false)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Ali", "Alijonov", "ali@gmail.com", "aa45", 19, "python", "Male", false)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Aziz", "Mansurov", "aziz@gmail.com", "aziz12", 23, "Golang", "Male", true)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Maruf", "Mansurov", "maruf@gmail.com", "maruf12", 23, "Golanf and python", "Male", false)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Umidjon", "Husenov", "umid@gmail.com", "umid", 20, "dezayner", "Male", true)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Ravshan", "Ravshanov", "ravshan@gmail.com", "ravshan", 15, "c++", "Male", false)
	if err != nil {
		panic(err)
	}

	_, err = postgress.Insert("Hasan", "Hasanboyev", "hasan@gmail.com", "hasan", 28, "Java", "Male", true)
	if err != nil {
		panic(err)
	}
}
