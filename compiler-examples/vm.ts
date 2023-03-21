// If you use Node.js and not https://sourceacademy.org,
// uncomment the following two lines:

// Object.entries(require('sicp'))
//       .forEach(([name, exported]) => global[name] = exported);

/* ****************************************
 * Idealized virtual machine for Source §4-
 * ****************************************/

// how to run: Copy this program to your favorite
// JavaScript development environment. Use
// ECMAScript 2016 or higher in Node.js.
// Import the NPM package sicp from
// https://www.npmjs.com/package/sicp

// for syntax and semantics of Source §4,
// see https://docs.sourceacademy.org/source_4.pdf

// simplifications:
//
// (1) every statement produces a value
//
// In this evaluator, all statements produce
// a value, and declarations produce undefined,
// whereas JavaScript distinguishes value-producing
// statements. This makes a difference at the top 
// level, outside of function bodies. For example, 
// in JavaScript, the execution of
// 1; const x = 2;
// results in 1, whereas the evaluator gives undefined.
// For details on this see: 
// https://sourceacademy.org/sicpjs/4.1.2#ex-4.8
//
// (2) while and for loops produce undefined
//
// In this evaluator, all while loops produce
// the value undefined, whereas in JavaScript loops
// produce:
// * undefined if the loop body is not executed
// * otherwise: the result of the last body execution

/* *************
 * parse to JSON
 * *************/

const list_to_array = xs =>
    is_null(xs)
    ? []
    : [head(xs)].concat(list_to_array(tail(xs)))

// simplify parameter format
const parameters = xs =>
    map(x => head(tail(x)),
        xs)


// turn tagged list syntax from parse into JSON object
const ast_to_json = t => {
    switch (head(t)) {
        case "literal":
            return { tag: "lit", val: head(tail(t)) }
        case "name":
            return { tag: "nam", sym: head(tail(t)) }
        case "application":
            return {
                tag: "app",
                fun: ast_to_json(head(tail(t))),
                args: list_to_array(map(ast_to_json, head(tail(tail(t)))))
            }
        case "logical_composition":
            return {
                tag: "log",
                sym: head(tail(t)),
                frst: ast_to_json(head(tail(tail(t)))),
                scnd: ast_to_json(head(tail(tail(tail(t)))))
            }
        case "binary_operator_combination":
            return {
                tag: "binop",
                sym: head(tail(t)),
                frst: ast_to_json(head(tail(tail(t)))),
                scnd: ast_to_json(head(tail(tail(tail(t)))))
            }
        case "unary_operator_combination":
            return {
                tag: "unop",
                sym: head(tail(t)),
                frst: ast_to_json(head(tail(tail(t))))
            }
        case "lambda_expression":
            return {
                tag: "lam",
                prms: list_to_array(parameters(head(tail(t)))),
                body: ast_to_json(head(tail(tail(t))))
            }
        case "sequence":
            return {
                tag: "seq",
                stmts: list_to_array(map(ast_to_json, head(tail(t))))
            }
        case "block":
            return {
                tag: "blk",
                body: ast_to_json(head(tail(t)))
            }
        case "variable_declaration":
            return {
                tag: "let",
                sym: head(tail(head(tail(t)))),
                expr: ast_to_json(head(tail(tail(t))))
            }
        case "constant_declaration":
            return {
                tag: "const",
                sym: head(tail(head(tail(t)))),
                expr: ast_to_json(head(tail(tail(t))))
            }
        case "assignment":
            return {
                tag: "assmt",
                sym: head(tail(head(tail(t)))),
                expr: ast_to_json(head(tail(tail(t))))
            }
        case "conditional_statement":
            return {
                tag: "cond", // dont distinguish stmt and expr
                pred: ast_to_json(head(tail(t))),
                cons: ast_to_json(head(tail(tail(t)))),
                alt: ast_to_json(head(tail(tail(tail(t)))))
            }
        case "conditional_expression":
            return {
                tag: "cond", // dont distinguish stmt and expr
                pred: ast_to_json(head(tail(t))),
                cons: ast_to_json(head(tail(tail(t)))),
                alt: ast_to_json(head(tail(tail(tail(t)))))
            }
        case "function_declaration":
            return {
                tag: "fun",
                sym: head(tail(head(tail(t)))),
                prms: list_to_array(parameters(head(tail(tail(t))))),
                body: ast_to_json(head(tail(tail(tail(t)))))
            }
        case "return_statement":
            return {
                tag: "ret",
                expr: ast_to_json(head(tail(t)))
            }
        case "while_loop":
            return {
                tag: "while",
                pred: ast_to_json(head(tail(t))),
                body: ast_to_json(head(tail(tail(t))))
            }            
       default:
            error(t, "unknown syntax:")
    }
}

