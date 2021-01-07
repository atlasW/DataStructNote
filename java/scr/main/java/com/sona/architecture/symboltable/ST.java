package com.sona.architecture.symboltable;

import edu.princeton.cs.algs4.LinkedQueue;

/**
 * TODO
 *
 * @author renfakai
 * @version 1.0
 * @since 2021/1/6
 */
public interface ST<Key, Value> {

    int size();

    boolean isEmpty();

    boolean containsKey(Key k);

    Value get(Key k);

    void put(Key k, Value v);

    void delete(Key k);

    Iterable<Key> keys();


}
