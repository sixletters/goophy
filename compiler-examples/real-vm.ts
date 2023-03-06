/* *********************************************
 * More realistic virtual machine for Source §3-
 * *********************************************/

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
// (2) no loops and arrays

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
const peek = (array, address) =>
    array.slice(-1 - address)[0]

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

/* *************************
 * HEAP
 * *************************/

// HEAP is an array of bytes (JS ArrayBuffer)

const word_size = 8
const mega = 2 ** 20

// heap_make allocates a heap of given size (in megabytes)
// and returns a DataView of that, see 
// https://www.javascripture.com/DataView
const heap_make = mega_bytes => {
    const data = new ArrayBuffer(mega * mega_bytes)
    const view = new DataView(data)
    return view
}

// we randomly pick a heap size of 10 megabytes
const HEAP = heap_make(10)

// free is the next free address;
// we keep allocating as if there was no tomorrow
let free = 0

// for debugging: display all bits of the heap
const heap_display = () => {
    display("", "heap:")
    for (let i = 0; i < free; i++) {
        display(
            word_to_string(
                heap_get_word_at_index(i)), 
            stringify(i) + " " +
            stringify(heap_get_word_at_index(i)) +
            " ")
    }
}

const heap_get_word_at_index = index =>
    HEAP.getFloat64(index * word_size)

const heap_set_word_at_index = (index, x) => {
        return HEAP.setFloat64(index * word_size, x)
    }

// for debugging: return a string that shows the bits
// of a given word
const word_to_string = word => {
    const buf = new ArrayBuffer(8);
    const view = new DataView(buf);
    view.setFloat64(0, word);
    let binStr = '';
    for (let i = 0; i < 8; i++) {
        binStr += ('00000000' + view.getUint8(i).toString(2)).slice(-8) + ' ';
    }
    return binStr
}

// primitive values 

// primitive values are represented by words that meet 
// the IEEE 754 representation of NaN ("not-a-number").
// Within the NaN value, we encode the type of the value 
// using a tag, from bit 13 to bit 16, using the following 
// scheme.

const NaN_tag            = [0,0,0,0]

// heap words without payload (literal values)
const False_tag          = [1,0,0,0]
const True_tag           = [1,0,0,1]
const Null_tag           = [1,0,1,0]
const Unassigned_tag     = [1,0,1,1]
const Undefined_tag      = [1,1,0,0]

// heap nodes with child nodes
const Blockframe_tag     = [1,1,0,1]
const Callframe_tag      = [1,1,1,0]
const Closure_tag        = [0,0,0,1]
const Frame_tag          = [0,0,1,0]
const Environment_tag    = [0,0,1,1]
const Pair_tag           = [0,1,0,0]

// heap words with payload
const Builtin_tag        = [0,1,0,1]
const Address_tag        = [0,1,1,0]

const String_tag         = [1,1,1,1]

// the first 13 bits of a NaN word are 
// needed to make the word a NaN. The
// subsequent bits are used as tags
const NaN_tag_offset = 13

const heap_set_tagged_NaN_at_index = (index, tag) => {
    const the_NaN = make_tagged_NaN(tag)
    heap_set_word_at_index(index, the_NaN)
}

// some magic bit manipulation: in view v,
// set bit at index i to 1, where 0 <= i < 64
const set_bit = (x, i) => {
        const byte_index = Math.floor(i / 8)
        const current_byte = x.getUint8(byte_index);
        const bit_index = 7 - (i % 8)
        x.setUint8(byte_index, 
                   current_byte | 1 << bit_index)
    }
// some more magic bit manipulation: in view v,
// set bit at index i to 1, where 0 <= i < 64
const unset_bit = (x, i) => {
        const byte_index = Math.floor(i / 8)
        const current_byte = x.getUint8(byte_index);
        const bit_index = 7 - (i % 8)
        x.setUint8(byte_index, 
                   current_byte & ~(1 << bit_index))
    }

