C = [17, 22, 24, 30, 34]

def to_fahrenheit(x):
    return float(9/5)*x+32

def to_celsius(x):
    return (float(5)/9)*(x-32)
    
F = list(map(lambda x: to_fahrenheit(x), C))
print(F)

C = list(map(lambda x: to_celsius(x), F))
print(C)

