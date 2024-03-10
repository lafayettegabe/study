// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var current *ListNode = &ListNode{Val: 0, Next: nil}
  var result  *ListNode = current

  for i:=0; l1 != nil || l2 != nil; i++ {
    if l1 != nil {
      current.Val += l1.Val
      l1 = l1.Next
    }
    if l2 != nil {
      current.Val += l2.Val
      l2 = l2.Next
    }
    if l1 != nil || l2 != nil {
      current.Next = &ListNode{Val: 0, Next: nil}
      if current.Val > 9 {
        current.Val -= 10
        current.Next.Val = 1
      }
      current = current.Next
    } else if current.Val > 9 {
      current.Next = &ListNode{Val: 1, Next: nil}
      current.Val -= 10
    }
  }
  return result
}
// @leet end
