# Go Programming Language

## Why go?
- Strong and statically typed
- Excellent community
- Key features
  - Simplicity
  - Fast compile times
  - Garbage collected
  - Built-in concurrency
  - Compile to standalone binaries

## Variable declaration
- Declare variable `var i int` or `var foo int = 42` or `i:=42` 
- **Value will be 0 if not initialised**
- We can also declare variables at package level but without the ":=" syntax.
- We can also declare variables like
  ```go
  var (
    actorName string = "Leonardio"
    actorAge int = 12
  )
- We can redeclare variable at diffrent scopes but at a single scope the variable will be declared only once. **Shadowing**
- Variables declared and not used will give error.
- Lower case first letter variables are spoced to the package, Locally scoped
- Upper case first letter will be exposed to the outside, Globally scoped
- Block scoped, inside the function
- Naming conventions
  - Pascal or camelCase
    - capitalized acronyms
  - As short as reasonable
    - longer names for longer lives
- We can convert the variables from one type to other by wrapping them:
  - destinationType(variable)
  - use strconv package for strings
  - We'll get an error if we try to convert type and losing the information, eg converting float to int. We'll have to do it manually.
  - **Go is very much hesitant about implicit data conversion**

## Primitive Types
- **Boolean type**
  <br/>Used as flags, logical test, 
- Numeric types
  - Integer
    - uint8, uint64, int8, int64, uintptr (both signed and unsigned, aliases are involved too)
    - There are signed counterparts for all of these as well like, int8, int64 etc.
    - signed integers are 8bit throught 64 bit and unsigned integers are 8bit through 32bit.
    - We can't perform the operation of diffrent types.
  - Floating point
    - follow IEEE-754 standards
    - 32 and 64 bit versions
    - literal styles
      - decimal(3.14)
      - exponential(13e18 or 2E10)
      - mixed(17.7e12)
  - Complex numbers
    - complex64 and complex128
    - `var n complex64 = 1+2i`
    - available operations are +, -, *, /
    - zero value is 0+0i
    - real(`<complex_number>`) and imag(`<complex_number>`).
    - `var n complex128 = complex(5, 12)`
- Text types
  - **strings**(UTF-8) are immutable. we can perform the concatenation
  - we can convert the string into bytes array for manipulation as it is very much flexible.
  - **Rune** represent UTF32 character.
  - Decalaring rune `r := 'a'`


## Constants
- `const <constant_name>` constant name follows the same name convention as that of variable naming
- constant can be shadowed as well.
- replaced by compiler at compile time. value must be calculated at compile time.
- ```go
  //This throws an error as typed constant works like immutable variables
  const a  int = 42
  var b int16  = 27
  fmt.Printf("%v, %T", a+b, a+b)

  //The program below works fine af
  const a = 42
  var b int16 = 27
  fmt.Printf("%v, %T", a+b, a+b)
  
- Enumerated Constants using iota but we need to watch out for constant values that mathches with 0. iota starts at 0 and block scoped.
  ```go
  const (
    a = iota
    b = iota
    c = iota
  )
  const (
    a2 = iota
  )

  func main() {
    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
    fmt.Printf("%v\n", a2)
  }

**Output** <br/>
0 <br/>
1 <br/>
2 <br/>
0 <br/>
This helps us to see that the declarations are block scoped.

## Array and Slices
- Declare `name := [sizeOfArray]typeOfData{value1, value2...}`
- ```go
  a := [3]int{1, 2, 3}
  a := [...]int{1, 2, 3}
  var a [3]int
- An array can be of any types say array of an array
`var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}`

**Slice**
- Declared as `a := []int{1, 2, 3}`
- Length and capacity of a slice can be measured using len() and cap().
- ```go
  a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  b := a[:] //slice all of the elements
  c := a[3:] //slice from 4th to the end
  d := a[:6] //slice first 6
  e := a[3:6] //slice [4->6]
- **Every operation with slice has an undelying array to it and the values are copied**
- `a := make([]int, 3, capacity<optional>)` to create a new slice.
- ```go
  a := [] int{}
  a = append(a, 1) //append(a, element1, element2...elementN);
  a = append(a, []int{2, 3, 4, 5}) // Throws error has appending type array instead of int  
  a = append(a, []int{2, 3, 4, 5}...) //Works as we're spreading the values over the entire array

- ```go
  a := []int{1, 2, 3, 4, 5}
  b := a[1:] //remove first ele
  c := a[:len(a)-1] //remove last ele
  d := append(a[:2], a[3:]...) //remove ele at 4

## Maps and Structs

### Maps
```go
  statePopulation := map[string]int {
    "California": 39250017,
    "Texas": 27862596,
  }
```
- Collectin of values that are accessed via keys
- Created via literals or via **make** function
- Members accessed via [key] syntax
- Check for presence with "value, ok" form of result
- Maps are **passed by reference** so the changes done in the copies of the data will be reflected in original data as well.
- An array is a valid keytype but a slice is not a valid keytype.

### Struct
- Collections of disparate data types that describes a single concept
- Keyed by named fields
- Normally created as types, but anonymous structs are allowed
- Struct are value types
- No inheritance but can use composition via embedding
- Tags can be added to struct fields to describe field

## Defer, Panic and Recover

### Defer
- Defer is executed in LIFO order.
- Executed after the main function has finished execution but before it returns.
- Defer can be used to have opening of resource and close of resource right next to each other. *Be careful in loops*.
- Arguments evaluated at time defer is executed, not at the time of called function execution.

### Panic
- Handling the state which can't be recovered.
  - Don't use when file can't be opened, unless it is critical.
  - Use for unrecoverable events.
- Function will stop executiong but defer function will still fire.
- If nothing handles panic, program will exit.
  
### Recover
- Used to recover from panics
- Only useful in deferred functions
- Current function will not attempt to continue, but higher functions will start to execute.

## Pointer
- Pointer type are declared using *** and assigned values using the *&*
- We can use the addressof operator(&) if value type already exists
  - ms := myStruct{foo: 42}
  - p := &ms
- Use addressof before initializer
  - &myStruct{foo: 42}
- Use the new keyword
  - Can't initialize fields at the same time.
- Types with internal pointers
  - All assignment operations in Go are copy operations.
  - Slices and maps contain internal pointers.

## Functions
- Basic syntax
  - ```go
      func foo() {
        ...
      }

- Parameters
    - Comma delimited list of variables and types
      - func foo(bar string, baz int)
    - Parameters of same type list type once
      - func foo(bar, baz int)
    - When pointers are passed in, the function can change the value in the caller
      - This is always true for data of slices and maps
- Return Values
  - Single return value
  - Multiple return value
    - func foo() (int, error)
    - the (result type, error) paradigm is a very common idiom
  
- Anonymous Function

## Interfaces
- Interface are a type
- type Writer interface {
      
  }