
//line parser.rl:1
package uri

import (
	"fmt"
)


//line parser.rl:8

//line parser_rl.go:13
const uri_start int = 1
const uri_first_final int = 120
const uri_error int = 0

const uri_en_uri int = 1


//line parser.rl:9

func RagelParse(str string) (*URI, error) {
	uri := &URI{}
	data := str
	cs := 0
	limit := len(data)
	p := 0 // data pointer
	m := 0 // marker for matching start position
	pe := limit // data end pointer
	eof := limit // End of data

//line parser.rl:59

  
//line parser_rl.go:36
	{
	cs = uri_start
	}

//line parser.rl:61
	
//line parser_rl.go:43
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 120:
		goto st_case_120
	case 17:
		goto st_case_17
	case 121:
		goto st_case_121
	case 18:
		goto st_case_18
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 19:
		goto st_case_19
	case 127:
		goto st_case_127
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 128:
		goto st_case_128
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 129:
		goto st_case_129
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 133:
		goto st_case_133
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 134:
		goto st_case_134
	case 82:
		goto st_case_82
	case 135:
		goto st_case_135
	case 83:
		goto st_case_83
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 84:
		goto st_case_84
	case 141:
		goto st_case_141
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 142:
		goto st_case_142
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 143:
		goto st_case_143
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 144:
		goto st_case_144
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 145:
		goto st_case_145
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 146:
		goto st_case_146
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	}
	goto st_out
	st_case_1:
		if data[p] == 115 {
			goto st2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 105 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 112 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 58:
			goto tr4
		case 115:
			goto st119
		}
		goto st0
tr4:
//line parser.rl:21
 uri.scheme   = SIP      
	goto st5
tr128:
//line parser.rl:22
 uri.scheme   = SIPS     
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser_rl.go:403
		switch data[p] {
		case 33:
			goto tr6
		case 37:
			goto tr7
		case 59:
			goto tr6
		case 61:
			goto tr6
		case 63:
			goto tr6
		case 91:
			goto tr10
		case 95:
			goto tr6
		case 126:
			goto tr6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto tr6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr9
				}
			case data[p] >= 65:
				goto tr9
			}
		default:
			goto tr8
		}
		goto st0
tr6:
//line parser.rl:20
 m = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser_rl.go:449
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 58:
			goto st9
		case 61:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		goto st0
tr7:
//line parser.rl:20
 m = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser_rl.go:488
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 61:
			goto st9
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st9
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st11
			}
		default:
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st9
			}
		default:
			goto st9
		}
		goto st0
tr14:
//line parser.rl:23
 uri.userinfo = str[m:p] 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser_rl.go:602
		if data[p] == 91 {
			goto tr10
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr18
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto st0
tr18:
//line parser.rl:20
 m = p 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parser_rl.go:628
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st32
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 45 {
			goto st14
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st16
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
tr19:
//line parser.rl:20
 m = p 
	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line parser_rl.go:720
		switch data[p] {
		case 45:
			goto st17
		case 46:
			goto st121
		case 58:
			goto st18
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st120
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 45 {
			goto st17
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st120
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 58:
			goto st18
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st123
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st124
		}
		goto st0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st126
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		goto st0
tr131:
//line parser.rl:24
 uri.hostport = str[m:p] 
//line parser.rl:20
 m = p 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser_rl.go:885
		switch data[p] {
		case 33:
			goto st127
		case 37:
			goto st20
		case 93:
			goto st127
		case 95:
			goto st127
		case 126:
			goto st127
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto st127
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st127
				}
			case data[p] >= 65:
				goto st127
			}
		default:
			goto st127
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 33:
			goto st127
		case 37:
			goto st20
		case 59:
			goto st19
		case 61:
			goto st22
		case 63:
			goto tr139
		case 93:
			goto st127
		case 95:
			goto st127
		case 126:
			goto st127
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto st127
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st127
				}
			case data[p] >= 65:
				goto st127
			}
		default:
			goto st127
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st21
			}
		default:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st127
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st127
			}
		default:
			goto st127
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 33:
			goto st128
		case 37:
			goto st23
		case 93:
			goto st128
		case 95:
			goto st128
		case 126:
			goto st128
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto st128
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st128
				}
			case data[p] >= 65:
				goto st128
			}
		default:
			goto st128
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 33:
			goto st128
		case 37:
			goto st23
		case 59:
			goto st19
		case 63:
			goto tr139
		case 93:
			goto st128
		case 95:
			goto st128
		case 126:
			goto st128
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto st128
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st128
				}
			case data[p] >= 65:
				goto st128
			}
		default:
			goto st128
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st24
			}
		default:
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st128
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st128
			}
		default:
			goto st128
		}
		goto st0
