# Structs
## To declare Struct type:

    type Somestruct struct{
        fieldname fieldtype
        fieldname fieldtype
    }
## To construct Struct instance
    
    func newSomeStruct() {
        return SomeStruct{fieldname: x, fieldname: y}
    }

## To access fields

    someStruct.x =
    var z = someStruct.y

## To Log to console

    fmt.Printf("formatstring: %d", varD, ...)