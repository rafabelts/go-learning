package arrays
func Sum(numbers []int) int {
    var sum int
    
    /*
    range permite iterar sobre un arreglo.

    En cada iteracion devuelve dos valores, el indice y el valor, el indice se ignora con _
    */

    for _, number := range numbers {
        sum += number
    }

    return sum
}


func SumAll(numbersToSum ...[]int) []int {
//    lenghtOfNumbers := len(numbersToSum)

    /*
    
    make permite hacer un slice con una capacidad inicial.

    El length es la cantidad de elementos que puede contener un slice, mientras que
    la capacidad es el numero de elementos que puede contener en la matriz subyacente, ejemplo,
    make([]int, 0, 5) crea un slice con longitud 0 y capacidad 5
    */

    /*

    sums := make([]int, lenghtOfNumbers)

    for i, array := range numbersToSum {
        sums[i] = Sum(array)
    } 
    */

    /*
    Tambien se puede usar el append el cual toma como argumentos un slice y un nuevo valor, despues
    retorna un nuevo slice con todos los elementos
    */

    var sums []int

    for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers))
    }

	return sums
}

func SumAllTails(arrays ...[]int)[]int{
    var tailsSum []int

    for _, numbers := range arrays {
        
        if len(numbers) == 0 {
            tailsSum = append(tailsSum, 0)
            continue
        }


        tails := numbers[1:]
        
        tailsSum = append(tailsSum, Sum(tails))

    }

    return tailsSum
}
