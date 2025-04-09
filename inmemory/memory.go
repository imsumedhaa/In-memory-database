package inmemory


import(
	"fmt"
	"os"
	"strings"
	"bufio"
	
)

var(
	store= make(map[string]string)
 reader= bufio.NewReader(os.Stdin)
)

func NewInMemory(){
	
}

func Create(){

	fmt.Println("Enter the key:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			fmt.Println("Enter the value:")
			val,_:=reader.ReadString('\n')
			val = strings.TrimSpace(val)

			if key=="" || val==""{                               
				fmt.Println("Require the key and value.")
				os.Exit(0)
			}

			if existValue, exists := store[key]; exists {
				fmt.Println("Key already exists. Use 'update' to change the value.")               ///
				fmt.Println(existValue)
			} else {
				store[key] = val
				fmt.Println("Created successfully.")
			}

}

func Get(){
	fmt.Println("Enter the key:")
			key,_:= reader.ReadString('\n')
			key= strings.TrimSpace(key)

			if key==""{
				fmt.Println("Enter the key properly.")
				os.Exit(0)
			}

			if val,ok:=store[key];ok{
				fmt.Printf("Value is %s\n",val)
			}else{
				fmt.Println("key not found")
			}
}

func Update(){
	fmt.Println("Enter the key:")
	key,_:= reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key==""{
		fmt.Println("require the key.")
		os.Exit(0)
	}

	if _,ok:=store[key];ok{
		fmt.Println("Enter new value:")
		value,_:=reader.ReadString('\n')
		value= strings.TrimSpace(value)
		store[key]=value
		}else{
		fmt.Println("Key not found.")
	}
}

func Delete(){
	fmt.Println("Enter the key you want to delete:")
	key,_:=reader.ReadString('\n')
	key=strings.TrimSpace(key)

	if key==""{
		fmt.Println("require the key.")
		}else{
			delete(store,key)
			fmt.Println("Succesfully deleted")
		}
}

func Show(){
	fmt.Println("The full map is:")
	fmt.Println(store)
}

func Exit(){
	fmt.Println("Exiting program.")
    os.Exit(0)
}