// parse, turn into json (using ast_to_json), 
// and wrap in a block
const parse_to_json = program_text =>
    ({tag: "blk",
      body: ast_to_json(parse(program_text))});

/* **********************
 * using arrays as stacks
 * **********************/

// add values destructively to the end of 
// given array; return the array
const push = (array, ...items) => {
    array.splice(array.length, 0, ...items)
    return array 
}

// return the last element of given array
// without changing the array
const peek = array =>
    array.slice(-1)[0]


/* ********
 * compiler
 * ********/

// scanning out the declarations from (possibly nested)
// sequences of statements, ignoring blocks
const scan = comp => 
    comp.tag === 'seq'
    ? comp.stmts.reduce((acc, x) => acc.concat(scan(x)),
                        [])
    : ['let', 'const', 'fun'].includes(comp.tag)
    ? [comp.sym]
    : []

const compile_sequence = seq => {
    if (seq.length === 0) 
        return instrs[wc++] = {tag: "LDC", val: undefined}
    let first = true
    for (let comp of seq) {
        first ? first = false
              : instrs[wc++] = {tag: 'POP'}
        compile(comp)
    }
}

    
// wc: write counter
let wc
// instrs: instruction array
let instrs
    
const compile_comp = {
lit:
    comp => {
        instrs[wc++] = { tag: "LDC", val: comp.val }
    },
nam:
    comp => {
        instrs[wc++] = { tag: "LD", sym: comp.sym }
    },
unop:
    comp => {
        compile(comp.frst)
        instrs[wc++] = {tag: 'UNOP', sym: comp.sym}
    },
binop:
    comp => {
        compile(comp.frst)
        compile(comp.scnd)
        instrs[wc++] = {tag: 'BINOP', sym: comp.sym}
    },
log:
    comp => {
        compile(comp.sym == '&&' 
                ? {tag: 'cond_expr', 
                   pred: comp.frst, 
                   cons: {tag: 'lit', val: true},
                   alt: comp.scnd}
                : {tag: 'cond_expr',  
                   pred: cmd.frst,
                   cons: cmd.scnd, 
                   alt: {tag: 'lit', val: false}})
    },
cond: 
    comp => {
        compile(comp.pred)
        const jump_on_false_instruction = {tag: 'JOF'}
        instrs[wc++] = jump_on_false_instruction
        compile(comp.cons)
        const goto_instruction = { tag: 'GOTO' }
        instrs[wc++] = goto_instruction;
        const alternative_address = wc;
        jump_on_false_instruction.addr = alternative_address;
        compile(comp.alt)
        goto_instruction.addr = wc
    },
app: 
    comp => {
        compile(comp.fun)
        for (let arg of comp.args) {
            compile(arg)
        }
        instrs[wc++] = {tag: 'CALL', arity: comp.args.length}
    },
assmt: 
    comp => {
        compile(comp.expr)
        instrs[wc++] = {tag: 'ASSIGN', sym: comp.sym}
    },
lam:
    comp => {
        instrs[wc++] = {tag: 'LDF', prms: comp.przms, addr: wc + 1};
        // jump over the body of the lambda expression
        const goto_instruction = {tag: 'GOTO'}
        instrs[wc++] = goto_instruction
        compile(comp.body)
        instrs[wc++] = {tag: 'LDC', val: undefined}
        instrs[wc++] = {tag: 'RESET'}
        goto_instruction.addr = wc;
    },
seq: 
    comp => compile_sequence(comp.stmts),
blk:
    comp => {
        const locals = scan(comp.body)
        instrs[wc++] = {tag: 'ENTER_SCOPE', syms: locals}
        compile(comp.body)
        instrs[wc++] = {tag: 'EXIT_SCOPE'}
    },
let: 
    comp => {
        compile(comp.expr)
        instrs[wc++] = {tag: 'ASSIGN', sym: comp.sym}
    },
const:
    comp => {
        compile(comp.expr)
        instrs[wc++] = {tag: 'ASSIGN', sym: comp.sym}
    },
ret:
    comp => {
        compile(comp.expr)
        if (comp.expr.tag === 'app') {
            // tail call: turn CALL into TAILCALL
            instrs[wc - 1].tag = 'TAIL_CALL'
        } else {
            instrs[wc++] = {tag: 'RESET'}
        }
    },
fun:
    comp => {
        compile(
            {tag:  'const',
             sym:  comp.sym,
             expr: {tag: 'lam', prms: comp.prms, body: comp.body}})
    }
}

