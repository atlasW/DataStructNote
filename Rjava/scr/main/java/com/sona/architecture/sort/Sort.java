package com.sona.architecture.sort;

/**
 * 排序的接口
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/10
 */
public interface Sort {

    /**
     * 需要排序的数据
     *
     * @param data 数据
     */
    void sort(Comparable[] data);
}
