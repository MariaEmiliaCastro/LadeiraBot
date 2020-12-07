package main

import (
	//"flag"
	"fmt"
	"reflect"
	"strings"
	//"os"
	//"os/signal"
	//"strconv"
	//"strings"
	//"syscall"

)

//baseado nisso aqui https://golangcode.com/check-if-element-exists-in-slice/
func Find(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}


func main() {


	comando := "!corona world"

	comandos := [...]string{"world","country","states"}
	estados := [...]string{"AC", "AL", "AM", "AP", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RO", "RS", "RR", "SC", "SE", "SP", "TO"}

	fmt.Println(comando)

	query:=strings.Split(comando, " ")
	fmt.Printf("%q\n", query)


	//if len(query) >1 {
	//	if(!Find(comandos, query[1])){
	//		
	//		fmt.Printf("Comando não encontrado!")
	//		
	//	}

	//	else {
	//		//restante to codigo
	//	}
	//}


	if len(query)==1{
		fmt.Println("Query do Brasil")
	} else if len(query)==2{

		fmt.Println(query[1])
	}

	if Find(estados, "LS") {
		fmt.Println("Achou!")
	}

}

