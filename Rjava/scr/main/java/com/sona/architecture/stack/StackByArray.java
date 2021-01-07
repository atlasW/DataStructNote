package com.sona.architecture.stack;

import java.util.Arrays;

/**
 * 王铮算法的demo
 *
 * @author renfakai
 * @since :2020/11/16 14:52
 */
public class StackByArray {

    // array
    public String[] items;

    // 栈中元素的个数
    private int count;

    // 栈的size
    private int n;


    public StackByArray(int size) {
        this.n = size;
        this.count = 0;
        items = new String[size];
    }

    public boolean push(String item) {
        if (count == n) return false;
        items[count] = item;
        count++;
        return true;
    }

    public boolean pushI(String item) {
        if (count == n) return false;
        items[count++] = item;
        return true;
    }


    public String pop() {
        if (count == 0) return null;
        String item = items[count - 1];
        count--;
        return item;
    }

    @Override
    public String toString() {
        return "StackByArray{" +
                "items=" + Arrays.toString(items) +
                ", count=" + count +
                ", n=" + n +
                '}';
    }

    public static void main(String[] args) {
        final StackByArray stackByArray = new StackByArray(4);
        stackByArray.push("1");
        stackByArray.push("2");
        stackByArray.push("3");
        stackByArray.push("4");
        stackByArray.push("5");
        System.out.println(stackByArray);
        System.out.println(stackByArray.pop());
        System.out.println(stackByArray);


        final StackByArray stackByArray2 = new StackByArray(4);
        stackByArray2.pushI("1");
        stackByArray2.pushI("2");
        stackByArray2.pushI("3");
        stackByArray2.pushI("4");
        System.out.println(stackByArray2);
        System.out.println(stackByArray2.pop());
        System.out.println(stackByArray2);
    }
}
