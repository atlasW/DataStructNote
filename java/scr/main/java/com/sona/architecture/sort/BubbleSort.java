package com.sona.architecture.sort;

import java.util.Arrays;

/**
 * 冒泡排序
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/10
 */
public class BubbleSort implements Sort {

    @Override
    public void sort(Comparable[] data) {
        int length = data.length;
        for (int i = 0; i < length - 1; i++) {
            for (int j = 0; j < length - i - 1; j++) {
                // 向后冒泡
                if (SortUtil.less(data[j + 1], data[j])) {
                    SortUtil.exch(data, j + 1, j);
                }
            }
        }
    }


    public static void main(String[] args) {
        Integer[] needSort = {5, 1, 10, 2, 7, 9, 4};
        new BubbleSort().sort(needSort);
        System.out.println(Arrays.deepToString(needSort));
    }
}
