package com.sona.architecture.algorithms;

import java.lang.Comparable;

public interface QueueS<E extends Comparable<E>> {

    void insert(E e);


    E max();


    E delMax();


    boolean isEmpty();


    int size();

}
