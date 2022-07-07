// Generated from Styx_refactoring.g4 by ANTLR 4.10.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Styx_refactoringParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.10.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, INT=11, ID=12, SEMICOLON=13, COMMA=14, WS=15;
	public static final int
		RULE_prog = 0, RULE_statement = 1, RULE_expression = 2, RULE_multiplyExpr = 3, 
		RULE_statementsBlock = 4, RULE_functionArgs = 5, RULE_functionDefinition = 6, 
		RULE_functionCall = 7, RULE_assignment = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "statement", "expression", "multiplyExpr", "statementsBlock", 
			"functionArgs", "functionDefinition", "functionCall", "assignment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'*'", "'/'", "'{'", "'}'", "'('", "')'", "'proc'", 
			"'='", null, null, "';'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, "INT", 
			"ID", "SEMICOLON", "COMMA", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Styx_refactoring.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Styx_refactoringParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(Styx_refactoringParser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(Styx_refactoringParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(Styx_refactoringParser.SEMICOLON, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterProg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitProg(this);
		}
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(24);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
				case 1:
					{
					setState(18);
					statement();
					setState(19);
					match(SEMICOLON);
					}
					break;
				case 2:
					{
					setState(21);
					expression();
					setState(22);
					match(SEMICOLON);
					}
					break;
				}
				}
				setState(26); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << ID))) != 0) );
			setState(28);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public FunctionDefinitionContext functionDefinition() {
			return getRuleContext(FunctionDefinitionContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitStatement(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(33);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(30);
				functionDefinition();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				functionCall();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(32);
				assignment();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public List<MultiplyExprContext> multiplyExpr() {
			return getRuleContexts(MultiplyExprContext.class);
		}
		public MultiplyExprContext multiplyExpr(int i) {
			return getRuleContext(MultiplyExprContext.class,i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitExpression(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			multiplyExpr();
			setState(36);
			_la = _input.LA(1);
			if ( !(_la==T__0 || _la==T__1) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(37);
			multiplyExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MultiplyExprContext extends ParserRuleContext {
		public List<TerminalNode> INT() { return getTokens(Styx_refactoringParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(Styx_refactoringParser.INT, i);
		}
		public List<TerminalNode> ID() { return getTokens(Styx_refactoringParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Styx_refactoringParser.ID, i);
		}
		public MultiplyExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multiplyExpr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterMultiplyExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitMultiplyExpr(this);
		}
	}

	public final MultiplyExprContext multiplyExpr() throws RecognitionException {
		MultiplyExprContext _localctx = new MultiplyExprContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_multiplyExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			_la = _input.LA(1);
			if ( !(_la==INT || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(40);
			_la = _input.LA(1);
			if ( !(_la==T__2 || _la==T__3) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(41);
			_la = _input.LA(1);
			if ( !(_la==INT || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementsBlockContext extends ParserRuleContext {
		public List<TerminalNode> SEMICOLON() { return getTokens(Styx_refactoringParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(Styx_refactoringParser.SEMICOLON, i);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public StatementsBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementsBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterStatementsBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitStatementsBlock(this);
		}
	}

	public final StatementsBlockContext statementsBlock() throws RecognitionException {
		StatementsBlockContext _localctx = new StatementsBlockContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_statementsBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			match(T__4);
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << ID))) != 0)) {
				{
				{
				setState(46);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(44);
					statement();
					}
					break;
				case 2:
					{
					setState(45);
					expression();
					}
					break;
				}
				setState(48);
				match(SEMICOLON);
				}
				}
				setState(54);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(55);
			match(T__5);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionArgsContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(Styx_refactoringParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Styx_refactoringParser.ID, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(Styx_refactoringParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(Styx_refactoringParser.COMMA, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FunctionArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionArgs; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterFunctionArgs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitFunctionArgs(this);
		}
	}

	public final FunctionArgsContext functionArgs() throws RecognitionException {
		FunctionArgsContext _localctx = new FunctionArgsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_functionArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(T__6);
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==INT || _la==ID) {
				{
				setState(63);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
				case 1:
					{
					setState(58);
					match(ID);
					setState(59);
					match(COMMA);
					}
					break;
				case 2:
					{
					setState(60);
					expression();
					setState(61);
					match(COMMA);
					}
					break;
				}
				}
				setState(67);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(68);
			match(T__7);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionDefinitionContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Styx_refactoringParser.ID, 0); }
		public FunctionArgsContext functionArgs() {
			return getRuleContext(FunctionArgsContext.class,0);
		}
		public StatementsBlockContext statementsBlock() {
			return getRuleContext(StatementsBlockContext.class,0);
		}
		public FunctionDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDefinition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterFunctionDefinition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitFunctionDefinition(this);
		}
	}

	public final FunctionDefinitionContext functionDefinition() throws RecognitionException {
		FunctionDefinitionContext _localctx = new FunctionDefinitionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_functionDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			match(T__8);
			setState(71);
			match(ID);
			setState(72);
			functionArgs();
			setState(73);
			statementsBlock();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Styx_refactoringParser.ID, 0); }
		public FunctionArgsContext functionArgs() {
			return getRuleContext(FunctionArgsContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterFunctionCall(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitFunctionCall(this);
		}
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_functionCall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			match(ID);
			setState(76);
			functionArgs();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(Styx_refactoringParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Styx_refactoringParser.ID, i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode INT() { return getToken(Styx_refactoringParser.INT, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).enterAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Styx_refactoringListener ) ((Styx_refactoringListener)listener).exitAssignment(this);
		}
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_assignment);
		try {
			setState(87);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				match(ID);
				setState(79);
				match(T__9);
				setState(80);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(81);
				match(ID);
				setState(82);
				match(T__9);
				setState(83);
				match(ID);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(84);
				match(ID);
				setState(85);
				match(T__9);
				setState(86);
				match(INT);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u000fZ\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0004\u0000\u0019\b\u0000\u000b\u0000\f\u0000\u001a\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\"\b"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004/\b\u0004\u0001\u0004\u0001\u0004\u0005\u00043\b\u0004\n\u0004\f"+
		"\u00046\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005@\b\u0005\n\u0005"+
		"\f\u0005C\t\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003"+
		"\bX\b\b\u0001\b\u0000\u0000\t\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0000\u0003\u0001\u0000\u0001\u0002\u0001\u0000\u000b\f\u0001\u0000\u0003"+
		"\u0004Z\u0000\u0018\u0001\u0000\u0000\u0000\u0002!\u0001\u0000\u0000\u0000"+
		"\u0004#\u0001\u0000\u0000\u0000\u0006\'\u0001\u0000\u0000\u0000\b+\u0001"+
		"\u0000\u0000\u0000\n9\u0001\u0000\u0000\u0000\fF\u0001\u0000\u0000\u0000"+
		"\u000eK\u0001\u0000\u0000\u0000\u0010W\u0001\u0000\u0000\u0000\u0012\u0013"+
		"\u0003\u0002\u0001\u0000\u0013\u0014\u0005\r\u0000\u0000\u0014\u0019\u0001"+
		"\u0000\u0000\u0000\u0015\u0016\u0003\u0004\u0002\u0000\u0016\u0017\u0005"+
		"\r\u0000\u0000\u0017\u0019\u0001\u0000\u0000\u0000\u0018\u0012\u0001\u0000"+
		"\u0000\u0000\u0018\u0015\u0001\u0000\u0000\u0000\u0019\u001a\u0001\u0000"+
		"\u0000\u0000\u001a\u0018\u0001\u0000\u0000\u0000\u001a\u001b\u0001\u0000"+
		"\u0000\u0000\u001b\u001c\u0001\u0000\u0000\u0000\u001c\u001d\u0005\u0000"+
		"\u0000\u0001\u001d\u0001\u0001\u0000\u0000\u0000\u001e\"\u0003\f\u0006"+
		"\u0000\u001f\"\u0003\u000e\u0007\u0000 \"\u0003\u0010\b\u0000!\u001e\u0001"+
		"\u0000\u0000\u0000!\u001f\u0001\u0000\u0000\u0000! \u0001\u0000\u0000"+
		"\u0000\"\u0003\u0001\u0000\u0000\u0000#$\u0003\u0006\u0003\u0000$%\u0007"+
		"\u0000\u0000\u0000%&\u0003\u0006\u0003\u0000&\u0005\u0001\u0000\u0000"+
		"\u0000\'(\u0007\u0001\u0000\u0000()\u0007\u0002\u0000\u0000)*\u0007\u0001"+
		"\u0000\u0000*\u0007\u0001\u0000\u0000\u0000+4\u0005\u0005\u0000\u0000"+
		",/\u0003\u0002\u0001\u0000-/\u0003\u0004\u0002\u0000.,\u0001\u0000\u0000"+
		"\u0000.-\u0001\u0000\u0000\u0000/0\u0001\u0000\u0000\u000001\u0005\r\u0000"+
		"\u000013\u0001\u0000\u0000\u00002.\u0001\u0000\u0000\u000036\u0001\u0000"+
		"\u0000\u000042\u0001\u0000\u0000\u000045\u0001\u0000\u0000\u000057\u0001"+
		"\u0000\u0000\u000064\u0001\u0000\u0000\u000078\u0005\u0006\u0000\u0000"+
		"8\t\u0001\u0000\u0000\u00009A\u0005\u0007\u0000\u0000:;\u0005\f\u0000"+
		"\u0000;@\u0005\u000e\u0000\u0000<=\u0003\u0004\u0002\u0000=>\u0005\u000e"+
		"\u0000\u0000>@\u0001\u0000\u0000\u0000?:\u0001\u0000\u0000\u0000?<\u0001"+
		"\u0000\u0000\u0000@C\u0001\u0000\u0000\u0000A?\u0001\u0000\u0000\u0000"+
		"AB\u0001\u0000\u0000\u0000BD\u0001\u0000\u0000\u0000CA\u0001\u0000\u0000"+
		"\u0000DE\u0005\b\u0000\u0000E\u000b\u0001\u0000\u0000\u0000FG\u0005\t"+
		"\u0000\u0000GH\u0005\f\u0000\u0000HI\u0003\n\u0005\u0000IJ\u0003\b\u0004"+
		"\u0000J\r\u0001\u0000\u0000\u0000KL\u0005\f\u0000\u0000LM\u0003\n\u0005"+
		"\u0000M\u000f\u0001\u0000\u0000\u0000NO\u0005\f\u0000\u0000OP\u0005\n"+
		"\u0000\u0000PX\u0003\u0004\u0002\u0000QR\u0005\f\u0000\u0000RS\u0005\n"+
		"\u0000\u0000SX\u0005\f\u0000\u0000TU\u0005\f\u0000\u0000UV\u0005\n\u0000"+
		"\u0000VX\u0005\u000b\u0000\u0000WN\u0001\u0000\u0000\u0000WQ\u0001\u0000"+
		"\u0000\u0000WT\u0001\u0000\u0000\u0000X\u0011\u0001\u0000\u0000\u0000"+
		"\b\u0018\u001a!.4?AW";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}