package crudproduct

import (
	model "Projects/N11/modelUser"
	"database/sql"
)

type CrudProductRepo struct{
	DB *sql.DB
}

func NewCrudProductRepo(DB *sql.DB) *CrudProductRepo{
	return &CrudProductRepo{DB}
}

func (i *CrudProductRepo) Insertproduct(name,description string,price float64,stock_quantity int) ([]model.Products,error) {
	tr,err:=i.DB.Begin()
	if err!=nil{
		return nil,err
	}
	defer tr.Commit()

	_,err=i.DB.Exec("insert into products(name,description,price,stock_quantity) values($1,$2,$3,$4)",name,description,price,stock_quantity)
	if err!=nil{
		return nil,err
	}
	return []model.Products{},nil
}

func (r *CrudProductRepo) GetProduct() ([]model.Products,error){
	tr,err:=r.DB.Begin()
	if err!=nil{
		return nil,err
	}
	defer tr.Commit()

	row,err:=r.DB.Query("select * from products")
	if err!=nil{
		return nil,err
	}

	var modelsProducts []model.Products
	for row.Next(){
		var modelProduct model.Products
		err:=row.Scan(&modelProduct.Id,&modelProduct.Name,&modelProduct.Description,&modelProduct.Price,&modelProduct.Stock_quantity)
		if err!=nil{
			return nil,err
		}
		modelsProducts=append(modelsProducts, modelProduct)
	}
	return modelsProducts,nil
}

func (up *CrudProductRepo) Updateproduct(id int,name string) ([]model.Products,error){
	tr,err:=up.DB.Begin()
	if err!=nil{
		return nil,err
	}
	defer tr.Commit()

	_,err=up.DB.Exec("update products set name=$1 where id=$2",name,id)
	if err!=nil{
		return nil,err
	}

	return []model.Products{},nil
}

func (d *CrudProductRepo) Deleteproduct(id int) ([]model.Products,error){
	tr,err:=d.DB.Begin()
	if err!=nil{
		return nil,err
	}
	defer tr.Commit()

	_,err=d.DB.Exec("delete from products where id=$1",id)
	if err!=nil{
		return nil,err
	}
	return []model.Products{},nil
}