// returns a NaN that is tagged as specified
const make_tagged_NaN = tag => {
        const buf = new ArrayBuffer(8);
        const view = new DataView(buf);
        view.setFloat64(0, NaN);
        for (let i = 0; i < tag.length; i ++) {
             (tag[i] ? set_bit : unset_bit)
             (view, NaN_tag_offset + i)
        }
        return view.getFloat64(0)
    }

// get bit at position i from view v
const get_bit = (v, i) => {
        const byteIndex = Math.floor(i / 8)
        const bitIndex = 7 - (i % 8)
        return (v.getUint8(byteIndex) >> bitIndex) & 1;
    }

// check that word x is a NaN tagged as specified
const check_tag = (x, tag) => {
        if (! isNaN(x)) return false
        const buf = new ArrayBuffer(8);
        const view = new DataView(buf);
        view.setFloat64(0, x)
        let answer = true
        for (let i = 0; i < tag.length; i ++) {
             answer &&= (get_bit(
                            view, 
                            NaN_tag_offset + i) === tag[i])
        }
        return answer
    }

// check if word x is tagged with any non-zero tag
const is_tagged = x =>
    isNaN(x) && ! check_tag(x, NaN_tag)

// numbers include the untagged NaN
const is_Number = x =>
    is_number(x) && ! is_tagged(x)

// literal values are encoded with NaNs 
// as follows

const False = make_tagged_NaN(False_tag)
const is_False = x => 
    check_tag(x, False_tag)
    
const True = make_tagged_NaN(True_tag)
const is_True = x => 
    check_tag(x, True_tag)
    
const is_Boolean = x =>
    is_True(x) || is_False(x)

const Null = make_tagged_NaN(Null_tag)
const is_Null = x => 
    check_tag(x, Null_tag)

const Unassigned = make_tagged_NaN(Unassigned_tag)
const is_Unassigned = x => 
    check_tag(x, Unassigned_tag)

const Undefined = make_tagged_NaN(Undefined_tag)
const is_Undefined = x => 
    check_tag(x, Undefined_tag)

const is_String = x =>{
    return check_tag(x,String_tag)
}

// heap addresses: index encoded
// as payload of heap word (using NaN-boxing)

const is_Address = x =>
    check_tag(x, Address_tag)

// the function returns a heap word that
// encodes the heap index using the right-most 32 bits.
// note that this limits the heap size to 4 Terabytes.
const make_Address = index => {
        const address = make_tagged_NaN(Address_tag)
        const buf = new ArrayBuffer(8)
        const view = new DataView(buf)
        view.setFloat64(0, address)
        view.setInt32(4, index)
        return view.getFloat64(0)
    }

// retrieve heap index from Address
const get_index_from_Address = 
    address => {
        const buf = new ArrayBuffer(8)
        const view = new DataView(buf)
        view.setFloat64(0, address)
        return view.getInt32(4)
    }

// access the heap with a given address
const heap_Address_deref =
    address => 
    heap_get_word_at_index(
        get_index_from_Address(address))
                
// builtins: builtin id is encoded
// as payload of heap word (using NaN-boxing)

const is_Builtin = x =>
    check_tag(x, Builtin_tag) 



// the function returns a heap word that
// encodes the index using the right-most 32 bits
const make_Builtin = id => {
        const address = make_tagged_NaN(Builtin_tag)
        const buf = new ArrayBuffer(8)
        const view = new DataView(buf)
        view.setFloat64(0, address)
        view.setInt32(4, id)
        return view.getFloat64(0)
    }

// retrieve the heap index from Address
const get_id_from_Builtin = 
    builtin => {
        const buf = new ArrayBuffer(8)
        const view = new DataView(buf)
        view.setFloat64(0, builtin)
        return view.getInt32(4)
    }
    
let string_pool = []

let string_pool_idx = {}

const make_string_word = index => {
    const string_tag = make_tagged_NaN(String_tag)
    const buf = new ArrayBuffer(8)
    const view = new DataView(buf)
    view.setFloat64(0, string_tag)
    view.setInt32(4, index)
    return view.getFloat64(0)
}

