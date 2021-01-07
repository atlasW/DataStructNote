package com.sona.architecture.symboltable;

import edu.princeton.cs.algs4.LinkedQueue;


/**
 * @author renfakai
 * @version 1.0
 * @since 2021/1/5
 */
public class SequentialSearchST<Key, Value> implements ST<Key, Value> {

    private int n;

    private Node first;

    //  内部数据结构
    private class Node {
        private final Key k;
        private Value v;
        private Node next;

        public Node(Key k, Value v, Node next) {
            this.k = k;
            this.v = v;
            this.next = next;
        }
    }

    @Override
    public int size() {
        return n;
    }

    @Override
    public boolean isEmpty() {
        return size() == 0;
    }

    @Override
    public boolean containsKey(Key k) {
        if (null == k) throw new IllegalArgumentException("key must not null");
        return get(k) != null;
    }

    @Override
    public Value get(Key k) {
        if (null == k) throw new IllegalArgumentException("key must not null");
        //  first 可能为空
        for (Node x = first; x != null; x = x.next) {
            if (x.k == k) {
                return x.v;
            }
        }
        return null;
    }

    @Override
    public void put(Key k, Value v) {
        // 为什么我插入数据的时候老想着从后面插入呢，囧
        if (null == k) throw new IllegalArgumentException("key must not null");
        if (null == v) {
            delete(k);
            return;
        }
        // 如果已经存在数据
        for (Node x = first; x != null; x = x.next) {
            if (x.k.equals(k)) {
                x.v = v;
                return;
            }
        }
        first = new Node(k, v, first);
        n++;
    }

    @Override
    public void delete(Key k) {
        if (null == k) throw new IllegalArgumentException("key must not null");
        first = delete(first, k);
    }

    private Node delete(Node x, Key k) {
        if (x == null) return null;
        if (k.equals(x.k)) {
            n--;
            return x.next;
        }
        x.next = delete(x.next, k);
        return x;
    }


    @Override
    public Iterable<Key> keys() {
        LinkedQueue<Key> keysQ = new LinkedQueue<>();
        for (Node x = first; x != null; x = x.next) {
            keysQ.enqueue(x.k);
        }
        return keysQ;
    }


    public static void main(String[] args) {
        SequentialSearchST<Integer, Integer> s = new SequentialSearchST<>();
        for (int i = 0; i < 10; i++) {
            s.put(i, i);
        }
        System.out.println(s.keys());
    }
}
