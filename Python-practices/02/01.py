# Question:
# Write a program that computes the value of a+aa+aaa+aaaa with a given digit as the value of a.
# Suppose the following input is supplied to the program:
# 9
# Then, the output should be:
# 11106
def funcEx04(number: int) -> int:
    result = 0
    for i in range(1, number):
        value = number ** i
        result += value
    
    return result


print(funcEx04(5))