const get_string_index = word => {
        const buf = new ArrayBuffer(8)
        const view = new DataView(buf)
        view.setFloat64(0, word)
        return view.getInt32(4)
}
// adds string to string pool and returns the word
// if string already exists, will not be added
const add_string = x => {
    if(!(x in string_pool_idx)){
        string_pool.push(x)
        let idx = string_pool.length - 1
        string_pool_idx[x] = idx
    }
    return make_string_word(string_pool_idx[x])
}
const word_to_JS_value = x =>
    is_Boolean(x)
    ? (is_True(x) ? true : false)
    : is_Undefined(x)
    ? undefined
    : is_Unassigned(x) 
    ? "<unassigned>" 
    : is_Null(x) 
    ? null 
    : is_Pair(x)
    ? [
        word_to_JS_value(heap_get_head(x)),
        word_to_JS_value(heap_get_tail(x))
        ]
    : is_Address(x)
    ? word_to_JS_value(heap_Address_deref(x))

    : is_Closure(x)
    ? "<closure>"
    : is_Builtin(x)
    
    ? "<builtin>"
    : is_String(x)
    ? string_pool[get_string_index(x)]
    : isNaN(x)
    ? word_to_string(x)
    : x
    
// closure 

const Closure_pc_offset = 1
const Closure_arity_offset = 2
const Closure_environment_offset = 3
const Closure_size = 4

const heap_allocate_Closure = (pc, arity, env) => {
        const closure_index = free
        free += Closure_size
        heap_set_tagged_NaN_at_index(
            closure_index, Closure_tag)
        heap_set_word_at_index(
            closure_index + Closure_pc_offset, pc)
        heap_set_word_at_index(
            closure_index + Closure_arity_offset, arity)
        heap_set_word_at_index(
            closure_index + Closure_environment_offset, env)
        return make_Address(closure_index)
    }

const heap_get_Closure_pc = address =>
    heap_get_word_at_index(
                   get_index_from_Address(address) + 
                   Closure_pc_offset)

const heap_get_Closure_arity = address => 
    heap_get_word_at_index(
                   get_index_from_Address(address) + 
                   Closure_arity_offset)

const heap_get_Closure_environment = address =>
    heap_get_word_at_index(
                   get_index_from_Address(address) + 
                   Closure_environment_offset)

const is_Closure = x => {
        return (check_tag(x, Address_tag) &&
                check_tag(
                    heap_Address_deref(x),
                    Closure_tag))
    }

// block frame 

const Blockframe_environment_offset = 1
const Blockframe_size = 2

const heap_allocate_Blockframe = (env) => {
        const frame_index = free
        free += Blockframe_size
        heap_set_tagged_NaN_at_index(
            frame_index, Blockframe_tag)
        heap_set_word_at_index(
            frame_index + Blockframe_environment_offset, env)
        return make_Address(frame_index)
    }

const heap_get_Blockframe_environment = address =>
    heap_get_word_at_index(
                   get_index_from_Address(address) + 
                   Blockframe_environment_offset)

const is_Blockframe = x => {
        return (check_tag(x, Address_tag) &&
                check_tag(
                    heap_Address_deref(x),
                    Blockframe_tag))
    }

// call frame 

const Callframe_environment_offset = 1
const Callframe_pc_offset = 2
const Callframe_size = 3

const heap_allocate_Callframe = (env, pc) => {
        const frame_index = free
        free += Callframe_size
        heap_set_tagged_NaN_at_index(
            frame_index, Callframe_tag)
        heap_set_word_at_index(
            frame_index + Callframe_environment_offset, env)
        heap_set_word_at_index(
            frame_index + Callframe_pc_offset, pc)
        return make_Address(frame_index)
    }

const heap_get_Callframe_environment = address =>
    heap_get_word_at_index(
                   get_index_from_Address(address) + 
                   Callframe_environment_offset)

const heap_get_Callframe_pc = address =>
    heap_get_word_at_index(
                   get_index_from_Address(address) + 
                   Callframe_pc_offset)

