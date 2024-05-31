package postgress

import (
	
	struct_test "example.com/n11/struct"
)

func Insert(firstName string,lastname string,email string,password string,age int,field string,gender string,isemploye bool) (*struct_test.User,error){
	user:=&struct_test.User{FirstName: firstName,LastName: lastname,Email: email,Password: password,Age: age,Field: field,Gender: gender,IsEmployee: isemploye}
	result:=db.Create(user)

	return user,result.Error
}