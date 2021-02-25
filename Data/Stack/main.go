package main

import "fmt"

func main() {
	a := isValid("([)]")
	fmt.Println("Valid is", a)
}

func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    if len(s) == 1 {
        return false
    }
    st := &stack {
        r : rune(s[0]),
    }
    for i := 1; i < len(s); i++ {
        switch rune(s[i]) {
            case '(' : 
				fmt.Println("pushing ( ")
                push('(', st)
            case ')' : 
                ret := pop(st)
				fmt.Println("poped", rune(ret))
                if ret != '(' {
                    return false
                }
            case '[' : 
				fmt.Println("pushing [ ")
                push('[', st)
            case ']' : 
                ret := pop(st)
				fmt.Println("poped", rune(ret))
                if ret != '[' {
                    return false
                }
            case '{' : 
				fmt.Println("pushing { ")
                push('{', st)
            case '}' : 
                ret := pop(st)
				fmt.Println("poped", rune(ret))
                if ret != '{' {
                    return false
                }
        }
    }
    return true
}

type stack struct {
    r       rune
    last    *stack
}

func push(ru rune, s *stack) {
    p := &stack {
        r : ru,
        last : s,
    }
    s = p
}

func pop(s *stack) rune{
    if s == nil {
        return '-'
    }
    ret := s.r
    s = s.last
    return ret
}