// compile component into instruction array instrs, 
// starting at wc (write counter)
const compile = comp => {
    compile_comp[comp.tag](comp)
    instrs[wc] = {tag: 'DONE'}
} 

// compile program into instruction array instrs, 
// after initializing wc and instrs
const compile_program = program => {
    wc = 0
    instrs = []
    compile(program)
} 

/* *************************
 * values of the machine
 * *************************/

// for numbers, strings, booleans, undefined, null
// we use the value directly

// closures aka function values
const is_closure = x =>
    x !== null && 
    typeof x === "object" &&
    x.tag === 'CLOSURE'

const is_builtin = x =>
    x !== null &&
    typeof x === "object" && 
    x.tag == 'BUILTIN'

// catching closure and builtins to get short displays
const value_to_string = x => 
     is_closure(x)
     ? '<closure>'
     : is_builtin(x)
     ? '<builtin: ' + x.sym + '>'
     : stringify(x)

/* **********************
 * operators and builtins
 * **********************/

const binop_microcode = {
    '+': (x, y)   => (is_number(x) && is_number(y)) ||
                     (is_string(x) && is_string(y))
                     ? x + y 
                     : error([x,y], "+ expects two numbers" + 
                                    " or two strings, got:"),
    // todo: add error handling to JS for the following, too
    '*':   (x, y) => x * y,
    '-':   (x, y) => x - y,
    '/':   (x, y) => x / y,
    '%':   (x, y) => x % y,
    '<':   (x, y) => x < y,
    '<=':  (x, y) => x <= y,
    '>=':  (x, y) => x >= y,
    '>':   (x, y) => x > y,
    '===': (x, y) => x === y,
    '!==': (x, y) => x !== y
}

// v2 is popped before v1
const apply_binop = (op, v2, v1) => binop_microcode[op](v1, v2)

const unop_microcode = {
    '-unary': x => - x,
    '!'     : x => is_boolean(x) 
                   ? ! x 
                   : error(x, '! expects boolean, found:')
}

const apply_unop = (op, v) => unop_microcode[op](v)

const builtin_mapping = {
    display       : display,
    get_time      : get_time,
    stringify     : stringify,
    error         : x => { PC = instrs.length - 1; return x; },
    prompt        : prompt,
    is_number     : is_number,
    is_string     : is_string,
    is_function   : x => typeof x === 'object' &&
                         (x.tag == 'BUILTIN' ||
                          x.tag == 'CLOSURE'),
    is_boolean    : is_boolean,
    is_undefined  : is_undefined,
    parse_int     : parse_int,
    char_at       : char_at,
    arity         : x => typeof x === 'object' 
                         ? x.arity
                         : error(x, 'arity expects function, received:'),
    math_abs      : math_abs,
    math_acos     : math_acos,
    math_acosh    : math_acosh,
    math_asin     : math_asin,
    math_asinh    : math_asinh,
    math_atan     : math_atan,
    math_atanh    : math_atanh,
    math_atan2    : math_atan2,
    math_ceil     : math_ceil,
    math_cbrt     : math_cbrt,
    math_expm1    : math_expm1,
    math_clz32    : math_clz32,
    math_cos      : math_cos,
    math_cosh     : math_cosh,
    math_exp      : math_exp,
    math_floor    : math_floor,
    math_fround   : math_fround,
    math_hypot    : math_hypot,
    math_imul     : math_imul,
    math_log      : math_log,
    math_log1p    : math_log1p,
    math_log2     : math_log2,
    math_log10    : math_log10,
    math_max      : math_max,
    math_min      : math_min,
    math_pow      : math_pow,
    math_random   : math_random,
    math_round    : math_round,
    math_sign     : math_sign,
    math_sin      : math_sin,
    math_sinh     : math_sinh,
    math_sqrt     : math_sqrt,
    math_tanh     : math_tanh,
    math_trunc    : math_trunc,
    pair          : pair,
    is_pair       : is_pair,
    head          : head,
    tail          : tail,
    is_null       : is_null,
    set_head      : set_head,
    set_tail      : set_tail,
    array_length  : array_length,
    is_array      : is_array,
    list          : list,
    is_list       : is_list,
    display_list  : display_list,
    // from list libarary
    equal         : equal,
    length        : length,
    list_to_string: list_to_string,
    reverse       : reverse,
    append        : append,
    member        : member,
    remove        : remove,
    remove_all    : remove_all,
    enum_list     : enum_list,
    list_ref      : list_ref,
    // misc
    draw_data     : draw_data,
    parse         : parse,
    tokenize      : tokenize,
    apply_in_underlying_javascript: apply_in_underlying_javascript
}

