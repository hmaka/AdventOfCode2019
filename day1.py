sum = 0
with open("day1.txt", 'r') as file:
    for line in file:
        
        n = int(line)
        sum += (n//3) - 2
print(sum)    
