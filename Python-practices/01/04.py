"""
This program is uqual to 50-golang-and-python-practices/go/01/04.go
"""

s = input("please enter a text: ")
d={"UPPER CASE":0, "LOWER CASE":0}
for c in s:
	if c.isupper():
		d["UPPER CASE"]+=1
	elif c.islower():
		d["LOWER CASE"]+=1
	else:
		pass
print ("UPPER CASE", d["UPPER CASE"])
print ("LOWER CASE", d["LOWER CASE"])