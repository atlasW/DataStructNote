package com.sona.architecture.algorithms;

import java.util.Arrays;

/**
 * 算法书中更简单的讲解了数据的exchange
 * doug lea更加牛逼的子令集的减少
 * {@link java.util.PriorityQueue}
 * {@see https://algs4.cs.princeton.edu/24pq/ }
 */
public class PriorityQueueS implements QueueS {

    public PriorityQueueS() {
        size = 0;
        this.queue = new Comparable[5];
    }

    private int size;

    public Comparable[] queue;

    public PriorityQueueS(Comparable[] c) {
        this.queue = c;
        int leafStart = (c.length >> 1) - 1;
        // 堆话
        for (int i = leafStart; i >= 0; i--) {
            shiftDown(i, queue[i]);
        }
    }

    public static void main(String[] args) {
        PriorityQueueS priorityQueueS = new PriorityQueueS();
        priorityQueueS.insert("1");
        priorityQueueS.insert("3");
        priorityQueueS.insert("5");
        priorityQueueS.insert("2");
        priorityQueueS.insert("2");
        for (int length = priorityQueueS.queue.length - 1; length >= 0; length--) {
            priorityQueueS.insertAt(length, priorityQueueS.delMax());
        }

        System.out.println(Arrays.toString(priorityQueueS.queue));


    }

    @Override
    public void insert(Comparable insertItem) {
        if (null == insertItem) {
            throw new IllegalArgumentException("insert item is must not null");
        }

        if (size == 0) {
            queue[0] = insertItem;
        } else {
            if (size > queue.length) {
                queue = Arrays.copyOf(queue, queue.length << 1);
            }
            shiftUp(size, insertItem);
        }
        size++;
    }

    private void insertAt(int index, Comparable it) {
        queue[index] = it;
    }

    private void shiftUp(int size, Comparable insertItem) {
        while (size > 0) {
            int father = (size - 1) >> 1;
            // 父亲比这个节点小不再交换，如果已经大了，这个就是正确位置
            // 堆是弱序性的，只要求根到叶子的每一条路径是有序的，
            if (queue[father].compareTo(insertItem) >= 0) {
                break;
            }
            queue[size] = queue[father];
            size = father;
        }
        queue[size] = insertItem;
    }

    @Override
    public Comparable max() {
        return size == 0 ? null : queue[0];
    }

    @Override
    public Comparable delMax() {
        if (size == 0) return null;
        Comparable max = queue[0];
        int lastOne = --size;
        queue[0] = queue[lastOne];
        queue[lastOne] = null;
        if (size != 0)
            shiftDown(0, queue[0]);
        return max;
    }

    private void shiftDown(int index, Comparable comparable) {
        int leaf = size >> 1;
        while (index < leaf) {
            // 2n+1
            int child = (index << 1) + 1;
            Comparable childValue = queue[child];
            // 2n+2
            int rightLeaf = child + 1;
            // 这里是一个大堆，所以我要找到子节点的最大元素进行上浮
            //
            if (rightLeaf < size && childValue.compareTo(queue[rightLeaf]) < 0) {
                childValue = queue[child = rightLeaf];
            }

            // 如果数据已经大于当前的子节点则跳出
            if (comparable.compareTo(childValue) >= 0) {
                break;
            }
            // 当前节点不比孩子大
            queue[index] = childValue;
            index = child;
        }
        queue[index] = comparable;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    @Override
    public int size() {
        return size;
    }
}
