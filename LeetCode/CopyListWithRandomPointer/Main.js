/**
 * Definition for singly-linked list with a random pointer.
 * function RandomListNode(label) {
 *     this.label = label;
 *     this.next = this.random = null;
 * }
 */

/**
 * @param {RandomListNode} head
 * @return {RandomListNode}
 */
var copyRandomList = function (head) {
    if (!head) return null;
    for (let curr = head, copy; curr != null; curr = curr.next.next) {
        copy = new RandomListNode(curr.label);
        copy.next = curr.next;
        curr.next = copy;
    }
    for (let curr = head; curr != null; curr = curr.next.next) {
        if (!curr.random) continue;
        curr.next.random = curr.random.next
    }
    let copy = head.next;
    for (let curr = copy, prev = head; curr != null; prev = curr, curr = curr.next) {
        prev.next = curr.next;
    }
    return copy
};

function RandomListNode(label) {
    this.label = label;
    this.next = this.random = null;
}