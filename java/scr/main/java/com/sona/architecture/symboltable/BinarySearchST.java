package com.sona.architecture.symboltable;

import edu.princeton.cs.algs4.LinkedQueue;

/**
 * {@see https://algs4.cs.princeton.edu/31elementary/BinarySearchST.java.html}
 * 由于这个是有序的所以每次都对有序做检查吗？这样性能问题
 *
 * @author renfakai
 * @version 1.0
 * @since 2021/1/6
 */
public class BinarySearchST<Key, Value> implements ST<Key, Value> {

    private static final int INIT_CAPACITY = 2;

    private int n;

    private Key[] keys;

    private Value[] values;

    public BinarySearchST() {
        this(INIT_CAPACITY);
    }

    public BinarySearchST(int capacity) {
        if (capacity <= 0) throw new IllegalArgumentException("capacity must po");
        keys = (Key[]) new Comparable[capacity];
        values = (Value[]) new Object[capacity];
    }

    @Override
    public int size() {
        return n;
    }

    @Override
    public boolean isEmpty() {
        return size() == 0;
    }

    @Override
    public boolean containsKey(Key k) {
        return null != get(k);
    }

    @Override
    public Value get(Key k) {
        if (null == k) throw new IllegalArgumentException("key must not null");
        if (isEmpty()) {
            return null;
        }
        int index = rank(k);
        if (index < n && keys[index].equals(k)) {
            return values[index];
        } else {
            return null;
        }
    }

    private int rank(Key k) {
        if (k == null) throw new IllegalArgumentException("argument to rank() is null");
        int lo = 0;
        int high = n;
        while (lo < high) {
            int mid = lo + (n - lo) / 2;
            int result = ((Comparable) keys[mid]).compareTo(k);
            if (result > 0) {
                high = mid;
            } else if (result < 0) {
                lo = mid;
            } else {
                return mid;
            }
        }
        return lo;
    }

    @Override
    public void put(Key k, Value v) {
        if (null == k) throw new IllegalArgumentException("key must not null");
        if (v == null) {
            delete(k);
            return;
        }
        int rank = rank(k);
        if (rank < n && ((Comparable) k).compareTo(keys[rank]) == 0) {
            values[rank] = v;
            return;
        }
        if (n == keys.length) resize(2 * keys.length);
        for (int j = n; j > rank; j--) {
            keys[j] = keys[j - 1];
            values[j] = values[j - 1];
        }
        keys[rank] = k;
        values[rank] = v;
        n++;
    }

    private void resize(int i) {
        if (i < n) throw new IllegalArgumentException(".....");
        Key[] newKeys = (Key[]) new Comparable[i];
        Value[] newValues = (Value[]) new Object[i];
        for (int i1 = 0; i1 < keys.length; i1++) {
            newKeys[i1] = keys[i1];
            newValues[i1] = values[i1];
        }
        keys = newKeys;
        values = newValues;
    }

    @Override
    public void delete(Key k) {

        if (null == k) throw new IllegalArgumentException("key must not null");

        if (isEmpty()) return;

        int i = rank(k);
        if (i == n && keys[i] != k) {
            return;
        }
        for (int j = i; j < n; j++) {
            keys[j] = keys[j + 1];
            values[j] = values[j + 1];
        }

        // clear it
        n--;
        keys[n] = null;
        values[n] = null;
        if (n > 0 && n == keys.length / 4) resize((keys.length / 2));
    }


    public void deleteMin() {
        if (isEmpty()) return;
        delete(min());
    }

    public void deleteMax() {
        if (isEmpty()) return;
        delete(max());
    }

    public Key min() {
        return keys[0];
    }

    public Key max() {
        return keys[n - 1];
    }

    public Key select(int k) {
        if (k < 0 || k > size()) throw new IllegalArgumentException("k is error");
        return keys[k];
    }

    public Key floor(Key k) {
        if (k == null) throw new IllegalArgumentException("argument to floor() is null");
        int i = rank(k);
        if (i < n && keys[i] == k)
            return keys[i];
        return i == 0 ? null : keys[i - 1];
    }


    public int size(Key lo, Key hi) {
        if (lo == null) throw new IllegalArgumentException("lo must not null");
        if (hi == null) throw new IllegalArgumentException("hi must not null");
        int rankLow = rank(lo);
        int rankHigh = rank(hi);
        if (rankHigh <= rankLow) return 0;
        if (containsKey(hi)) return rankHigh - rankLow + 1;
        else return rankHigh - rankLow;
    }


    public Key ceiling(Key k) {
        if (k == null) throw new IllegalArgumentException("argument to ceiling() is null");
        int i = rank(k);
        if (i == n) {
            return null;
        } else {
            return keys[i];
        }
    }


    // 获取所有的key
    public Iterable<Key> keys() {
        return keys(min(), max());
    }

    // 获取特定区间的key
    private Iterable<Key> keys(Key min, Key max) {
        if (min == null) throw new IllegalArgumentException("first argument to keys() is null");
        if (max == null) throw new IllegalArgumentException("second argument to keys() is nul");
        int rankMin = rank(max);
        int rankMax = rank(max);
        LinkedQueue<Key> linkedQueue = new LinkedQueue<>();
        if (rankMax >= rankMin) {
            for (; rankMin < rankMax; rankMin++) {
                linkedQueue.enqueue(keys[rankMin]);
            }
        }
        return linkedQueue;
    }


    private boolean check() {
        return isSort() && rankCheck();
    }

    private boolean rankCheck() {
        for (int i = 0; i < keys.length; i++) {
            if (i != rank(select(i))) {
                return false;
            }

            if (((Comparable) keys[i]).compareTo(select(rank(keys[i]))) != 0) {
                return false;
            }
        }
        return true;
    }

    private boolean isSort() {
        for (int i = 0; i < keys.length; i++) {
            // 如果key
            if (((Comparable) keys[i]).compareTo(keys[i + 1]) > 0) {
                return false;
            }
        }
        return true;
    }
}
