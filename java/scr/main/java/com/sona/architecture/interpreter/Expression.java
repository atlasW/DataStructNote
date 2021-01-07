package com.sona.architecture.interpreter;

/**
 * 表达式上下文
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/5
 */
public interface Expression {

    boolean interpret(String context);
}
