package main

import (
	crudproduct "Projects/N11/crudProduct"
	crudusers "Projects/N11/crudUser"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main(){
	db,err:=sql.Open("postgres","postgres://postgres:salom@localhost:5432/n11?sslmode=disable")
	if err!=nil{
		panic(err)
	}
	defer db.Close()

	tr,err:=db.Begin()
	if  err!=nil {
		panic(err)
	}

	defer tr.Commit()

	u:=crudusers.NewCrudUsersRepo(db)
	p:=crudproduct.NewCrudProductRepo(db)

	// _,err=u.Insertuser("Ozodbek","ozodbek@gmail.com","ozodbek")
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=u.Insertuser("Azizbek","aziz123@gmail.com","azizbek123")
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=u.Insertuser("Abbos","abbos2321@gmail.com","abbos2321")
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=u.Insertuser("Saidazim","saidazim12@gmail.com","saidazim12")
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=u.Insertuser("Asilbek","asilbek1911@gmail.com","asilbek1911")
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=p.Insertproduct("Laptop", "High-performance laptop with 16GB RAM and 512GB SSD", 1299.99, 50)
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=p.Insertproduct("Smartphone", "Latest model smartphone with 5G support and 128GB storage", 799.49, 150)
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=p.Insertproduct("Headphones", "Wireless noise-cancelling headphones with 30-hour battery life", 199.95, 200)
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=p.Insertproduct("Smartwatch", "Smartwatch with heart rate monitor and GPS", 149.99, 100)
	// if err!=nil{
	// 	panic(err)
	// }

	// _,err=p.Insertproduct("Tablet", "10-inch tablet with 64GB storage and 8MP camera", 349.99, 80)
	// if err!=nil{
	// 	panic(err)
	// }

	// getuser,_:=u.Readuser()
	// fmt.Println(getuser)

	// fmt.Println()

	getproduct,_:=p.GetProduct()
	fmt.Println(getproduct)

	_,err=u.UpdateUser("ozodbek1235",1)
	if err!=nil{
		panic(err)
	}

	_,err=p.Updateproduct(2,"Tablet1232")
	if err!=nil{
		fmt.Println("salom")
		panic(err)
	}

	_,err=u.Deleteuser(3)
	if err!=nil{
		fmt.Println("salom1")
		panic(err)
	}

	_,err=p.Deleteproduct(5)
	if err!=nil{
		fmt.Println("salom2")
		panic(err)
	}
}