const is_Callframe = x => {
        return (check_tag(x, Address_tag) &&
                check_tag(
                    heap_Address_deref(x),
                    Callframe_tag))
    }


// environment frame

const Frame_size_offset = 1
const Frame_values_offset = 2

// size is number of words to be reserved
// for values
const heap_allocate_Frame = size => {
        const frame_index = free
        free += Frame_values_offset + size
        heap_set_tagged_NaN_at_index(frame_index, Frame_tag)
        heap_set_word_at_index(frame_index + Frame_size_offset, size)
        return make_Address(frame_index)
    }

const heap_get_Frame_size = address =>
    heap_get_word_at_index(
        get_index_from_Address(frame_address) + 
        Frame_size_offset)

const heap_get_Frame_value = 
    (frame_address, value_index) =>
    heap_get_word_at_index(
        get_index_from_Address(frame_address) + 
        Frame_values_offset + 
        value_index)

const heap_set_Frame_value = 
    (frame_address, value_index, value) =>
    heap_set_word_at_index(
        get_index_from_Address(frame_address) + 
        Frame_values_offset + 
        value_index,
        value)

const heap_Frame_display = frame_address => {
        display("", "Frame:")
        const size = heap_get_Frame_size(frame_address)
        display(size, "frame size:")
        for (let i = 0; i < size; i++) {
            display(i, "value index:")
            const value = 
                  heap_get_Frame_value(frame_address, i)
            display(value, "value:")
            display(word_to_string(value), "value word:")
        }
    }
        
// environment

// environments are heap nodes that contain 
// addresses of frames
const Environment_size_offset = 1
const Environment_frames_offset = 2

const heap_allocate_Environment = size => {
        const env_index = free
        free += Environment_frames_offset + size
        heap_set_tagged_NaN_at_index(
            env_index, Environment_tag)
        heap_set_word_at_index(
            env_index + Environment_size_offset, size)
        return make_Address(env_index)
    }

const heap_empty_Environment = heap_allocate_Environment(0)

const heap_get_Environment_size = env_address =>
    heap_get_word_at_index(
        get_index_from_Address(env_address) + 
        Environment_size_offset)

// access environment given by address 
// using a "position", i.e. a pair of 
// frame index and value index
const heap_get_Environment_value = 
    (env_address, position) => {
        const [frame_index, value_index] = position
        const frame_address =
            heap_get_word_at_index(
                get_index_from_Address(env_address) + 
                Environment_frames_offset + 
                frame_index)
        return heap_get_Frame_value(
                   frame_address, value_index)
    }

const heap_set_Environment_value =
    (env_address, position, value) => {
        const [frame_index, value_index] = position
        const frame_address =
            heap_get_word_at_index(
                get_index_from_Address(env_address) + 
                Environment_frames_offset + 
                frame_index)
        heap_set_Frame_value(
            frame_address, value_index, value)
    }

// get the whole frame at given frame index
const heap_get_Environment_frame = 
    (env_address, frame_index) =>
    heap_get_word_at_index(
         get_index_from_Address(env_address) + 
         Environment_frames_offset + 
         frame_index)

// set the whole frame at given frame index
const heap_set_Environment_frame =
    (env_address, frame_index, frame) =>
         heap_set_word_at_index(
                get_index_from_Address(env_address) + 
                Environment_frames_offset + 
                frame_index,
                frame)

// extend a given environment by a new frame: 
// create a new environment that is bigger by 1
// frame slot than the given environment.
// copy the frame address to the new environment.
// enter the address of the new frame to end 
// of the new environment
const heap_Environment_extend =
    (frame_address, env_address) => {
        const old_size = 
            heap_get_Environment_size(env_address)
        const new_env_address =
            heap_allocate_Environment(old_size + 1)
        let i
        for (i = 0; i < old_size; i++) {
            heap_set_Environment_frame(
                new_env_address, i,
                heap_get_Environment_frame(
                    env_address, i))
        }
        heap_set_Environment_frame(
            new_env_address, i, frame_address)
        return new_env_address
    }

