def fizzBuzz(number):
    assert type(number) == int, "Fizz Buzz only takes integer arguments"
    if number % 3 == 0 and number % 5 ==0:
        return "Fizz Buzz"
    elif number % 3 == 0:
        return "Fizz"
    elif number % 5 ==0:
        return "Buzz"
    else:
        return number
    

# Testing a Function  
def test_fizz_buzz_multiples_3():
    for item in [3,6,9,12,33,-10]:
        assert fizzBuzz(item) == 'Fizz', f"{item}"



print(test_fizz_buzz_multiples_3())