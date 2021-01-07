package com.sona.architecture.interpreter.policy;

import com.sona.architecture.interpreter.Expression;

/**
 * 终端表达式
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/5
 */
public class TerminalExpression implements Expression {

    private final String data;

    public TerminalExpression(String data){
        this.data = data;
    }

    @Override
    public boolean interpret(String context) {
        return context.contains(data);
    }
}