// for debuggging: display environment
const heap_Environment_display =
    env_address => {
        const size = heap_get_Environment_size(env_address)
        display("", "Environment:")
        display(size, "environment size:")
        for (let i = 0; i < size; i++) {
            display(i, "frame index:")
            const frame = 
                  heap_get_Environment_frame(env_address, i)
            heap_Frame_display(frame)
        }
    }
    
// pair

const Pair_head_offset = 1
const Pair_tail_offset = 2
const Pair_size = 3

const heap_allocate_Pair = (hd, tl) => {
        const pair_index = free
        free += Pair_size
        heap_set_tagged_NaN_at_index(pair_index, Pair_tag)
        heap_set_word_at_index(pair_index + Pair_head_offset, hd)
        heap_set_word_at_index(pair_index + Pair_tail_offset, tl)

        return make_Address(pair_index)
    }

const heap_get_head = address =>
    heap_get_word_at_index(get_index_from_Address(address) + 
                  Pair_head_offset)

const heap_get_tail = address =>
    heap_get_word_at_index(get_index_from_Address(address) + 
                  Pair_tail_offset)

const heap_set_head = (address, val) =>
    heap_set_word_at_index(address + Pair_head_offset, val)

const heap_set_tail = (address, val) =>
    heap_set_word_at_index(address + Pair_tail_offset, val)

const is_Pair = x => {
    return (check_tag(x, Address_tag) &&
            check_tag(
                    heap_Address_deref(x),
                    Pair_tag))
    }

/* ************************
 * compile-time environment
 * ************************/
 
// a compile-time environment is an array of 
// compile-time frames, and a compile-time frame 
// is an array of symbols

// find the position [frame-index, value-index] 
// of a given symbol x
const compile_time_environment_position = (env, x) => {
    let frame_index = env.length
    while (value_index(env[--frame_index], x) === -1) {}
    return [frame_index, 
            value_index(env[frame_index], x)]
}

const value_index = (frame, x) => {
  for (let i = 0; i < frame.length; i++) {
    if (frame[i] === x) return i
  }
  return -1;
}

// in this machine, the builtins take their
// arguments directly from the operand stack,
// to save the creation of an intermediate 
// argument array
const builtin_object = {
    display       : () => display(OS.pop()),
    get_time      : () => get_time(),
    stringify     : () => stringify(OS.pop()),
    error         : () => error(OS.pop()),
    prompt        : () => prompt(OS.pop()),
    is_number     : () => is_Number(OS.pop()) ? True : False,
    is_string     : () => is_String(OS.pop()) ? True : False,
    is_function   : () => is_Closure(OS.pop()) ? True : False,
    is_boolean    : () => is_Boolean(OS.pop()) ? True : False,
    is_undefined  : () => is_Undefined(OS.pop()) ? True : False,
    math_sqrt     : () => math_sqrt(OS.pop()),
    pair          : () => {
                        const tl = OS.pop()
                        const hd = OS.pop()
                        return heap_allocate_Pair(hd, tl)
                    },
    is_pair       : () => is_Pair(OS.pop()) ? True : False,
    head          : () => heap_get_head(OS.pop()),
    tail          : () => heap_get_tail(OS.pop()),
    is_null       : () => is_Null(OS.pop()) ? True : False,
    set_head      : () => {
                        const val = OS.pop()
                        const p = OS.pop()
                        heap_set_head(p, val)
                    },
    set_tail      : () => {
                        const val = OS.pop()
                        const p = OS.pop()
                        heap_set_tail(p, val)
                    }
}

const primitive_object = {}
const builtin_array = []
{
    let i = 0
    for (const key in builtin_object) {
        primitive_object[key] = { tag:   'BUILTIN', 
                                  id:    i,
                                  arity: arity(builtin_object[key])
                                }
        builtin_array[i++] = builtin_object[key]
    }
}

const constants = {
    undefined     : Undefined,
    math_E        : math_E,
    math_LN10     : math_LN10,
    math_LN2      : math_LN2,
    math_LOG10E   : math_LOG10E,
    math_LOG2E    : math_LOG2E,
    math_PI       : math_PI,
    math_SQRT1_2  : math_SQRT1_2,
    math_SQRT2    : math_SQRT2 
}

