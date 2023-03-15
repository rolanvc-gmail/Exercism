# Slices
## instantiating with literals

    slice = []int{0,1,2,3,4,}
    
## indexing

    fmt.Println("%d", slice[1]) // '1'

## append()

    slice = []int{1,2,3,4,}
    result=append([]int, values...)
    result = append(result, slice...)