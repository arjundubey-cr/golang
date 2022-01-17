# GoLang

1. go run :- run the go program
2. go help :- This is helpful if you want to read docs and all or if you want to preview any command.
3. **Lexer and Lexical Elements :** 
   The formal grammar uses semicolons ";" as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:
      1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is 
         1. an identifier
         2. an integer, floating-point, imaginary, rune, or string literal
         3. one of the keywords break, continue, fallthrough, or return 
         4. one of the operators and punctuation ++, --, ), ], or }
      2. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing ")" or "}".
4. **Types in Golang**
   1. Case insensitive; almost
   2. Variable type should be knows in advance
   3. Everything is a type
   4. Umm.. What are types?
      1. String
      2. Bool
      3. Integer: uint8, uint64, int8, int64, uintptr (both signed and unsigned, aliases are involved too)
      4. Floating: float32, float64
      5. Complex 
      6. Array, Slice, Maps, Structs, Pointers
      7. Functions, Channels etc.