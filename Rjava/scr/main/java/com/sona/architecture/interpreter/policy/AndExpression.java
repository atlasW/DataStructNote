package com.sona.architecture.interpreter.policy;

import com.sona.architecture.interpreter.Expression;

/**
 * TODO
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/5
 */
public class AndExpression implements Expression {

    private Expression o1 = null;
    private Expression o2 = null;

    public AndExpression(Expression o1, Expression o2) {
        this.o1 = o1;
        this.o2 = o2;
    }

    @Override
    public boolean interpret(String context) {
        return o1.interpret(context) && o2.interpret(context);
    }
}