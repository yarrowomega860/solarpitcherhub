package main
import ("errors";"fmt";"strings")
var moduleName = "config-vault-bd78e8"
type AppError struct{Op string;Err error}
func (e *AppError) Error() string{return fmt.Sprintf("[%s] %s: %v",moduleName,e.Op,e.Err)}
func (e *AppError) Unwrap() error{return e.Err}
var ErrNotFound=errors.New("not found")
var ErrInvalid=errors.New("invalid input")
func lookup(key string) (string,error){if key==""{return "",&AppError{Op:"lookup",Err:ErrInvalid}};if !strings.HasPrefix(key,"k-"){return "",&AppError{Op:"lookup",Err:ErrNotFound}};return fmt.Sprintf("value-for-%s",key),nil}
func main(){for _,k:=range []string{"k-abc","","xyz"}{v,err:=lookup(k);if err!=nil{fmt.Println(err);if errors.Is(err,ErrNotFound){fmt.Printf("  -> retryable\n")}}else{fmt.Printf("[%s] %s = %s\n",moduleName,k,v)}}}
