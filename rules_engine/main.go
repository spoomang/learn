package main

import (
	"fmt"
	"time"

	"github.com/buptmiao/parallel"
)

var p *parallel.Parallel = parallel.NewParallel()

type E struct {
	Msg string
}

func (e E) Error() string {
	return e.Msg
}

func (e E) Wrap() {

}

func testJobA(i int) (string, error) {
	return fmt.Sprintf("job_%d", i), nil
}

func testJobB(args ...interface{}) (int, error) {

	x := args[0].(int)
	y := args[1].(int)
	z := args[2].(int)

	return x + y + z, E{"This is an error"}
}

func testJobC(x int) (int, error) {
	return -x, nil
}

type Functions struct {
	F interface{}
	I []interface{}
	O []interface{}
}

type Rule struct {
	Name     string
	Function interface{}
	InputLen int
}

type RuleEngine struct {
	Rules []*Rule
}

func NewRuleEngine() *RuleEngine {
	ruleEngine := &RuleEngine{}
	ruleEngine.Rules = make([]*Rule, 0)
	return ruleEngine
}

func (r *RuleEngine) AddRule(name string, f interface{}, inputLen int) {
	rule := &Rule{
		Name:     name,
		Function: f,
		InputLen: inputLen,
	}
	r.Rules = append(r.Rules, rule)
}

func (r *RuleEngine) Execute(inputs [][]interface{}, outputs [][]interface{}) error {
	numOfRules := len(r.Rules)
	if numOfRules != len(inputs) || numOfRules != len(outputs) {
		return E{
			Msg: "Number of inputs / outputs does not match the number of rules",
		}
	}

	p := parallel.NewParallel()

	for i := 0; i < numOfRules; i++ {
		input := inputs[i]

		if len(input) != r.Rules[i].InputLen || len(outputs[i]) != 2 {
			return E{
				Msg: fmt.Sprintf("Number of input / output does not match for rule %s", r.Rules[i].Name),
			}
		}

		p.Register(r.Rules[i].Function, input...).SetReceivers(&outputs[i][0], &outputs[i][1])
	}

	p.Run()

	return nil
}

func main() {
	fmt.Println("Hello, playground", time.Now().Add(time.Hour*24).UTC().Format(time.RFC3339))
	ruleEngine := NewRuleEngine()
	ruleEngine.AddRule("A", testJobA, 1)
	ruleEngine.AddRule("B", testJobB, 3)
	ruleEngine.AddRule("C", testJobC, 1)

	for i := 0; i < 3; i++ {
		go func(i int) {

			// functions := make([]*Functions, 0)
			// o1 := make([]interface{}, 2)
			// i1 := make([]interface{}, 0)
			// f1 := &Functions{
			// 	F: testJobA,
			// 	O: o1,
			// 	I: i1,
			// }
			// functions = append(functions, f1)

			// o2 := make([]interface{}, 2)
			// i2 := make([]interface{}, 0)
			// i2 = append(i2, 10)
			// i2 = append(i2, 20)
			// i2 = append(i2, 30)
			// f2 := &Functions{
			// 	F: testJobB,
			// 	I: i2,
			// 	O: o2,
			// }
			// functions = append(functions, f2)

			// myRuleEngine(functions)
			// fmt.Println(functions[0].O)
			// fmt.Println(functions[1].O)

			i1 := returnInterfaceArray(i)
			i2 := returnInterfaceArray(10, 20, i)
			i3 := returnInterfaceArray(i)

			o1 := make([]interface{}, 2)
			o2 := make([]interface{}, 2)
			o3 := make([]interface{}, 2)

			inputs := returnIterfacesArray(i1, i2, i3)
			outputs := returnIterfacesArray(o1, o2, o3)

			err := ruleEngine.Execute(inputs, outputs)

			fmt.Println(err)
			fmt.Println(outputs)
		}(i)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("done")
}

func returnInterfaceArray(args ...interface{}) []interface{} {
	return args
}

func returnIterfacesArray(args ...[]interface{}) [][]interface{} {
	return args
}

func myRuleEngine(args []*Functions) {
	p := parallel.NewParallel()
	for i, _ := range args {
		p.Register(args[i].F, args[i].I...).SetReceivers(&args[i].O[0], &args[i].O[1])
	}

	// block here
	p.Run()
}
