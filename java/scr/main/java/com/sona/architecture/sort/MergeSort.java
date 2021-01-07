package com.sona.architecture.sort;

import java.util.Arrays;

/**
 * {@see https://algs4.cs.princeton.edu/22mergesort/Merge.java.html}
 * up-down 排序
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/16
 */
public class MergeSort {


    static int count = 0;


    public static void merge(Comparable[] org, Comparable[] aux, int lo, int mid, int hi) {
        // copy it into aux
        for (int i = lo; i <= hi; i++) {
            aux[i] = org[i];
        }
        count++;
        System.err.println("..........----M- " + count + "------lo:" + lo + "---- hi:" + hi);
        //  [lo,mid] [mid,hi]
        int i = lo, j = mid + 1;
        for (int k = lo; k <= hi; k++) {
            // [lo,mid] 已经处理完毕了
            if (i > mid)
                //  i大于中间，说明左边已经取完了
                org[k] = aux[j++];
            else if (j > hi)
                // j > hi 说明右边取完了
                org[k] = aux[i++];
                // 分别推荐左边和右边
            else if (aux[j].compareTo(aux[i]) < 0)
                org[k] = aux[j++];
            else
                org[k] = aux[i++];
        }

    }




    public static void sort(Comparable[] a, Comparable[] aux, int lo, int hi) {
//        System.err.println("------------lo:" + lo + "---- hi:" + hi);
        // 无法进行切割
        if (hi <= lo) {
//            System.err.println("------------lo:" + lo + "---- hi:" + hi + " return");
            return;
        }
        count++;
        System.err.println("---------" + count + "---lo:" + lo + "---- hi:" + hi);
        int mid = lo + ((hi - lo) >> 1);
        // 左边排序
        sort(a, aux, lo, mid);
        // 右边排序
        sort(a, aux, mid + 1, hi);
        // 合并
        merge(a, aux, lo, mid, hi);

    }


    public static void main(String[] args) {
        Comparable[] a = {1, 5, 9, 2, 6, 8, 3, 0};
        Comparable[] aux = new Comparable[a.length];
        sort(a, aux, 0, a.length - 1);
        System.out.println(Arrays.toString(Arrays.stream(a).toArray()));
    }
}
