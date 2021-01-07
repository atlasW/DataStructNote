package com.sona.architecture.sort;

import java.util.Arrays;

/**
 * sortUtil
 * {@see https://algs4.cs.princeton.edu/21elementary/}
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/10
 */
public abstract class SortUtil {


    public static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }


    public static void exch(Comparable[] data, int indexOne, int indexTwo) {
        Comparable temp = data[indexOne];
        data[indexOne] = data[indexTwo];
        data[indexTwo] = temp;
    }


    public static void main(String[] args) {
        System.out.println(less(3, 4));
        System.out.println(less("a", "b"));
        Integer[] test = {1, 2, 3, 5};
        exch(test, 0, 1);
        System.out.println(Arrays.deepToString(test));
    }


}
