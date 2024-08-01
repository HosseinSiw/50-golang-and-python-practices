# A robot moves in a plane starting from the original point (0,0). The robot can move toward UP,
# DOWN, LEFT and RIGHT with a given steps. The trace of robot movement is shown as the following:
# UP 5
# DOWN 3
# LEFT 3
# RIGHT 2
# ¡­
# The numbers after the direction are steps. Please write a program to compute the distance from current position after a sequence of movement and original point. If the distance is a float, then just print the nearest integer.
# Example:
# If the following tuples are given as input to the program:
# UP 5
# DOWN 3
# LEFT 3
# RIGHT 2
# Then, the output of the program should be:
# 2

# Hints:
# In case of input data being supplied to the question, it should be assumed to be a console input.

# Solution
import math
__name__ = "robut"
class Robut:
    def __init__(self, init_cordinate: tuple=None) -> None:
        self.init_cordinate = (0, 0) if init_cordinate is None else init_cordinate
        self.cordintate = list(self.init_cordinate)
        
    def move_left(self, x):
        self.cordintate[0] -= x
    
    def move_right(self, x):
        self.cordintate[0] += x
    
    def move_up(self, y):
        self.cordintate[1] += y
    
    def move_down(self, y):
        self.cordintate[1] -= y
    
    def calculate_movement(self) -> int:
        xdir_movement = math.pow(self.cordintate[0] - self.init_cordinate[0], 2)
        ydir_movement = math.pow(self.cordintate[1] - self.init_cordinate[1], 2)
        movement = math.sqrt(xdir_movement + ydir_movement) 
        return movement
    
    def __str__(self) -> str:
        return f'The current cordinat is {self.cordintate}, and the total movement is {self.calculate_movement()}'
    
    
r = Robut()
r.move_down(10)
r.move_right(20)
print(r)
print(abs.__doc__)
    
    
    