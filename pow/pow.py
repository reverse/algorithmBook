def PowX(x, n):
    if n == 0:
        return 0
    
    if n == 1:
        return x 

    if n < 0:
        return 1 / PowX(x, -n)
    
    if n%2==0:
        return PowX(square(x), n/2)

    return x * PowX(x, n-1)


def square(x):
    return x*x