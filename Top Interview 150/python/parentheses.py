def isValid(s):
    map_balancing = {')' : '(', ']' : '[', '}' : '{'}
    stack = ""

    for i in range(s.__len__()):
        if s[i] == '(' or s[i] == '[' or s[i] == '{':
            stack += s[i]
        elif s[i] == ')' or s[i] == ']' or s[i] == '}':
            if stack.__len__() == 0:
                return False

            par_tmp = map_balancing[s[i]]
            if stack[stack.__len__()-1] != par_tmp:
                return False
            else:
                stack = stack[:-1]

    return stack.__len__() == 0

s_1 = "()"
print(isValid(s_1))

s_2 = ")()"
print(isValid(s_2))

s_3 = ")("
print(isValid(s_3))

s_4 = "{[()]}"
print(isValid(s_4))