const apply_builtin = (builtin_symbol, args) =>
    builtin_mapping[builtin_symbol](...args)

/* ************
 * environments
 * ************/

// Frames are objects that map symbols (strings) to values.

const global_frame = {}

// fill global frame with built-in objects
for (const key in builtin_mapping) 
    global_frame[key] = { tag:   'BUILTIN', 
                          sym:   key, 
                          arity: arity(builtin_mapping[key])
                        }
// fill global frame with built-in constants
global_frame.undefined    = undefined
global_frame.math_E       = math_E
global_frame.math_LN10    = math_LN10
global_frame.math_LN2     = math_LN2
global_frame.math_LOG10E  = math_LOG10E
global_frame.math_LOG2E   = math_LOG2E
global_frame.math_PI      = math_PI
global_frame.math_SQRT1_2 = math_SQRT1_2
global_frame.math_SQRT2   = math_SQRT2

// An environment is null or a pair whose head is a frame 
// and whose tail is an environment.
const empty_environment = null
const global_environment = pair(global_frame, empty_environment)

const lookup = (x, e) => {
    if (is_null(e)) 
        error(x, 'unbound name:')
    if (head(e).hasOwnProperty(x)) {
        const v = head(e)[x]
        if (is_unassigned(v))
            error(cmd.sym, 'unassigned name:')
        return v
    }
    return lookup(x, tail(e))
}

const assign_value = (x, v, e) => {
    if (is_null(e))
        error(x, 'unbound name:')
    if (head(e).hasOwnProperty(x)) {
        head(e)[x] = v
    } else {
        assign_value(x, v, tail(e))
    }
}

const extend = (xs, vs, e) => {
    if (vs.length > xs.length) error('too many arguments')
    if (vs.length < xs.length) error('too few arguments')
    const new_frame = {}
    for (let i = 0; i < xs.length; i++) 
        new_frame[xs[i]] = vs[i]
    return pair(new_frame, e)
}

// At the start of executing a block, local 
// variables refer to unassigned values.
const unassigned = { tag: 'unassigned' }

const is_unassigned = v => {
    return v !== null && 
    typeof v === "object" && 
    v.hasOwnProperty('tag') &&
    v.tag === 'unassigned'
} 

/* *******
 * machine
 * *******/

let OS
let PC
let E
let RTS

const microcode = {
LDC:
    instr => {
        PC++
        push(OS, instr.val);
    },
UNOP:
    instr => {
        PC++
        push(OS, apply_unop(instr.sym, OS.pop()))
    },
BINOP:
    instr => {
        PC++
        push(OS, apply_binop(instr.sym, OS.pop(), OS.pop()))
    },
POP: 
    instr => {
        PC++
        OS.pop()
    },
JOF: 
    instr => {
        PC = OS.pop() ? PC + 1 : instr.addr
    },
GOTO:
    instr => {
        PC = instr.addr
    },
ENTER_SCOPE: 
    instr => {
        PC++
        push(RTS, {tag: 'BLOCK_FRAME', env: E})
        const locals = instr.syms
        const unassigneds = locals.map(_ => unassigned)
        E = extend(locals, unassigneds, E)
    }, 
EXIT_SCOPE: 
    instr => {
        PC++
        E = RTS.pop().env
    },
LD: 
    instr => {
        PC++
        push(OS, lookup(instr.sym, E))
    },
ASSIGN: 
    instr => {
        PC++
        assign_value(instr.sym, peek(OS), E)
    },
LDF: 
    instr => {
        PC++
        push(OS, {tag: 'CLOSURE', prms: instr.prms, 
                   addr: instr.addr, env: E})
    },
CALL: 
    instr => {
        const arity = instr.arity
        let args = []
        for (let i = arity - 1; i >= 0; i--)
            args[i] = OS.pop()
        const sf = OS.pop()
        if (sf.tag === 'BUILTIN') {
            PC++
            return push(OS, apply_builtin(sf.sym, args))
        }
        push(RTS, {tag: 'CALL_FRAME', addr: PC + 1, env: E})
        E = extend(sf.prms, args, sf.env)
        PC = sf.addr
    },
TAIL_CALL: 
    instr => {
        const arity = instr.arity
        let args = []
        for (let i = arity - 1; i >= 0; i--)
            args[i] = OS.pop()
        const sf = OS.pop()
        if (sf.tag === 'BUILTIN') {
            PC++
            return push(OS, apply_builtin(sf.sym, args))
        }
        // dont push on RTS here
        E = extend(sf.prms, args, sf.env)
        PC = sf.addr
    },
RESET : 
    instr => {
        // keep popping...
        const top_frame = RTS.pop()
        if (top_frame.tag === 'CALL_FRAME') {
            // ...until top frame is a call frame
            PC = top_frame.addr
            E = top_frame.env
        }
    }
}