for (const key in constants) 
    primitive_object[key] = constants[key]

const compile_time_environment_extend = (vs, e) => {
    //  make shallow copy of e
    return push([...e], vs)
}

// compile-time frames only need synbols (keys), no values
const global_compile_frame = Object.keys(primitive_object)
const global_compile_environment = [global_compile_frame]

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

const compile_sequence = (seq, ce) => {
    if (seq.length === 0) 
        return instrs[wc++] = {tag: "LDC", val: undefined}
    let first = true
    for (let comp of seq) {
        first ? first = false
              : instrs[wc++] = {tag: 'POP'}
        compile(comp, ce)
    }
}

    
// wc: write counter
let wc
// instrs: instruction array
let instrs
    
const compile_comp = {
lit:
    (comp, ce) => {
        instrs[wc++] = { tag: "LDC", 
                         val: 
                         is_null(comp.val)
                         ? Null 
                         : is_string(comp.val)
                         ? add_string(comp.val)
                         : is_boolean(comp.val)
                         ? (comp.val ? True : False)
                         : comp.val
        }
    },
nam:
    // store precomputed position information in LD instruction
    (comp, ce) => {
        instrs[wc++] = { tag: "LD", 
                         sym: comp.sym,
                         pos: compile_time_environment_position(
                                  ce, comp.sym)
                        }
    },
unop:
    (comp, ce) => {
        compile(comp.frst, ce)
        instrs[wc++] = {tag: 'UNOP', sym: comp.sym}
    },
binop:
    (comp, ce) => {
        compile(comp.frst, ce)
        compile(comp.scnd, ce)
        instrs[wc++] = {tag: 'BINOP', sym: comp.sym}
    },
log:
    (comp, ce) => {
        compile(comp.sym == '&&' 
                ? {tag: 'cond_expr', 
                   pred: comp.frst, 
                   cons: {tag: 'lit', val: true},
                   alt: comp.scnd}
                : {tag: 'cond_expr',  
                   pred: cmd.frst,
                   cons: cmd.scnd, 
                   alt: {tag: 'lit', val: false}},
	            ce)
    },
cond: 
    (comp, ce) => {
        compile(comp.pred, ce)
        const jump_on_false_instruction = {tag: 'JOF'}
        instrs[wc++] = jump_on_false_instruction
        compile(comp.cons, ce)
        const goto_instruction = { tag: 'GOTO' }
        instrs[wc++] = goto_instruction;
        const alternative_address = wc;
        jump_on_false_instruction.addr = alternative_address;
        compile(comp.alt, ce)
        goto_instruction.addr = wc
    },
while:
    (comp, ce) => {
        const loop_start = wc
        compile(comp.pred, ce)
        const jump_on_false_instruction = {tag: 'JOF'}
        instrs[wc++] = jump_on_false_instruction
        compile(comp.body, ce)
        instrs[wc++] = {tag: 'POP'}
        instrs[wc++] = {tag: 'GOTO', addr: loop_start}
        jump_on_false_instruction.addr = wc
        instrs[wc++] = {tag: 'LDC', val: Undefined}
    }, 
app:
    (comp, ce) => {
        compile(comp.fun, ce)
        for (let arg of comp.args) {
            compile(arg, ce)
        }
        instrs[wc++] = {tag: 'CALL', arity: comp.args.length}
    },
assmt:
    // store precomputed position information in ASSIGN instruction
    (comp, ce) => {
        compile(comp.expr, ce)
        instrs[wc++] = {tag: 'ASSIGN', 
                        pos: compile_time_environment_position(
                                 ce, comp.sym)}
    },
lam:
    (comp, ce) => {
        instrs[wc++] = {tag: 'LDF', arity: comp.arity, addr: wc + 1};
        // jump over the body of the lambda expression
        const goto_instruction = {tag: 'GOTO'}
        instrs[wc++] = goto_instruction
        // extend compile-time environment
        compile(comp.body,
		        compile_time_environment_extend(
		            comp.prms, ce))
        instrs[wc++] = {tag: 'LDC', val: Undefined}
        instrs[wc++] = {tag: 'RESET'}
        goto_instruction.addr = wc;
    },
seq: 
    (comp, ce) => compile_sequence(comp.stmts, ce),
blk:
    (comp, ce) => {
        const locals = scan(comp.body)
        instrs[wc++] = {tag: 'ENTER_SCOPE', num: locals.length}
        compile(comp.body,
                // extend compile-time environment
		        compile_time_environment_extend(
		            locals, ce))     
        instrs[wc++] = {tag: 'EXIT_SCOPE'}
    },
let: 
    (comp, ce) => {
        compile(comp.expr, ce)
        instrs[wc++] = {tag: 'ASSIGN', 
                        pos: compile_time_environment_position(
                                 ce, comp.sym)}
    },
const:
    (comp, ce) => {
        compile(comp.expr, ce)
        instrs[wc++] = {tag: 'ASSIGN', 
                        pos: compile_time_environment_position(
                                 ce, comp.sym)}
    },
ret:
    (comp, ce) => {
        compile(comp.expr, ce)
        if (comp.expr.tag === 'app') {
            // tail call: turn CALL into TAILCALL
            instrs[wc - 1].tag = 'TAIL_CALL'
        } else {
            instrs[wc++] = {tag: 'RESET'}
        }
    },
fun:
    (comp, ce) => {
        compile(
            {tag:  'const',
             sym:  comp.sym,
             expr: {tag: 'lam', 
                    prms: comp.prms, 
                    body: comp.body}},
	        ce)
    }
}

