# try:
#     pass
# except Exception:
#     pass
# else:
#     pass
# finally:
#     pass


print("Sample 1")
# Sample 1
try:
    f = open('Python/POC/defensive_programming/test_file.txt')
    print("This file does exit")
except Exception:
    print("Sorry this file don't exit")
    
print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
print("Sample 1")
# Sample 2
try:
    f = open('Python/POC/defensive_programming/test_file.txt')
    var = bad_var
except FileNotFoundError as e: # Added more specific exception at the top and more generic at the bottom
    print(e)

except Exception as e:
    print(e)
    
print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
print("Sample 1")
# Sample 3
try:
    f = open('Python/POC/defensive_programming/test_file.txt')
    
except FileNotFoundError as e: # Added more specific exception at the top and more generic at the bottom
    print(e)

except Exception as e:
    print(e)

else: # only runs if we don't throw an exception
    print(f.read())
    f.close()
    
finally: # this will run no matter if the code throws an exception or not.
    print("Executing Finally...")
    

print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
print("Sample 4")
# Sample 4 rasining your own expection.
try:
    f = open('Python/POC/defensive_programming/test_file.txt')
    if f.name == "Python/POC/defensive_programming/test_file.txt":
        raise Exception # manually rasining 
except FileNotFoundError as e: # Added more specific exception at the top and more generic at the bottom
    print(e)

except Exception:
    print("Error!!!!!!!")

else: # only runs if we don't throw an exception
    print(f.read())
    f.close()
    
finally: # this will run no matter if the code throws an exception or not.
    print("Executing Finally...")
    
print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
