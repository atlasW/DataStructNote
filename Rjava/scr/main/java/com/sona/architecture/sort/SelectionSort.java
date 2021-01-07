package com.sona.architecture.sort;

import java.util.Arrays;

/**
 * 选择排序
 * 自然排序和比较器
 * <p>
 * {@see https://algs4.cs.princeton.edu/21elementary/Selection.java.html}
 * 首先，找到数组中最小的那个元素，其次，将它和数组的第一个元素交换位置（如果第一个元素就是最小元素那么它就和自己交换）。
 * 再次，在剩下的元素中找到最小的元素，将它与数组的第二个元素交换位置。如此往复，直到将整个数组排序。这种方法叫做选择排序，
 * 因为它在不断地选择剩余元素之中的最小者。
 * 数据为n,对比(n-1)+(n-2)+...+1 = n^2/2
 * exch = n次
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/10
 */
public class SelectionSort implements Sort {

    @Override
    public void sort(Comparable[] data) {
        int length = data.length;
        // 使用第i角标与后面的数据进行对比，找到比较小的放到第i位置
        // 缺点，数据i已经为最小元素，仍然要进行对比(length - i )
        for (int i = 0; i < length; i++) {
            // 认为i为最小值，然后一一对比
            int minItem = i;
            // 哈哈哈哈，这里写了个最垃圾的代码，data[i].com(data[i])
            // -,-
            // 修正   for (int j = i+1; j < n; j++) {  可以减少一次判断
            for (int j = i; j < length; j++) {
                if (SortUtil.less(data[j], data[minItem])) {
                    minItem = j;
                }
            }
            SortUtil.exch(data, i, minItem);
        }
    }

    public static void main(String[] args) {
        Integer[] needSort = {5, 1, 10, 2, 7, 9, 4};
        new SelectionSort().sort(needSort);
        System.out.println(Arrays.deepToString(needSort));
    }
}
