# Things i learned
To iterate over s string, convert it first into an array of runes. A run is essentially an alias of i32, treated as a unicode character.

   ```
   r = []rune(str)
   ```

`r` is now an array of runes.
So now, we can iterate over the first half of the string from 0 to len/2 using one variable, and another variable going from len-1 to len/2. Then we swap the values of the 2 variables. 