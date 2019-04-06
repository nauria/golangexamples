package golangexamples
import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat[]byte) string {

	var concat_string string
	
	for _,x:= range sliceToConcat {
		
		concat_string=concat_string+"-"+string(x)
		
	}
	if len(sliceToConcat)!=0{
	return concat_string[1:]
	}else{
		return	concat_string
	}
}


func Encrypt(sliceToEncrypt* []byte,ceaserCount int){
	slice:=*sliceToEncrypt
	A:=int('A')
	Z:=int('Z')
	a:=int('a')
	z:=int('z')
	
	for i,_ := range slice{
		if int(slice[i])>=A && int(slice[i])<=Z {
			slice[i]=byte((int(slice[i])-A+ceaserCount)%26+A)
		}
		if int(slice[i])>=a && int(slice[i])<=z {
			slice[i]=byte((int(slice[i])-a+ceaserCount)%26+a)
		}
	}
}

func EZGreetings(name string) string{
	return greetings.PrintGreetings(name)
}