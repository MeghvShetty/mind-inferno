def two_fer(name="you"):
    """Given name take a string in put as a parra 

    Args:
        :param name (string): name of the user 
        :returns : string
    """
    #name=name.replace("'", "''") 
    name=name.replace("<","/")
    return("One for "+name+", one for me.")
def main():
    value = "<script>this is XSS</script>"
    print(two_fer(value))
    

if __name__ == '__main__':
    main()

