package com.sona.architecture.dynamicconnectivity;

import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.UF;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;

/**
 * debug  edu.princeton.cs.algs4.UF
 * <p>
 * {@see https://algs4.cs.princeton.edu/15uf/UF.java.html}
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/9
 */
public class SonaUf {

    public static void main(String[] args) throws Exception {
        BufferedReader bufferedReader = new BufferedReader(new FileReader(new File(args[0])));
        String s = bufferedReader.readLine();
        UF uf = new UF(Integer.parseInt(s));
        String line;
        while ((line = bufferedReader.readLine()) != null) {
            String[] lineItems = line.split(" ");
            int p = Integer.parseInt(lineItems[0]);
            int q = Integer.parseInt(lineItems[1]);
            if (uf.find(p) == uf.find(q)) continue;
            uf.union(p, q);
            StdOut.println(p + " " + q);
        }
        StdOut.println(uf.count() + " components");
    }
}