tr132:
//line parser.rl:24
 uri.hostport = str[m:p] 
//line parser.rl:20
 m = p 
//line parser.rl:25
 uri.params   = str[m:p] 
	goto st25
tr139:
//line parser.rl:25
 uri.params   = str[m:p] 
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parser_rl.go:1120
		switch data[p] {
		case 33:
			goto tr34
		case 36:
			goto tr34
		case 37:
			goto tr35
		case 63:
			goto tr34
		case 93:
			goto tr34
		case 95:
			goto tr34
		case 126:
			goto tr34
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto tr34
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr34
				}
			case data[p] >= 65:
				goto tr34
			}
		default:
			goto tr34
		}
		goto st0
tr34:
//line parser.rl:20
 m = p 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parser_rl.go:1164
		switch data[p] {
		case 33:
			goto st26
		case 36:
			goto st26
		case 37:
			goto st27
		case 61:
			goto st129
		case 63:
			goto st26
		case 93:
			goto st26
		case 95:
			goto st26
		case 126:
			goto st26
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto st26
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st26
				}
			case data[p] >= 65:
				goto st26
			}
		default:
			goto st26
		}
		goto st0
tr35:
//line parser.rl:20
 m = p 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parser_rl.go:1210
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st28
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st26
			}
		default:
			goto st26
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 33:
			goto st129
		case 37:
			goto st29
		case 38:
			goto st31
		case 63:
			goto st129
		case 93:
			goto st129
		case 95:
			goto st129
		case 126:
			goto st129
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto st129
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st129
				}
			case data[p] >= 65:
				goto st129
			}
		default:
			goto st129
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st30
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st30
			}
		default:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st129
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st129
			}
		default:
			goto st129
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 33:
			goto st26
		case 36:
			goto st26
		case 37:
			goto st27
		case 63:
			goto st26
		case 93:
			goto st26
		case 95:
			goto st26
		case 126:
			goto st26
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto st26
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto st26
				}
			case data[p] >= 65:
				goto st26
			}
		default:
			goto st26
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st33
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st36
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st130
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st120
			}
		default:
			goto st120
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st16
		case 58:
			goto st18
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st131
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st16
		case 58:
			goto st18
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st132
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st16
		case 58:
			goto st18
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st36
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st38
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st36
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st40
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st32
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 45:
			goto st14
		case 46:
			goto st32
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
tr10:
//line parser.rl:20
 m = p 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line parser_rl.go:1701
		if data[p] == 58 {
			goto st77
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st44
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st44
			}
		default:
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 58:
			goto st48
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st45
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st45
			}
		default:
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 58:
			goto st48
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st46
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st46
			}
		default:
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 58:
			goto st48
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st47
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 58:
			goto st48
		case 93:
			goto st133
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 58 {
			goto st64
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st49
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st44
			}
		default:
			goto st44
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 46:
			goto st50
		case 58:
			goto st48
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st62
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st45
			}
		default:
			goto st45
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if 48 <= data[p] && data[p] <= 57 {
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if data[p] == 46 {
			goto st52
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st60
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if 48 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 46 {
			goto st54
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st58
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if 48 <= data[p] && data[p] <= 57 {
			goto st55
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 93 {
			goto st133
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 93 {
			goto st133
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 93 {
			goto st133
		}
		goto st0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 58:
			goto st18
		case 59:
			goto tr131
		case 63:
			goto tr132
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 46 {
			goto st54
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 46 {
			goto st54
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if data[p] == 46 {
			goto st52
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 46 {
			goto st52
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 46:
			goto st50
		case 58:
			goto st48
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st63
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st46
			}
		default:
			goto st46
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 46:
			goto st50
		case 58:
			goto st48
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st47
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 58:
			goto st73
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st65
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st65
			}
		default:
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 58:
			goto st69
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st66
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st66
			}
		default:
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 58:
			goto st69
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st67
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st67
			}
		default:
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 58:
			goto st69
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st68
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st68
			}
		default:
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 58:
			goto st69
		case 93:
			goto st133
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st70
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st65
			}
		default:
			goto st65
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 46:
			goto st50
		case 58:
			goto st69
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st71
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st66
			}
		default:
			goto st66
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 46:
			goto st50
		case 58:
			goto st69
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st72
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st67
			}
		default:
			goto st67
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 46:
			goto st50
		case 58:
			goto st69
		case 93:
			goto st133
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st68
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st68
			}
		default:
			goto st68
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if 48 <= data[p] && data[p] <= 57 {
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 46 {
			goto st50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st75
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 46 {
			goto st50
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st76
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		if data[p] == 46 {
			goto st50
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		if data[p] == 58 {
			goto st64
		}
		goto st0
tr8:
//line parser.rl:20
 m = p 
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line parser_rl.go:2305
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st108
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st117
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st81
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st80
		}
		goto st0
tr9:
//line parser.rl:20
 m = p 
	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line parser_rl.go:2492
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st82
		case 46:
			goto st135
		case 58:
			goto st83
		case 59:
			goto tr146
		case 61:
			goto st6
		case 63:
			goto tr147
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st134
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st82
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st134
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 58:
			goto st83
		case 59:
			goto tr146
		case 61:
			goto st6
		case 63:
			goto tr147
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st80
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 61:
			goto st9
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st136
		}
		goto st0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 59:
			goto tr131
		case 61:
			goto st9
		case 63:
			goto tr132
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st137
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 59:
			goto tr131
		case 61:
			goto st9
		case 63:
			goto tr132
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st138
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 59:
			goto tr131
		case 61:
			goto st9
		case 63:
			goto tr132
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st139
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 59:
			goto tr131
		case 61:
			goto st9
		case 63:
			goto tr132
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st140
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 33:
			goto st9
		case 37:
			goto st10
		case 59:
			goto tr131
		case 61:
			goto st9
		case 63:
			goto tr132
		case 64:
			goto tr14
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st9
		}
		goto st0
tr146:
//line parser.rl:24
 uri.hostport = str[m:p] 
//line parser.rl:20
 m = p 
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line parser_rl.go:2876
		switch data[p] {
		case 33:
			goto st141
		case 37:
			goto st85
		case 44:
			goto st6
		case 58:
			goto st142
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 91:
			goto st127
		case 93:
			goto st127
		case 95:
			goto st141
		case 126:
			goto st141
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st141
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st141
			}
		default:
			goto st141
		}
		goto st0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 33:
			goto st141
		case 37:
			goto st85
		case 44:
			goto st6
		case 58:
			goto st142
		case 59:
			goto st84
		case 61:
			goto st92
		case 63:
			goto tr154
		case 64:
			goto tr14
		case 91:
			goto st127
		case 93:
			goto st127
		case 95:
			goto st141
		case 126:
			goto st141
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st141
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st141
			}
		default:
			goto st141
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st86
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st86
			}
		default:
			goto st86
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st141
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st141
			}
		default:
			goto st141
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 33:
			goto st142
		case 37:
			goto st87
		case 44:
			goto st9
		case 47:
			goto st127
		case 58:
			goto st127
		case 59:
			goto st19
		case 61:
			goto st89
		case 63:
			goto tr139
		case 64:
			goto tr14
		case 91:
			goto st127
		case 93:
			goto st127
		case 95:
			goto st142
		case 126:
			goto st142
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st142
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st142
			}
		default:
			goto st142
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st88
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st88
			}
		default:
			goto st88
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st142
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st142
			}
		default:
			goto st142
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 33:
			goto st143
		case 37:
			goto st90
		case 44:
			goto st9
		case 47:
			goto st128
		case 58:
			goto st128
		case 61:
			goto st9
		case 64:
			goto tr14
		case 91:
			goto st128
		case 93:
			goto st128
		case 95:
			goto st143
		case 126:
			goto st143
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st143
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st143
			}
		default:
			goto st143
		}
		goto st0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 33:
			goto st143
		case 37:
			goto st90
		case 44:
			goto st9
		case 47:
			goto st128
		case 58:
			goto st128
		case 59:
			goto st19
		case 61:
			goto st9
		case 63:
			goto tr139
		case 64:
			goto tr14
		case 91:
			goto st128
		case 93:
			goto st128
		case 95:
			goto st143
		case 126:
			goto st143
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st143
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st143
			}
		default:
			goto st143
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st91
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st91
			}
		default:
			goto st91
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st143
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st143
			}
		default:
			goto st143
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 33:
			goto st144
		case 37:
			goto st93
		case 44:
			goto st6
		case 58:
			goto st143
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 91:
			goto st128
		case 93:
			goto st128
		case 95:
			goto st144
		case 126:
			goto st144
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st144
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st144
			}
		default:
			goto st144
		}
		goto st0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 33:
			goto st144
		case 37:
			goto st93
		case 44:
			goto st6
		case 58:
			goto st143
		case 59:
			goto st84
		case 61:
			goto st6
		case 63:
			goto tr154
		case 64:
			goto tr14
		case 91:
			goto st128
		case 93:
			goto st128
		case 95:
			goto st144
		case 126:
			goto st144
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st144
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st144
			}
		default:
			goto st144
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st94
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st94
			}
		default:
			goto st94
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st144
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st144
			}
		default:
			goto st144
		}
		goto st0