// compile component into instruction array instrs, 
// starting at wc (write counter)
const compile = (comp, ce) => {
    compile_comp[comp.tag](comp, ce)
} 

// compile program into instruction array instrs, 
// after initializing wc and instrs
const compile_program = program => {
    wc = 0
    instrs = []    
    compile(program, global_compile_environment)
    instrs[wc] = {tag: 'DONE'}
} 

/* **********************
 * operators and builtins
 * **********************/

const binop_microcode = {
    '+': (x, y)   => (is_Number(x) && is_Number(y))
                     ? x + y 
                     :(is_String(x) && is_String(y))
                     ? add_string(string_pool[get_string_index(x)] + string_pool[get_string_index(y)])
                     : error([x,y], "+ expects two numbers" + 
                                    " or two strings, got:"),
    // todo: add error handling to JS for the following, too
    '*':   (x, y) => x * y,
    '-':   (x, y) => x - y,
    '/':   (x, y) => x / y,
    '%':   (x, y) => x % y,
    '<':   (x, y) => x < y ? True : False,
    '<=':  (x, y) => x <= y ? True : False,
    '>=':  (x, y) => x >= y ? True : False,
    '>':   (x, y) => x > y ? True : False,
    '===': (x, y) => x === y ? True : False,
    '!==': (x, y) => x !== y ? True : False
}

// v2 is popped before v1
const apply_binop = (op, v2, v1) => binop_microcode[op](v1, v2)

const unop_microcode = {
    '-unary': x => - x,
    '!'     : x => is_Boolean(x) 
                   ? (is_True(x) ? False : True)
                   : error(x, '! expects boolean, found:')
}

const apply_unop = (op, v) => unop_microcode[op](v)

const apply_builtin = (builtin_id) => {
    const result = builtin_array[builtin_id]()
    OS.pop() // pop fun
    push(OS, result)
}

// creating global runtime environment
const primitive_values = Object.values(primitive_object)
const frame_address = 
            heap_allocate_Frame(primitive_values.length)
for (let i = 0; i < primitive_values.length; i++) {
    const primitive_value = primitive_values[i];
    if (typeof primitive_value === "object" && 
        primitive_value.hasOwnProperty("id")) {
        heap_set_Frame_value(
            frame_address, 
            i, 
            make_Builtin(primitive_value.id))
    } else {
        heap_set_Frame_value(
            frame_address, 
            i,
            primitive_value)
    }
}

