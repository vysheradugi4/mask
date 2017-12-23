# Mask
![alt text](https://github.com/vysheradugi4/mask/blob/master/Gomask.png)

This is generator contains two steps:
1. Create mask from exists word with use keyword (part of word) replace keyword to
'key'.
2. From mask generate many words with use keyword and replace regular expressions
to chars or numbers.

## Installation guide

This project for Go (golang). For install Mask into your project run in console:

```
go get https://github.com/vysheradugi4/mask
```

## Users guide

1. Creating mask:

```
// Part of word what need to save into results.
keyword := "winter"

// So this is word.
word := "winter18"

// Mask is the mask, err is the error.
mask, err := mask.CreateMask(code, key)
```
Examples of mask: key\d\d, \d\dkey[a-z][A-Z]
Be careful, the system is case sensitive. And in result may give diffrent keys:

key			keyword in word in lowercase
Key			keyword in word with first character in uppercase
KEY			keyword in word with all characters in uppercase
kEy			keyword in word with camelcase

And supported regexp:
\d			digit char
[a-z]		lowercase character (NOT SUPPORTED)
[A-Z]		uppercase character (NOT SUPPORTED)


2. Creating codes, now its easy:

```
// Last argument is how many digits creating recursive function.
result := mask.GenerateCodesFromMask(mask, keyword, 3)
```

Don't forget import this package.

```
import (
    mask
    )
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
