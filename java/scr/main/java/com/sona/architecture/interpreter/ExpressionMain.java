package com.sona.architecture.interpreter;

import com.sona.architecture.interpreter.policy.AndExpression;
import com.sona.architecture.interpreter.policy.TerminalExpression;
import com.sona.architecture.interpreter.policy.OrExpression;

/**
 * 一般情况下别人写代码喜欢用责任链代替解释器
 * 责任链全部都是and 活着 or 好处理如果
 * <p>
 * (e1 or e2) and e3 and ( e4 or e5)
 * 还是解释器牛逼啊
 *
 * @author renfakai
 * @version 1.0
 * @since 2020/12/5
 * (a or b ) and  (c and d)
 */
public class ExpressionMain {


    //Rule: Robert and John are male
    public static Expression getMaleExpression() {
        Expression robert = new TerminalExpression("Robert");
        Expression john = new TerminalExpression("John");
        return new OrExpression(robert, john);
    }

    //Rule: Julie is a married women
    public static Expression getMarriedWomanExpression() {
        Expression julie = new TerminalExpression("Julie");
        Expression married = new TerminalExpression("Married");
        return new AndExpression(julie, married);
    }

//    //Rule: Julie is a married women
//    public static Expression orAnd() {
//        Expression maleExpression = getMaleExpression();
//        Expression marriedWomanExpression = getMarriedWomanExpression();
//        return new AndExpression(maleExpression, marriedWomanExpression);
//    }


    public static void main(String[] args) {
        Expression isMale = getMaleExpression();
        Expression isMarriedWoman = getMarriedWomanExpression();

        Expression ops = new AndExpression(isMale, isMarriedWoman);

        System.out.println("John is male? " + isMale.interpret("John"));
        System.out.println("Julie is a married women? " + isMarriedWoman.interpret("Married Julie"));
        System.out.println("Julie is a married women? " + ops.interpret("Married Julie"));
        System.out.println("John is male?  Julie is a married women? " + ops.interpret("John is male?  Married Julie"));
    }


}
