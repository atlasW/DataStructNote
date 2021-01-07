package com.sona.architecture;

/**
 * @author renfakai
 * @since :2020/11/14 14:56
 */
public class SentinelC {


    // 正常程序员开发
    static int findByIndex(char[] data, int n, char key) {

        // data check
        if (null == data || data.length == 0 || n <= 0) {
            return -1;
        }

        // check array index and equals
        int index = 0;
        while (index < n) {
            if (data[index] == key) {
                return index;
            }
            index++;
        }

        // return result
        return -1;
    }

    static int findBySentinel(char[] data, int n, char key) {
        if (null == data || data.length == 0 || n <= 0) {
            return -1;
        }
        //  if key equals lastOne return
        if (data[n - 1] == key) {
            return n - 1;
        }

        // cache it
        char last = data[n - 1];
        data[n - 1] = key;

        // reduce index com
        int i = 0;
        while (data[i] != key) {
            i++;
        }
        data[n - 1] = last;


        if (i == (n - 1)) {
            return -1;
        } else {
            return i;
        }

    }
}
