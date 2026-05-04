package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v2"

	"github.com/google/skylark"
	"github.com/google/skylark/resolve"
)

type entry struct {
	globals skylark.StringDict
	err     error
}

var (
	moduleCacheMu sync.Mutex
	moduleCache             = make(map[string]*entry)
	outputWriter  io.Writer = os.Stdout
)

// load is a simple sequential implementation of module loading.
func load(thread *skylark.Thread, module string) (skylark.StringDict, error) {
	moduleCacheMu.Lock()
	e, ok := moduleCache[module]
	if e == nil {
		if ok {
			moduleCacheMu.Unlock()
			// request for package whose loading is in progress
			return nil, fmt.Errorf("cycle in load graph")
		}

		// Add a placeholder to indicate "load in progress".
		moduleCache[module] = nil
		moduleCacheMu.Unlock()

		// Load it.
		loadThread := &skylark.Thread{Load: load}
		loadedGlobals, err := skylark.ExecFile(loadThread, module, nil, globals)
		e = &entry{globals: loadedGlobals, err: err}

		// Update the cache.
		moduleCacheMu.Lock()
		moduleCache[module] = e
		moduleCacheMu.Unlock()
		return e.globals, e.err
	}
	moduleCacheMu.Unlock()
	return e.globals, e.err
}

type k8sObject struct {
	ApiVersion skylark.String         `json:"apiVersion"`
	Kind       skylark.String         `json:"kind"`
	Metadata   map[string]interface{} `json:"metadata"`
	Spec       map[string]interface{} `json:"spec"`
}

func convertDict(d *skylark.Dict) map[string]interface{} {
	r := map[string]interface{}{}

	for _, k := range d.Keys() {
		key, _ := skylark.AsString(k)
		v, _, _ := d.Get(k)
		r[key] = convert(v)
	}
	return r
}

func convertArray(l *skylark.List) []interface{} {
	elems := []interface{}{}
	for i := 0; i < l.Len(); i++ {
		v := l.Index(i)
		elems = append(elems, convert(v))
	}
	return elems
}

func convert(v skylark.Value) interface{} {
	switch v.(type) {
	case *skylark.Dict:
		d, _ := v.(*skylark.Dict)
		return convertDict(d)
	case *skylark.List:
		l, _ := v.(*skylark.List)
		return convertArray(l)
	case skylark.Int:
		i, _ := v.(skylark.Int)
		if i64, ok := i.Int64(); ok {
			return i64
		}
		return i.String()
	case skylark.Float:
		f, _ := skylark.AsFloat(v)
		return f
	case skylark.String:
		s, _ := skylark.AsString(v)
		return s
	case skylark.Bool:
		b, _ := v.(skylark.Bool)
		return bool(b)
	case skylark.NoneType:
		return nil
	default:
		return v.String()
	}
}

func outputType(t *skylark.Thread, b *skylark.Builtin, args skylark.Tuple, kwargs []skylark.Tuple) (skylark.Value, error) {
	if len(args) < 2 {
		return skylark.None, fmt.Errorf("output_type requires apiVersion and kind")
	}

	var s, m *skylark.Dict
	for _, kw := range kwargs {
		st, _ := skylark.AsString(kw[0])
		if st == "spec" {
			s, _ = kw[1].(*skylark.Dict)
		} else if st == "metadata" {
			m, _ = kw[1].(*skylark.Dict)
		}
	}

	apiVersion, _ := args[0].(skylark.String)
	kind, _ := args[1].(skylark.String)
	if m == nil {
		m = &skylark.Dict{}
	}
	if s == nil {
		s = &skylark.Dict{}
	}

	obj := k8sObject{
		ApiVersion: apiVersion,
		Kind:       kind,
		Metadata:   convertDict(m),
		Spec:       convertDict(s),
	}

	by, err := yaml.Marshal(obj)
	if err != nil {
		return skylark.None, err
	}

	if _, err := fmt.Fprintln(outputWriter, string(by)); err != nil {
		return skylark.None, err
	}

	return skylark.None, nil
}

var globals = skylark.StringDict{
	"output_type": skylark.NewBuiltin("output_type", outputType),
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-o output.yaml] <file.sky>\n", os.Args[0])
		flag.PrintDefaults()
	}

	outputFile := flag.String("o", "", "write generated YAML to file")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to create output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		outputWriter = f
	}

	resolve.AllowFloat = true
	resolve.AllowLambda = true
	thread := &skylark.Thread{Load: load}

	if _, err := skylark.ExecFile(thread, flag.Arg(0), nil, globals); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