const global_environment = 
      heap_Environment_extend(frame_address, 
                              heap_empty_Environment)
          
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
        PC = is_True(OS.pop()) ? PC + 1 : instr.addr
    },
GOTO:
    instr => {
        PC = instr.addr
    },
ENTER_SCOPE: 
    instr => {
        PC++
        push(RTS, heap_allocate_Blockframe(E))
        const frame_address = heap_allocate_Frame(instr.num)
        E = heap_Environment_extend(frame_address, E)
    }, 
EXIT_SCOPE:
    instr => {
        PC++
        E = heap_get_Blockframe_environment(RTS.pop())
    },
LD: 
    instr => {
        PC++
        push(OS, heap_get_Environment_value(E, instr.pos))
    },
ASSIGN: 
    instr => {
        PC++
        heap_set_Environment_value(E, instr.pos, peek(OS,0))
    },
LDF: 
    instr => {
        PC++
        const closure_address = 
                  heap_allocate_Closure(
                      instr.addr, instr.arity, E)
        push(OS, closure_address)
    },
CALL: 
    instr => {
        const arity = instr.arity
        const fun = peek(OS, arity)
        if (is_Builtin(fun)) {
            PC++
            return apply_builtin(get_id_from_Builtin(fun))
        }
        const frame_address = heap_allocate_Frame(arity)
        for (let i = arity - 1; i >= 0; i--) {
            
            heap_set_Frame_value(frame_address, i, OS.pop())
        }
        OS.pop() // pop fun
        push(RTS, heap_allocate_Callframe(E, PC + 1))
        E = heap_Environment_extend(
                frame_address, 
                heap_get_Closure_environment(fun))
        PC = heap_get_Closure_pc(fun)
    },
TAIL_CALL: 
    instr => {
        const arity = instr.arity
        const fun = peek(OS, arity)
        if (is_True(is_Builtin(fun))) {
            PC++
            return apply_builtin(get_id_from_Builtin(fun))
        }
        const frame_address = heap_allocate_Frame(arity)
        for (let i = arity - 1; i >= 0; i--) {
            
            heap_set_Frame_value(frame_address, i, OS.pop())
        }
        OS.pop() // pop fun
        // dont push on RTS here
        E = heap_Environment_extend(
                frame_address, 
                heap_get_Closure_environment(fun))
        PC = heap_get_Closure_pc(fun)
    },
RESET : 
    instr => {
        // keep popping...
        const top_frame = RTS.pop()
        if (is_Callframe(top_frame)) {
            // ...until top frame is a call frame
            PC = heap_get_Callframe_pc(top_frame)
            E = heap_get_Callframe_environment(top_frame)
        }
    }
}

function run() {
    OS = []
    PC = 0
    E = global_environment
    RTS = []    
    //print_code()
    
    while (! (instrs[PC].tag === 'DONE')) {
        //display(PC, "PC: ")
        //print_OS("\noperands:            ");
        //print_RTS("\nRTS:            ");
        const instr = instrs[PC]
        //display(instrs[PC].tag, "next instruction: ")
        microcode[instr.tag](instr)
    }
    //display(OS, "\nfinal operands:           ")
    //print_OS()
    //display(word_to_string(peek(OS, 0)))
    return peek(OS, 0)
} 

// debugging

const print_code = () => {
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
                    word_to_JS_value(val) 
                    )
    }
}

// parse_compile_run on top level
// * parse input to json syntax tree
// * compile syntax tree into code
// * run code

const parse_compile_run = program => {
    compile_program(parse_to_json(program))
    return run()
}

//
// testing
//

const test = (program, expected) => {
    display("", `
    
****************
Test case: ` + program + "\n")
    const result = parse_compile_run(program)
    if (stringify(word_to_JS_value(result)) === stringify(expected)) {
        display(word_to_JS_value(result), "success:")
    } else {
        display(expected, "FAILURE! expected:")
        error(word_to_JS_value(result), "result:")
    }
}

// test("'x'+'y';", 'xy')