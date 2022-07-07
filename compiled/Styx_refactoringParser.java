// Generated from Styx_refactoring.g4 by ANTLR 4.9
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
	static { RuntimeMetaData.checkVersion("4.9", RuntimeMetaData.VERSION); }

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
		public List<TerminalNode> ID() { return getTokens(Styx_refactoringParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Styx_refactoringParser.ID, i);
		}
		public List<TerminalNode> INT() { return getTokens(Styx_refactoringParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(Styx_refactoringParser.INT, i);
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
			setState(42);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
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
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
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
				if ( !(_la==T__0 || _la==T__1) ) {
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
			setState(44);
			_la = _input.LA(1);
			if ( !(_la==INT || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(45);
			_la = _input.LA(1);
			if ( !(_la==T__2 || _la==T__3) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(46);
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
			setState(48);
			match(T__4);
			setState(57);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << ID))) != 0)) {
				{
				{
				setState(51);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
				case 1:
					{
					setState(49);
					statement();
					}
					break;
				case 2:
					{
					setState(50);
					expression();
					}
					break;
				}
				setState(53);
				match(SEMICOLON);
				}
				}
				setState(59);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(60);
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
			setState(62);
			match(T__6);
			setState(70);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==INT || _la==ID) {
				{
				setState(68);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(63);
					match(ID);
					setState(64);
					match(COMMA);
					}
					break;
				case 2:
					{
					setState(65);
					expression();
					setState(66);
					match(COMMA);
					}
					break;
				}
				}
				setState(72);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(73);
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
			setState(75);
			match(T__8);
			setState(76);
			match(ID);
			setState(77);
			functionArgs();
			setState(78);
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
			setState(80);
			match(ID);
			setState(81);
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
			setState(92);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(83);
				match(ID);
				setState(84);
				match(T__9);
				setState(85);
				expression();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(ID);
				setState(87);
				match(T__9);
				setState(88);
				match(ID);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(89);
				match(ID);
				setState(90);
				match(T__9);
				setState(91);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\21a\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\6\2\33\n\2\r\2\16\2\34\3\2\3\2\3\3\3\3\3\3\5\3$\n\3\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\5\4-\n\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\5\6\66\n\6"+
		"\3\6\3\6\7\6:\n\6\f\6\16\6=\13\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\7\7G"+
		"\n\7\f\7\16\7J\13\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n_\n\n\3\n\2\2\13\2\4\6\b\n\f\16\20\22\2"+
		"\5\3\2\3\4\3\2\r\16\3\2\5\6\2b\2\32\3\2\2\2\4#\3\2\2\2\6,\3\2\2\2\b.\3"+
		"\2\2\2\n\62\3\2\2\2\f@\3\2\2\2\16M\3\2\2\2\20R\3\2\2\2\22^\3\2\2\2\24"+
		"\25\5\4\3\2\25\26\7\17\2\2\26\33\3\2\2\2\27\30\5\6\4\2\30\31\7\17\2\2"+
		"\31\33\3\2\2\2\32\24\3\2\2\2\32\27\3\2\2\2\33\34\3\2\2\2\34\32\3\2\2\2"+
		"\34\35\3\2\2\2\35\36\3\2\2\2\36\37\7\2\2\3\37\3\3\2\2\2 $\5\16\b\2!$\5"+
		"\20\t\2\"$\5\22\n\2# \3\2\2\2#!\3\2\2\2#\"\3\2\2\2$\5\3\2\2\2%&\5\b\5"+
		"\2&\'\t\2\2\2\'(\5\b\5\2(-\3\2\2\2)*\t\3\2\2*+\t\2\2\2+-\t\3\2\2,%\3\2"+
		"\2\2,)\3\2\2\2-\7\3\2\2\2./\t\3\2\2/\60\t\4\2\2\60\61\t\3\2\2\61\t\3\2"+
		"\2\2\62;\7\7\2\2\63\66\5\4\3\2\64\66\5\6\4\2\65\63\3\2\2\2\65\64\3\2\2"+
		"\2\66\67\3\2\2\2\678\7\17\2\28:\3\2\2\29\65\3\2\2\2:=\3\2\2\2;9\3\2\2"+
		"\2;<\3\2\2\2<>\3\2\2\2=;\3\2\2\2>?\7\b\2\2?\13\3\2\2\2@H\7\t\2\2AB\7\16"+
		"\2\2BG\7\20\2\2CD\5\6\4\2DE\7\20\2\2EG\3\2\2\2FA\3\2\2\2FC\3\2\2\2GJ\3"+
		"\2\2\2HF\3\2\2\2HI\3\2\2\2IK\3\2\2\2JH\3\2\2\2KL\7\n\2\2L\r\3\2\2\2MN"+
		"\7\13\2\2NO\7\16\2\2OP\5\f\7\2PQ\5\n\6\2Q\17\3\2\2\2RS\7\16\2\2ST\5\f"+
		"\7\2T\21\3\2\2\2UV\7\16\2\2VW\7\f\2\2W_\5\6\4\2XY\7\16\2\2YZ\7\f\2\2Z"+
		"_\7\16\2\2[\\\7\16\2\2\\]\7\f\2\2]_\7\r\2\2^U\3\2\2\2^X\3\2\2\2^[\3\2"+
		"\2\2_\23\3\2\2\2\13\32\34#,\65;FH^";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}