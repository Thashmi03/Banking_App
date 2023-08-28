package interfaces

type Itransact interface{
	Transfer(from int64,to int64,amount int64)(string,error)
}