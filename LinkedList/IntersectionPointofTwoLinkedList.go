/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode] int, )
    temp := headA 
    for temp != nil {
            m[temp] = 1
            temp = temp.Next
    }
    tempA := headB
    for tempA != nil {
        if _, ok := m[tempA]; ok {
           return tempA
        } else {
                m[tempA] = 1
            tempA = tempA.Next
        }
    }
    return nil
}

//Input: 
/*
8
[4,1,8,4,5]
[5,6,1,8,4,5]
2
3
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
*/