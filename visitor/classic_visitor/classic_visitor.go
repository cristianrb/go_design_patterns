package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
	VisitSubtractionExpression(e *SubtractionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type SubtractionExpression struct {
	left, right Expression
}

func (a *SubtractionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitSubtractionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) VisitSubtractionExpression(e *SubtractionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('-')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func (ev *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ev.result = e.value
}

func (ev *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ev)
	x := ev.result
	e.right.Accept(ev)
	x += ev.result
	ev.result = x
}

func (ev *ExpressionEvaluator) VisitSubtractionExpression(e *SubtractionExpression) {
	e.left.Accept(ev)
	x := ev.result
	e.right.Accept(ev)
	x -= ev.result
	ev.result = x
}

func main() {
	// 1 + (2+3) - 4 = 2
	e := &SubtractionExpression{
		left: &AdditionExpression{
			left: &DoubleExpression{1},
			right: &AdditionExpression{
				left:  &DoubleExpression{2},
				right: &DoubleExpression{3},
			},
		},
		right: &DoubleExpression{4},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g", ep, ee.result)
}
