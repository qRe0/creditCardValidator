# Introduction
Here you can find my small pet projet. It uses `Luhn algorithm` to check validity of entered Card Number

## Language
This project was written on `Golang`

## Used packages:
- `buffio` to read input data
- `fmt` - standar console dialog package
- `os` - suuport package for `buffio`
- `regexp` to check format of CC number
- `strings` to cut all spaces from string to use Lyhn alg.

## Luhn Algorithm
1. The digits of the sequence to be checked are numbered from right to left.
2. The digits on odd places remain unchanged.
3. The digits on the even places are multiplied by 2.
4. If such multiplication results in a number greater than 9, it is replaced by the sum of the digits of the resulting product - a one-digit number, i.e. a digit.
5. All the digits obtained as a result of the conversion are added together. If the sum is a multiple of 10, the original data are correct.

## Future of the project
The project will soon be migrated from console to web application. The project will also be linked to a database (`Postgres`).
