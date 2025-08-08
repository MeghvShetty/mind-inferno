'''
The Collatz Conjecture or 3x+1 problem can be summarized as follows:

Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely. The conjecture states that no matter which number you start with, you will always reach 1 eventually.

Given a number n, return the number of steps required to reach 1.'''
def steps(n):
    # validate the input value 
    if n<= 0 :
        raise ValueError("Only positive integers are allowed")
    
    stepsCount = 0  # iterative count of the steps

    while n!= 1:
        if (n% 2) == 0:
            n = n / 2
            stepsCount +=1
        else:
            n = n*3+1
            stepsCount += 1 
    
    return stepsCount

def main():
    print(steps(16)) 


if __name__ == "__main__":
    main()
