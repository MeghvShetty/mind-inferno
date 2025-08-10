import ast
data_file = "python/AST.py"
f = open(data_file)
data =f.read()
code = "print('Hellowworld')"
tree = ast.parse(data)
print (ast.dump(tree, indent=4))