function run() {
    OS = []
    PC = 0
    E = global_environment
    RTS = []    
    //print_code(instrs)
    while (! (instrs[PC].tag === 'DONE')) {
        //display("next instruction: ")
        //print_code([instrs[PC]]) 
        //display(PC, "PC: ")
        //print_OS("\noperands:            ");
        //print_RTS("\nRTS:            ");
        const instr = instrs[PC]
        microcode[instr.tag](instr)
    }
    return peek(OS)
} 

// parse_compile_run on top level
// * parse input using parse_to_json
// * compile syntax tree into vm code
// * run code

const parse_compile_run = program => {
    compile_program(parse_to_json(program))
    return run()
}

/* *********
 * debugging
 * *********/

const print_code = (instrs) => {
    for (let i = 0; i < instrs.length; i = i + 1) {
        const instr = instrs[i]
        display("", stringify(i) + ": " + instr.tag +
                    " " +
                    (instr.tag === 'GOTO' 
                     ? stringify(instr.addr)
                     : "")  +
                    (instr.tag === 'LDC' 
                     ? stringify(instr.val)
                     : "")  +
                    (instr.tag === 'ASSIGN' 
                     ? stringify(instr.sym)
                     : "") +
                    (instr.tag === 'LD' 
                     ? stringify(instr.sym)
                     : "")
                     )
    }
}

const print_RTS = (x) => {
    display("",x)
    for (let i = 0; i < RTS.length; i = i + 1) {
        const f = RTS[i]
        display("", stringify(i) + ": " + f.tag)
    }
}

const print_OS = (x) => {
    display("",x)
    for (let i = 0; i < OS.length; i = i + 1) {
        const val = OS[i]
        display("", stringify(i) + ": " +
                    (is_closure(val) 
                    ? "<closure>"
                    : val)
                    )
    }
}

/* *******
 * testing
 * *******/

const test = (program, expected) => {
    display("", `
    
****************
Test case: ` + program + "\n")
    const result = parse_compile_run(program)
    if (stringify(result) === stringify(expected)) {
        display(result, "success:")
    } else {
        display(expected, "FAILURE! expected:")
        error(result, "result:")
    }
}


test("1;", 1);

test("2 + 3;", 5);

test("1; 2; 3;", 3);

test("false ? 2 : 3;", 3);

test("8 + 34; true ? 1 + 2 : 17;", 3);

test(`
const y = 4; 
{
    const x = y + 7; 
    x * 2;
}
`, 22);

test(`
function f() {
    return 1;
}
f();
`, 1);

test(`
function f(x) {
    return x;
}
f(33);
`, 33);

test(`
function f(x, y) {
    return x - y;
}
f(33, 22);
`, 11);

test(`
function fact(n) {
    return n === 1 ? 1 : n * fact(n - 1);
}
fact(10);
`, 3628800);

test("error(1); 2;", 1)

test("error(1 + 2); 4;", 3)

test(`
function fact(n) {
    return fact_iter(n, 1, 1);
}
function fact_iter(n, i, acc) {
    return i > n
           ? error(100)
           : fact_iter(n, i + 1, acc * i);
}
fact(4);
`, 100);