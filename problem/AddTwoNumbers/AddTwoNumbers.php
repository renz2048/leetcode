<?php
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $dummyHead = new ListNode();
        $p = $l1;
        $q = $l2;
        $cunn = $dummyHead;
        $carry = 0;
        while (!empty((array)$p) || !empty((array)$q)) {
            echo "haha";
            $x = (!empty((array)$p)) ? $p->val : 0;
            $y = (!empty((array)$q)) ? $q->var : 0;
            $sum = $carry + $x + $y;
            $carry = $sum / 10;
            $curr->next = new ListNode($sum % 10);
            $curr = $curr->next;
            if (!empty((array)$p))
                $p = $p->next;
            if (!empty((array)$q))
                $q = $q->next;
        }
        if ($carry > 0) {
            $curr->next = new ListNode($carry);
        }
        return $dummyHead->next;
    }
}
?>
