/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    nums := []int{}

    for head != nil{
        nums = append(nums,head.Val)
        head = head.Next
    }

    l,r := 0,len(nums)-1
    for l < r {
        if nums[l] != nums[r]{
        return false
        }else{
            l++
            r-- 
        }
    }
    return true
}