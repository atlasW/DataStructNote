package com.sona.architecture.sort;

import java.util.Arrays;

/**
 * 希尔排序,主要解决插入排序移动过多问题
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/10
 */
public class ShellSort implements Sort {


    @Override
    public void sort(Comparable[] data) {
        int length = data.length;
        int h = 1;

        // 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
        while (h < length / 3)
            h = (h * 3) + 1;

        while (h >= 1) {
            for (int i = 0; i < length; i++) {
                // 当前j小于j-h的数据，
                // 当i=9 h=4时候，则index = 9与index =5 交换后,可能会在向前交换
                for (int j = i; j >= h && SortUtil.less(data[j], data[j - h]); j -= h) {
                    SortUtil.exch(data, j, j - h);
                }
            }
            h = h / 3;
        }
    }

    public static void main(String[] args) {
        Integer[] needSort = {5, 1, 10, 2, 7, 9, 4, 3, 5, 6};
        new ShellSort().sort(needSort);
        System.out.println(Arrays.deepToString(needSort));
    }
}
