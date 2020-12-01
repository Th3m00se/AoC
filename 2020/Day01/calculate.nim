# Written for https://adventofcode.com/2020/day/1

import streams, strutils

var 
    #InputStream = newFileStream("sample_input.txt", fmRead)
    InputStream = newFileStream("full_input.txt", fmRead)
    line = ""
    Val1 = 0
    Val2 = 0
    Val3 = 0 # Added for Part 2
    InputArray: seq[int]

if not isNil(InputStream):
    while InputStream.readLine(line):
        # echo line
        InputArray.add(parseInt(line))
InputStream.close()

for i, value1 in InputArray:
    Val1 = value1

    for j, value2 in InputArray:
        Val2 = value2
        
        # Addint another loop for part 2
        for k, value3 in InputArray:
            Val3 = value3

            if (Val1+Val2+Val3) == 2020:
                echo "index ", $i, "+", $j, "+", $k, "=2020"
                echo "so the multiplied product is ", intToStr(Val1*Val2*Val3)
                break
