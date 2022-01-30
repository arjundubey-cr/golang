# GoLang

`go run` :- run the go program <br/>
`go help` :- This is helpful if you want to read docs and all or if you want to preview any command.
### 1. Lexer and Lexical Elements 
   The formal grammar uses semicolons ";" as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:
   - When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is 
     - an identifier
     - an integer, floating-point, imaginary, rune, or string literal
     - one of the keywords break, continue, fallthrough, or return 
     - one of the operators and punctuation ++, --, ), ], or }
   - To allow complex statements to occupy a single line, a semicolon may be omitted before a closing ")" or "}".
### 2. Types in Golang
   - Case insensitive; almost
   - Variable type should be knows in advance
   - Everything is a type
   - So.. What are types?
      - String
      - Bool
      - Integer: uint8, uint64, int8, int64, uintptr (both signed and unsigned, aliases are involved too)
      - Floating: float32, float64
      - Complex 
      - Array, Slice, Maps, Structs, Pointers
      - Functions, Channels etc.
   - We can declare variable in the following syntax:
            ```var <VARIABLE_NAME> <VARIABLE_TYPE> = <VALUE>``` <br/>
                        OR<br/>
            ```var <VARIABLE_NAME> = <VALUE>``` <br/>
                        OR<br/>
            ```<VARIABLE_NAME> := <VALUE>``` <br/>
### 3. User Inputs 
   - We use libraries to take user input. We declare a reader and use the same to handle the user input.
   - Concerned libraries are **bufio** and **os**
   - comma-okay syntax or err-ok syntax is used to manage the input taken from the user.
### 4. Type Conversion
   -  Just converting your string input directly to the float or number won't work because the '\n', i.e. "Enter Key" is also recorded in the input so you'll recieve an error like `strconv.ParseFloat: parsing "4\n": invalid syntax`
   -  The packages used here are "strconv" to conver the string and the "strings" to eliminate the trailing space at the end of the string.

### 5. Handling Time
   - There are some constants in time values which are used in Go to specify formats, eg:
   - ```
      Default format: 2015-02-25 11:06:39 -0800 PST
      Unix format: Wed Feb 25 11:06:39 PST 2015
      Same, in UTC: Wed Feb 25 19:06:39 UTC 2015

      Formats:

      Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
      Basic short date "2006/01/02" gives "2015/02/25"
      AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
      No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
      0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
      9s for fraction  "15:04:05.99999999" gives "11:06:39.1234" 