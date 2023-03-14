package machine

var NaN_tag           = [4]byte{0,0,0,0}

// heap words without payload (literal values)
var False_tag         = [4]byte{1,0,0,0}
var True_tag          = [4]byte{1,0,0,1}
var Null_tag          = [4]byte{1,0,1,0}
var Unassigned_tag    = [4]byte{1,0,1,1}
var Undefined_tag     = [4]byte{1,1,0,0}

// heap nodes with child nodes
var Blockframe_tag    = [4]byte{1,1,0,1}
var Callframe_tag     = [4]byte{1,1,1,0}
var Closure_tag       = [4]byte{0,0,0,1}
var Frame_tag         = [4]byte{0,0,1,0}
var Environment_tag   = [4]byte{0,0,1,1}
var Pair_tag          = [4]byte{0,1,0,0}

// heap words with payload
var Builtin_tag       = [4]byte{0,1,0,1}
var Address_tag       = [4]byte{0,1,1,0}

var String_tag        = [4]byte{1,1,1,1}

const NaN_tag_offset = 13