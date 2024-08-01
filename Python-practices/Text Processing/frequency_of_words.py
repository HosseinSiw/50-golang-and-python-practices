# This exercise 
text = input("Enter you msg: ")

def check_words(text :str = 'Hello Hello golang py py python py') -> dict:
    # Creating a list of seperate words
    text = text.split(" ")
    set_of_words = set(text)  
    d = dict()
    for word in set_of_words:
        d[word] = 0
    
    for word in text:
        if word in set_of_words:
            d[word] += 1
    return d

print(check_words()) 