# Question: Write a programm that gives an input from the console and calculates the frequency of words within that.


# Hint: For example you give "Hello Hello World Python Rust World Python" as the input
# then the output should be:
# [Hello]: 2
# [World]: 2
# [Python]: 2
# [Rust]: 1

text = input("Enter your msg: ")

def check_words(text :str) -> dict:
    # Creating a list of seperate words
    text = text.split(" ")
    set_of_words = set(text)  
    d = dict()
    for word in set_of_words:
        d[word] = 0
    
    for word in text:
        if word in set_of_words:
            d[word] += 1
    
    for k, v in d.items():
        print(f"[{k}]: {v}", )


if __name__ == "__main__":
    check_words(text) 