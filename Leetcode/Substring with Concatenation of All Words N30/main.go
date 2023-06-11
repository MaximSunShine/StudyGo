package main
import "fmt"

func find(s string, a[]string) bool{
	
	if len(s) > 0 {
		
	}
	return true
}

func is(s string, a []string) (int, bool){
	for i, val := range a {
		if s == val {
			return i, true
		}
	}
	return -1, false
	
}

func f(s string, a []string) []int {
	
	if i, ok := is(s[:3],a); ok{
		
		fmt.Println("-> ", s[3:], append(a[:i], a[i+1:]...) )
		//f(s[3:],a)
	}
	
	hash := make(map[string]struct{})
	
	for _, val := range a {
		hash[val] = struct{}{}
	}
	
	for val := range hash {
		fmt.Println(val)
	}
	
	return []int{1}
}
func main() {
	fmt.Println( f("abcabc", []string{"abc","def"}) )

}