tr147:
//line parser.rl:24
 uri.hostport = str[m:p] 
//line parser.rl:20
 m = p 
//line parser.rl:25
 uri.params   = str[m:p] 
	goto st95
tr154:
//line parser.rl:25
 uri.params   = str[m:p] 
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line parser_rl.go:3343
		switch data[p] {
		case 33:
			goto tr105
		case 37:
			goto tr106
		case 38:
			goto st6
		case 44:
			goto st6
		case 58:
			goto tr107
		case 59:
			goto st6
		case 61:
			goto st6
		case 64:
			goto tr14
		case 91:
			goto tr34
		case 93:
			goto tr34
		case 95:
			goto tr105
		case 126:
			goto tr105
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 57 {
				goto tr105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr105
			}
		default:
			goto tr105
		}
		goto st0
tr105:
//line parser.rl:20
 m = p 
	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line parser_rl.go:3392
		switch data[p] {
		case 33:
			goto st96
		case 37:
			goto st97
		case 38:
			goto st6
		case 44:
			goto st6
		case 58:
			goto st99
		case 59:
			goto st6
		case 61:
			goto st146
		case 64:
			goto tr14
		case 91:
			goto st26
		case 93:
			goto st26
		case 95:
			goto st96
		case 126:
			goto st96
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 57 {
				goto st96
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
tr106:
//line parser.rl:20
 m = p 
	goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line parser_rl.go:3441
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st98
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st98
			}
		default:
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st96
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
tr107:
//line parser.rl:20
 m = p 
	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line parser_rl.go:3482
		switch data[p] {
		case 33:
			goto st99
		case 37:
			goto st100
		case 38:
			goto st9
		case 44:
			goto st9
		case 47:
			goto st26
		case 58:
			goto st26
		case 61:
			goto st145
		case 63:
			goto st26
		case 64:
			goto tr14
		case 91:
			goto st26
		case 93:
			goto st26
		case 95:
			goto st99
		case 126:
			goto st99
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st99
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st101
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st101
			}
		default:
			goto st101
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st99
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 33:
			goto st145
		case 37:
			goto st102
		case 38:
			goto st104
		case 44:
			goto st9
		case 47:
			goto st129
		case 58:
			goto st129
		case 61:
			goto st9
		case 63:
			goto st129
		case 64:
			goto tr14
		case 91:
			goto st129
		case 93:
			goto st129
		case 95:
			goto st145
		case 126:
			goto st145
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st145
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st145
			}
		default:
			goto st145
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st103
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st103
			}
		default:
			goto st103
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st145
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st145
			}
		default:
			goto st145
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 33:
			goto st99
		case 37:
			goto st100
		case 38:
			goto st9
		case 44:
			goto st9
		case 47:
			goto st26
		case 58:
			goto st26
		case 61:
			goto st9
		case 63:
			goto st26
		case 64:
			goto tr14
		case 91:
			goto st26
		case 93:
			goto st26
		case 95:
			goto st99
		case 126:
			goto st99
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto st99
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st99
			}
		default:
			goto st99
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 33:
			goto st146
		case 37:
			goto st105
		case 38:
			goto st107
		case 44:
			goto st6
		case 58:
			goto st145
		case 59:
			goto st6
		case 61:
			goto st6
		case 64:
			goto tr14
		case 91:
			goto st129
		case 93:
			goto st129
		case 95:
			goto st146
		case 126:
			goto st146
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 57 {
				goto st146
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st146
			}
		default:
			goto st146
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st106
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st106
			}
		default:
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st146
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st146
			}
		default:
			goto st146
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 33:
			goto st96
		case 37:
			goto st97
		case 38:
			goto st6
		case 44:
			goto st6
		case 58:
			goto st99
		case 59:
			goto st6
		case 61:
			goto st6
		case 64:
			goto tr14
		case 91:
			goto st26
		case 93:
			goto st26
		case 95:
			goto st96
		case 126:
			goto st96
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 57 {
				goto st96
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st109
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st110
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st115
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st112
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st113
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st134
				}
			case data[p] >= 65:
				goto st134
			}
		default:
			goto st147
		}
		goto st0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st81
		case 58:
			goto st83
		case 59:
			goto tr146
		case 61:
			goto st6
		case 63:
			goto tr147
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st148
		}
		goto st0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st81
		case 58:
			goto st83
		case 59:
			goto tr146
		case 61:
			goto st6
		case 63:
			goto tr147
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st149
		}
		goto st0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st81
		case 58:
			goto st83
		case 59:
			goto tr146
		case 61:
			goto st6
		case 63:
			goto tr147
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st112
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st112
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st110
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st116
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st110
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st108
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st118
		}
		goto st0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 33:
			goto st6
		case 37:
			goto st7
		case 45:
			goto st79
		case 46:
			goto st108
		case 58:
			goto st9
		case 59:
			goto st6
		case 61:
			goto st6
		case 63:
			goto st6
		case 64:
			goto tr14
		case 95:
			goto st6
		case 126:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 47 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st80
				}
			case data[p] >= 65:
				goto st80
			}
		default:
			goto st80
		}
		goto st0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		if data[p] == 58 {
			goto tr128
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 127, 128, 141, 142, 143, 144:
//line parser.rl:25
 uri.params   = str[m:p] 
		case 129, 145, 146:
//line parser.rl:26
 uri.headers  = str[m:p] 
		case 120, 121, 122, 123, 124, 125, 126, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 147, 148, 149:
//line parser.rl:24
 uri.hostport = str[m:p] 
//line parser.rl:20
 m = p 
//line parser.rl:25
 uri.params   = str[m:p] 
//line parser_rl.go:4633
		}
	}

	_out: {}
	}

//line parser.rl:62

	if cs >= uri_first_final {
		return uri, nil
	}
	return nil, fmt.Errorf("Invalid URI '%s'.", str)
}

/* vim: set filetype=go : */
