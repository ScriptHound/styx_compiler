// Generated from Styx_refactoring.g4 by ANTLR 4.9
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link Styx_refactoringParser}.
 */
public interface Styx_refactoringListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(Styx_refactoringParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(Styx_refactoringParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(Styx_refactoringParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(Styx_refactoringParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(Styx_refactoringParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(Styx_refactoringParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#multiplyExpr}.
	 * @param ctx the parse tree
	 */
	void enterMultiplyExpr(Styx_refactoringParser.MultiplyExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#multiplyExpr}.
	 * @param ctx the parse tree
	 */
	void exitMultiplyExpr(Styx_refactoringParser.MultiplyExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#statementsBlock}.
	 * @param ctx the parse tree
	 */
	void enterStatementsBlock(Styx_refactoringParser.StatementsBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#statementsBlock}.
	 * @param ctx the parse tree
	 */
	void exitStatementsBlock(Styx_refactoringParser.StatementsBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#functionArgs}.
	 * @param ctx the parse tree
	 */
	void enterFunctionArgs(Styx_refactoringParser.FunctionArgsContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#functionArgs}.
	 * @param ctx the parse tree
	 */
	void exitFunctionArgs(Styx_refactoringParser.FunctionArgsContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#functionDefinition}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDefinition(Styx_refactoringParser.FunctionDefinitionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#functionDefinition}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDefinition(Styx_refactoringParser.FunctionDefinitionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(Styx_refactoringParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(Styx_refactoringParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link Styx_refactoringParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(Styx_refactoringParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link Styx_refactoringParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(Styx_refactoringParser.AssignmentContext ctx);
}