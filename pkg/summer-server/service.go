package summerservice

type SummerService struct {

}

func (s SummerService) AddTuple(a int, b int) int{
	return a + b
}

func (s SummerService) AddTriple(a int, b int, c int) int{
	return a + b + c
}
