package com.sona.architecture.sort;

import java.util.Arrays;

/**
 * 插入排序
 * 在计算机的实现中，为了给要插入的元素腾出空间，我们需要将其余所有元素在插入之前都向右移动一位,这种算法叫做插入排序。
 * 数组中每个元素距离它的最终位置都不远；这样3n就不会特别大
 * 一个有序的大数组接一个小数组；
 * 数组中只有几个元素的位置不正确。
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/10
 */
public class InsertionSort implements Sort {

    /**
     * {@see https://algs4.cs.princeton.edu/21elementary/Insertion.java.html}
     * 代码很清晰，但是数据移动需要3n
     *
     * @param data 数据
     */
    @Override
    public void sort(Comparable[] data) {
        int length = data.length;
        for (int i = 0; i < length; i++) {
            for (int j = i; j > 0 && SortUtil.less(data[j], data[j - 1]); j--) {
                SortUtil.exch(data, j, j - 1);
            }
        }
    }

    // 本来想用二分查找进行优化的，后来发现找不到 key==data[i]
    public void sortWithRank(Comparable[] data) {
        int length = data.length;
        for (int i = 0; i < length; i++) {
            // 保存数据应该要保存到的位置
            int findIndex = i;
            // 将i前面的数据全部都移动
            int needHandler = i;
            Comparable datum = data[i];
            for (int j = i; j > 0; j--) {
                //  如果需要数组查找的话，必须得判断如果数据比第一个数据还小的话
                if (j - 1 == 0 && SortUtil.less(datum, data[j - 1])) {
                    findIndex = j - 1;
                }
                // 数据不比第一个小，找到[x,y]中索引,把后面的数据全部都后移一位
                if (j - 1 > 0 && SortUtil.less(data[j - 1], datum)
                        && !SortUtil.less(data[j], datum)) {
                    findIndex = j;
                }
            }

            // 找到了数据的位置，将数据全部右移动 数据移动n+1次
            while (needHandler > findIndex) {
                data[needHandler] = data[--needHandler];
            }
            // 赋值
            data[needHandler] = datum;
        }
    }


    public static void main(String[] args) {
        Integer[] needSort = {5, 1, 10, 2, 7, 9, 4};
//        new InsertionSort().sort(needSort);
//        System.out.println(Arrays.deepToString(needSort));

        new InsertionSort().sortWithRank(needSort);
        System.out.println(Arrays.deepToString(needSort));

    }

}
