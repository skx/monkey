package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skx/monkey/ast"
	"github.com/skx/monkey/lexer"
	"github.com/skx/monkey/token"
)

// prefix Parse function
// infix parse function
// postfix parse function
type (
	prefixParseFn  func() ast.Expression
	infixParseFn   func(ast.Expression) ast.Expression
	postfixParseFn func() ast.Expression
)

// precedence order
const (
	_ int = iota
	LOWEST
	ASSIGN      // =
	EQUALS      // == or !=
	LESSGREATER // > or <
	SUM         // + or -
	PRODUCT     // * or /
	POWER       // **
	MOD         // %
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index], map[key]
)

// each token precedence
var precedences = map[token.TokenType]int{
	token.ASSIGN:          ASSIGN,
	token.EQ:              EQUALS,
	token.NOT_EQ:          EQUALS,
	token.LT:              LESSGREATER,
	token.LT_EQUALS:       LESSGREATER,
	token.GT:              LESSGREATER,
	token.GT_EQUALS:       LESSGREATER,
	token.PLUS:            SUM,
	token.PLUS_EQUALS:     SUM,
	token.MINUS:           SUM,
	token.MINUS_EQUALS:    SUM,
	token.SLASH:           PRODUCT,
	token.SLASH_EQUALS:    PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.ASTERISK_EQUALS: PRODUCT,
	token.POW:             POWER,
	token.MOD:             MOD,
	token.LPAREN:          CALL,
	token.PERIOD:          CALL,
	token.LBRACKET:        INDEX,
}

// Parser object
type Parser struct {
	// l is our lexer
	l *lexer.Lexer

	// prevToken holds the previous token from our lexer.
	// (used for "++" + "--")
	prevToken token.Token

	// curToken holds the current token from our lexer.
	curToken token.Token

	// peekToken holds the next token which will come from the lexer.
	peekToken token.Token

	// errors holds parsing-errors.
	errors []string

	// prefixParseFns holds a map of parsing methods for
	// prefix-based syntax.
	prefixParseFns map[token.TokenType]prefixParseFn

	// infixParseFns holds a map of parsing methods for
	// infix-based syntax.
	infixParseFns map[token.TokenType]infixParseFn

	// postfixParseFns holds a map of parsing methods for
	// postfix-based syntax.
	postfixParseFns map[token.TokenType]postfixParseFn
}

// New returns our new parser-object.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FOR, p.parseForLoopExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.DEFINE_FUNCTION, p.parseFunctionDefinition)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.BACKTICK, p.parseBacktickLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.ASSIGN, p.parseAssignExpression)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.POW, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LT_EQUALS, p.parseInfixExpression)
	p.registerInfix(token.GT_EQUALS, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	p.registerInfix(token.PERIOD, p.parseMethodCallExpression)
	p.registerInfix(token.PLUS_EQUALS, p.parseAssignExpression)
	p.registerInfix(token.MINUS_EQUALS, p.parseAssignExpression)
	p.registerInfix(token.ASTERISK_EQUALS, p.parseAssignExpression)
	p.registerInfix(token.SLASH_EQUALS, p.parseAssignExpression)

	p.postfixParseFns = make(map[token.TokenType]postfixParseFn)
	p.registerPostfix(token.PLUS_PLUS, p.parsePostfixExpression)
	p.registerPostfix(token.MINUS_MINUS, p.parsePostfixExpression)
	return p
}

// registerPrefix registers a function for handling a prefix-based statement
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix registers a function for handling a infix-based statement
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// registerPostfix registers a function for handling a postfix-based statement
func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

// Errors return stored errors
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError raises an error if the next token is not the expected type.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.curToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken moves to our next token from the lexer.
func (p *Parser) nextToken() {
	p.prevToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram used to parse the whole program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// parse statement
// it splits into three parts: 1: LET Statement; 2: RETURN Statement; 3: ExpressionStatement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.CONST:
		return p.parseConstStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement parses a let-statement.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseConstStatement parses a constant declaration.
func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseReturnStatement parses a return-statement.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// no prefix parse function error
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// parse Expression Statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	for p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	postfix := p.postfixParseFns[p.curToken.Type]
	if postfix != nil {
		return (postfix())
	}
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

// parse identifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parse integer literal
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	var value int64
	var err error

	if strings.HasPrefix(p.curToken.Literal, "0b") {
		value, err = strconv.ParseInt(p.curToken.Literal[2:], 2, 64)
	} else if strings.HasPrefix(p.curToken.Literal, "0x") {
		value, err = strconv.ParseInt(p.curToken.Literal[2:], 16, 64)
	} else {
		value, err = strconv.ParseInt(p.curToken.Literal, 10, 64)
	}

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

// parse float literal
func (p *Parser) parseFloatLiteral() ast.Expression {
	flo := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	flo.Value = value
	return flo
}

// parse boolean
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

// parse prefix expression
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

// parse postfix expression
func (p *Parser) parsePostfixExpression() ast.Expression {
	expression := &ast.PostfixExpression{
		Token:    p.prevToken,
		Operator: p.curToken.Literal,
	}
	return expression
}

// parse infix expression
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

// parse grouped expression
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

// parse if condition expression
func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}
	return expression
}

