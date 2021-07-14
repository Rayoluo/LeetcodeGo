package main

func isValid(s string) bool {
	var (
		i   int
		stk []byte
	)
	stk = make([]byte, 0)
	for i = 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stk = append(stk, s[i])
		} else {
			if s[i] == ')' {
				if len(stk) == 0 || stk[len(stk)-1] != '(' {
					break
				}
			} else if s[i] == '}' {
				if len(stk) == 0 || stk[len(stk)-1] != '{' {
					break
				}
			} else {
				if len(stk) == 0 || stk[len(stk)-1] != '[' {
					break
				}
			}
			stk = stk[:len(stk)-1]
		}
	}
	return i == len(s) && len(stk) == 0
}

// 使用栈
// 遍历字符串元素
// ---如果是左括号，则入栈
// ---如果是右括号，则和栈顶元素匹配，相同则将栈顶元素出栈，不相同则跳出for循环
// 返回i==len(s)&&len(stack)==0
