#!/usr/bin/python3

class Node:
    def __init__(self, data):
        self.data = data
        self.prev = None
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None
        self.tail = None

    def get(self, data):
        if self.head == None:
            return None, False

        node = self.head
        while node != None:
            if node.data == data:
                return node, True
            else:
                node = node.next

    def push(self, value):
        self.append(value)

    def pop(self):
        if self.tail == None:
            return None, False
        else:
            node = self.tail
            self.remove(node)
            return node, True

    def traverse(self):
        if self.head == None:
            return

        node = self.head
        while node != None:
            print(node.data)
            node = node.next

    def append(self, data):
        newNode = Node(data)

        if self.head == None:
            self.head, self.tail = newNode, newNode
        else:
            newNode.prev = self.tail
            self.tail.next = newNode
            self.tail = newNode

    def insertBefore(self, item, data):
        if self.head == None or item == None:
            return

        # Create the new node
        newNode = Node(data)
        newNode.next = item
        if item == self.head:
            self.head.prev = newNode
            self.head = newNode
        else:
            newNode.prev = item.prev
            item.prev.next = newNode
            item.prev = newNode

    def insertAfter(self, item, data):
        if self.tail == None or item == None:
            return

        # Create the new node
        newNode = Node(data)
        newNode.prev = item
        if item == self.tail:
            item.next = newNode
            self.tail = newNode
        else:
            item.next.prev = newNode
            newNode.next = item.next
            item.next = newNode

    def remove(self, item):
        if item == None:
            return

        if self.head == self.tail:
            self.head, self.tail = None, None
        elif self.head == item:
            item.next.prev = None
            self.head = item.next
        elif self.tail == item:
            item.prev.next = None
            self.tail = item.prev
        else:
            item.next.prev = item.prev
            item.prev.next = item.next

        del(item)