// parse for loop expression
func (p *Parser) parseForLoopExpression() ast.Expression {
	expression := &ast.ForLoopExpression{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()
	return expression
}

// parse block statement
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.curTokenIs(token.RBRACE) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

// parse function literal expression
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Defaults, lit.Parameters = p.parseFunctionParameters()
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = p.parseBlockStatement()
	return lit
}

// parse function definition
func (p *Parser) parseFunctionDefinition() ast.Expression {
	p.nextToken()
	lit := &ast.FunctionDefineLiteral{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Defaults, lit.Parameters = p.parseFunctionParameters()
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = p.parseBlockStatement()
	return lit
}

// parse function parameters
func (p *Parser) parseFunctionParameters() (map[string]ast.Expression, []*ast.Identifier) {

	// Any default parameters.
	m := make(map[string]ast.Expression)

	// The argument-definitions.
	identifiers := make([]*ast.Identifier, 0)

	// Is the next parameter ")" ?  If so we're done. No args.
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return m, identifiers
	}
	p.nextToken()

	// Keep going until we find a ")"
	for !p.curTokenIs(token.RPAREN) {

		// Get the identifier.
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
		p.nextToken()

		// If there is "=xx" after the name then that's
		// the default parameter.
		if p.curTokenIs(token.ASSIGN) {
			p.nextToken()
			// Save the default value.
			m[ident.Value] = p.parseExpressionStatement().Expression
			p.nextToken()
		}

		// Skip any comma.
		if p.curTokenIs(token.COMMA) {
			p.nextToken()
		}
	}

	return m, identifiers
}

// parse string literal
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// parse backtick-quoted string
func (p *Parser) parseBacktickLiteral() ast.Expression {
	return &ast.BacktickLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// parse array literal
func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

// parse array elements literal
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := make([]ast.Expression, 0)
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}
	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(end) {
		return nil
	}
	return list
}

// parse array index expression
func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}
	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	return exp
}

// parse (bare) assignment, without a `let`.
func (p *Parser) parseAssignExpression(name ast.Expression) ast.Expression {
	stmt := &ast.AssignStatement{Token: p.curToken}
	if n, ok := name.(*ast.Identifier); ok {
		stmt.Name = n
	} else {
		msg := fmt.Sprintf("expected assign token to be IDENT, got %s instead", name.TokenLiteral())
		p.errors = append(p.errors, msg)
	}

	oper := p.curToken
	p.nextToken()

	//
	// An assignment is generally:
	//
	//    variable = value
	//
	// But we cheat and reuse the implementation for:
	//
	//    i += 4
	//
	// In this case we record the "operator" as "+="
	//
	switch oper.Type {
	case token.PLUS_EQUALS:
		stmt.Operator = "+="
	case token.MINUS_EQUALS:
		stmt.Operator = "-="
	case token.SLASH_EQUALS:
		stmt.Operator = "/="
	case token.ASTERISK_EQUALS:
		stmt.Operator = "*="
	default:
		stmt.Operator = "="
	}
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

// parse function call
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

// parse hash literal
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)
		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()
		value := p.parseExpression(LOWEST)
		hash.Pairs[key] = value
		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return hash
}

// parse method-call, on an object.
func (p *Parser) parseMethodCallExpression(obj ast.Expression) ast.Expression {
	methodCall := &ast.ObjectCallExpression{Token: p.curToken, Object: obj}
	p.nextToken()
	name := p.parseIdentifier()
	p.nextToken()
	methodCall.Call = p.parseCallExpression(name)
	return methodCall
}

//
func (p *Parser) parseCallArguments() []ast.Expression {
	args := make([]ast.Expression, 0)
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}
	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return args
}

// determinate current token is t or not.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// determinate next token is t or not
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expect next token is t
// succeed: return true and forward token
// failed: return false and store error
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// look up next token precedence
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// look up current token precedence
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
