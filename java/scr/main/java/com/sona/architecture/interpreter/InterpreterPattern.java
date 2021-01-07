package com.sona.architecture.interpreter;

import java.util.Stack;

/**
 * 解释器设计模式
 * {@see https://algs4.cs.princeton.edu/13stacks/Evaluate.java.html}
 * {@see https://www.tutorialspoint.com/design_pattern/interpreter_pattern.htm}
 * <p>
 * <p>
 * java Evaluate
 * ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/5
 */
public class InterpreterPattern {


    public static void main(String[] args) {
        //
        Stack<String> ops = new Stack<>();
        Stack<Double> values = new Stack<>();

        String str = "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )";
        String[] s = str.split(" ");

        for (String item : s) {
            if (item.equals("(")) ;
            else if (item.equals("+")) ops.push(item);
            else if (item.equals("-")) ops.push(item);
            else if (item.equals("*")) ops.push(item);
            else if (item.equals("/")) ops.push(item);
            else if (item.equals("sqrt")) ops.push(item);
            else if (item.equals(")")) {
                String op = ops.pop();
                double v = values.pop();
                if (op.equals("+")) v = values.pop() + v;
                else if (op.equals("-")) v = values.pop() - v;
                else if (op.equals("*")) v = values.pop() * v;
                else if (op.equals("/")) v = values.pop() / v;
                else if (op.equals("sqrt")) v = Math.sqrt(v);
                values.push(v);
            } else values.push(Double.parseDouble(item));
        }
    }

}
