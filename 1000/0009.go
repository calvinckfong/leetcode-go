// 9. Palindrome Number
func isPalindrome(x int) bool {
    var reverse = 0
    for xx:=x; xx>0; xx/=10 {
        reverse = (xx%10) + reverse*10
    }
    return (reverse==x)
}
