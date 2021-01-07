package com.sona.architecture.q20;

import java.util.Stack;

/**
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * <p>
 * 有效字符串需满足：
 * <p>
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 注意空字符串可被认为是有效字符串。
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/valid-parentheses
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 * @author renfakai
 * @since :2020/11/16 15:40
 */
public class Stack20A1 {


    /**
     * 因为只有一次循环时间复杂度为O(n)，空间复杂度O(n)
     *
     * @param s need validate data
     * @return 是否有效
     */
    public static boolean isOk(String s) {
        final char[] chars = s.toCharArray();
        Stack<Character> stack = new Stack<>();
        for (char aChar : chars) {
            if (aChar == '{') {
                stack.push('}');
            } else if (aChar == '[') {
                stack.push(']');
            } else if (aChar == '(') {
                stack.push(')');
            } else if (!stack.isEmpty() && stack.pop() != aChar) {
                return false;
            }
        }
        return stack.isEmpty();
    }

    public static void main(String[] args) {
        System.out.println(isOk("()"));
        System.out.println(isOk("()[]{}"));
        System.out.println(isOk("(]"));

        System.out.println(isOk("([)]"));
        System.out.println(isOk("{[]}"));
    }
}
