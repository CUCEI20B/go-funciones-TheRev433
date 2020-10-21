package main

func Burbuja(s []int64)  {
	for i := 0; i <len(s); i++ {
		for j:= 0; j < len(s)-1; j++{
			e := s[j]
			if e > s[j+1]{
				s[j] = s[j+1]
				s[j+1] = e
			}
		}
	}
}

func main()  {
}